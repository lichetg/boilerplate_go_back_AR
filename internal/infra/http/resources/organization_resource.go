package resources

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
