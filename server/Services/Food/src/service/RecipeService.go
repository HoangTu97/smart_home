package service

import (
	// "Food/entity"
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

// FindPageByCateID return page entity.Recipe
func (s *recipeService) FindPageByCateID(cateID uint, pageable pagination.Pageable) page.Page {
	return s.recipeRepository.FindPageByCateID(cateID, pageable)
}
