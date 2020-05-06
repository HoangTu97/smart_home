package service

import (
	"Food/dto"
	"Food/repository"
	"Food/service/mapper"
)

// CategoryService ...
type CategoryService interface {
	Save(categoryDTO dto.CategoryDTO) dto.CategoryDTO
	FindOne(id uint) dto.CategoryDTO
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

func (s *categoryService) Save(categoryDTO dto.CategoryDTO) dto.CategoryDTO {
	category := s.categoryMapper.ToEntity(categoryDTO)
	category = s.categoryRepository.Save(category)
	return s.categoryMapper.ToDTO(category)
}

func (s *categoryService) FindOne(id uint) dto.CategoryDTO {
	category := s.categoryRepository.FindOne(id)
	return s.categoryMapper.ToDTO(category)
}

func (s *categoryService) FindAll() []dto.CategoryDTO {
	categories := s.categoryRepository.FindAll()
	return s.categoryMapper.ToDTOS(categories)
}

func (s *categoryService) Delete(id uint) bool {
	category := s.categoryRepository.FindOne(id)
	if category.ID == 0 {
		return false
	}
	s.categoryRepository.Delete(category)
	return true
}
