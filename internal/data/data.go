package data

import (
	"encoding/json"
	"fmt"
	"strings"
)

//go:generate go-bindata -pkg data -prefix "dataFiles/" ./dataFiles/...

type Template struct {
	Name string
	Key  string

	Vars        map[string]string
	Directories []string
	Files       map[string]string
	Commands    []struct {
		Command string
		Stdin   []string
	}
}

const templateDirName = "templates"

var AvailableTemplates []*Template

func init() {
	list, err := AssetDir(templateDirName)
	if err != nil {
		panic(err)
	}
	for _, f := range list {
		tpl, err := LoadTemplate(strings.Join([]string{templateDirName, f}, "/"))
		if err != nil {
			panic(fmt.Errorf("data.init: (%s) %s", f, err.Error()))
		}
		AvailableTemplates = append(AvailableTemplates, tpl)
	}
}

func LoadTemplate(templateName string) (*Template, error) {
	fCont, err := Asset(templateName)
	if err != nil {
		return nil, err
	}

	loadedTemplate := new(Template)
	return loadedTemplate, json.Unmarshal(fCont, loadedTemplate)
}
