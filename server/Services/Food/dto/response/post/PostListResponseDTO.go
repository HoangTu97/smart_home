package post

import (
	"Food/helpers/page"
	"Food/models"
)

// PostListResponseDTO godoc
type PostListResponseDTO struct {
	Items      []PostListItemResponseDTO `json:"items"`
	TotalItems int64                       `json:"totalItems"`
}

// CreatePostListResponseDTOFromPage create page from page models.Post
func CreatePostListResponseDTOFromPage(page page.Page) *PostListResponseDTO {
	result := make([]PostListItemResponseDTO, page.GetTotalElements())
	for i, v := range page.GetContent() {
		entity := v.(models.Post)
		user := entity.User
		recipe := entity.Recipe

		result[i] = PostListItemResponseDTO{
			ID:         entity.ID,

			UserID: user.ID,
			UserName: user.Name,
			CreatedAt: entity.CreatedAt,

			RecipePhoto: recipe.Image,
			RecipeTitle: recipe.Name,
			RecipeDesc: recipe.Description,
			// HashTags: nil,

			CountLike: 0,
		}
	}
	return &PostListResponseDTO{
		Items:      result,
		TotalItems: page.GetTotalItems(),
	}
}
