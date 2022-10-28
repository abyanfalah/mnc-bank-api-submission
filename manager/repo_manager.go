package manager

import "warung-makan/repository"

type repoManager struct {
	infra InfraManager
}

type RepoManager interface {
	CustomerRepo() repository.CustomerRepository
	TransactionRepo() repository.TransactionRepository
}

func (rm *repoManager) CustomerRepo() repository.CustomerRepository {
	return repository.NewCustomerRepository(rm.infra.GetSqlDb())
}

func (rm *repoManager) TransactionRepo() repository.TransactionRepository {
	return repository.NewTransactionRepository(rm.infra.GetSqlDb())
}

func NewRepoManager(infra InfraManager) RepoManager {
	return &repoManager{
		infra: infra,
	}
}
