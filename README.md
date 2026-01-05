# Go README Generator

## Description

This project is a Command-Line Interface (CLI) tool written in Go designed to automate the creation of high-quality `README.md` files. It works by intelligently scanning a specified project directory, collecting relevant file paths and their contents (while respecting both default and user-defined ignore lists). The collected information is then fed to the Google Gemini AI API, which generates a comprehensive `README.md` based on the project's structure and content, optionally augmented by a custom `prompt.md` file.

## Features

*   **AI-Powered README Generation**: Leverages the Google Gemini API to intelligently analyze project content and generate descriptive `README.md` files.
*   **Configurable Ignore Lists**: Users can specify folders and files to be excluded from the scanning process at runtime.
*   **Default Ignore Rules**: Includes a predefined set of common directories (e.g., `.git`, `node_modules`) and files (e.g., `.env`, `go.sum`) that are automatically skipped.
*   **Customizable AI Prompt**: Supports an optional `prompt.md` file in the root directory to provide specific instructions or context to the AI model for README generation.
*   **Project Directory Traversal**: Efficiently walks through the project's file system to gather necessary information.
*   **Interactive Input**: Prompts the user for ignore preferences during execution.
*   **Go-based CLI**: Built entirely in Go for performance and ease of deployment.
*   **ASCII Art Greeting**: Displays a stylized "GO README GENERATOR" message upon startup.

## Installation

To get started with the Go README Generator, follow these steps:

### Prerequisites

*   **Go**: Ensure you have Go version `1.25.5` or higher installed. You can download it from [golang.org](https://golang.org/dl/).

### Setup

1.  **Clone the Repository**:
    ```bash
    git clone github.com/EsanSamuel/go-readme-generator
    cd go-readme-generator
    ```

2.  **Configure Gemini API Key**:
    The tool requires an API key for the Google Gemini API.
    *   Obtain your `GEMINI_API_KEY` from the [Google AI Studio](https://aistudio.google.com/app/apikey).
    *   Create a `.env` file in the root directory of the `go-readme-generator` project and add your API key:
        ```env
        GEMINI_API_KEY="YOUR_GEMINI_API_KEY_HERE"
        ```

3.  **Install Dependencies**:
    The Go modules will be automatically downloaded when you run the application or you can explicitly download them:
    ```bash
    go mod tidy
    ```

## Usage

After installation and setup, you can run the CLI tool:

1.  **Run the application**:
    Navigate to the project's root directory in your terminal and execute:
    ```bash
    go run .
    ```

2.  **Interactive Prompts**:
    The CLI will guide you through the process:
    *   It will first display a "GO README GENERATOR" ASCII art.
    *   You will be prompted to "Enter folders you want to skip:". You can enter space-separated folder names (e.g., `tmp logs backup`). Press Enter if you have no additional folders to skip.
    *   Next, you will be prompted to "Enter files you want to skip:". Enter space-separated file names or patterns (e.g., `test.txt *.bak`). Press Enter if you have no additional files to skip.

3.  **Optional `prompt.md`**:
    For more tailored README generation, you can create a file named `prompt.md` in the same directory where you run the tool (the current working directory). The content of this file will be appended to the AI's prompt, allowing you to provide specific instructions or context. For example:
    ```markdown
    DESIGN ALL FILE STRUCTURE AND TEACH USERS HOW TO USE THE CLI TOOL
    ```

4.  **Output**:
    Upon successful completion, a new `README.md` file will be generated in your current working directory. You will see a success message like:
    ```
    ✨ README.md file generated successfully!
    ```

## Folder Structure Explanation

*   `ai.go`: This file contains the logic for interacting with the Google Gemini API. It handles loading the API key, constructing the AI request, and processing the AI's response to generate the README content.
*   `go.mod`: Defines the module path for the project (`github.com/EsanSamuel/go-readme-generator`) and lists all required external Go module dependencies.
*   `main.go`: The main entry point of the CLI application. It orchestrates the entire workflow: displaying the ASCII art, getting user input for ignore lists, initiating the directory walk, calling the AI generation function, and writing the final `README.md` file.
*   `prompt.md`: (Optional) A user-defined Markdown file. If present in the execution directory, its content is appended to the AI's prompt to guide the README generation with custom instructions.
*   `read.go`: Manages the core file system operations. It contains the `ReadmeIgnore` struct and methods to configure directories and files to skip, and the `WalkDir` function which recursively traverses the project directory to collect file paths and contents.

## Technologies

*   **Go**: The primary programming language used for building the CLI tool (Go 1.25.5).
*   **Google Gemini API**: Utilized for its advanced AI capabilities to generate the README content (`google.golang.org/genai`).
*   **`github.com/joho/godotenv`**: For loading environment variables from the `.env` file, specifically the `GEMINI_API_KEY`.
*   **`github.com/mbndr/figlet4go`**: Used to render the "GO README GENERATOR" ASCII art at the start of the application.

## License

No license information provided.