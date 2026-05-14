package requests

type RoomRequest struct {
	Name        string  `json:"name" validate:"required,gte=1,max=40"`
	Description *string `json:"description"`
}
