package category

import (
	"Food/entity"
	"Food/util/page"
)

// CategoryListItemResponseDTO
type CategoryListItemResponseDTO struct {
	ID            uint   `json:"id"`
	Name          string `json:"name"`
	Image         string `json:"image"`
	NumberRecipes int    `json:"numberRecipes"`
}

// CategoryListResponseDTO godoc
type CategoryListResponseDTO struct {
	Items      []CategoryListItemResponseDTO `json:"items"`
	TotalItems int                           `json:"totalItems"`
}

// CreateCategoryListResponseDTOFromPage create page from page entity.Category
func CreateCategoryListResponseDTOFromPage(page page.Page) *CategoryListResponseDTO {
	result := make([]CategoryListItemResponseDTO, page.GetTotalElements())
	for i, v := range page.GetContent() {
		entity := v.(entity.Category)
		recipes := entity.Recipes

		result[i] = CategoryListItemResponseDTO{
			ID:            entity.ID,
			Name:          entity.Name,
			Image:         entity.Image,
			NumberRecipes: len(recipes),
		}
	}
	return &CategoryListResponseDTO{
		Items:      result,
		TotalItems: page.GetTotalItems(),
	}
}
