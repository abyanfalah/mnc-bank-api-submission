package repository_test

import (
	"mnc-bank-api/model"
	"mnc-bank-api/repository"
	"mnc-bank-api/utils"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

const tableName = "customer"

type CustomerRepositoryTestSuite struct {
	suite.Suite
}

func (suite *CustomerRepositoryTestSuite) TestGetAll_Success() {
	repo := repository.NewCustomerRepository(tableName)

	actual, err := repo.GetAll()

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 4, len(actual))
}

func (suite *CustomerRepositoryTestSuite) TestUpdateList_Success() {
	repo := repository.NewCustomerRepository(tableName)

	oldList, err := repo.GetAll()

	newCust := model.Customer{
		Id:       utils.GenerateId(),
		Name:     utils.GenerateId(),
		Username: utils.GenerateId(),
		Password: "5baa61e4c9b93f3f0682250b6cf8331b7ee68fd8",
		Balance:  utils.GenerateClockSequence(),
	}

	newList := append(oldList, newCust)
	updateErr := repo.UpdateList(newList)

	assert.Nil(suite.T(), err)
	assert.Nil(suite.T(), updateErr)
	assert.True(suite.T(), len(newList) > len(oldList))
}

func (suite *CustomerRepositoryTestSuite) TestUpdateList_FailedNoNewList() {
	repo := repository.NewCustomerRepository(tableName)

	err := repo.UpdateList(nil)

	assert.NotNil(suite.T(), err)
}

func TestCustomerRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(CustomerRepositoryTestSuite))
}
