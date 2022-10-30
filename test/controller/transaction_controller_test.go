package controller_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"mnc-bank-api/controller"
	"mnc-bank-api/model"
	"mnc-bank-api/utils"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
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
		SenderId:   "sender1",
		ReceiverId: "receiver1",
		Amount:     85858,
		Created_at: time.Now(),
	},
}

var dummyTxCustomerList = []model.Customer{
	{
		Id:       "sender1",
		Name:     "test",
		Username: "test",
		Password: utils.Sha1("test"),
		Balance:  5000,
	},

	{
		Id:       "receiver1",
		Name:     "test2",
		Username: "test2",
		Password: utils.Sha1("test2"),
		Balance:  5000,
	},
}

var Cookie = &http.Cookie{
	Name:     "session",
	Value:    dummyTxCustomerList[0].Id,
	Path:     "/",
	Domain:   "localhost",
	MaxAge:   1,
	Secure:   true,
	HttpOnly: true,
}

type TransactionUsecaseMock struct {
	mock.Mock
}

func (cu *TransactionUsecaseMock) GetAll() ([]model.Transaction, error) {
	args := cu.Called()
	if args.Get(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.Transaction), nil
}

func (cu *TransactionUsecaseMock) GetById(id string) (model.Transaction, error) {
	args := cu.Called(id)
	if args.Get(1) != nil {
		return model.Transaction{}, args.Error(1)
	}
	return args.Get(0).(model.Transaction), nil
}

func (cu *TransactionUsecaseMock) Insert(transaction *model.Transaction) (model.Transaction, error) {
	args := cu.Called(transaction)
	if args.Get(1) != nil {
		return model.Transaction{}, args.Error(1)
	}
	return args.Get(0).(model.Transaction), nil
}

type TransactionControllerTestSuite struct {
	suite.Suite
	transactionUsecaseMock *TransactionUsecaseMock
	CustomerUsecaseMock    *CustomerUsecaseMock
	routerMock             *gin.Engine
}

func (suite *TransactionControllerTestSuite) SetupTest() {
	suite.routerMock = gin.Default()
	suite.CustomerUsecaseMock = new(CustomerUsecaseMock)
	suite.transactionUsecaseMock = new(TransactionUsecaseMock)
}

func (suite *TransactionControllerTestSuite) TestListTransactionApi_Success() {
	suite.transactionUsecaseMock.On("GetAll").Return(dummyTransactionList, nil)

	controller.NewTransactionController(suite.transactionUsecaseMock, suite.CustomerUsecaseMock, suite.routerMock)

	r := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodGet, "/transaction", nil)
	request.AddCookie(Cookie)
	suite.routerMock.ServeHTTP(r, request)

	var actual []model.Transaction
	response := r.Body.String()
	jsonerr := json.Unmarshal([]byte(response), &actual)

	assert.Equal(suite.T(), http.StatusOK, r.Code)
	assert.Nil(suite.T(), err)
	assert.Nil(suite.T(), jsonerr)
	assert.Equal(suite.T(), len(dummyTransactionList), len(actual))
}

func (suite *TransactionControllerTestSuite) TestListTransactionApi_Failed() {
	suite.transactionUsecaseMock.On("GetAll").Return(nil, errors.New("failed"))

	controller.NewTransactionController(suite.transactionUsecaseMock, suite.CustomerUsecaseMock, suite.routerMock)

	r := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodGet, "/transaction", nil)
	request.AddCookie(Cookie)
	suite.routerMock.ServeHTTP(r, request)

	var actual []model.Transaction
	response := r.Body.String()
	jsonerr := json.Unmarshal([]byte(response), &actual)

	assert.Equal(suite.T(), http.StatusInternalServerError, r.Code)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), jsonerr)
	assert.Equal(suite.T(), 0, len(actual))
}

func (suite *TransactionControllerTestSuite) TestGetByIdTransactionApi_Success() {
	transaction := dummyTransactionList[0]
	suite.transactionUsecaseMock.On("GetById", transaction.Id).Return(transaction, nil)

	controller.NewTransactionController(suite.transactionUsecaseMock, suite.CustomerUsecaseMock, suite.routerMock)

	r := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodGet, "/transaction/"+transaction.Id, nil)
	request.AddCookie(Cookie)
	suite.routerMock.ServeHTTP(r, request)

	var actual model.Transaction
	response := r.Body.String()
	jsonerr := json.Unmarshal([]byte(response), &actual)

	assert.Equal(suite.T(), http.StatusOK, r.Code)
	assert.Nil(suite.T(), err)
	assert.Nil(suite.T(), jsonerr)
	assert.Equal(suite.T(), transaction.Id, actual.Id)
	assert.Equal(suite.T(), transaction.SenderId, actual.SenderId)
	assert.Equal(suite.T(), transaction.ReceiverId, actual.ReceiverId)
	assert.Equal(suite.T(), transaction.Amount, actual.Amount)
}

func (suite *TransactionControllerTestSuite) TestGetByIdTransactionApi_Failed() {
	transaction := dummyTransactionList[0]
	suite.transactionUsecaseMock.On("GetById", transaction.Id).Return(model.Transaction{}, errors.New("failed"))

	controller.NewTransactionController(suite.transactionUsecaseMock, suite.CustomerUsecaseMock, suite.routerMock)

	r := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodGet, "/transaction/"+transaction.Id, nil)
	request.AddCookie(Cookie)
	suite.routerMock.ServeHTTP(r, request)

	var actual model.Transaction
	response := r.Body.String()
	json.Unmarshal([]byte(response), &actual)

	assert.Equal(suite.T(), http.StatusBadRequest, r.Code)
	assert.Nil(suite.T(), err)
	assert.NotEqual(suite.T(), transaction, actual)
}

func (suite *TransactionControllerTestSuite) TestCreateNewTransactionApi_Success() {
	sender := dummyTxCustomerList[0]
	receiver := dummyTxCustomerList[1]

	transaction := model.Transaction{
		Id:         "test2",
		SenderId:   sender.Id,
		ReceiverId: receiver.Id,
		Amount:     100,
	}
	suite.CustomerUsecaseMock.On("GetById", transaction.SenderId).Return(sender, nil)
	suite.CustomerUsecaseMock.On("GetById", transaction.ReceiverId).Return(receiver, nil)
	suite.CustomerUsecaseMock.On("UpdateBothBalance", transaction.Amount, sender.Id, receiver.Id).Return(nil)

	suite.transactionUsecaseMock.On("Insert", &transaction).Return(transaction, nil)

	controller.NewTransactionController(suite.transactionUsecaseMock, suite.CustomerUsecaseMock, suite.routerMock)

	r := httptest.NewRecorder()

	reqBody, _ := json.Marshal(transaction)
	request, err := http.NewRequest(http.MethodPost, "/transaction", bytes.NewBuffer(reqBody))
	request.AddCookie(Cookie)
	suite.routerMock.ServeHTTP(r, request)

	var actual model.Transaction
	response := r.Body.String()
	jsonerr := json.Unmarshal([]byte(response), &actual)

	assert.Equal(suite.T(), http.StatusOK, r.Code)
	assert.Nil(suite.T(), err)
	assert.Nil(suite.T(), jsonerr)
	assert.Equal(suite.T(), transaction, actual)
}

func (suite *TransactionControllerTestSuite) TestCreateNewTransactionApi_FailedExceedBalance() {
	sender := dummyTxCustomerList[0]
	receiver := dummyTxCustomerList[1]

	transaction := model.Transaction{
		Id:         "test2",
		SenderId:   sender.Id,
		ReceiverId: receiver.Id,
		Amount:     99999,
	}
	suite.CustomerUsecaseMock.On("GetById", transaction.SenderId).Return(sender, nil)
	suite.CustomerUsecaseMock.On("GetById", transaction.ReceiverId).Return(receiver, nil)
	suite.CustomerUsecaseMock.On("UpdateBothBalance", transaction.Amount, sender.Id, receiver.Id).Return(nil)

	suite.transactionUsecaseMock.On("Insert", &transaction).Return(model.Transaction{}, errors.New("exceed balance"))

	controller.NewTransactionController(suite.transactionUsecaseMock, suite.CustomerUsecaseMock, suite.routerMock)

	r := httptest.NewRecorder()

	reqBody, _ := json.Marshal(transaction)
	request, err := http.NewRequest(http.MethodPost, "/transaction", bytes.NewBuffer(reqBody))
	request.AddCookie(Cookie)
	suite.routerMock.ServeHTTP(r, request)

	var actual model.Transaction
	response := r.Body.String()
	jsonerr := json.Unmarshal([]byte(response), &actual)

	assert.Equal(suite.T(), http.StatusBadRequest, r.Code)
	assert.Nil(suite.T(), err)
	assert.Nil(suite.T(), jsonerr)
	assert.Equal(suite.T(), model.Transaction{}, actual)
}

func (suite *TransactionControllerTestSuite) TestCreateNewTransactionApi_FailedAmountIsZeroOrNegative() {
	sender := dummyTxCustomerList[0]
	receiver := dummyTxCustomerList[1]

	transaction := model.Transaction{
		Id:         "test2",
		SenderId:   sender.Id,
		ReceiverId: receiver.Id,
		Amount:     -1,
	}
	suite.CustomerUsecaseMock.On("GetById", transaction.SenderId).Return(sender, nil)
	suite.CustomerUsecaseMock.On("GetById", transaction.ReceiverId).Return(receiver, nil)
	suite.CustomerUsecaseMock.On("UpdateBothBalance", transaction.Amount, sender.Id, receiver.Id).Return(nil)

	suite.transactionUsecaseMock.On("Insert", &transaction).Return(model.Transaction{}, errors.New("invalid amount"))

	controller.NewTransactionController(suite.transactionUsecaseMock, suite.CustomerUsecaseMock, suite.routerMock)

	r := httptest.NewRecorder()

	reqBody, _ := json.Marshal(transaction)
	request, err := http.NewRequest(http.MethodPost, "/transaction", bytes.NewBuffer(reqBody))
	request.AddCookie(Cookie)
	suite.routerMock.ServeHTTP(r, request)

	var actual model.Transaction
	response := r.Body.String()
	jsonerr := json.Unmarshal([]byte(response), &actual)

	assert.Equal(suite.T(), http.StatusBadRequest, r.Code)
	assert.Nil(suite.T(), err)
	assert.Nil(suite.T(), jsonerr)
	assert.Equal(suite.T(), model.Transaction{}, actual)
}

func (suite *TransactionControllerTestSuite) TestCreateNewTransactionApi_FailedSendToSelf() {
	sender := dummyTxCustomerList[0]

	transaction := model.Transaction{
		Id:         "test2",
		SenderId:   sender.Id,
		ReceiverId: sender.Id,
		Amount:     100,
	}
	suite.CustomerUsecaseMock.On("GetById", transaction.SenderId).Return(sender, nil)
	suite.transactionUsecaseMock.On("Insert", &transaction).Return(model.Transaction{}, errors.New("cant send to self"))

	controller.NewTransactionController(suite.transactionUsecaseMock, suite.CustomerUsecaseMock, suite.routerMock)

	r := httptest.NewRecorder()

	reqBody, _ := json.Marshal(transaction)
	request, err := http.NewRequest(http.MethodPost, "/transaction", bytes.NewBuffer(reqBody))
	request.AddCookie(Cookie)
	suite.routerMock.ServeHTTP(r, request)

	var actual model.Transaction
	response := r.Body.String()
	jsonerr := json.Unmarshal([]byte(response), &actual)

	assert.Equal(suite.T(), http.StatusBadRequest, r.Code)
	assert.Nil(suite.T(), err)
	assert.Nil(suite.T(), jsonerr)
	assert.Equal(suite.T(), model.Transaction{}, actual)
}

func TestTransactionControllerTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionControllerTestSuite))
}
