package mapper

import (
	"Food/dto"
	"Food/entity"
)

// CategoryMapper godoc
type CategoryMapper interface {
	ToDTO(entity entity.Category) dto.CategoryDTO
	ToEntity(dto dto.CategoryDTO) entity.Category
	ToDTOS(entityList []entity.Category) []dto.CategoryDTO
	ToEntities(dtoList []dto.CategoryDTO) []entity.Category
}

type categoryMapper struct{}

// NewCategory godoc
func NewCategory() CategoryMapper {
	return &categoryMapper{}
}

func (mapper *categoryMapper) ToDTO(entity entity.Category) dto.CategoryDTO {
	return dto.CategoryDTO{
		ID:        entity.ID,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
		DeletedAt: entity.DeletedAt,
		Name:      entity.Name,
		Image:     entity.Image,
	}
}

func (mapper *categoryMapper) ToEntity(dto dto.CategoryDTO) entity.Category {
	return entity.Category{
		ID:        dto.ID,
		CreatedAt: dto.CreatedAt,
		UpdatedAt: dto.UpdatedAt,
		DeletedAt: dto.DeletedAt,
		Name:      dto.Name,
		Image:     dto.Image,
	}
}

func (mapper *categoryMapper) ToDTOS(entityList []entity.Category) []dto.CategoryDTO {
	dtos := []dto.CategoryDTO{}

	for _, entity := range entityList {
		dtos = append(dtos, mapper.ToDTO(entity))
	}

	return dtos
}

func (mapper *categoryMapper) ToEntities(dtoList []dto.CategoryDTO) []entity.Category {
	entities := []entity.Category{}

	for _, dto := range dtoList {
		entities = append(entities, mapper.ToEntity(dto))
	}

	return entities
}