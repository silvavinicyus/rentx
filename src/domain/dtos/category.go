package dtos

type CreateCategoryDto struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type DeleteCategoryDto struct {
	ID uint64
}

type FindCategoryByDto struct {
	UUID string
}

type UpdateCategoryDto struct {
	ID          uint64 `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type FindAllCategoriesDto struct {
	Name string
}
