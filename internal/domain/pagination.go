package domain

import "time"

type Pagination struct {
	Page         uint64
	CountPerPage uint64
}

type MeasurementFilters struct {
	DeviceId uint64

	CreatedDateFrom *time.Time
	CreatedDateTo   *time.Time

	Sort string
}

type EventFilters struct {
	DeviceId uint64

	CreatedDateFrom *time.Time
	CreatedDateTo   *time.Time

	Sort string
}
