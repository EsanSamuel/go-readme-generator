package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting root directory")
	}

	content, err := WalkDir(cwd)
	if err != nil {
		fmt.Println(err)
	}

	config, err := os.ReadFile(filepath.Join(cwd, "config.md"))
	//fmt.Println(string(config))

	if string(config) != "" {
		fmt.Println("Generating README file...")
		ai, err := Ai(string(content + string(config)))
		if err != nil {
			fmt.Println("Error generating readme with AI:", err.Error())
		}
		dst, _ := os.Create(filepath.Join(cwd, "README.md"))
		fmt.Println("Writing README file...")
		if _, err := dst.Write([]byte(ai)); err != nil {
			fmt.Println(err)
		}
		fmt.Println("README.md file generated successfully!")
	} else {
		fmt.Println("Generating README file...")
		ai, err := Ai(string(content))
		if err != nil {
			fmt.Println("Error generating readme with AI:", err.Error())
		}
		dst, _ := os.Create(filepath.Join(cwd, "README.md"))
		fmt.Println("Writing README file...")
		if _, err := dst.Write([]byte(ai)); err != nil {
			fmt.Println(err)
		}
		fmt.Println("README.md file generated successfully!")
	}

}
