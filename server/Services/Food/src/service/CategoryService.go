package service

import (
	"Food/dto"
	"Food/repository"
	"Food/service/mapper"
)

// CategoryService ...
type CategoryService interface {
	Save(categoryDTO dto.CategoryDTO) (dto.CategoryDTO, bool)
	FindOne(id uint) (dto.CategoryDTO, bool)
	FindAll() []dto.CategoryDTO
	Delete(id uint) bool
}

type categoryService struct {
	categoryRepository repository.CategoryRepository
	categoryMapper     mapper.CategoryMapper
}

// NewCategory ...
func NewCategory(categoryRepository repository.CategoryRepository, categoryMapper mapper.CategoryMapper) CategoryService {
	return &categoryService{
		categoryRepository: categoryRepository,
		categoryMapper:     categoryMapper,
	}
}

func (s *categoryService) Save(categoryDTO dto.CategoryDTO) (dto.CategoryDTO, bool) {
	category := s.categoryMapper.ToEntity(categoryDTO)
	var err error
	category, err = s.categoryRepository.Save(category)
	if err != nil {
		return categoryDTO, false
	}
	return s.categoryMapper.ToDTO(category), true
}

func (s *categoryService) FindOne(id uint) (dto.CategoryDTO, bool) {
	category, err := s.categoryRepository.FindOne(id)
	if err != nil {
		return dto.CategoryDTO{}, false
	}
	return s.categoryMapper.ToDTO(category), true
}

func (s *categoryService) FindAll() []dto.CategoryDTO {
	categories := s.categoryRepository.FindAll()
	return s.categoryMapper.ToDTOS(categories)
}

func (s *categoryService) Delete(id uint) bool {
	category, err := s.categoryRepository.FindOne(id)
	if err != nil {
		return false
	}
	s.categoryRepository.Delete(category)
	return true
}
