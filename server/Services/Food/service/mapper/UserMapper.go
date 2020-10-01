package mapper

import (
	"Food/domain"
	"Food/dto"
	"Food/helpers/converter"
	"Food/models"

	uuid "github.com/satori/go.uuid"
)

func ToUserDTO(entity models.User) dto.UserDTO {
	roleStrings, _ := converter.ArrStr(entity.Roles)
	roles := make([]domain.UserRole, len(roleStrings))
	for i, r := range roleStrings {
		ro, _ := domain.ParseUserRole(r)
		roles[i] = ro
	}

	return dto.UserDTO{
		ID:         entity.ID,
		UserID:     entity.UserID.String(),
		CreatedAt:  entity.CreatedAt,
		UpdatedAt:  entity.UpdatedAt,
		DeletedAt:  entity.DeletedAt,
		Name:       entity.Name,
		Password:   entity.Password,
		Roles:      roles,
		Address:    entity.Address,
		Age:        entity.Age,
		Gender:     entity.Gender,
		Occupation: entity.Occupation,
		Long:       entity.Long,
		Lat:        entity.Lat,
		ZipCode:    entity.ZipCode,
	}
}

func ToUser(dto dto.UserDTO) models.User {
	var id uuid.UUID
	if len(dto.UserID) > 0 {
		id = uuid.Must(uuid.FromString(dto.UserID))
	}

	roles := make([]string, len(dto.Roles))
	for i, r := range dto.Roles {
		roles[i] = r.String()
	}

	return models.User{
		ID:         dto.ID,
		UserID:     id,
		CreatedAt:  dto.CreatedAt,
		UpdatedAt:  dto.UpdatedAt,
		DeletedAt:  dto.DeletedAt,
		Name:       dto.Name,
		Password:   dto.Password,
		Roles:      converter.ToStr(roles),
		Address:    dto.Address,
		Age:        dto.Age,
		Gender:     dto.Gender,
		Occupation: dto.Occupation,
		Long:       dto.Long,
		Lat:        dto.Lat,
		ZipCode:    dto.ZipCode,
	}
}

func ToUserDTOS(entityList []models.User) []dto.UserDTO {
	dtos := make([]dto.UserDTO, len(entityList))

	for i, entity := range entityList {
		dtos[i] = ToUserDTO(entity)
	}

	return dtos
}

func ToUsers(dtoList []dto.UserDTO) []models.User {
	entities := make([]models.User, len(dtoList))

	for i, dto := range dtoList {
		entities[i] = ToUser(dto)
	}

	return entities
}
