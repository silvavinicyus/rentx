package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	usecases "rentx/src/business/useCases"
	"rentx/src/domain/dtos"
)

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	requestBody, erro := io.ReadAll(r.Body)
	if erro != nil {
		fmt.Println("Deu erro no request body", erro)
		return
	}

	var categoryDto dtos.CreateCategoryDto

	if erro = json.Unmarshal(requestBody, &categoryDto); erro != nil {
		fmt.Println("Deu erro no json unmarshall", erro)
		return
	}

	category, erro := usecases.CreateCategoryUseCase(categoryDto)
	if erro != nil {
		fmt.Println("Deu erro no create category use case", erro)
		return
	}
	fmt.Println(category)

	w.Write([]byte("Passou no create category"))
}
