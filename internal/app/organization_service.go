package app

import (
	"log"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/database"
)

type organizationService struct {
	orgRepo database.OrganizationRepository
}

type OrganizationService interface {
	Save(o domain.Organization) (domain.Organization, error)
}

func NewOrganizationService(or database.OrganizationRepository) OrganizationService {
	return organizationService{
		orgRepo: or,
	}
}

func (s organizationService) Save(o domain.Organization) (domain.Organization, error) {
	org, err := s.orgRepo.Save(o)
	if err != nil {
		log.Printf("organizationService.Save(s.orgRepo.Save): %s", err)
		return domain.Organization{}, err
	}

	return org, nil
}
