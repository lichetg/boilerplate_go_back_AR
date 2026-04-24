package resources

import (
	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
)

type OrganizationDto struct {
	Id          uint64  `json:"id"`
	UserId      uint64  `json:"userid"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	City        string  `json:"city"`
	Address     string  `json:"address"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
}

func (d OrganizationDto) DomainToDto(o domain.Organization) OrganizationDto {
	return OrganizationDto{
		Id:          o.Id,
		UserId:      o.UserId,
		Name:        o.Name,
		Description: o.Description,
		City:        o.City,
		Address:     o.Address,
		Lat:         o.Lat,
		Lon:         o.Lon,
	}
}
