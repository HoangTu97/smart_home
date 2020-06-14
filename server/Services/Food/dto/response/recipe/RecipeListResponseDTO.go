package recipe

import (
	"Food/entity"
	"Food/pkg/page"
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
		categories := entity.Categories

		cates := make([]string, len(categories))
		for k, cate := range categories {
			cates[k] = cate.Name
		}

		result[i] = RecipeListItemResponseDTO{
			ID:         entity.ID,
			Title:      entity.Name,
			PhotoURL:   entity.Image,
			Categories: cates,
		}
	}
	return &RecipeListResponseDTO{
		Items:      result,
		TotalItems: page.GetTotalItems(),
	}
}
