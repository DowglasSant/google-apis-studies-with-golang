package sentiment

import (
	"context"
	"log"

	language "cloud.google.com/go/language/apiv1"
	"cloud.google.com/go/language/apiv1/languagepb"
)

func AnalyzeSentiment(text []byte) {
	ctx := context.Background()

	languageClient, err := language.NewClient(ctx)
	if err != nil {
		log.Fatalf("Erro ao criar o cliente da API de linguagem: %v", err)
	}
	defer languageClient.Close()

	sentiment, err := languageClient.AnalyzeSentiment(ctx, &languagepb.AnalyzeSentimentRequest{
		Document: &languagepb.Document{
			Source: &languagepb.Document_Content{
				Content: string(text),
			},
			Type: languagepb.Document_PLAIN_TEXT,
		},
	})
	if err != nil {
		log.Fatalf("Erro ao analisar o sentimento do texto: %v", err)
	}

	mainSentiment := sentiment.GetDocumentSentiment()
	sentimentScore := mainSentiment.GetScore()
	sentimentMagnitude := mainSentiment.GetMagnitude()

	sentimentDescription, magnitudeDescription := getScoreAndGetMagnitude(sentimentScore, sentimentMagnitude)

	log.Printf("Sentimento principal do texto: %s, Magnitude: %s\n", sentimentDescription, magnitudeDescription)
}

func getScoreAndGetMagnitude(sentimentScore, sentimentMagnitude float32) (string, string) {
	var sentimentDescription string
	var magnitudeDescription string

	switch {
	case sentimentScore > -1.0 && sentimentScore < -0.5:
		sentimentDescription = "muito negativo"
	case sentimentScore >= -0.5 && sentimentScore < 0.0:
		sentimentDescription = "negativo"
	case sentimentScore == 0.0:
		sentimentDescription = "neutro"
	case sentimentScore > 0.0 && sentimentScore <= 0.5:
		sentimentDescription = "positivo"
	case sentimentScore > 0.5 && sentimentScore <= 1.0:
		sentimentDescription = "muito positivo"
	}

	switch {
	case sentimentMagnitude <= 1.0:
		magnitudeDescription = "muito fraca"
	case sentimentMagnitude > 1.0 && sentimentMagnitude <= 2.0:
		magnitudeDescription = "fraca"
	case sentimentMagnitude > 2.0 && sentimentMagnitude <= 3.0:
		magnitudeDescription = "moderada"
	case sentimentMagnitude > 3.0 && sentimentMagnitude <= 4.0:
		magnitudeDescription = "forte"
	default:
		magnitudeDescription = "muito forte"
	}

	return sentimentDescription, magnitudeDescription
}
