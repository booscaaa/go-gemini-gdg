/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/booscaaa/go-gemini-gdg/api/pkg/adapter/gemini"
	"github.com/booscaaa/go-gemini-gdg/api/pkg/adapter/postgres"
	"github.com/booscaaa/go-gemini-gdg/api/pkg/adapter/rest"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		geminiClient := gemini.Initialize()
		database := postgres.Initialize()

		rest.Initialize(geminiClient, database)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
