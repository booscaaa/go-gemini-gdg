/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/booscaaa/go-gemini-gdg/api/internals/core/contract"
	"github.com/booscaaa/go-gemini-gdg/api/pkg/adapter/postgres"
	"github.com/booscaaa/go-gemini-gdg/api/pkg/di"
	"github.com/playwright-community/playwright-go"
	"github.com/spf13/cobra"
)

var productCLI contract.ProductCLI

// updatePrice represents the serve command
var updatePrice = &cobra.Command{
	Use:   "update-price",
	Short: "",
	Long:  "",
	Run:   productCLI.SeedProducts,
}

func init() {
	pw, err := playwright.Run()
	db := postgres.Initialize()

	if err != nil {
		panic(err)
	}

	productCLI = di.InitializeProductCLIDI(pw, db)

	rootCmd.AddCommand(updatePrice)
}
