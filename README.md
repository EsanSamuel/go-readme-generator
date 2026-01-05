# Go README Generator CLI

## Description

This project provides a Command Line Interface (CLI) tool written in Go that automates the generation of `README.md` files. It works by scanning the current directory, collecting file and folder information, and then leveraging the Google Gemini AI model to generate a descriptive `README.md`. The tool is designed to help users document their projects by intelligently analyzing the code structure and content.

Specifically, it aims to design the file structure explanation and teach users how to use the CLI tool, as indicated by its internal prompt.

## Features

*   **Automated README Generation**: Scans project files and directories to gather information for README creation.
*   **Customizable File/Folder Exclusion**: Allows users to specify directories and files to skip during the scanning process, ensuring sensitive or irrelevant information is not included.
*   **AI-Powered Content Generation**: Integrates with the Google Gemini API (`gemini-2.5-flash` model) to create high-quality, contextual README content.
*   **Environment Variable Support**: Loads API keys from a `.env` file using `godotenv`.
*   **Optional Custom Prompt**: Supports an optional `prompt.md` file in the root directory to provide specific instructions or context to the AI for generating the README.
*   **Colored Output**: Provides visual feedback in the terminal with colored success messages.

## Installation

To install and run this CLI tool, you need to have Go (version 1.25.5 or later) installed on your system.

1.  **Clone the repository:**
    ```bash
    git clone github.com/EsanSamuel/go-readme-generator
    cd go-readme-generator
    ```

2.  **Set up environment variables:**
    Create a `.env` file in the root of the project and add your Google Gemini API key:
    ```
    GEMINI_API_KEY="YOUR_GEMINI_API_KEY"
    ```
    You can obtain a Gemini API key from the [Google AI Studio](https://ai.google.dev/).

3.  **Download dependencies and build:**
    ```bash
    go mod tidy
    go build -o go-readme-generator .
    ```
    This will create an executable named `go-readme-generator` (or `go-readme-generator.exe` on Windows) in the current directory.

## Usage

To generate a `README.md` file for your project:

1.  Navigate to the root directory of your project where you want the `README.md` to be generated.
2.  Run the executable:
    ```bash
    ./go-readme-generator
    ```

3.  The CLI will prompt you to enter folders and files you wish to skip:
    ```
    Enter folders you want to skip:
    ```
    Enter folder names separated by spaces (e.g., `node_modules .git build`). Press Enter when done.

    ```
    Enter files you want to skip:
    ```
    Enter file names or patterns separated by spaces (e.g., `.env *.log temp.txt`). Press Enter when done.

4.  The tool will then scan your project, communicate with the Gemini AI, and generate a `README.md` file in the current directory. You will see progress messages in the terminal:
    ```
    Generating README file...
    Writing README file...
    ✨ README.md file generated successfully!
    ```

**Using a Custom Prompt:**
You can guide the AI's README generation by placing a file named `prompt.md` in your project's root directory. The content of this file will be prepended to the AI's internal prompt, allowing you to provide specific instructions or context relevant to your project.

Example `prompt.md` content:
```
This project is a web server built with Gin Gonic. Focus on API endpoints.
```

## Folder Structure Explanation

*   `ai.go`: Contains the core logic for interacting with the Google Gemini API. It handles setting up the client, sending prompts, and receiving responses for README generation.
*   `main.go`: The main entry point of the CLI application. It orchestrates user input for ignore rules, initiates directory traversal, calls the AI function, and writes the generated `README.md` file.
*   `read.go`: Manages directory scanning, file content collection, and defines ignore rules. It includes functions to configure custom skip directories and files, and to recursively walk the file system.
*   `go.mod`: Defines the Go module path and lists all direct and indirect dependencies required by the project.
*   `prompt.md` (Optional): An external file that can be created by the user to provide custom instructions or context to the AI model during README generation.

## Technologies

*   **Go**: The programming language used for developing the CLI tool (Go 1.25.5).
*   **Google Gemini API (`google.golang.org/genai`)**: Utilized for generating text content for the README.
*   **godotenv (`github.com/joho/godotenv`)**: Used for loading environment variables from a `.env` file.