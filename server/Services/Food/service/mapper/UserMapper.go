package mapper

import (
	"Food/dto"
	"Food/entity"

	uuid "github.com/satori/go.uuid"
)

func ToUserDTO(entity entity.User) dto.UserDTO {
	return dto.UserDTO{
		ID:         entity.ID.String(),
		CreatedAt:  entity.CreatedAt,
		UpdatedAt:  entity.UpdatedAt,
		DeletedAt:  entity.DeletedAt,
		Name:       entity.Name,
		Password:   entity.Password,
		Address:    entity.Address,
		Age:        entity.Age,
		Gender:     entity.Gender,
		Occupation: entity.Occupation,
		Long:       entity.Long,
		Lat:        entity.Lat,
		ZipCode:    entity.ZipCode,
	}
}

func ToUser(dto dto.UserDTO) entity.User {
	var id uuid.UUID
	if len(dto.ID) > 0 {
		id = uuid.Must(uuid.FromString(dto.ID))
	}

	return entity.User{
		ID:         id,
		CreatedAt:  dto.CreatedAt,
		UpdatedAt:  dto.UpdatedAt,
		DeletedAt:  dto.DeletedAt,
		Name:       dto.Name,
		Password:   dto.Password,
		Address:    dto.Address,
		Age:        dto.Age,
		Gender:     dto.Gender,
		Occupation: dto.Occupation,
		Long:       dto.Long,
		Lat:        dto.Lat,
		ZipCode:    dto.ZipCode,
	}
}

func ToUserDTOS(entityList []entity.User) []dto.UserDTO {
	dtos := make([]dto.UserDTO, len(entityList))

	for i, entity := range entityList {
		dtos[i] = ToUserDTO(entity)
	}

	return dtos
}

func ToUsers(dtoList []dto.UserDTO) []entity.User {
	entities := make([]entity.User, len(dtoList))

	for i, dto := range dtoList {
		entities[i] = ToUser(dto)
	}

	return entities
}
