package usecase

import (
	"errors"
	"mnc-bank-api/model"
	"mnc-bank-api/repository"
	"mnc-bank-api/utils"
)

type customerUsecase struct {
	customerRepository repository.CustomerRepository
}

type CustomerUsecase interface {
	GetAll() ([]model.Customer, error)
	GetById(id string) (model.Customer, error)
	GetByUsername(username string) model.Customer
	GetByCredentials(username, password string) (model.Customer, error)

	Insert(customer *model.Customer) (model.Customer, error)
	UpdateBothBalance(payAmount int, senderId, receiverId string) error
}

func (usecase *customerUsecase) GetAll() ([]model.Customer, error) {
	return usecase.customerRepository.GetAll()
}

func (usecase *customerUsecase) GetById(id string) (model.Customer, error) {
	tableName := "customer"
	list, err := usecase.customerRepository.GetAll()
	if err != nil {
		return model.Customer{}, errors.New("unable to read json file from table " + tableName + " : " + err.Error())
	}

	for index, each := range list {
		if each.Id == id {
			return list[index], nil
		}
	}

	return model.Customer{}, errors.New("unable to find customer " + id)
}

func (usecase *customerUsecase) GetByUsername(username string) model.Customer {
	list, _ := usecase.GetAll()
	for _, each := range list {
		if each.Username == username {
			return each
		}
	}

	return model.Customer{}
}

func (usecase *customerUsecase) GetByCredentials(username, password string) (model.Customer, error) {
	list, err := usecase.GetAll()
	if err != nil {
		return model.Customer{}, err
	}

	for index, each := range list {
		if each.Username == username && each.Password == utils.Sha1(password) {
			return list[index], nil
		}
	}

	return model.Customer{}, errors.New("invalid credential")
}

func (usecase *customerUsecase) Insert(newCustomer *model.Customer) (model.Customer, error) {
	newCustomer.Id = utils.GenerateId()
	newCustomer.Password = utils.Sha1(newCustomer.Password)

	list, _ := usecase.GetAll()
	list = append(list, *newCustomer)

	err := usecase.customerRepository.UpdateList(list)
	if err != nil {
		return model.Customer{}, nil
	}

	return *newCustomer, nil
}

func (usecase *customerUsecase) UpdateBothBalance(payAmount int, senderId, receiverId string) error {
	list, _ := usecase.GetAll()

	for index, each := range list {
		if each.Id == senderId {
			list[index].Balance -= payAmount
		}

		if each.Id == receiverId {
			list[index].Balance += payAmount
		}
	}

	return usecase.customerRepository.UpdateList(list)
}

func NewCustomerUsecase(customerRepository repository.CustomerRepository) CustomerUsecase {
	usecase := new(customerUsecase)
	usecase.customerRepository = customerRepository
	return usecase
}
