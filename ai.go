package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/genai"
)

func Ai(prompt string) (string, error) {
	ctx := context.Background()
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Cannot find .env file")
	}
	api_key := os.Getenv("GEMINI_API_KEY")
	config := &genai.ClientConfig{
		APIKey: api_key,
	}
	client, err := genai.NewClient(ctx, config)
	if err != nil {
		log.Fatal(err)
	}

	promptDetail := `
Act as a senior software engineer and technical writer.

Generate a high-quality README.md using ONLY the information provided.
Do not assume or hallucinate missing features.

README structure:
1. Project Title
2. Description
3. Features
4. Installation
5. Usage
6. Folder Structure Explanation
7. Technologies
8. License (if mentioned)

Formatting rules:
- Use proper Markdown headings
- Use code blocks where appropriate
- Keep explanations simple and clear

Project data:
` + prompt
	fmt.Println(promptDetail)

	result, err := client.Models.GenerateContent(
		ctx,
		"gemini-2.5-flash",
		genai.Text(promptDetail),
		nil,
	)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	//fmt.Println(result.Text())
	return result.Text(), nil
}
