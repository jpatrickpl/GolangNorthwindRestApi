package customer

import "github.com/GolangNorthwindRestApi/helper"

type Service interface {
	GetCustomers(param *getCustomerRequest) (*CustomerList, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s service) GetCustomers(param *getCustomerRequest) (*CustomerList, error) {
	customers, err := s.repo.GetCustomers(param)
	helper.Catch(err)
	totalCustomers, err := s.repo.GetTotalCustomers()
	helper.Catch(err)

	return &CustomerList{Data: customers, TotalRecords: totalCustomers}, nil

}
