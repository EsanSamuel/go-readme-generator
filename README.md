# Go README Generator CLI

## Description
This is a Go command-line interface (CLI) tool designed to automate the creation of `README.md` files. It traverses a project directory, collects code and content, and leverages the Google Gemini API to generate a `README.md`. It can be optionally guided by a `config.md` file for tailored instructions.

## Features
*   Automatic `README.md` generation for projects.
*   Integration with the Google Gemini API for AI-powered content creation.
*   Customizable generation process via an optional `config.md` file.
*   Intelligent file and directory exclusion during content collection, skipping common development artifacts.
*   Built using Go.

## Installation
1.  **Prerequisites:** Ensure Go (version 1.25.5 or newer) is installed on your system.
2.  **Clone the repository:**
    ```bash
    git clone https://github.com/EsanSamuel/Go_cli.git
    cd Go_cli
    ```
3.  **Install dependencies:**
    ```bash
    go mod tidy
    ```
4.  **Build the executable:**
    ```bash
    go build -o go_readme_gen
    ```
5.  **Configure API Key:** Create a `.env` file in the root of the `Go_cli` project and add your Gemini API key:
    ```
    GEMINI_API_KEY=YOUR_GEMINI_API_KEY
    ```

## Usage
To generate a README for your project:
1.  Place the compiled `go_readme_gen` executable (or run `go run .`) in the root directory of the project for which you want to generate a README.
2.  (Optional) Create a `config.md` file in the same directory. This file can contain specific instructions or context to guide the AI in generating your README.
3.  Run the tool from your project's root directory:
    ```bash
    ./go_readme_gen
    ```
    or
    ```bash
    go run .
    ```
    A `README.md` file will be generated and saved in the current directory.

## Folder Structure Explanation
*   `main.go`: The entry point and orchestrator of the application. It handles reading the current directory, invoking the AI, and writing the final `README.md`.
*   `ai.go`: Contains the logic for interacting with the Google Gemini API, sending prompts, and receiving generated content.
*   `read.go`: Implements the directory traversal functionality, including defining and applying rules to skip specific files and directories (e.g., `.git`, `node_modules`, `.env`).
*   `.env`: Stores environment variables, specifically the `GEMINI_API_KEY`.
*   `config.md` (optional): A user-defined file that can provide additional instructions or context to the AI for README generation.

## Technologies
*   **Go**: The primary programming language used for development (version 1.25.5).
*   **Google Gemini API**: Utilized for its generative AI capabilities to create README content.
*   **github.com/joho/godotenv**: A Go package used for loading environment variables from `.env` files.