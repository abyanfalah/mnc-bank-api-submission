package manager

import "warung-makan/usecase"

type usecaseManager struct {
	repo RepoManager
}

type UsecaseManager interface {
	CustomerUsecase() usecase.CustomerUsecase
	TransactionUsecase() usecase.TransactionUsecase
}

func (um *usecaseManager) CustomerUsecase() usecase.CustomerUsecase {
	return usecase.NewCustomerUsecase(um.repo.CustomerRepo())
}

func (um *usecaseManager) TransactionUsecase() usecase.TransactionUsecase {
	return usecase.NewTransactionUsecase(um.repo.TransactionRepo())
}

func NewUsecaseManager(repo RepoManager) UsecaseManager {
	return &usecaseManager{
		repo: repo,
	}
}
