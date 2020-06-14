package recipe

// RecipeListItemResponseDTO godoc
type RecipeListItemResponseDTO struct {
	ID         uint     `json:"id"`
	Title      string   `json:"title"`
	PhotoURL   string   `json:"photo_url"`
	Categories []string `json:"categories"`
}
