package repository

import (
	"fmt"
	"mnc-bank-api/model"
	"mnc-bank-api/utils/jsonio"

	"github.com/jmoiron/sqlx"
)

type customerRepository struct {
	tableName string
}

type CustomerRepository interface {
	GetAll() (interface{}, error)
	GetById(id string) (interface{}, error)
	// GetById(id string) (model.Customer, error)
	// GetByCredentials(customername, password string) (interface{}, error)

	// Insert(customer *interface{}) (interface{}, error)
	// Update(customer *interface{}) (interface{}, error)
	// Delete(id string) error
}

func (repo *customerRepository) GetAll() (interface{}, error) {
	customers, err := jsonio.JsonReadData(repo.tableName)
	if err != nil {
		return nil, err
	}
	return customers, nil
}

// func (repo *customerRepository) GetById(id string) (model.Customer, error) {
func (repo *customerRepository) GetById(id string) (interface{}, error) {
	// var datas []model.Customer
	list, _ := jsonio.JsonReadData(repo.tableName)
	// for index, each := range list {

	// }

	fmt.Println(list)

	return model.Customer{}, nil

}

// func (repo *customerRepository) GetByName(name string) ([]interface{}, error) {
// 	// var customer []interface{}
// 	// err := p.db.Select(&customer, utils.USER_GET_BY_NAME, "%"+name+"%")
// 	// if err != nil {
// 	// 	return nil, err
// 	// }
// 	// return customer, nil
// }

// func (repo *customerRepository) GetByCredentials(customername, password string) (interface{}, error) {
// 	// var customer interface{}
// 	// err := p.db.Get(&customer, utils.USER_GET_BY_CREDENTIALS, customername, password)
// 	// if err != nil {
// 	// 	return nil, err
// 	// }
// 	// return customer, nil
// }

// func (repo *customerRepository) Insert(newCustomer *interface{}) (interface{}, error) {
// 	_, err := p.db.NamedExec(utils.USER_INSERT, newCustomer)
// 	// if err != nil {
// 	// 	return nil, err
// 	// }
// 	// customer := newCustomer
// 	// return *customer, nil
// }

// func (repo *customerRepository) Update(newData *interface{}) (interface{}, error) {
// 	_, err := p.db.NamedExec(utils.USER_UPDATE, newData)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return *newData, nil
// }

// func (repo *customerRepository) Delete(id string) error {
// 	_, err := p.db.Exec(utils.USER_DELETE, id)
// 	return err
// }

func NewCustomerRepository(db *sqlx.DB) CustomerRepository {
	repo := new(customerRepository)
	repo.tableName = "customer"
	return repo
}
