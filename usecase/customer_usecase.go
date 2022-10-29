package usecase

import (
	"mnc-bank-api/model"
	"mnc-bank-api/repository"
)

type customerUsecase struct {
	customerRepository repository.CustomerRepository
}

type CustomerUsecase interface {
	GetAll() ([]model.Customer, error)
	GetById(id string) (model.Customer, error)
	// GetByName(name string) (interface{}, error)
	GetByCredentials(username, password string) (model.Customer, error)

	Insert(customer *model.Customer) (model.Customer, error)
	UpdateBothBalance(payAmount int, senderId, receiverId string) error

	// Update(customer *model.Customer) (model.Customer, error)
	// Delete(id string) error
}

func (usecase *customerUsecase) GetAll() ([]model.Customer, error) {
	return usecase.customerRepository.GetAll()
}

func (usecase *customerUsecase) GetById(id string) (model.Customer, error) {
	return usecase.customerRepository.GetById(id)
}

// func (usecase *customerUsecase) GetByName(name string) (interface{}, error) {
// 	return usecase.customerRepository.GetByName(name)
// }

func (usecase *customerUsecase) GetByCredentials(username, password string) (model.Customer, error) {
	return usecase.customerRepository.GetByCredentials(username, password)
}

func (usecase *customerUsecase) Insert(newCustomer *model.Customer) (model.Customer, error) {
	return usecase.customerRepository.Insert(newCustomer)
}

func (usecase *customerUsecase) UpdateBothBalance(payAmount int, senderId, receiverId string) error {
	// list, _ := usecase.GetAll()

	// for index, each := range list {

	// }

	return nil

}

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
