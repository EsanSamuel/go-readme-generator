# Go README Generator

## Description

The Go README Generator is a command-line interface (CLI) tool written in Go that automatically generates a `README.md` file for a given project. It works by traversing the project's directory, collecting file paths and their contents, and then feeding this information, along with an optional custom prompt, to a Generative AI model (Gemini). The tool allows users to specify folders and files to skip during the directory traversal, ensuring only relevant project data is sent for README generation.

## Features

*   **Automated README Generation**: Generates a `README.md` file by analyzing project content.
*   **Generative AI Integration**: Leverages the Gemini AI model to intelligently create descriptive READMEs.
*   **Configurable Directory Traversal**: Allows users to specify additional folders and files to ignore during the content collection process.
*   **Default Ignore List**: Comes with a predefined list of common directories (e.g., `.git`, `node_modules`) and files (e.g., `.env`, `README.md`) to skip.
*   **Custom Prompt Support**: Supports an optional `prompt.md` file in the project's root directory to provide specific instructions or context to the AI.
*   **CLI Interface**: Provides an interactive command-line experience for inputting skip patterns.

## Installation

To install and run the Go README Generator, follow these steps:

1.  **Clone the repository (if applicable)**:
    Since the project data is provided directly, we'll assume the source code is available.

2.  **Ensure Go is installed**:
    This project requires Go version `1.25.5` or higher.
    You can download Go from [golang.org](https://golang.org/dl/).

3.  **Set up your environment variable**:
    The tool requires a `GEMINI_API_KEY` to interact with the Generative AI.
    Create a `.env` file in the root directory of the project and add your API key:

    ```ini
    GEMINI_API_KEY="YOUR_GEMINI_API_KEY"
    ```

    You can obtain a Gemini API key from the Google AI Studio or Google Cloud Console.

4.  **Install dependencies**:
    Navigate to the project's root directory in your terminal and install the Go modules:

    ```bash
    go mod tidy
    ```

## Usage

1.  **Navigate to your project directory**:
    Change your current working directory to the root of the project for which you want to generate a `README.md`.

2.  **Run the application**:
    Execute the main Go program:

    ```bash
    go run .
    ```

3.  **Provide input when prompted**:
    The CLI will guide you through the process:
    *   It will display a welcome banner.
    *   You will be asked to enter folders to skip (e.g., `test data docs`). Separate multiple folders with spaces.
    *   You will be asked to enter files to skip (e.g., `config.json .DS_Store`). Separate multiple files with spaces.

    Example interaction:

    ```
        _.-._
    _.-'-._'-._
    '-._'-._'-._
    _.-'_.-'_.-'
    '-._'-._'-._
    _.-'_.-'_.-'
    '-._'-._'-._
        '-._'-._'-._
    _.-'_.-'_.-'
    '-._'-._'-._
    _.-'_.-'_.-'
    '-._'-._'-._
        _.-'_.-'_.-'
        '-._'-._'-._
            _.-'_.-'_.-'
            '-._'-._'-._
                _.-'_.-'_.-'
                '-._'-._'-._
                    _.-'_.-'_.-'
                    '-._'-._'-._
                        _.-'_.-'_.-'
                        '-._'-._'-._
                            _.-'_.-'_.-'
                            '-._'-._'-._
    _.-._           _.-._        _.-._
    '-._'-._'-._'-._'-._'-._'-._'-._'-._'-._'-._'-._'-._'-._'-._'-._'-._
        _.-'_.-'_.-'_.-'_.-'_.-'_.-'_.-'_.-'_.-'_.-'_.-'_.-'_.-'_.-'_.-'_.-'
        '-._'-._'-._'-._'-._'-._'-._'-._'-._'-._'-._'-._'-._'-._'-._'-._'-._
            _.-'_.-'_.-'_.-'_.-'_.-'_.-'_.-'_.-'_.-'_.-'_.-'_.-'_.-'_.-'_.-'_.-'
            '-._'-._'-._'-._'-._'-._'-._'-._'-._'-._'-._'-._'-._'-._'-._'-._'-._
    _.-._           _.-._        _.-._            _.-._           _.-._
    '-._'-._'-._'-._'-._'-._'-._'-._'-._'-._'-._'-._'-._'-._'-._'-._'-._'-._'-._'-._
        _.-'_.-'_.-'_.-'_.-'_.-'_.-'_.-'_.-'_.-'_.-'_.-'_.-'_.-'_.-'_.-'_.-'_.-'_.-'_.-'
        '-._'-._'-._'-._'-._'-._'-._'-._'-._'-._'-._'-._'-._'-._'-._'-._'-._'-._'-._'-._
            _.-'_.-'_.-'_.-'_.-'_.-'_.-'_.-'_.-'_.-'_.-'_.-'_.-'_.-'_.-'_.-'_.-'_.-'_.-'_.-'
            '-._'-._'-._'-._'-._'-._'-._'-._'-._'-._'-._'-._'-._'-._'-._'-._'-._'-._'-._'-._
    _.-._           _.-._        _.-._            _.-._           _.-._
    '-._'-._'-._'-._'-._'-._'-._'-._'-._'-._'-._'-._'-._'-._'-._'-._'-._'-._'-._'-._
        _.-'_.-'_.-'_.-'_.-'_.-'_.-'_.-'_.-'_.-'_.-'_.-'_.-'_.-'_.-'_.-'_.-'_.-'_.-'_.-'
        '-._'-._'-._'-._'-._'-._'-._'-._'-._'-._'-._'-._'-._'-._'-._'-._'-._'-._'-._'-._
            _.-'_.-'_.-'_.-'_.-'_.-'_.-'_.-'_.-'_.-'_.-'_.-'_.-'_.-'_.-'_.-'_.-'_.-'_.-'_.-'
            '-._'-._'-._'-._'-._'-._'-._'-._'-._'-._'-._'-._'-._'-._'-._'-._'-._'-._'-._'-._
    GO README GENERATOR
    Enter folders you want to skip: .vscode
    Enter files you want to skip: .gitkeep
    Generating README file...
    Writing README file...
    ✨ README.md file generated successfully!
    ```

4.  **Customizing the AI prompt**:
    If you want to provide specific instructions to the AI beyond the project content, create a file named `prompt.md` in the root of your project directory. The content of this file will be appended to the project data before being sent to the AI.

    For example, `prompt.md` could contain:
    ```markdown
    DESIGN ALL FILE STRUCTURE AND TEACH USERS HOW TO USE THE CLI TOOL
    ```

## Folder Structure Explanation

The project has the following structure:

```
Cli/
├── ai.go                 # Contains the logic for interacting with the Gemini AI model.
├── go.mod                # Defines the module path and lists project dependencies.
├── main.go               # The main entry point of the CLI tool, handles user input, directory traversal, and README generation.
├── prompt.md             # An optional file to provide custom instructions to the AI for README generation.
└── read.go               # Contains logic for directory traversal and managing files/folders to ignore.
```

## Technologies

*   **Go (Golang)**: The primary programming language used for the CLI tool.
*   **Google Gemini API**: Used for generating high-quality README content.
    *   `google.golang.org/genai`: Go client library for the Gemini API.
*   **`github.com/joho/godotenv`**: For loading environment variables from a `.env` file.
*   **`github.com/mbndr/figlet4go`**: For rendering ASCII art in the CLI.

## License

*(No license information was provided in the project data.)*