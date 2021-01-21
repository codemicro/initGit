package data

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/codemicro/initGit/internal/input"
)

//go:generate python ../../scripts/loadGitignores.py ./datafiles/gitignores.json
//go:generate go-bindata -pkg data -prefix "dataFiles/" ./dataFiles/...

type Template struct {
	Name string
	Key  string

	Vars []struct {
		Key         string
		Description string
	}
	Directories []string
	Files       map[string]string
	Commands    []struct {
		Command string
		Stdin   []string
	}
}

func GetTemplateVariableValues(t Template) map[string]string {
	o := make(map[string]string)
	for _, varDef := range t.Vars {
		o[varDef.Key] = input.Prompt(varDef.Description + ": ")
	}
	return o
}

type Licence struct {
	Spdx    string
	Name    string
	Content string
}

const (
	templateDirName = "templates"
	licencesFile    = "licences.json"
	gitignoresFile  = "gitignores.json"
)

var (
	Templates  []Template
	Gitignores map[string]string
	Licences   []Licence
)

func init() {

	// Load all templates
	list, err := AssetDir(templateDirName)
	if err != nil {
		panic(err)
	}
	for _, f := range list {
		var tpl Template
		err := LoadResource(strings.Join([]string{templateDirName, f}, "/"), &tpl)
		if err != nil {
			panic(fmt.Errorf("data.init: (%s) %s", f, err.Error()))
		}
		Templates = append(Templates, tpl)
	}

	// Load gitignore fragments and licences
	LoadResource(gitignoresFile, &Gitignores)
	LoadResource(licencesFile, &Licences)

}

// LoadResource loads a given filename from the in-memory data store and unmarshalls the JSON of it into a given thing
func LoadResource(filename string, out interface{}) error {
	fCont, err := Asset(filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(fCont, out)
}

// MakeFullGitignore selects gitignore fragments from the specified options and combines them into a single string
// Option keys that cannot be found are silently ignored
func MakeFullGitignore(opts []string) string {
	var contents []string
	for _, opt := range opts {
		cont, ok := Gitignores[strings.ToLower(opt)]
		if ok {
			contents = append(contents, cont)
		}
	}
	return strings.Join(contents, "\n")
}
