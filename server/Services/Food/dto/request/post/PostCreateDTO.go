package post

type PostCreateDTO struct {
	Photo       string   `json:"photo"`
	Description string   `json:"description"`
	HashTags    []string `json:"tags"`
	RecipeID    uint     `json:"recipeId"`
}
