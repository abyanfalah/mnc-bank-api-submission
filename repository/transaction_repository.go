package repository

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"mnc-bank-api/model"
	"mnc-bank-api/utils"
	"mnc-bank-api/utils/jsonrw"
	"time"
)

type transactionRepository struct {
	tableName string
}

type TransactionRepository interface {
	GetAll() ([]model.Transaction, error)
	Insert(transaction *model.Transaction) (model.Transaction, error)
}

func (repo *transactionRepository) GetAll() ([]model.Transaction, error) {
	var list []model.Transaction

	file, err := ioutil.ReadFile("database/" + repo.tableName + ".json")
	if err != nil {
		return nil, errors.New("unable to read json file from table " + repo.tableName + " : " + err.Error())
	}

	json.Unmarshal(file, &list)
	return list, nil
}

func (repo *transactionRepository) Insert(newTransaction *model.Transaction) (model.Transaction, error) {

	if newTransaction.SenderId == "" {
		return model.Transaction{}, errors.New("no sender")
	}

	newTransaction.Id = utils.GenerateId()
	newTransaction.Created_at = time.Now()

	err := jsonrw.JsonWriteData(repo.tableName, newTransaction)
	if err != nil {
		return model.Transaction{}, errors.New("unable to write to json table " + repo.tableName + " : " + err.Error())
	}

	return *newTransaction, nil
}

func NewTransactionRepository(tableName string) TransactionRepository {
	repo := new(transactionRepository)
	repo.tableName = tableName
	return repo
}
