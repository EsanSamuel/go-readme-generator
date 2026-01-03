# Go_cli: Automated README Generation 📝✨

## Description

`Go_cli` serves as an intelligent assistant for developers, simplifying the often time-consuming task of writing and maintaining project documentation. It addresses the common challenge of keeping READMEs up-to-date by dynamically analyzing your project's structure and content.

The primary goal of `Go_cli` is to provide a robust, automated solution for generating `README.md` files that are:
*   **Accurate:** Directly derived from the project's source code and structure.
*   **Comprehensive:** Covering essential sections like description, features, installation, usage, folder structure, and technologies.
*   **Customizable:** Allowing users to define specific directories and files to include or exclude from the scan.
*   **AI-Powered:** Utilizing the Gemini API to craft human-like, well-structured explanations.

This tool is ideal for any software project, regardless of scale, aiming to streamline documentation workflows and ensure that project insights are readily available and consistent. 🚀

## Features

`Go_cli` offers a suite of features designed to make README generation seamless and effective:

*   **Intelligent Directory Traversal** 🌲: Automatically scans the current working directory, identifying relevant files and folders to include in the README generation process.
*   **Configurable Skip Lists** 🚫:
    *   **Default Exclusions:** Comes pre-configured with a list of common directories (`.git`, `node_modules`, `vendor`, `dist`, `build`, `coverage`, `.vscode`, etc.) and files (`.env`, `.gitignore`, `*.lock` files, `*.log`, etc.) to automatically ignore during scanning.
    *   **User-Defined Exclusions:** Allows users to specify additional directories and files to skip, providing fine-grained control over the content included in the prompt.
*   **AI-Powered Content Generation** 🤖: Integrates with the Google Gemini API to process collected project data and generate detailed, well-structured, and clear `README.md` content.
*   **Optional `config.md` Integration** ⚙️: Supports an optional `config.md` file in the project root, enabling users to provide additional context, specific instructions, or custom sections to the AI for a more tailored README output.
*   **Structured README Output** 📄: Generates `README.md` files adhering to a predefined, high-quality structure including: Project Title, Description, Features, Installation, Usage, Folder Structure Explanation, Technologies, and License.
*   **Dynamic Prompt Construction** 📝: Builds a comprehensive AI prompt by concatenating scanned project content with the optional `config.md` and explicit instructions for the AI on README structure and formatting.
*   **Automated File Output** 💾: Writes the AI-generated content directly to a `README.md` file in the project's root directory.

## Installation

To get `Go_cli` up and running on your system, follow these steps:

### 1. Prerequisites 📦

Ensure you have the following installed on your machine:
*   **Go:** Version 1.25.5 or higher. You can download it from [golang.org/dl](https://golang.org/dl/).

### 2. Environment Configuration 🔑

This tool relies on the Google Gemini API. You need to provide your API key:
*   Create a `.env` file in the root directory where you will run `Go_cli`.
*   Add your Gemini API key to this file:
    ```dotenv
    GEMINI_API_KEY="YOUR_GEMINI_API_KEY_HERE"
    ```
    Replace `"YOUR_GEMINI_API_KEY_HERE"` with your actual API key obtained from the Google AI Studio.

### 3. Clone the Repository 📥

First, clone the `Go_cli` repository to your local machine:

```bash
git clone https://github.com/EsanSamuel/Go_cli.git
cd Go_cli
```

### 4. Build the Executable 🛠️

Once inside the project directory, build the Go executable:

```bash
go build -o Go_cli .
```
This command will create an executable file named `Go_cli` (or `Go_cli.exe` on Windows) in your current directory.

### 5. Install Dependencies (Optional, handled by `go build`) 🔗

Go automatically handles dependency resolution during the `go build` step using the `go.mod` file. You generally don't need a separate `go mod tidy` or `go mod download` step unless you are developing the project.

## Usage

Using `Go_cli` is straightforward. Navigate to your project's root directory (the one for which you want to generate a README) and run the executable.

### 1. Run the CLI Tool ▶️

Execute the built CLI tool from your project's root directory:

```bash
./Go_cli
```

### 2. Configure Exclusions (Interactive) 🚶‍♂️

The tool will prompt you interactively to specify any folders or files you wish to explicitly skip during the content scanning phase.

*   **Enter folders you want to skip:**
    ```
    Enter folders you want to skip: .vscode temp
    ```
    (Enter space-separated folder names. You can press Enter directly to accept default skips only.)

*   **Enter files you want to skip:**
    ```
    Enter files you want to skip: my_secret.txt old_log.txt
    ```
    (Enter space-separated file names, including patterns like `*.tmp`. Press Enter for default skips.)

### 3. (Optional) Provide `config.md` for Enhanced Context 📝

You can place an optional `config.md` file in the root of your project. This file's content will be appended to the project's source code content and sent to the AI, allowing you to give specific instructions, highlight key areas, or add custom sections that the AI should include in the generated README.

**Example `config.md` content:**
```markdown
# Additional Context for README Generation

**Project Goal:** This project aims to provide a robust solution for real-time chat functionality with end-to-end encryption. Focus on the security aspects in the description.

**Key Features to Highlight:**
- User registration and login
- Encrypted messaging (mention AES-256)
- Group chat support
- File sharing

**Audience:** Primarily developers and security researchers.
```

### 4. README Generation & Output 📄

After you provide the skip configurations, the tool will:
1.  Scan your project's files and folders (excluding those specified).
2.  Combine the scanned content with any provided `config.md`.
3.  Send this aggregated data to the Gemini AI with instructions to generate a `README.md`.
4.  Once the AI response is received, it will write the content to a new `README.md` file in your project's root directory.

You will see status messages in the console:
```
Generating README file...
Writing README file...
README.md file generated successfully!
```

## Folder Structure Explanation

`Go_cli` itself is structured as a typical Go project. Understanding its internal layout helps in comprehension and contribution:

```
Go_cli/
├── .env                  # Environment file for storing the GEMINI_API_KEY 🔑
├── ai.go                 # Handles communication with the Google Gemini API, constructing and sending prompts 🤖
├── go.mod                # Go module definition file, listing project dependencies and Go version 📦
├── go.sum                # Go checksums for modules to ensure integrity 🔐
├── ignore.go             # Defines the ReadmeIgnore struct and methods for configuring file/folder skips 🚫
├── main.go               # The main entry point for the CLI application, orchestrating user input, scanning, and AI calls 🖥️
└── README.md             # This file! (And the output of the CLI tool when run in another project) ✨
```

When `Go_cli` is run *within another project*, the `config.md` (optional) and the generated `README.md` will reside in that project's root directory.

## Technologies

`Go_cli` is built using a modern Go stack with powerful third-party integrations:

*   **Core Language:** Go (version 1.25.5) 🐿️
*   **AI Integration:**
    *   `google.golang.org/genai`: Google Gemini API client for advanced content generation. 🤖
*   **Environment Variables:**
    *   `github.com/joho/godotenv`: For loading environment variables from `.env` files. 🔑
*   **File System Operations:**
    *   `os`: Standard library for interacting with the operating system, including file and directory operations. 📂
    *   `path/filepath`: Standard library for manipulating file paths. 🌳
*   **Input/Output Handling:**
    *   `bufio`: Standard library for buffered I/O, used for reading user input. 📝
*   **String Manipulation:**
    *   `strings`: Standard library for string operations. 🧵
*   **Dependency Management:**
    *   Go Modules (`go.mod`, `go.sum`): Standard Go dependency management system. 📦

## License

No specific license information was mentioned in the provided project data. Therefore, the project's default copyright rules apply, meaning all rights are reserved unless a license is explicitly stated.

If you are the project owner, please consider adding a license (e.g., MIT, Apache 2.0, GPLv3) to this `README.md` to define how others can use, distribute, and contribute to your work. This is crucial for both open-source collaboration and proprietary clarity. ⚖️