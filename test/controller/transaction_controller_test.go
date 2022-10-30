package controller_test

import (
	"encoding/json"
	"errors"
	"fmt"
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
		SenderId:   "sender2",
		ReceiverId: "receiver2",
		Amount:     85858,
		Created_at: time.Now(),
	},
}

var cookie = &http.Cookie{
	Name:  "session",
	Value: utils.GenerateId(),
	Path:  "/",
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
	customerUsecaseMock    *CustomerUsecaseMock
	routerMock             *gin.Engine
}

func (suite *TransactionControllerTestSuite) SetupTest() {
	suite.routerMock = gin.Default()
	suite.transactionUsecaseMock = new(TransactionUsecaseMock)
}

func (suite *TransactionControllerTestSuite) TestListTransactionApi_Success() {
	suite.transactionUsecaseMock.On("GetAll").Return(dummyTransactionList, nil)

	controller.NewTransactionController(suite.transactionUsecaseMock, suite.customerUsecaseMock, suite.routerMock)

	r := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodGet, "/transaction", nil)
	suite.routerMock.ServeHTTP(r, request)
	request.AddCookie(cookie)

	var actual []model.Transaction
	response := r.Body.String()
	jsonerr := json.Unmarshal([]byte(response), &actual)

	utils.Line()
	fmt.Println(response)
	utils.Line()

	assert.Equal(suite.T(), http.StatusOK, r.Code)
	return
	assert.Nil(suite.T(), err)
	assert.Nil(suite.T(), jsonerr)
	assert.Equal(suite.T(), dummyTransactionList, actual)
}

func (suite *TransactionControllerTestSuite) TestListTransactionApi_Failed() {
	suite.transactionUsecaseMock.On("GetAll").Return(nil, errors.New("failed"))

	controller.NewTransactionController(suite.transactionUsecaseMock, suite.customerUsecaseMock, suite.routerMock)

	r := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodGet, "/transaction", nil)
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

	controller.NewTransactionController(suite.transactionUsecaseMock, suite.customerUsecaseMock, suite.routerMock)

	r := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodGet, "/transaction/"+transaction.Id, nil)
	suite.routerMock.ServeHTTP(r, request)

	var actual model.Transaction
	response := r.Body.String()
	jsonerr := json.Unmarshal([]byte(response), &actual)

	assert.Equal(suite.T(), http.StatusOK, r.Code)
	assert.Nil(suite.T(), err)
	assert.Nil(suite.T(), jsonerr)
	assert.Equal(suite.T(), transaction, actual)
}

func (suite *TransactionControllerTestSuite) TestGetByIdTransactionApi_Failed() {
	transaction := dummyTransactionList[0]
	suite.transactionUsecaseMock.On("GetById", transaction.Id).Return(model.Transaction{}, errors.New("failed"))

	controller.NewTransactionController(suite.transactionUsecaseMock, suite.customerUsecaseMock, suite.routerMock)

	r := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodGet, "/transaction/"+transaction.Id, nil)
	suite.routerMock.ServeHTTP(r, request)

	var actual model.Transaction
	response := r.Body.String()
	jsonerr := json.Unmarshal([]byte(response), &actual)

	assert.Equal(suite.T(), http.StatusNotFound, r.Code)
	assert.Nil(suite.T(), err)
	assert.Nil(suite.T(), jsonerr)
	assert.Equal(suite.T(), model.Transaction{}, actual)
}

// func (suite *TransactionControllerTestSuite) TestCreateNewTransactionApi_Success() {
// 	transaction := model.Transaction{
// 		Id:         "test2",
// 		SenderId:   "sender2",
// 		ReceiverId: "receiver2",
// 		Amount:     85858,
// 		Created_at: time.Now(),
// 	}

// 	suite.transactionUsecaseMock.On("GetByUsername", transaction.Username).Return(model.Transaction{})
// 	suite.transactionUsecaseMock.On("Insert", &transaction).Return(transaction, nil)

// 	controller.NewTransactionController(suite.transactionUsecaseMock, suite.customerUsecaseMock, suite.routerMock)

// 	r := httptest.NewRecorder()

// 	reqBody, _ := json.Marshal(transaction)
// 	request, err := http.NewRequest(http.MethodPost, "/transaction", bytes.NewBuffer(reqBody))
// 	suite.routerMock.ServeHTTP(r, request)

// 	var actual model.Transaction
// 	response := r.Body.String()
// 	jsonerr := json.Unmarshal([]byte(response), &actual)

// 	assert.Equal(suite.T(), http.StatusOK, r.Code)
// 	assert.Nil(suite.T(), err)
// 	assert.Nil(suite.T(), jsonerr)
// 	assert.Equal(suite.T(), transaction, actual)
// }

// func (suite *TransactionControllerTestSuite) TestCreateNewTransactionApi_FailedTakenUsername() {
// 	transaction := model.Transaction{
// 		Id:       utils.GenerateId(),
// 		Name:     "new cust",
// 		Username: "new",
// 		Password: "new password",
// 		Balance:  123,
// 	}

// 	suite.transactionUsecaseMock.On("GetByUsername", transaction.Username).Return(transaction)
// 	suite.transactionUsecaseMock.On("Insert", &transaction).Return(model.Transaction{}, nil)

// 	controller.NewTransactionController(suite.transactionUsecaseMock, suite.customerUsecaseMock, suite.routerMock)

// 	r := httptest.NewRecorder()

// 	reqBody, _ := json.Marshal(transaction)
// 	request, err := http.NewRequest(http.MethodPost, "/transaction", bytes.NewBuffer(reqBody))
// 	suite.routerMock.ServeHTTP(r, request)

// 	var actual model.Transaction
// 	response := r.Body.String()
// 	jsonerr := json.Unmarshal([]byte(response), &actual)

// 	assert.Equal(suite.T(), http.StatusBadRequest, r.Code)
// 	assert.Nil(suite.T(), err)
// 	assert.Nil(suite.T(), jsonerr)
// }

// func (suite *TransactionControllerTestSuite) TestCreateNewTransactionApi_FailedEmptyStruct() {
// 	transaction := model.Transaction{}

// 	suite.transactionUsecaseMock.On("GetByUsername", transaction.Username).Return(model.Transaction{})
// 	suite.transactionUsecaseMock.On("Insert", &transaction).Return(model.Transaction{}, nil)

// 	controller.NewTransactionController(suite.transactionUsecaseMock, suite.customerUsecaseMock, suite.routerMock)

// 	r := httptest.NewRecorder()

// 	reqBody, _ := json.Marshal(transaction)
// 	request, err := http.NewRequest(http.MethodPost, "/transaction", bytes.NewBuffer(reqBody))
// 	suite.routerMock.ServeHTTP(r, request)

// 	var actual model.Transaction
// 	response := r.Body.String()
// 	jsonerr := json.Unmarshal([]byte(response), &actual)

// 	assert.Equal(suite.T(), http.StatusBadRequest, r.Code)
// 	assert.Nil(suite.T(), err)
// 	assert.Nil(suite.T(), jsonerr)
// }

func TestTransactionControllerTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionControllerTestSuite))
}
