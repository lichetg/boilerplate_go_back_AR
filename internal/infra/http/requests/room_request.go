package requests

import "github.com/BohdanBoriak/boilerplate-go-back/internal/domain"

type RoomRequest struct {
	Name        string  `json:"name" validate:"required,gte=1,max=40"`
	Description *string `json:"description"`
}

func (r RoomRequest) ToDomainModel() (interface{}, error) {
	return domain.Room{
		Name:        r.Name,
		Description: r.Description,
	}, nil
}
