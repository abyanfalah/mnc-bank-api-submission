package repository_test

import (
	"mnc-bank-api/model"
	"mnc-bank-api/repository"
	"mnc-bank-api/utils/migration"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CustomerRepositoryTestSuite struct {
	suite.Suite
	tableName string
	tablePath string
}

func (suite *CustomerRepositoryTestSuite) SetupTest() {
	suite.tableName = "customer"
	suite.tablePath = "database/" + suite.tableName + ".json"
}

func (suite *CustomerRepositoryTestSuite) TestGetAll_Success() {
	repo := repository.NewCustomerRepository(suite.tableName)

	migration.AddDummyCustomer()
	actual, err := repo.GetAll()

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 2, len(actual))
}

func (suite *CustomerRepositoryTestSuite) TestUpdateList_Success() {
	repo := repository.NewCustomerRepository(suite.tableName)

	oldList, err := repo.GetAll()

	newCust := model.Customer{
		Id:       "test",
		Name:     "test",
		Username: "test",
		Password: "test",
		Balance:  123,
	}

	newList := append(oldList, newCust)
	updateErr := repo.UpdateList(newList)

	assert.Nil(suite.T(), err)
	assert.Nil(suite.T(), updateErr)
	assert.True(suite.T(), len(newList) > len(oldList))

	os.Truncate(suite.tablePath, 0)

}

func (suite *CustomerRepositoryTestSuite) TestUpdateList_FailedNoNewList() {
	repo := repository.NewCustomerRepository(suite.tableName)

	err := repo.UpdateList(nil)

	assert.NotNil(suite.T(), err)
}

func TestCustomerRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(CustomerRepositoryTestSuite))
}
