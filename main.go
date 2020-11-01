package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

var scanner = bufio.NewScanner(os.Stdin)

func prompt(prompt string) string {
	fmt.Print(prompt)
	scanner.Scan()
	r := scanner.Text()
	fmt.Println()
	return r
}

func promptMulti(prompt string) (items []string) {
	fmt.Println(prompt)
	fmt.Println("(Enter a blank line to continue)")
	for {
		fmt.Print("> ")
		scanner.Scan()
		inputted := scanner.Text()

		if inputted == "" {
			fmt.Println()
			return
		}

		items = append(items, inputted)
	}

}

func promptSelect(prompt string, options []string) (selected int, chosenItem string) {
	fmt.Println(prompt)
	for i, v := range options {
		fmt.Printf("  %d: %s\n", i+1, v)
	}
	for {
		fmt.Print("> ")
		scanner.Scan()
		userInput := scanner.Text()
		userSelected, err := strconv.Atoi(userInput)
		if err != nil {
			fmt.Println("Not an integer.")
			continue
		}

		userSelected -= 1

		if userSelected < 0 || userSelected >= len(options) {
			fmt.Println("Out of bounds.")
			continue
		}

		selected = userSelected
		chosenItem = options[userSelected]

		fmt.Println()

		return
	}
}

func httpGet(url string) (respBody []byte, err error) {

	var hasFetched bool

	go func() {
		time.Sleep(time.Second)
		if hasFetched {
			return
		} else {
			fmt.Printf("Loading %s...\n", url)
		}
	}()

	resp, err := http.Get(url)

	hasFetched = true

	if err != nil {
		return
	}

	respBody, err = ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	return
}

type licence struct {
	Key  string `json:"key"`
	Name string `json:"name"`
	Url  string `json:"url"`
	Body string `json:"body"`
}

var filesToWrite = make(map[string][]byte)

func main() {
	// --- Get licence to use ---

	// Get a list of available licences
	rawLicencesBody, err := httpGet("https://api.github.com/licenses")
	checkError(err)

	var availLicences []licence
	err = json.Unmarshal(rawLicencesBody, &availLicences)
	checkError(err)

	// Ask the user which one they want
	var licenceOptions []string
	for _, v := range availLicences {
		licenceOptions = append(licenceOptions, fmt.Sprintf("%s (%s)", v.Name, v.Key))
	}

	selectedLicenceIndex, _ := promptSelect("Select a licence:", licenceOptions)
	selectedLicence := availLicences[selectedLicenceIndex]

	// Fetch licence text
	rawSelectedLicenceResp, err := httpGet(selectedLicence.Url)
	checkError(err)

	err = json.Unmarshal(rawSelectedLicenceResp, &selectedLicence)
	checkError(err)

	if strings.Contains(selectedLicence.Body, "[year]") {
		selectedLicence.Body = strings.ReplaceAll(selectedLicence.Body, "[year]", strconv.Itoa(time.Now().Year()))
	}

	if strings.Contains(selectedLicence.Body, "[fullname]") {
		fmt.Println("This licence requires the name(s) of the copyright holder(s).")
		name := prompt("Enter copyright holder name: ")
		selectedLicence.Body = strings.ReplaceAll(selectedLicence.Body, "[fullname]", name)
	}

	filesToWrite["LICENCE"] = []byte(selectedLicence.Body)

	fmt.Println("LICENCE generated")
	fmt.Println()

	// --- Make .gitignore ---

	langs := promptMulti("Enter the languages you're going to be using (leave blank for no .gitignore):")

	if len(langs) != 0 {
		gitignoreContent, err := httpGet(fmt.Sprintf("https://www.toptal.com/developers/gitignore/api/%s", strings.Join(langs, ",")))
		checkError(err)

		filesToWrite[".gitignore"] = gitignoreContent

		fmt.Println(".gitignore generated")

	} else {
		fmt.Println("Skipping .gitignore creation...")
	}

	fmt.Println()

	// --- Make .github/README.md ---

	currentWorkingDirectory, err := os.Getwd()
	checkError(err)
	splitCwd := strings.Split(currentWorkingDirectory, string(filepath.Separator))
	currentDir := splitCwd[len(splitCwd)-1]

	filesToWrite[filepath.Join(".github", "README.md")] = []byte(fmt.Sprintf("# %s\n", currentDir))

	fmt.Println(".github/README.md generated")
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

		originUrl := prompt("Enter the URL of the origin remote (leave blank for none): ")

		fmt.Println("Initialising git...")

		err = exec.Command("git", "init").Run()
		if err != nil {
			_ = os.RemoveAll(".git")
			checkError(err)
		}

		if originUrl != "" {

			fmt.Println("Adding specified remote...")

			err = exec.Command("git", "remote", "add", "origin", originUrl).Run()
			checkError(err)
		}

	} else {
		fmt.Println(".git folder already found. git init will not be run.")
	}

	fmt.Println()

	// --- Create new files ---

	fmt.Println("Creating new files...")

	var filenames []string

	for filename, fileContent := range filesToWrite {
		path := strings.Split(filename, string(filepath.Separator))

		if len(path) > 1 { // if there's more than just the file name
			// If there's an error, let's just assume that means the directory already existed and move on with our day
			_ = os.MkdirAll(filepath.Join(path[:len(path)-1]...), os.ModePerm)
		}

		err = ioutil.WriteFile(filename, fileContent, 0644)
		checkError(err)

		filenames = append(filenames, filename)
	}

	// --- Commit ---

	fmt.Println("Tracking files...")

	err = exec.Command("git", append([]string{"add"}, filenames...)...).Run()
	checkError(err)

	fmt.Println("Committing files...")

	err = exec.Command("git", "commit", "-sm", "Initial commit").Run()
	checkError(err)

}
