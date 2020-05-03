package response

// FoodListItemResponseDTO ...
type FoodListItemResponseDTO struct {
	ID         uint     `json:"id"`
	Name       string   `json:"name"`
	Categories []string `json:"categories"`
}
