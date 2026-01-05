package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mbndr/figlet4go"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Println(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}

func main() {
	options := figlet4go.NewRenderOptions()
	options.FontColor = []figlet4go.Color{
		figlet4go.ColorCyan,
	}

	ascii := figlet4go.NewAsciiRender()

	renderStr, _ := ascii.RenderOpts("GO README GENERATOR", options)
	fmt.Print(renderStr)
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting root directory")
	}
	const (
		colorRed   = "\033[31m"
		colorGreen = "\033[32m"
		colorReset = "\033[0m"
	)
	ignore := AddReadmeIgnore()
	reader := bufio.NewReader(os.Stdin)
	directory, _ := getInput("Enter folders you want to skip:", reader)

	directories := strings.Split(directory, " ")

	ignore.ConfigFolderToSkip(directories)

	file, _ := getInput("Enter files you want to skip:", reader)

	files := strings.Split(file, " ")

	ignore.ConfigFileToSkip(files)

	result := make(chan string, 5)

	go func() {
		WalkDir(cwd, ignore, result)
		close(result) // Close after all sends are done
	}()

	content := <-result

	config, err := os.ReadFile(filepath.Join(cwd, "prompt.md"))
	//fmt.Println(string(config))

	if string(config) != "" {
		fmt.Println("Generating README file...")
		prompt := strings.TrimSpace(content + string(config))
		ai, err := Ai(prompt)
		if err != nil {
			fmt.Println("Error generating readme with AI:", err.Error())
		}
		dst, _ := os.Create(filepath.Join(cwd, "README.md"))
		fmt.Println("Writing README file...")
		if _, err := dst.Write([]byte(ai)); err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%s✨ README.md file generated successfully!%s\n", colorGreen, colorReset)
	} else {
		fmt.Println("Generating README file...")
		prompt := strings.TrimSpace(content)
		ai, err := Ai(prompt)
		if err != nil {
			fmt.Println("Error generating readme with AI:", err.Error())
		}
		dst, _ := os.Create(filepath.Join(cwd, "README.md"))
		fmt.Println("Writing README file...")
		if _, err := dst.Write([]byte(ai)); err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%s✨ README.md file generated successfully!%s\n", colorGreen, colorReset)

	}

}
