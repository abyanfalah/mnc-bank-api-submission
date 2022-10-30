package usecase_test

import (
	"errors"
	"mnc-bank-api/model"
	"mnc-bank-api/usecase"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

var dummyTransactionList = []model.Transaction{
	{
		Id:         "test",
		SenderId:   "sender1",
		ReceiverId: "receiver1",
		Amount:     12345,
		Created_at: time.Now(),
	},

	{
		Id:         "test2",
		SenderId:   "sender2",
		ReceiverId: "receiver2",
		Amount:     85858,
		Created_at: time.Now(),
	},
}

type transactionRepoMock struct {
	mock.Mock
}

type TransactionUsecaseTestSuite struct {
	suite.Suite
	transactionRepoMock *transactionRepoMock
}

func (r *transactionRepoMock) GetAll() ([]model.Transaction, error) {
	args := r.Called()
	if args.Get(1) != nil {
		return nil, args.Error(1)
	}

	return args[0].([]model.Transaction), nil
}

func (r *transactionRepoMock) Insert(newTransaction *model.Transaction) (model.Transaction, error) {
	args := r.Called(newTransaction)
	if args.Get(1) != nil {
		return model.Transaction{}, args.Error(1)
	}

	return *newTransaction, nil
}

func (suite *TransactionUsecaseTestSuite) TestGetAll_Success() {
	suite.transactionRepoMock.On("GetAll").Return(dummyTransactionList, nil)

	TransactionUsecaseTest := usecase.NewTransactionUsecase(suite.transactionRepoMock)
	actual, err := TransactionUsecaseTest.GetAll()

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), dummyTransactionList, actual)
}

func (suite *TransactionUsecaseTestSuite) TestGetAll_Failed() {
	suite.transactionRepoMock.On("GetAll").Return(nil, errors.New("failed to get transaction list"))

	TransactionUsecaseTest := usecase.NewTransactionUsecase(suite.transactionRepoMock)
	actual, err := TransactionUsecaseTest.GetAll()

	assert.NotNil(suite.T(), err)
	assert.Nil(suite.T(), actual)
}

func (suite *TransactionUsecaseTestSuite) TestGetById_Success() {
	suite.transactionRepoMock.On("GetAll").Return(dummyTransactionList, nil)

	TransactionUsecaseTest := usecase.NewTransactionUsecase(suite.transactionRepoMock)
	actual, err := TransactionUsecaseTest.GetById(dummyTransactionList[0].Id)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), dummyTransactionList[0], actual)
}

func (suite *TransactionUsecaseTestSuite) TestGetById_Failed() {
	suite.transactionRepoMock.On("GetAll").Return(nil, errors.New("failed to get customer "+dummyTransactionList[0].Id))

	TransactionUsecaseTest := usecase.NewTransactionUsecase(suite.transactionRepoMock)
	actual, err := TransactionUsecaseTest.GetById(dummyTransactionList[0].Id)

	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), model.Transaction{}, actual)
}

func (suite *TransactionUsecaseTestSuite) TestInsert_Success() {
	newTransaction := model.Transaction{
		SenderId:   "sender2",
		ReceiverId: "receiver2",
		Amount:     85858,
	}

	suite.transactionRepoMock.On("Insert", &newTransaction).Return(newTransaction, nil)

	TransactionUsecaseTest := usecase.NewTransactionUsecase(suite.transactionRepoMock)
	actual, err := TransactionUsecaseTest.Insert(&newTransaction)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), newTransaction.SenderId, actual.SenderId)
	assert.Equal(suite.T(), newTransaction.ReceiverId, actual.ReceiverId)
	assert.Equal(suite.T(), newTransaction.Amount, actual.Amount)
}

func (suite *TransactionUsecaseTestSuite) TestInsert_Failed() {
	newTransaction := model.Transaction{
		SenderId:   "sender2",
		ReceiverId: "receiver2",
		Amount:     85858,
	}

	suite.transactionRepoMock.On("Insert", &newTransaction).Return(model.Transaction{}, errors.New("failed"))

	TransactionUsecaseTest := usecase.NewTransactionUsecase(suite.transactionRepoMock)
	actual, err := TransactionUsecaseTest.Insert(&newTransaction)

	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), model.Transaction{}, actual)
}

func (suite *TransactionUsecaseTestSuite) SetupTest() {
	suite.transactionRepoMock = new(transactionRepoMock)
}

func TestTransactionomerUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionUsecaseTestSuite))
}
