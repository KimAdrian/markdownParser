package main

import (
	"github.com/russross/blackfriday/v2"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatal("Usage: markdownParser input.md output.html")
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	err := parseMarkdown(inputFile, outputFile)
	if err != nil {
		log.Fatal("Error: ", err)
	}
	log.Printf("Successfully parsed %s to %s", inputFile, outputFile)
}

func parseMarkdown(inputFile, outputFile string) error {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		return err
	}

	htmlData := blackfriday.Run(data)

	err = os.WriteFile(outputFile, htmlData, 0644)
	if err != nil {
		return err
	}

	return nil
}
