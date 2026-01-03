# Go README Generator CLI

This `README.md` describes the `Go README Generator CLI`, a robust command-line interface tool built with Go to automate the creation of `README.md` files.

## Description

The Go README Generator CLI is a powerful Go-based command-line interface tool meticulously designed to streamline the process of generating `README.md` files for your projects. It intelligently traverses a specified project directory, efficiently collects relevant code and content, and then harnesses the advanced capabilities of the Google Gemini API to produce a comprehensive and high-quality `README.md`. For enhanced customization and tailored output, the generation process can be optionally guided by a `config.md` file, allowing users to provide specific instructions and context.

## Features

*   **⚡ Automatic README.md Generation:** Automates the creation of `README.md` files for any given project directory.
*   **🧠 AI-Powered Content Creation:** Integrates seamlessly with the Google Gemini API to leverage advanced AI for generating insightful and well-structured README content.
*   **⚙️ Customizable Generation:** Offers flexibility through an optional `config.md` file, enabling users to provide specific instructions and context to guide the AI.
*   **🚫 Intelligent Exclusion:** Implements smart file and directory exclusion during content collection, ensuring common development artifacts (e.g., `.git`, `node_modules`, `.env`) are skipped to maintain focus on relevant project content.
*   **🚀 Built with Go:** Developed using the Go programming language, ensuring high performance, reliability, and ease of deployment.

## Installation

To get started with the Go README Generator CLI, follow these simple steps:

1.  **Prerequisites:**
    Ensure that Go (version 1.25.5 or newer) is installed and correctly configured on your system.

2.  **Clone the Repository:**
    Open your terminal and clone the project repository:
    ```bash
    git clone https://github.com/EsanSamuel/Go_cli.git
    cd Go_cli
    ```

3.  **Install Dependencies:**
    Navigate into the cloned directory and install the necessary Go modules:
    ```bash
    go mod tidy
    ```

4.  **Build the Executable:**
    Compile the Go source code to create an executable binary:
    ```bash
    go build -o go_readme_gen
    ```
    This will create an executable named `go_readme_gen` in your current directory.

5.  **Configure API Key:**
    Create a new file named `.env` in the root of the `Go_cli` project directory. Add your Google Gemini API key to this file in the following format:
    ```
    GEMINI_API_KEY=YOUR_GEMINI_API_KEY
    ```
    Replace `YOUR_GEMINI_API_KEY` with your actual Gemini API key.

## Usage

To generate a `README.md` file for your project using this tool:

1.  **Placement:**
    Place the compiled `go_readme_gen` executable (or run `go run .`) in the root directory of the project for which you intend to generate a `README.md`.

2.  **Optional `config.md`:**
    If you wish to provide specific instructions or context to guide the AI during README generation, create a `config.md` file in the same root directory. This file can contain any tailored prompts or information you deem useful.

3.  **Run the Tool:**
    Execute the tool from your project's root directory using one of the following commands:
    ```bash
    ./go_readme_gen
    ```
    or
    ```bash
    go run .
    ```

Upon successful execution, a `README.md` file will be generated and saved in the current directory, containing the AI-generated content based on your project's structure and any provided `config.md`.

## Folder Structure Explanation

The project is organized into a clear and modular structure:

*   `main.go`:
    This file serves as the application's entry point and primary orchestrator. It is responsible for initiating the reading of the current project directory, invoking the AI module for content generation, and ultimately writing the final `README.md` file.

*   `ai.go`:
    Encapsulates all the logic required for interaction with the Google Gemini API. This includes sending carefully constructed prompts and processing the generated content received from the API.

*   `read.go`:
    Manages the directory traversal functionality. It defines and applies intelligent rules to skip specific files and directories (e.g., `.git`, `node_modules`, `.env`, various lock files, build artifacts) during content collection, ensuring only relevant project data is processed.

*   `.env`:
    A configuration file used to store environment variables, most notably the `GEMINI_API_KEY` required for authenticating with the Google Gemini API.

*   `config.md` (optional):
    A user-defined Markdown file that can provide additional instructions, context, or specific guidelines to the AI, allowing for highly customized README generation.

## Technologies

The Go README Generator CLI is built upon the following key technologies:

*   **Go**: The primary programming language (version 1.25.5) used for developing the entire CLI tool, chosen for its efficiency and strong concurrency support.
*   **Google Gemini API**: Utilized for its advanced generative AI capabilities, providing the core intelligence to create rich and detailed `README.md` content.
*   **github.com/joho/godotenv**: A robust Go package specifically used for conveniently loading environment variables from `.env` files, simplifying API key management.

---

No license information was provided.