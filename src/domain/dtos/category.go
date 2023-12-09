package dtos

type CreateCategoryDto struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type DeleteCategoryDto struct {
	UUID string
}

type FindCategoryByDto struct {
	ID uint64
}

type UpdateCategoryDto struct {
	ID          uint64 `json:"id;omitempty"`
	Name        string `json:"name;omitempty"`
	Description string `json:"description;omitempty"`
}

type FindAllCategoriesDto struct {
	Name string
}
