package service

import (
	"Food/dto"
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
	FindIDsByName(name string) []uint
	FindOne(id uint) (dto.RecipeDTO, bool)
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

func (s *recipeService) FindIDsByName(name string) []uint {
	recipes := s.recipeRepository.FindByName(name)
	ids := make([]uint, len(recipes))
	for i, v := range recipes {
		ids[i] = v.ID
	}
	return ids
}

func (s *recipeService) FindOne(id uint) (dto.RecipeDTO, bool) {
	recipe, err := s.recipeRepository.FindOne(id)
	if err != nil {
		return dto.RecipeDTO{}, false
	}
	return s.recipeMapper.ToDTO(recipe), true
}

func (s *recipeService) CountByCateID(cateID uint) int {
	return s.recipeRepository.CountByCateID(cateID)
}
