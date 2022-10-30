package repository_test

import (
	"mnc-bank-api/model"
	"mnc-bank-api/repository"
	"mnc-bank-api/utils/migration"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TransactionRepositoryTestSuite struct {
	suite.Suite
	tableName string
	tablePath string
}

func (suite *TransactionRepositoryTestSuite) SetupTest() {
	suite.tableName = "transaction"
	suite.tablePath = "database/" + suite.tableName + ".json"
}

func (suite *TransactionRepositoryTestSuite) TestGetAll_Success() {
	repo := repository.NewTransactionRepository(suite.tableName)

	migration.AddDummyTransaction()
	actual, err := repo.GetAll()

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 2, len(actual))
}

func (suite *TransactionRepositoryTestSuite) TestInsert_Success() {
	repo := repository.NewTransactionRepository(suite.tableName)

	oldList, _ := repo.GetAll()

	newTx := model.Transaction{
		Id:         "test",
		SenderId:   "test",
		ReceiverId: "test",
		Amount:     85858,
		Created_at: time.Now(),
	}

	repo.Insert(&newTx)
	newList, _ := repo.GetAll()

	assert.True(suite.T(), len(newList) > len(oldList))

	os.Truncate(suite.tablePath, 0)
}

func (suite *TransactionRepositoryTestSuite) TestUpdateList_FailedNoNewList() {
	repo := repository.NewTransactionRepository(suite.tableName)

	actual, err := repo.Insert(&model.Transaction{})

	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), model.Transaction{}, actual)
}

func TestTransactionRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionRepositoryTestSuite))
}
