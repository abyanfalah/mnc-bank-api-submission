package usecase

import (
	"mnc-bank-api/model"
	"mnc-bank-api/repository"
)

type transactionUsecase struct {
	transactionRepository repository.TransactionRepository
}

type TransactionUsecase interface {
	GetAll() ([]model.Transaction, error)
	// GetAllPaginated(page int, rows int) ([]model.Transaction, error)
	// GetById(id string) (model.Transaction, error)

	Insert(transaction *model.Transaction) (model.Transaction, error)
	// Update(transaction *model.Transaction) (model.Transaction, error)
	// Delete(id string) error
}

func (tu *transactionUsecase) GetAll() ([]model.Transaction, error) {
	return tu.transactionRepository.GetAll()
}

// func (tu *transactionUsecase) GetAllPaginated(page int, rows int) ([]model.Transaction, error) {
// 	return tu.transactionRepository.GetAllPaginated(page, rows)
// }

// func (tu *transactionUsecase) GetById(id string) (model.Transaction, error) {
// 	return tu.transactionRepository.GetById(id)
// }

func (tu *transactionUsecase) Insert(newTransaction *model.Transaction) (model.Transaction, error) {
	return tu.transactionRepository.Insert(newTransaction)
}

// func (tu *transactionUsecase) Update(newTransaction *model.Transaction) (model.Transaction, error) {
// 	return tu.transactionRepository.Update(newTransaction)
// }

// func (tu *transactionUsecase) Delete(id string) error {
// 	return tu.transactionRepository.Delete(id)
// }

func NewTransactionUsecase(transactionRepository repository.TransactionRepository) TransactionUsecase {
	usecase := new(transactionUsecase)
	usecase.transactionRepository = transactionRepository
	return usecase
}
