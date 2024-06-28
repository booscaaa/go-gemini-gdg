/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/booscaaa/go-gemini-gdg/api/pkg/adapter/postgres"
	"github.com/booscaaa/go-gemini-gdg/api/pkg/di"
	"github.com/playwright-community/playwright-go"
	"github.com/spf13/cobra"
)

// updatePrice represents the serve command
var updatePrice = &cobra.Command{
	Use:   "update-price",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		pw, err := playwright.Run()
		db := postgres.Initialize()

		if err != nil {
			panic(err)
		}

		productCLI := di.InitializeProductCLIDI(pw, db)

		productCLI.SeedProducts(cmd.Context())
	},
}

func init() {

	rootCmd.AddCommand(updatePrice)
}
