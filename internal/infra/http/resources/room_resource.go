package resources

import "github.com/BohdanBoriak/boilerplate-go-back/internal/domain"

type RoomDto struct {
	Id             uint64  `json:"id"`
	OrganizationId uint64  `json:"organization_id"`
	Name           string  `json:"name"`
	Description    *string `json:"description,omitempty"`
}

func (d RoomDto) DomainToDto(r domain.Room) RoomDto {
	return RoomDto{
		Id:             d.Id,
		OrganizationId: d.OrganizationId,
		Name:           d.Name,
		Description:    d.Description,
	}
}
