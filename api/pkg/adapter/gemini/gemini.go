package gemini

import (
	"context"
	"log"

	"github.com/google/generative-ai-go/genai"
	"github.com/spf13/viper"
	"google.golang.org/api/option"
)

func Initialize() *genai.Client {
	client, err := genai.NewClient(
		context.Background(),
		option.WithAPIKey(viper.GetString("GEMINI_API_KEY")),
	)
	if err != nil {
		log.Fatal(err)
	}

	return client
}
