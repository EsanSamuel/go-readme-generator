# go-readme-generator

## Description

The `go-readme-generator` is a Go-based command-line tool designed to automate the creation of high-quality `README.md` files for your projects. It achieves this by scanning the project's directory structure, collecting relevant file contents (while intelligently skipping specified files and folders), and then leveraging Google's Gemini Generative AI to craft a comprehensive README based on the collected data and an optional custom prompt.

## Features

*   **Automated Project Content Collection**: Recursively walks the project directory to gather file contents for AI processing.
*   **Configurable Exclusion List**: Comes with a default list of directories (e.g., `.git`, `node_modules`) and files (e.g., `.env`, `*.log`) to skip, preventing irrelevant data from being sent to the AI.
*   **User-Defined Exclusions**: Allows users to specify additional folders and files to skip during runtime via interactive prompts.
*   **Google Gemini AI Integration**: Utilizes the `gemini-2.5-flash` model via the `google.golang.org/genai` library to generate the README content.
*   **Customizable AI Prompt**: Supports an optional `prompt.md` file in the project root to provide additional, specific instructions or context to the AI for README generation.
*   **Output to `README.md`**: Generates and writes the AI-produced content directly into a `README.md` file in the current working directory.
*   **Environment Variable Support**: Loads the Gemini API key from a `.env` file.

## Installation

To install and run `go-readme-generator`, you need to have Go installed on your system.

1.  **Go Version**: Ensure you have Go version `1.25.5` or later installed.
    ```bash
    go version
    ```
2.  **Clone the Repository**:
    ```bash
    git clone https://github.com/EsanSamuel/Go_cli.git go-readme-generator
    cd go-readme-generator
    ```
3.  **Install Dependencies**: Go will automatically download the necessary dependencies defined in `go.mod`.
    ```bash
    go mod tidy
    ```
4.  **API Key Configuration**: Create a `.env` file in the root of the `go-readme-generator` directory with your Google Gemini API key.
    ```
    # .env
    GEMINI_API_KEY=YOUR_GEMINI_API_KEY_HERE
    ```
    You can obtain a `GEMINI_API_KEY` from the [Google AI Studio](https://ai.google.dev/).

## Usage

Follow these steps to generate a `README.md` for your project:

1.  **Navigate to your project directory**: Change your current working directory to the root of the project for which you want to generate a README.
2.  **Run the generator**:
    ```bash
    go run main.go
    ```
3.  **Interact with prompts**: The application will prompt you to enter folders and files you wish to skip. You can enter multiple items separated by spaces.
    ```
    Enter folders you want to skip: .vscode temp
    Enter files you want to skip: config.yaml test_data.json
    ```
    *Note: There are default folders and files already skipped (e.g., `.git`, `node_modules`, `.env`). Your input adds to this list.*
4.  **Optional: Provide a custom prompt**: If you want to give the AI specific instructions beyond just the project's code, create a file named `prompt.md` in your project's root directory. The content of this file will be appended to the code content sent to the AI.
    For example, `prompt.md` could contain:
    ```markdown
    ALSO TEACH PEOPLE HOW TO USE THIS PROJECT WITH WELL EXPLAINED DETAILS. The project name is now go-readme-generator and the config file is now prompt.md
    ```
    If `prompt.md` is present and has content, it will be used. Otherwise, the AI will generate the README based solely on the scanned project files.
5.  **Review the generated `README.md`**: Once the process is complete, a `README.md` file will be created or updated in your project's root directory with the AI-generated content.

    ```
    Generating README file...
    Writing README file...
    ✨ README.md file generated successfully!
    ```

## Folder Structure Explanation

The project has the following key files:

*   **`ai.go`**: Contains the core logic for interacting with the Google Gemini API. It handles loading the API key, configuring the AI client, and sending prompts to generate content.
*   **`go.mod`**: Defines the Go module path (`github.com/EsanSamuel/Go_cli`) and lists all direct and indirect dependencies required by the project, including `google.golang.org/genai` and `github.com/joho/godotenv`.
*   **`main.go`**: The entry point of the application. It orchestrates the entire README generation process, including getting user input, initiating directory scanning, reading the `prompt.md` (if present), calling the AI function, and writing the final `README.md` file.
*   **`prompt.md`**: An optional configuration file. If present in the project root, its content is appended to the scanned project code to provide additional instructions or context to the AI for README generation.
*   **`read.go`**: Manages the directory walking and file/folder exclusion logic. It defines the `ReadmeIgnore` struct and functions to configure default and user-specified items to skip during the scanning process.

## Technologies

*   **Go**: The primary programming language used for the entire application (version `1.25.5`).
*   **Google Generative AI (Gemini)**: Utilized through the `google.golang.org/genai` Go client library to power the AI-driven README generation. Specifically, the `gemini-2.5-flash` model is used.
*   **`github.com/joho/godotenv`**: A Go package used for loading environment variables from `.env` files.