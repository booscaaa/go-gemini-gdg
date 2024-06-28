package rest

import (
	"fmt"
	"net/http"

	"github.com/booscaaa/go-gemini-gdg/api/pkg/di"
	"github.com/google/generative-ai-go/genai"
	"github.com/jmoiron/sqlx"
)

func Initialize(geminiClient *genai.Client, database *sqlx.DB) {
	productController := di.InitializeProductControllerDI(geminiClient, database)
	mux := http.NewServeMux()

	mux.HandleFunc("POST /alexa", productController.SearchForTips)

	fmt.Println("SERVIDOR NA PORTA: 8000")
	http.ListenAndServe(":8000", mux)
}
