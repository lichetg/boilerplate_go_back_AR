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
		Id:             r.Id,
		OrganizationId: r.OrganizationId,
		Name:           r.Name,
		Description:    r.Description,
	}
}

func (d RoomDto) RoomDomainToDtoCollection(rms []domain.Room) []RoomDto {
	rmsDto := make([]RoomDto, len(rms))
	for i, _ := range rms {
		rmsDto[i] = d.DomainToDto(rms[i])
	}
	return rmsDto
}
