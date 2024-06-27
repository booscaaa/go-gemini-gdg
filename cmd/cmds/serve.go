/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/booscaaa/go-gemini-gdg/pkg/adapter/postgres"
	"github.com/booscaaa/go-gemini-gdg/pkg/di"
	"github.com/google/generative-ai-go/genai"
	"github.com/spf13/cobra"
	"google.golang.org/api/option"
)

// Define structs to represent Alexa responses
type AlexaResponse struct {
	Version  string               `json:"version"`
	Response AlexaResponseContent `json:"response"`
}

type AlexaResponseContent struct {
	OutputSpeech     AlexaOutputSpeech `json:"outputSpeech"`
	ShouldEndSession bool              `json:"shouldEndSession"`
}

type AlexaOutputSpeech struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		client, err := genai.NewClient(
			context.Background(),
			option.WithAPIKey(""),
		)
		if err != nil {
			log.Fatal(err)
		}

		database := postgres.Initialize()

		productUseCase := di.InitializeProductDI(client, database)

		mux := http.NewServeMux()

		mux.HandleFunc("POST /alexa", func(w http.ResponseWriter, r *http.Request) {
			output, _ := productUseCase.GetMenu(cmd.Context())

			response := createAlexaResponse(*output)
			sendAlexaResponse(w, response)
		})

		http.ListenAndServe(":8000", mux)

	},
}

func sendAlexaResponse(w http.ResponseWriter, response AlexaResponse) {
	// Convert the response to JSON
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to create JSON response", http.StatusInternalServerError)
		return
	}

	// Set the content type and send the response
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func createAlexaResponse(message string) AlexaResponse {
	// Create an Alexa response object
	return AlexaResponse{
		Version: "1.0",
		Response: AlexaResponseContent{
			OutputSpeech: AlexaOutputSpeech{
				Type: "PlainText",
				Text: message,
			},
			ShouldEndSession: true, // Set to true if this response ends the session
		},
		// You can add other optional elements like card, directives, etc. here
	}
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
