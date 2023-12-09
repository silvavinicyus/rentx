package repository

import (
	"database/sql"
	"fmt"
	"rentx/src/domain/dtos"
	"rentx/src/domain/entities"
	"strings"
)

type categories struct {
	db *sql.DB
}

func CategoryRepository(db *sql.DB) *categories {
	return &categories{db}
}

func (repository *categories) Create(category entities.CategoryEntity) (uint64, error) {
	statement, statementError := repository.db.Prepare("INSERT INTO categories (uuid, name, description) VALUES (?, ?, ?)")

	if statementError != nil {
		return 0, statementError
	}

	defer statement.Close()

	result, execError := statement.Exec(category.UUID, category.Name, category.Description)

	if execError != nil {
		return 0, execError
	}

	categoryId, lastIdError := result.LastInsertId()
	if lastIdError != nil {
		return 0, lastIdError
	}

	return uint64(categoryId), nil
}

func (repository *categories) FindById(findByDto dtos.FindCategoryByDto) (entities.CategoryEntity, error) {
	lines, repositoryError := repository.db.Query("SELECT * FROM categories WHERE uuid = ?", findByDto.UUID)
	if repositoryError != nil {
		return entities.CategoryEntity{}, repositoryError
	}
	defer lines.Close()

	var category entities.CategoryEntity

	if lines.Next() {
		if repositoryError := lines.Scan(&category.ID, &category.UUID, &category.Name, &category.Description, &category.CreatedAt); repositoryError != nil {
			return entities.CategoryEntity{}, repositoryError
		}
	}

	return category, nil
}

func (repository *categories) FindAll(findAllDto dtos.FindAllCategoriesDto) ([]entities.CategoryEntity, error) {
	nameLikeClause := fmt.Sprintf("%%%s%%", findAllDto.Name)

	lines, repositoryError := repository.db.Query(`SELECT * FROM categories WHERE name LIKE ?`, nameLikeClause)
	if repositoryError != nil {
		return nil, repositoryError
	}
	defer lines.Close()

	var categories []entities.CategoryEntity

	for lines.Next() {
		var category entities.CategoryEntity

		if repositoryError := lines.Scan(&category.ID, &category.UUID, &category.Name, &category.Description, &category.CreatedAt); repositoryError != nil {
			return nil, repositoryError
		}

		categories = append(categories, category)
	}

	return categories, nil
}

func (repository *categories) Delete(deleteDto dtos.DeleteCategoryDto) error {
	statement, repositoryError := repository.db.Prepare("DELETE FROM categories WHERE id = ?")
	if repositoryError != nil {
		return repositoryError
	}
	defer statement.Close()

	_, repositoryError = statement.Exec(deleteDto.ID)
	if repositoryError != nil {
		return repositoryError
	}

	return nil
}

func (repository *categories) Update(updateDto dtos.UpdateCategoryDto) error {
	var updateFields []string

	if updateDto.Name != "" {
		updateFields = append(updateFields, "name = ?")
	}
	if updateDto.Description != "" {
		updateFields = append(updateFields, "description = ?")
	}

	updateFieldsJoined := strings.Join(updateFields, ",")

	fmt.Println(updateFieldsJoined)

	sqlClause := fmt.Sprintf("UPDATE categories SET %s WHERE id = ?", updateFieldsJoined)

	statement, repositoryError := repository.db.Prepare(sqlClause)
	if repositoryError != nil {
		return repositoryError
	}
	defer statement.Close()

	_, repositoryError = statement.Exec(updateDto.Name, updateDto.Description, updateDto.ID)
	if repositoryError != nil {
		return repositoryError
	}

	return nil
}
