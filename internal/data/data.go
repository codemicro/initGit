package data

import (
	"encoding/json"
	"fmt"
	"strings"
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
	AvailableTemplates []Template
	Gitignores         map[string]string
	Licences           []Licence
)

func init() {

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
		AvailableTemplates = append(AvailableTemplates, tpl)
	}

	LoadResource(gitignoresFile, &Gitignores)
	LoadResource(licencesFile, &Licences)

}

func LoadResource(filename string, out interface{}) error {
	fCont, err := Asset(filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(fCont, out)
}

func MakeFullGitignore(opts []string) (string, error) {
	var contents []string
	for _, opt := range opts {
		cont, ok := Gitignores[strings.ToLower(opt)]
		if ok {
			contents = append(contents, cont)
		}
	}
	return strings.Join(contents, "\n"), nil
}
