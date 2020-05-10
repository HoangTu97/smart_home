package service

import (
	"Food/entity"
	"Food/util/page"
	"Food/util/pagination"
	"Food/repository"
	"Food/service/mapper"
)

// RecipeService godoc
type RecipeService interface {
	FindPageByCateID(cateID uint, pageable pagination.Pageable) page.Page
	FindPageByCates(cates []entity.Category, pageable pagination.Pageable) page.Page
	FindPageByName(name string, pageable pagination.Pageable) page.Page
	FindPageByIngredientID(ingredientID uint, pageable pagination.Pageable) page.Page
	FindPageByIngredientIDIn(ingredientIDs []uint, pageable pagination.Pageable) page.Page
	CountByCateID(cateID uint) int
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

func (s *recipeService) FindPageByCates(cates []entity.Category, pageable pagination.Pageable) page.Page {
	return s.recipeRepository.FindPageByCates(cates, pageable)
}

func (s *recipeService) FindPageByName(name string, pageable pagination.Pageable) page.Page {
	return s.recipeRepository.FindPageByName(name, pageable)
}

func (s *recipeService) FindPageByIngredientID(ingredientID uint, pageable pagination.Pageable) page.Page {
	return s.recipeRepository.FindPageByIngredientID(ingredientID, pageable)
}

func (s *recipeService) FindPageByIngredientIDIn(ingredientIDs []uint, pageable pagination.Pageable) page.Page {
	return s.recipeRepository.FindPageByIngredientIDIn(ingredientIDs, pageable)
}

func (s *recipeService) CountByCateID(cateID uint) int {
	return s.recipeRepository.CountByCateID(cateID)
}
