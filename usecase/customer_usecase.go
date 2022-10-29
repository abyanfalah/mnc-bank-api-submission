package usecase

import (
	"mnc-bank-api/model"
	"mnc-bank-api/repository"
)

type customerUsecase struct {
	customerRepository repository.CustomerRepository
}

type CustomerUsecase interface {
	GetAll() (interface{}, error)
	GetById(id string) (interface{}, error)
	// GetById(id string) (model.Customer, error)
	// GetByName(name string) (interface{}, error)
	// GetByCredentials(customername, password string) (model.Customer, error)

	// Insert(customer *model.Customer) (model.Customer, error)
	// Update(customer *model.Customer) (model.Customer, error)
	// Delete(id string) error
}

func (usecase *customerUsecase) GetAll() (interface{}, error) {
	return usecase.customerRepository.GetAll()
}

// func (usecase *customerUsecase) GetById(id string) (interface{}, error) {
// 	return usecase.customerRepository.GetById(id)
// }

func (usecase *customerUsecase) GetById(id string) (model.Customer, error) {
	return usecase.customerRepository.GetById(id)
}

// func (usecase *customerUsecase) GetByName(name string) (interface{}, error) {
// 	return usecase.customerRepository.GetByName(name)
// }

// func (usecase *customerUsecase) GetByCredentials(customername, password string) (model.Customer, error) {
// 	return usecase.customerRepository.GetByCredentials(customername, password)
// }

// func (usecase *customerUsecase) Insert(newCustomer *model.Customer) (model.Customer, error) {
// 	return usecase.customerRepository.Insert(newCustomer)
// }

// func (usecase *customerUsecase) Update(newCustomer *model.Customer) (model.Customer, error) {
// 	return usecase.customerRepository.Update(newCustomer)
// }

// func (usecase *customerUsecase) Delete(id string) error {
// 	return usecase.customerRepository.Delete(id)
// }

func NewCustomerUsecase(customerRepository repository.CustomerRepository) CustomerUsecase {
	usecase := new(customerUsecase)
	usecase.customerRepository = customerRepository
	return usecase
}
