package repository

import (
	"warung-makan/model"
	"warung-makan/utils"

	"github.com/jmoiron/sqlx"
)

type transactionRepository struct {
	db *sqlx.DB
}

type TransactionRepository interface {
	GetAllPaginated(rows int, page int) ([]model.Transaction, error)
	GetAll() ([]model.Transaction, error)
	GetAllTest() ([]model.Transaction, error)

	GetById(id string) (model.Transaction, error)
	// GetByIdTest(id string) (model.TransactionTest, error)

	Insert(transaction *model.Transaction) (model.Transaction, error)
	// InsertTest(transaction *model.TransactionTest) (model.TransactionTest, error)
}

func (p *transactionRepository) GetAll() ([]model.Transaction, error) {
	var transactions []model.Transaction

	err := p.db.Select(&transactions, utils.TRANSACTION_GET_ALL+" order by created_at desc")
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (p *transactionRepository) GetAllPaginated(rows int, page int) ([]model.Transaction, error) {
	var transactions []model.Transaction
	limit := rows
	offset := rows * (page - 1)

	err := p.db.Select(&transactions, utils.TRANSACTION_GET_ALL_PAGINATED+" order by created_at desc", limit, offset)
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (p *transactionRepository) GetAllTest() ([]model.Transaction, error) {
	var transactions []model.Transaction

	err := p.db.Select(&transactions, utils.TRANSACTION_GET_ALL+" order by created_at desc")
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

// func (p *transactionRepository) GetAllTransaction() ([]model.TransactionTest, error) {
// 	var transactions []model.TransactionTest

// 	err := p.db.Select(&transactions, utils.TRANSACTION_GET_ALL+" order by created_at desc")
// 	if err != nil {
// 		return nil, err
// 	}

// 	return transactions, nil
// }

// func (p *transactionRepository) GetAllPaginated(page int, rows int) ([]model.Transaction, error) {
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

func (p *transactionRepository) GetById(id string) (model.Transaction, error) {
	var transaction model.Transaction
	err := p.db.Get(&transaction, utils.TRANSACTION_GET_BY_ID, id)
	if err != nil {
		return model.Transaction{}, err
	}

	return transaction, nil
}

// func (p *transactionRepository) GetByIdTest(id string) (model.TransactionTest, error) {
// 	var transaction model.TransactionTest
// 	err := p.db.Get(&transaction, utils.TRANSACTION_GET_BY_ID, id)
// 	if err != nil {
// 		return model.TransactionTest{}, err
// 	}

// 	return transaction, nil
// }

func (p *transactionRepository) Insert(newTransaction *model.Transaction) (model.Transaction, error) {
	// ===================================
	tx, err := p.db.Beginx()
	if err != nil {
		return model.Transaction{}, err
	}
	_, err = tx.NamedExec(utils.TRANSACTION_INSERT, newTransaction)
	if err != nil {
		return model.Transaction{}, err
	}

	err = tx.Commit()
	if err != nil {
		return model.Transaction{}, err
	}
	// ===================================

	if err != nil {
		return model.Transaction{}, err
	}

	return *newTransaction, nil
}

// func (p *transactionRepository) InsertTest(newTransaction *model.TransactionTest) (model.TransactionTest, error) {

// 	_, err := p.db.NamedExec(utils.TRANSACTION_INSERT, newTransaction)
// 	if err != nil {
// 		return model.TransactionTest{}, nil
// 	}

// 	return *newTransaction, nil
// }

// func (p *transactionRepository) Update(newData *model.Transaction) (model.Transaction, error) {
// 	_, err := p.db.NamedExec(utils.TRANSACTION_UPDATE, newData)
// 	if err != nil {
// 		return model.Transaction{}, err
// 	}
// 	return *newData, nil
// }

// func (p *transactionRepository) Delete(id string) error {
// 	_, err := p.db.Exec(utils.TRANSACTION_DELETE, id)
// 	return err
// }

func NewTransactionRepository(db *sqlx.DB) TransactionRepository {
	repo := new(transactionRepository)
	repo.db = db
	return repo
}
