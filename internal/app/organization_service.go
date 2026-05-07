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
	FindList(uId uint64) ([]domain.Organization, error)
	Find(id uint64) (interface{}, error)
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

func (s organizationService) FindList(uId uint64) ([]domain.Organization, error) {
	orgs, err := s.orgRepo.FindList(uId)
	if err != nil {
		log.Printf("organizationService.FindList(s.orgRepo.FindList): %s", err)
		return nil, err
	}

	return orgs, nil
}

func (s organizationService) Find(id uint64) (interface{}, error) {
	org, err := s.orgRepo.Find(id)
	if err != nil {
		log.Printf("organizationService.Find(s.orgRepo.Find): %s", err)
		return nil, err
	}

	return org, nil
}