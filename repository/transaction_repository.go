package repository

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"mnc-bank-api/model"
	"mnc-bank-api/utils"
	"mnc-bank-api/utils/jsonrw"
)

type transactionRepository struct {
	tableName string
}

type TransactionRepository interface {
	// GetAllPaginated(rows int, page int) ([]model.Transaction, error)
	GetAll() ([]model.Transaction, error)
	// GetAllTest() ([]model.Transaction, error)

	// GetById(id string) (model.Transaction, error)
	// GetByIdTest(id string) (model.TransactionTest, error)

	Insert(transaction *model.Transaction) (model.Transaction, error)
	// InsertTest(transaction *model.TransactionTest) (model.TransactionTest, error)
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

// func (repo *transactionRepository) GetAllPaginated(rows int, page int) ([]model.Transaction, error) {
// 	var transactions []model.Transaction
// 	limit := rows
// 	offset := rows * (page - 1)

// 	err := p.db.Select(&transactions, utils.TRANSACTION_GET_ALL_PAGINATED+" order by created_at desc", limit, offset)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return transactions, nil
// }

// func (repo *transactionRepository) GetAllTest() ([]model.Transaction, error) {
// 	var transactions []model.Transaction

// 	err := p.db.Select(&transactions, utils.TRANSACTION_GET_ALL+" order by created_at desc")
// 	if err != nil {
// 		return nil, err
// 	}

// 	return transactions, nil
// }

// func (repo *transactionRepository) GetAllTransaction() ([]model.TransactionTest, error) {
// 	var transactions []model.TransactionTest

// 	err := p.db.Select(&transactions, utils.TRANSACTION_GET_ALL+" order by created_at desc")
// 	if err != nil {
// 		return nil, err
// 	}

// 	return transactions, nil
// }

// func (repo *transactionRepository) GetAllPaginated(page int, rows int) ([]model.Transaction, error) {
// 	var transactions []model.Transaction
// 	limit := rows
// 	offset := limit * (page - 1)

// 	err := p.db.Select(&transactions, utils.TRANSACTION_GET_ALL_PAGINATED, limit, offset)
// 	if err != nil {
// 		return nil, err
// 	}

// 	tdRepo := NewTransactionDetailRepository(p.db)

// 	for i, transaction := range transactions {
// 		items, err := tdRepo.GetByTrasactionId(transaction.Id)
// 		if err != nil {
// 			panic(err)
// 		}
// 		transactions[i].Items = items
// 	}

// 	return transactions, nil
// }

// func (repo *transactionRepository) GetById(id string) (model.Transaction, error) {
// 	var transaction model.Transaction

// 	return transaction, nil
// }

// func (repo *transactionRepository) GetByIdTest(id string) (model.TransactionTest, error) {
// 	var transaction model.TransactionTest
// 	err := p.db.Get(&transaction, utils.TRANSACTION_GET_BY_ID, id)
// 	if err != nil {
// 		return model.TransactionTest{}, err
// 	}

// 	return transaction, nil
// }

func (repo *transactionRepository) Insert(newTransaction *model.Transaction) (model.Transaction, error) {
	newTransaction.Id = utils.GenerateId()

	err := jsonrw.JsonWriteData(repo.tableName, newTransaction)
	if err != nil {
		return model.Transaction{}, errors.New("unable to write to json table " + repo.tableName + " : " + err.Error())
	}

	return *newTransaction, nil
}

// func (repo *transactionRepository) InsertTest(newTransaction *model.TransactionTest) (model.TransactionTest, error) {

// 	_, err := p.db.NamedExec(utils.TRANSACTION_INSERT, newTransaction)
// 	if err != nil {
// 		return model.TransactionTest{}, nil
// 	}

// 	return *newTransaction, nil
// }

// func (repo *transactionRepository) Update(newData *model.Transaction) (model.Transaction, error) {
// 	_, err := p.db.NamedExec(utils.TRANSACTION_UPDATE, newData)
// 	if err != nil {
// 		return model.Transaction{}, err
// 	}
// 	return *newData, nil
// }

// func (repo *transactionRepository) Delete(id string) error {
// 	_, err := p.db.Exec(utils.TRANSACTION_DELETE, id)
// 	return err
// }

func NewTransactionRepository(tableName string) TransactionRepository {
	repo := new(transactionRepository)
	repo.tableName = tableName
	return repo
}
