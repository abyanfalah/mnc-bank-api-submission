package usecase

import (
	"warung-makan/model"
	"warung-makan/repository"
)

type customerUsecase struct {
	customerRepository repository.CustomerRepository
}

type CustomerUsecase interface {
	GetAll() ([]model.Customer, error)
	GetById(id string) (model.Customer, error)
	GetByName(name string) ([]model.Customer, error)
	GetByCredentials(customername, password string) (model.Customer, error)

	Insert(customer *model.Customer) (model.Customer, error)
	Update(customer *model.Customer) (model.Customer, error)
	Delete(id string) error
}

func (p *customerUsecase) GetAll() ([]model.Customer, error) {
	return p.customerRepository.GetAll()

}

func (p *customerUsecase) GetById(id string) (model.Customer, error) {
	return p.customerRepository.GetById(id)
}

func (p *customerUsecase) GetByName(name string) ([]model.Customer, error) {
	return p.customerRepository.GetByName(name)
}

func (p *customerUsecase) GetByCredentials(customername, password string) (model.Customer, error) {
	return p.customerRepository.GetByCredentials(customername, password)
}

func (p *customerUsecase) Insert(newCustomer *model.Customer) (model.Customer, error) {
	return p.customerRepository.Insert(newCustomer)
}

func (p *customerUsecase) Update(newCustomer *model.Customer) (model.Customer, error) {
	return p.customerRepository.Update(newCustomer)
}

func (p *customerUsecase) Delete(id string) error {
	return p.customerRepository.Delete(id)
}

func NewCustomerUsecase(customerRepository repository.CustomerRepository) CustomerUsecase {
	usecase := new(customerUsecase)
	usecase.customerRepository = customerRepository
	return usecase
}
