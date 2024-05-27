package main

import (
	"google-apis-studies/internal/sentiment"
	"google-apis-studies/internal/translation"
	"log"
	"path/filepath"
)

func main() {
	inputFilePath := filepath.Join("data", "input.txt")

	translatedText, err := translation.TranslateText(inputFilePath)
	if err != nil {
		log.Fatalf("Erro ao traduzir o texto: %v", err)
	}

	sentiment.AnalyzeSentiment(translatedText)
}
