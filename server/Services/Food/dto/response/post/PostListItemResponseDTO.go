package post

import "time"

type PostListItemResponseDTO struct {
	ID uint `json:"id"`

	UserID    uint      `json:"user_id"`
	UserName  string    `json:"user_name"`
	CreatedAt time.Time `json:"created_at"`

	RecipePhoto string   `json:"recipe_photo"`
	RecipeTitle string   `json:"recipe_title"`
	RecipeDesc  string   `json:"recipe_desc"`
	HashTags    []string `json:"tags"`

	CountLike uint `json:"count_like"`
}
