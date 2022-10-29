package repository

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"mnc-bank-api/model"
	"mnc-bank-api/utils/jsonrw"
)

type customerRepository struct {
	tableName string
}

type CustomerRepository interface {
	GetAll() ([]model.Customer, error)
	UpdateList(newList []model.Customer) error
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

func (repo *customerRepository) UpdateList(newList []model.Customer) error {
	err := jsonrw.JsonUpdateList(repo.tableName, newList)
	if err != nil {
		return err
	}

	return nil
}

func NewCustomerRepository(tableName string) CustomerRepository {
	repo := new(customerRepository)
	repo.tableName = tableName
	return repo
}
