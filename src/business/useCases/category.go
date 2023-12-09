package usecases

import (
	"rentx/src/business/repository"
	"rentx/src/domain/dtos"
	"rentx/src/domain/entities"
	"rentx/src/utils/database"

	"github.com/google/uuid"
)

func CreateCategoryUseCase(props dtos.CreateCategoryDto) (entities.CategoryEntity, error) {
	db, erro := database.Connect()

	if erro != nil {
		return entities.CategoryEntity{}, erro
	}
	defer db.Close()

	var categoryEntity entities.CategoryEntity

	categoryEntity.Name = props.Name
	categoryEntity.Description = props.Description
	categoryEntity.UUID = uuid.NewString()

	categoryRepository := repository.CategoryRepository(db)

	categoryEntity.ID, erro = categoryRepository.Create(categoryEntity)
	if erro != nil {
		return entities.CategoryEntity{}, erro
	}

	return categoryEntity, nil
}

func FindCategoryByUseCase(props dtos.FindCategoryByDto) (entities.CategoryEntity, error) {
	db, erro := database.Connect()

	if erro != nil {
		return entities.CategoryEntity{}, erro
	}
	defer db.Close()

	categoryRepository := repository.CategoryRepository(db)

	categoryEntity, erro := categoryRepository.FindById(props)
	if erro != nil {
		return entities.CategoryEntity{}, erro
	}

	return categoryEntity, nil
}

func FindAllCategoriesUseCase(props dtos.FindAllCategoriesDto) ([]entities.CategoryEntity, error) {
	db, erro := database.Connect()

	if erro != nil {
		return nil, erro
	}
	defer db.Close()

	categoryRepository := repository.CategoryRepository(db)

	categories, erro := categoryRepository.FindAll(props)
	if erro != nil {
		return nil, erro
	}

	return categories, nil
}

func DeleteCategoryUseCase(props dtos.DeleteCategoryDto) error {
	db, erro := database.Connect()

	if erro != nil {
		return erro
	}
	defer db.Close()

	categoryRepository := repository.CategoryRepository(db)

	erro = categoryRepository.Delete(props)
	if erro != nil {
		return erro
	}

	return nil
}

func UpdateCategoryUseCase(props dtos.UpdateCategoryDto) error {
	db, erro := database.Connect()

	if erro != nil {
		return erro
	}
	defer db.Close()

	categoryRepository := repository.CategoryRepository(db)

	erro = categoryRepository.Update(props)
	if erro != nil {
		return erro
	}

	return nil
}
