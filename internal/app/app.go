package app

import (
	"bytes"
	"fmt"
	"github.com/codemicro/initGit/internal/directoryTree"
	"github.com/codemicro/initGit/internal/substitutions"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/codemicro/initGit/internal/data"
	"github.com/codemicro/initGit/internal/input"
)

func Run() {

	filesToWrite := make(map[string][]byte)
	var directoriesToMake []string
	var createdFilenames []string

	// --- Get licence to use ---
	spdx := getLicense(filesToWrite)
	fmt.Println()
	// --- Make .gitignore ---
	getGitignore(filesToWrite)
	fmt.Println()
	// --- Prompt for template ---
	template := pickTemplate()
	if template != nil {
		createdFilenames = append(createdFilenames, executeTemplate(*template, filesToWrite, &directoriesToMake, spdx)...)
	}
	fmt.Println()
	// --- Check if there's already a .git directory ---
	var gitDirAlreadyExists bool
	if _, err := os.Stat(".git"); err != nil {
		if os.IsNotExist(err) {
		} else {
			gitDirAlreadyExists = true
		}
	}
	// --- Init repo ---
	if !gitDirAlreadyExists {
		runGitInit()
	} else {
		fmt.Println(".git directory already found. git init will not be run.")
	}
	fmt.Println()
	// --- Create new files ---
	createdFilenames = append(createdFilenames, createDiskObjects(filesToWrite, directoriesToMake)...)
	fmt.Println()
	// --- Commit ---
	runGitTrackAndCommit(createdFilenames)
}

func getLicense(filesToWrite map[string][]byte) string {
	// Ask the user which licence they want
	var licenceOptions []string
	for _, v := range data.Licences {
		licenceOptions = append(licenceOptions, fmt.Sprintf("%s (%s)", v.Name, v.Spdx))
	}

	selectedLicenceIndex, _ := input.PromptSelect("Select a licence (leave blank for none):", licenceOptions, true)

	if selectedLicenceIndex != -1 {

		selectedLicence := data.Licences[selectedLicenceIndex]

		// Replace year and name sections if applicable

		if strings.Contains(selectedLicence.Content, "{year}") {
			selectedLicence.Content = strings.ReplaceAll(selectedLicence.Content, "{year}", strconv.Itoa(time.Now().Year()))
		}

		if strings.Contains(selectedLicence.Content, "{name}") {
			fmt.Println("\nThis licence requires the name(s) of the copyright holder(s).")
			name := input.Prompt("Enter copyright holder name: ")
			selectedLicence.Content = strings.ReplaceAll(selectedLicence.Content, "{name}", name)
		}

		filesToWrite["LICENCE"] = []byte(selectedLicence.Content)

		fmt.Println("LICENCE generated")

		return selectedLicence.Spdx
	}
	return ""
}

func getGitignore(filesToWrite map[string][]byte) {
	langs := input.PromptMulti("Enter the languages you're going to be using (leave blank for no .gitignore):")
	if len(langs) != 0 {
		filesToWrite[".gitignore"] = []byte(data.MakeFullGitignore(langs))
		fmt.Println(".gitignore generated")
	} else {
		fmt.Println("Skipping .gitignore creation...")
	}
}

func pickTemplate() *data.Template {
	// Ask the user which licence they want
	var templateOptions []string
	for _, v := range data.Templates {
		templateOptions = append(templateOptions, fmt.Sprintf("%s", v.Name))
	}

	selectedTemplateIndex, _ := input.PromptSelect("Select a project template (leave blank for none):", templateOptions, true)
	if selectedTemplateIndex == -1 {
		return nil
	}
	selectedTemplate := data.Templates[selectedTemplateIndex]
	return &selectedTemplate
}

func executeTemplate(template data.Template, filesToWrite map[string][]byte, directoriesToMake *[]string, spdxId string) (filenames []string) {

	currentWorkingDirectory, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	splitCwd := strings.Split(currentWorkingDirectory, string(filepath.Separator))
	currentDir := splitCwd[len(splitCwd)-1]

	variableVals := make(map[string]string)
	variableVals["spdxLicense"] = spdxId
	variableVals["dirName"] = currentDir

	variableVals = data.GetTemplateVariableValues(template, variableVals)

	// make directories
	for _, dirName := range template.Directories {
		*directoriesToMake = append(*directoriesToMake, substitutions.SubVariables(dirName, variableVals))
	}
	// make files
	for fName, fContent := range template.Files {
		filesToWrite[substitutions.SubVariables(fName, variableVals)] = []byte(substitutions.SubVariables(fContent, variableVals))
	}
	// run commands
	for _, command := range template.Commands {
		modCommand := strings.Split(substitutions.SubVariables(command.Command, variableVals), " ")
		cmd := exec.Command(modCommand[0], modCommand[1:]...)

		fmt.Printf("Running %s", strings.Join(modCommand, " "))

		if len(command.Stdin) != 0 {
			wc, err := cmd.StdinPipe()
			if err != nil {
				panic(err)
			}
			go func() {
				defer wc.Close()
				_, _ = io.WriteString(wc, substitutions.SubVariables(strings.Join(command.Stdin, "\n"), variableVals))
			}()
		}


		diff, err := directoryTree.NewDifference(currentWorkingDirectory)
		if err != nil {
			panic(err)
		}

		cmd.Stderr = new(bytes.Buffer)
		err = cmd.Run()
		if err != nil {
			fmt.Printf(" - failed to run command (exit code %d)\n%s\n", cmd.ProcessState.ExitCode(), string(cmd.Stderr.(*bytes.Buffer).Bytes()))
		} else {
			fmt.Println()
		}

		x, err := diff.Get()
		if err != nil {
			panic(err)
		}
		filenames = append(filenames, x...)
	}
	return
}

func runGitInit() {
	originUrl := input.Prompt("Enter the URL of the origin remote (leave blank for none): ")

	fmt.Println("Initialising git...")

	err := exec.Command("git", "init").Run()
	if err != nil {
		_ = os.RemoveAll(".git")
		panic(err)
	}

	if originUrl != "" {

		fmt.Println("Adding specified remote...")

		err = exec.Command("git", "remote", "add", "origin", originUrl).Run()
		if err != nil {
			panic(err)
		}
	}
}

func createDiskObjects(filesToWrite map[string][]byte, directoriesToMake []string) (filenames []string) {
	fmt.Println("Creating new directories...")
	for _, dirname := range directoriesToMake {
		// If there's an error, let's just assume that means the directory already existed and move on with our day
		_ = os.MkdirAll(dirname, os.ModePerm)
	}

	fmt.Println("Creating new files...")

	for filename, fileContent := range filesToWrite {
		filename = strings.Trim(filename, string(filepath.Separator))
		filename = strings.Trim(filename, "/")
		path := strings.Split(strings.ReplaceAll(filename, string(filepath.Separator), "/"), "/")

		if len(path) > 1 { // if there's more than just the file name
			_ = os.MkdirAll(filepath.Join(path[:len(path)-1]...), os.ModePerm)
		}

		err := ioutil.WriteFile(filename, fileContent, 0644)
		if err != nil {
			panic(err)
		}

		filenames = append(filenames, filename)
	}

	return filenames
}

func runGitTrackAndCommit(filenames []string) {
	fmt.Println("Tracking files...")
	err := exec.Command("git", append([]string{"add"}, filenames...)...).Run()
	if err != nil {
		panic(err)
	}

	fmt.Println("Committing files...")
	err = exec.Command("git", "commit", "-sm", "Setup commit").Run()
	if err != nil {
		panic(err)
	}
}
