package repository

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"text/tabwriter"

	"github.com/booscaaa/go-gemini-gdg/api/internals/core/contract"
	"github.com/booscaaa/go-gemini-gdg/api/internals/core/domain"
	"github.com/google/generative-ai-go/genai"
)

type geminiRepository struct {
	client *genai.Client
}

// GetMenu implements contract.ProductLLMRepository.
func (repository *geminiRepository) GetMenu(ctx context.Context, products []domain.Product) (*string, error) {
	model := repository.client.GenerativeModel("gemini-1.5-pro")

	chatSession := model.StartChat()

	var b bytes.Buffer

	writer := tabwriter.NewWriter(&b, 0, 8, 1, '\t', tabwriter.AlignRight)
	for _, product := range products {
		fmt.Fprintf(writer, "ID: %v\tNome: %s\tPreço: %v\tEmpresa: %s\n", product.ID, product.Name, product.Price, product.Company)
	}
	writer.Flush()

	prompt := fmt.Sprintf(` 
	Você é um atendente de restaurante.
	Sempre monte um cardápio para pedir em qualquer hora segundo essa lista de produtos: %s.
	Informe o que seria ideal pedir para comer e o preço total da compra.
	Detalhe de qual empresa é cada produto, bem como seu nome e preço.
	Mostre três opções de pedidos.
	Não dê mais informações do que o necessário.
	Sempre seja gentil.
	Seja sucinto!
	Você pode misturar os produtos de todas as empresas, se quiser.
	Diversifique sempre os pedidos para não ser toda vez a mesma coisa.
	Coloque uma pequena frase legal no final.
	Escreva os números por extenso sempre.
	Escreva tudo por extenso para leitura da Alexa.
	Escreva tudo em uma frase, respeitando o portugues.
	Não coloque caracteres especiais.
	Fale sempre em primeira pessoa.
												 `, b.String())

	res, err := chatSession.SendMessage(ctx, genai.Text(prompt))
	if err != nil {
		log.Fatal(err)
	}

	output := ""

	for _, cand := range res.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				output = fmt.Sprintf("%s %s", output, part)
			}
		}
	}

	fmt.Println(output)

	return &output, nil
}

func NewGeminiRepository(client *genai.Client) contract.ProductLLMRepository {
	return &geminiRepository{
		client: client,
	}
}
