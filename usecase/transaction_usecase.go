package usecase

import (
	"errors"
	"mnc-bank-api/model"
	"mnc-bank-api/repository"
)

type transactionUsecase struct {
	transactionRepository repository.TransactionRepository
}

type TransactionUsecase interface {
	GetAll() ([]model.Transaction, error)
	GetById(id string) (model.Transaction, error)
	Insert(transaction *model.Transaction) (model.Transaction, error)
}

func (usecase *transactionUsecase) GetAll() ([]model.Transaction, error) {
	return usecase.transactionRepository.GetAll()
}

func (usecase *transactionUsecase) GetById(id string) (model.Transaction, error) {
	tableName := "transaction"
	list, err := usecase.transactionRepository.GetAll()
	if err != nil {
		return model.Transaction{}, errors.New("unable to read json file from table " + tableName + " : " + err.Error())
	}

	for _, each := range list {
		if each.Id == id {
			return each, nil
		}
	}

	return model.Transaction{}, errors.New("unable to find transaction " + id)
}

func (usecase *transactionUsecase) Insert(newTransaction *model.Transaction) (model.Transaction, error) {
	return usecase.transactionRepository.Insert(newTransaction)
}

func NewTransactionUsecase(transactionRepository repository.TransactionRepository) TransactionUsecase {
	usecase := new(transactionUsecase)
	usecase.transactionRepository = transactionRepository
	return usecase
}
