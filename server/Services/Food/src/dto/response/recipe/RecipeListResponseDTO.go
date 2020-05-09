package recipe

import (
	"Food/entity"
	"Food/util/page"
)

// RecipeListResponseDTO godoc
type RecipeListResponseDTO struct {
	Items      []RecipeListItemResponseDTO `json:"items"`
	TotalItems int                         `json:"totalItems"`
}

// CreateRecipeListResponseDTOFromPage create page from page entity.Recipe
func CreateRecipeListResponseDTOFromPage(page page.Page) *RecipeListResponseDTO {
	result := make([]RecipeListItemResponseDTO, page.GetTotalElements())
	for i, v := range page.GetContent() {
		entity := v.(entity.Recipe)
		result[i] = RecipeListItemResponseDTO{
			ID:       entity.ID,
			Title:    entity.Name,
			PhotoURL: entity.Image,
		}
	}
	return &RecipeListResponseDTO{
		Items:      result,
		TotalItems: page.GetTotalItems(),
	}
}
