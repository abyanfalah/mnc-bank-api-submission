package manager

import "mnc-bank-api/repository"

type repoManager struct {
	customerTableName    string
	transactionTableName string
}

type RepoManager interface {
	CustomerRepo() repository.CustomerRepository
	TransactionRepo() repository.TransactionRepository
}

func (rm *repoManager) CustomerRepo() repository.CustomerRepository {
	return repository.NewCustomerRepository(rm.customerTableName)
}

func (rm *repoManager) TransactionRepo() repository.TransactionRepository {
	return repository.NewTransactionRepository(rm.transactionTableName)
}

func NewRepoManager() RepoManager {
	return &repoManager{
		customerTableName:    "customer",
		transactionTableName: "transaction",
	}
}
