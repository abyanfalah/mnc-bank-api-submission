package repository

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"mnc-bank-api/model"
	"mnc-bank-api/utils"
	"mnc-bank-api/utils/jsonrw"
)

type customerRepository struct {
	tableName string
}

type CustomerRepository interface {
	GetAll() ([]model.Customer, error)
	GetById(id string) (model.Customer, error)
	GetByCredentials(customername, password string) (model.Customer, error)

	Insert(customer *model.Customer) (model.Customer, error)
	// Update(customer *model.Customer) (model.Customer, error)
	UpdateList(newList []model.Customer) error
	// Delete(id string) error
}

func (repo *customerRepository) GetAll() ([]model.Customer, error) {
	var list []model.Customer

	file, err := ioutil.ReadFile("database/" + repo.tableName + ".json")
	if err != nil {
		return nil, errors.New("unable to read json file from table " + repo.tableName + " : " + err.Error())
	}

	json.Unmarshal(file, &list)
	return list, nil
}

func (repo *customerRepository) GetById(id string) (model.Customer, error) {
	list, err := repo.GetAll()
	if err != nil {
		return model.Customer{}, errors.New("unable to read json file from table " + repo.tableName + " : " + err.Error())
	}

	for index, each := range list {
		if each.Id == id {
			return list[index], nil
		}
	}

	return model.Customer{}, errors.New("unable to find customer " + id + " from table " + repo.tableName)

}

func (repo *customerRepository) GetByCredentials(username, password string) (model.Customer, error) {
	list, err := repo.GetAll()
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

func (repo *customerRepository) Insert(newCustomer *model.Customer) (model.Customer, error) {
	newCustomer.Id = utils.GenerateId()
	newCustomer.Password = utils.Sha1(newCustomer.Password)
	err := jsonrw.JsonWriteData(repo.tableName, newCustomer)
	if err != nil {
		return model.Customer{}, errors.New("unable to write to json table " + repo.tableName + " : " + err.Error())
	}

	return *newCustomer, nil
}

func (repo *customerRepository) UpdateList(newList []model.Customer) error {
	err := jsonrw.JsonUpdateList(repo.tableName, newList)
	if err != nil {
		return err
	}

	return nil
}

// func (repo *customerRepository) Update(newData *model.Customer) (model.Customer, error) {

// }

// func (repo *customerRepository) Delete(id string) error {
// 	_, err := p.db.Exec(utils.USER_DELETE, id)
// 	return err
// }

func NewCustomerRepository(tableName string) CustomerRepository {
	repo := new(customerRepository)
	repo.tableName = tableName
	return repo
}
