package controller

import (
	"encoding/json"
	"net/http"

	"github.com/booscaaa/go-gemini-gdg/api/internals/core/contract"
)

type productController struct {
	productUseCase contract.ProductUseCase
}

// SearchForTips implements contract.ProductController.
func (controller *productController) SearchForTips(response http.ResponseWriter, request *http.Request) {
	output, err := controller.productUseCase.SearchForTips(request.Context())
	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(response).Encode(output)
}

func NewProductController(productUseCase contract.ProductUseCase) contract.ProductController {
	return &productController{
		productUseCase: productUseCase,
	}
}
