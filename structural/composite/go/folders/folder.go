package main

import (
	"fmt"
	"strings"
)

type Folder struct {
	Name    string
	Content []Component
}

type File struct {
	Name string
}

func (folder Folder) GetPath(indent int) string {
	var sb strings.Builder
	indented := strings.Repeat(" ", indent)
	sb.WriteString(fmt.Sprintf("%s %s", indented, folder.Name))
	for _, component := range folder.Content {
		sb.WriteString(fmt.Sprintf("\n%s", component.GetPath(indent+2)))
	}
	return sb.String()
}

func (folder Folder) Search(target string) string {
	if folder.Name == target {
		return folder.Name
	}
	path := ""
	for _, component := range folder.Content {
		if found := component.Search(target); found != "" {
			path = fmt.Sprintf("%s/%s", folder.Name, found)
			break
		}
	}
	return path
}

func (file File) Search(target string) string {
	if target == file.Name {
		return file.Name
	}
	return ""
}

func (file File) GetPath(indent int) string {
	return fmt.Sprintf("%s %s", strings.Repeat(" ", indent), file.Name)
}
