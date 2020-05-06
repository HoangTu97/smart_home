package service

import (
	"Food/repository"
	"Food/service/mapper"
)

// IngredientService godoc
type IngredientService interface {
}

type ingredientService struct {
	ingredientRepository repository.IngredientRepository
	ingredientMapper     mapper.IngredientMapper
}

// NewIngredient godoc
func NewIngredient(ingredientRepository repository.IngredientRepository, ingredientMapper mapper.IngredientMapper) IngredientService {
	return &ingredientService{
		ingredientRepository: ingredientRepository,
		ingredientMapper:     ingredientMapper,
	}
}
