package controller

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	usecases "rentx/src/business/useCases"
	"rentx/src/domain/dtos"
	"rentx/src/domain/entities"
	"rentx/src/utils/response"
	"strings"

	"github.com/gorilla/mux"
)

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	requestBody, erro := io.ReadAll(r.Body)
	if erro != nil {
		response.Error(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var categoryDto dtos.CreateCategoryDto

	if erro = json.Unmarshal(requestBody, &categoryDto); erro != nil {
		response.Error(w, http.StatusBadRequest, erro)
		return
	}

	category, erro := usecases.CreateCategoryUseCase(categoryDto)
	if erro != nil {
		response.Error(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusCreated, category)
}

func FindCategoryBy(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	categoryUuid := parameters["uuid"]

	findByDto := dtos.FindCategoryByDto{
		UUID: categoryUuid,
	}

	category, useCaseError := usecases.FindCategoryByUseCase(findByDto)
	if useCaseError != nil {
		response.Error(w, http.StatusInternalServerError, useCaseError)
		return
	}

	if (category == entities.CategoryEntity{}) {
		response.Error(w, http.StatusNotFound, nil)
		return
	}

	response.JSON(w, http.StatusOK, category)
}

func FindAllCategories(w http.ResponseWriter, r *http.Request) {
	name := strings.ToLower(r.URL.Query().Get("name"))

	categories, useCaseError := usecases.FindAllCategoriesUseCase(dtos.FindAllCategoriesDto{Name: name})
	if useCaseError != nil {
		response.Error(w, http.StatusInternalServerError, useCaseError)
		return
	}

	response.JSON(w, http.StatusOK, categories)
}

func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	categoryUuid := parameters["uuid"]

	requestBody, bodyError := io.ReadAll(r.Body)
	if bodyError != nil {
		response.Error(w, http.StatusUnprocessableEntity, bodyError)
		return
	}

	var updateDto dtos.UpdateCategoryDto

	if jsonError := json.Unmarshal(requestBody, &updateDto); jsonError != nil {
		response.Error(w, http.StatusBadRequest, jsonError)
		return
	}

	categoryExists, useCaseError := usecases.FindCategoryByUseCase(dtos.FindCategoryByDto{UUID: categoryUuid})
	if useCaseError != nil {
		response.Error(w, http.StatusInternalServerError, useCaseError)
		return
	}

	if (categoryExists == entities.CategoryEntity{}) {
		response.Error(w, http.StatusNotFound, errors.New("Category not found!"))
		return
	}

	updateDto.ID = categoryExists.ID

	updateError := usecases.UpdateCategoryUseCase(updateDto)
	if updateError != nil {
		response.Error(w, http.StatusInternalServerError, updateError)
	}

	categoryExists.Name = updateDto.Name
	categoryExists.Description = updateDto.Description

	response.JSON(w, http.StatusOK, categoryExists)
}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	categoryUuid := parameters["uuid"]

	categoryExists, useCaseError := usecases.FindCategoryByUseCase(dtos.FindCategoryByDto{UUID: categoryUuid})
	if useCaseError != nil {
		response.Error(w, http.StatusInternalServerError, useCaseError)
		return
	}

	if (categoryExists == entities.CategoryEntity{}) {
		response.Error(w, http.StatusNotFound, errors.New("Category not found!"))
		return
	}

	deleteError := usecases.DeleteCategoryUseCase(dtos.DeleteCategoryDto{ID: categoryExists.ID})

	if deleteError != nil {
		response.Error(w, http.StatusInternalServerError, deleteError)
	}

	response.JSON(w, http.StatusNoContent, nil)
}
