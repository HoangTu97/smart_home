package service

import (
	"Food/dto"
	"Food/repository"
	"Food/service/mapper"
)

// IngredientService godoc
type IngredientService interface {
	FindOne(id uint) (dto.IngredientDTO, bool)
	FindIDByName(name string) []uint
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

func (s *ingredientService) FindOne(id uint) (dto.IngredientDTO, bool) {
	ingredient, err := s.ingredientRepository.FindOne(id)
	if err != nil {
		return dto.IngredientDTO{}, false
	}
	return s.ingredientMapper.ToDTO(ingredient), true
}

func (s *ingredientService) FindIDByName(name string) []uint {
	ingredients := s.ingredientRepository.FindByName(name)

	result := make([]uint, len(ingredients))
	for i, v := range ingredients {
		result[i] = v.ID
	}

	return result
}
