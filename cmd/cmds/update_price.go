/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/booscaaa/go-gemini-gdg/pkg/adapter/postgres"
	"github.com/booscaaa/go-gemini-gdg/pkg/di"
	"github.com/playwright-community/playwright-go"
	"github.com/spf13/cobra"
)

// updatePrice represents the serve command
var updatePrice = &cobra.Command{
	Use:   "update-price",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		pw, err := playwright.Run()
		db := postgres.Initialize()

		if err != nil {
			panic(err)
		}

		productScraperDI := di.InitializeProductScraperDI(pw, db)

		products, err := productScraperDI.SeedProducts(cmd.Context())
		if err != nil {
			panic(err)
		}

		fmt.Println("Produtos Importados com Sucesso")
		fmt.Println("----------------------------------------------------------------------")
		writer := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', tabwriter.AlignRight)
		for _, product := range products {
			fmt.Fprintf(writer, "ID: %v\tNome: %s\tPreço: %v\tEmpresa: %s\n", product.ID, product.Name, product.Price, product.Company)
		}
		writer.Flush()
	},
}

func init() {
	rootCmd.AddCommand(updatePrice)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updatePrice.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updatePrice.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
