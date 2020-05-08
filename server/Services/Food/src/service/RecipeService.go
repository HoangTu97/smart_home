package service

import (
	"Food/util/page"
	"Food/util/pagination"
	"Food/repository"
	"Food/service/mapper"
)

// RecipeService godoc
type RecipeService interface {
	FindPageByCateID(cateID uint, pageable pagination.Pageable) page.Page
}

type recipeService struct {
	recipeRepository repository.RecipeRepository
	recipeMapper     mapper.RecipeMapper
}

// NewRecipe godoc
func NewRecipe(recipeRepository repository.RecipeRepository, recipeMapper mapper.RecipeMapper) RecipeService {
	return &recipeService{
		recipeRepository: recipeRepository,
		recipeMapper:     recipeMapper,
	}
}

func (s *recipeService) FindPageByCateID(cateID uint, pageable pagination.Pageable) page.Page {
	return nil
}
