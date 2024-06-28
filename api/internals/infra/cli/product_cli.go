package cli

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/booscaaa/go-gemini-gdg/api/internals/core/contract"
	"github.com/spf13/cobra"
)

type productCLI struct {
	productScraperUseCase contract.ProductScraperUseCase
}

// SeedProducts implements contract.ProductCLI.
func (cli *productCLI) SeedProducts(cmd *cobra.Command, args []string) {
	products, err := cli.productScraperUseCase.SeedProducts(cmd.Context())
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
}

func NewProcuctCLI(productScraperUseCase contract.ProductScraperUseCase) contract.ProductCLI {
	return &productCLI{
		productScraperUseCase: productScraperUseCase,
	}
}
