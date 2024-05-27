package translation

import (
	"context"
	"log"
	"os"
	"path/filepath"
	"time"

	translate "cloud.google.com/go/translate"
	"golang.org/x/text/language"
)

func TranslateText(filePath string) ([]byte, error) {
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "C:\\Example\\golang-studies-69d707a440e3.json")

	text, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	client, err := translate.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	detection, err := client.DetectLanguage(ctx, []string{string(text)})
	if err != nil {
		return nil, err
	}

	if detection[0][0].Language.String() != "pt" {
		translations, err := client.Translate(ctx,
			[]string{string(text)}, language.BrazilianPortuguese,
			&translate.Options{
				Source: language.English,
				Format: translate.Text,
			})
		if err != nil {
			return nil, err
		}
		text = []byte(translations[0].Text)

		outputFileName := "traduzido.txt"
		outputFilePath := filepath.Join("..", "data", outputFileName)
		outputFile, err := os.Create(outputFilePath)
		if err != nil {
			return nil, err
		}
		defer outputFile.Close()

		log.Printf("Tradução do arquivo: %s", filePath)
		log.Printf("Data da tradução: %s\n", time.Now().Format("2006-01-02 15:04:05"))
		if _, err := outputFile.Write(text); err != nil {
			return nil, err
		}

		log.Printf("Tradução concluída. Arquivo criado: %s", outputFileName)
	} else {
		log.Println("Texto já se encontra em português")
	}

	return text, nil
}
