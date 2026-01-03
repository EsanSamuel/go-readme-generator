package main

import (
	"fmt"
	"os"
	"path/filepath"
)

type ReadmeIgnore struct {
	SkipDirectories []string
	SkipFiles       []string
}

func AddReadmeIgnore() ReadmeIgnore {
	ignore := ReadmeIgnore{
		SkipDirectories: []string{
			".git",
			"node_modules",
			"vendor",
			"dist",
			"build",
			"out",
			"bin",
			"obj",
			"target",
			"coverage",
			".vscode",
			".idea",
			".cache",
			".tmp",
			".temp",
			"tmp",
		},
		SkipFiles: []string{
			".env",
			".env.*",
			".gitignore",
			"yarn.lock",
			"pnpm-lock.yaml",
			"package-lock.json",
			"bun.lockb",
			"go.sum",
			".DS_Store",
			"*.log",
			"*.pem",
			"*.key",
			"*.crt",
			"*.zip",
		},
	}

	return ignore
}

func WalkDir(root string, ignore ReadmeIgnore, contentChan chan<- string) (string, error) {
	entries, err := os.ReadDir(root)
	if err != nil {
		return "", err
	}
	var content []byte
	for _, entry := range entries {
		fullpath := filepath.Join(root, entry.Name())
		//fmt.Println("Fullpath:", fullpath)
		if entry.IsDir() {
			skip := false
			for _, dir := range ignore.SkipDirectories {
				match, err := filepath.Match(dir, entry.Name())
				if err != nil {
					fmt.Println("invalid pattern:", dir)
					continue
				}
				if match {
					fmt.Println("skipping dir:", entry.Name())
					skip = true
					break
				}
			}

			if skip {
				continue
			}
			WalkDir(fullpath, ignore, contentChan)
		} else {
			skip := false
			for _, pattern := range ignore.SkipFiles {
				match, err := filepath.Match(pattern, entry.Name())
				if err != nil {
					fmt.Println("invalid pattern:", pattern)
					continue
				}
				if match {
					fmt.Println("skipping file:", entry.Name())
					skip = true
					break
				}
			}
			if skip {
				continue
			}
			file, _ := os.ReadFile(fullpath)
			//fmt.Println("File content:", string(file))
			content = append(content, file...)
		}
	}

	contentChan <- string(content)
	return string(content), nil
}

func (s *ReadmeIgnore) ConfigFolderToSkip(dirs []string) {
	for _, dir := range dirs {
		s.SkipDirectories = append(s.SkipDirectories, dir)
	}
	//fmt.Println(s.SkipDirectories)
}

func (s *ReadmeIgnore) ConfigFileToSkip(files []string) {
	for _, file := range files {
		s.SkipFiles = append(s.SkipFiles, file)
	}
	//fmt.Println(s.SkipFiles)
}
