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

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

var dummyCustomerList = []model.Customer{
	{
		Id:       "test",
		Name:     "test",
		Username: "test",
		Password: utils.Sha1("test"),
		Balance:  5000,
	},

	{
		Id:       "test2",
		Name:     "test2",
		Username: "test2",
		Password: utils.Sha1("test2"),
		Balance:  5000,
	},
}

type CustomerUsecaseMock struct {
	mock.Mock
}

func (cu *CustomerUsecaseMock) GetAll() ([]model.Customer, error) {
	args := cu.Called()
	if args.Get(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.Customer), nil
}

func (cu *CustomerUsecaseMock) GetById(id string) (model.Customer, error) {
	args := cu.Called(id)
	if args.Get(1) != nil {
		return model.Customer{}, args.Error(1)
	}
	return args.Get(0).(model.Customer), nil
}

func (cu *CustomerUsecaseMock) GetByUsername(username string) model.Customer {
	args := cu.Called(username)
	if args.Get(0).(model.Customer).Id == "" {
		return model.Customer{}
	}
	return args.Get(0).(model.Customer)
}

func (cu *CustomerUsecaseMock) GetByCredentials(username, password string) (model.Customer, error) {
	args := cu.Called(username, password)
	if args.Get(1) != nil {
		return model.Customer{}, args.Error(1)
	}
	return args.Get(0).(model.Customer), nil
}

func (cu *CustomerUsecaseMock) Insert(customer *model.Customer) (model.Customer, error) {
	args := cu.Called(customer)
	if args.Get(1) != nil {
		return model.Customer{}, args.Error(1)
	}
	return args.Get(0).(model.Customer), nil
}

func (cu *CustomerUsecaseMock) UpdateBothBalance(payAmount int, senderId, receiverId string) error {
	args := cu.Called(payAmount, senderId, receiverId)
	return args.Error(0)
}

type CustomerControllerTestSuite struct {
	suite.Suite
	CustomerUsecaseMock *CustomerUsecaseMock
	routerMock          *gin.Engine
}

func (suite *CustomerControllerTestSuite) SetupTest() {
	suite.routerMock = gin.Default()
	suite.CustomerUsecaseMock = new(CustomerUsecaseMock)
}

func (suite *CustomerControllerTestSuite) TestListCustomerApi_Success() {
	suite.CustomerUsecaseMock.On("GetAll").Return(dummyCustomerList, nil)

	controller.NewCustomerController(suite.CustomerUsecaseMock, suite.routerMock)

	r := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodGet, "/customer", nil)
	suite.routerMock.ServeHTTP(r, request)

	var actual []model.Customer
	response := r.Body.String()
	jsonerr := json.Unmarshal([]byte(response), &actual)

	assert.Equal(suite.T(), http.StatusOK, r.Code)
	assert.Nil(suite.T(), err)
	assert.Nil(suite.T(), jsonerr)
	assert.Equal(suite.T(), dummyCustomerList, actual)
}

func (suite *CustomerControllerTestSuite) TestListCustomerApi_Failed() {
	suite.CustomerUsecaseMock.On("GetAll").Return(nil, errors.New("failed"))

	controller.NewCustomerController(suite.CustomerUsecaseMock, suite.routerMock)

	r := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodGet, "/customer", nil)
	suite.routerMock.ServeHTTP(r, request)

	var actual []model.Customer
	response := r.Body.String()
	jsonerr := json.Unmarshal([]byte(response), &actual)

	assert.Equal(suite.T(), http.StatusInternalServerError, r.Code)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), jsonerr)
	assert.Equal(suite.T(), 0, len(actual))
}

func (suite *CustomerControllerTestSuite) TestGetByIdCustomerApi_Success() {
	customer := dummyCustomerList[0]
	suite.CustomerUsecaseMock.On("GetById", customer.Id).Return(customer, nil)

	controller.NewCustomerController(suite.CustomerUsecaseMock, suite.routerMock)

	r := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodGet, "/customer/"+customer.Id, nil)
	suite.routerMock.ServeHTTP(r, request)

	var actual model.Customer
	response := r.Body.String()
	jsonerr := json.Unmarshal([]byte(response), &actual)

	assert.Equal(suite.T(), http.StatusOK, r.Code)
	assert.Nil(suite.T(), err)
	assert.Nil(suite.T(), jsonerr)
	assert.Equal(suite.T(), customer, actual)
}

func (suite *CustomerControllerTestSuite) TestGetByIdCustomerApi_Failed() {
	customer := dummyCustomerList[0]
	suite.CustomerUsecaseMock.On("GetById", customer.Id).Return(model.Customer{}, errors.New("failed"))

	controller.NewCustomerController(suite.CustomerUsecaseMock, suite.routerMock)

	r := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodGet, "/customer/"+customer.Id, nil)
	suite.routerMock.ServeHTTP(r, request)

	var actual model.Customer
	response := r.Body.String()
	jsonerr := json.Unmarshal([]byte(response), &actual)

	assert.Equal(suite.T(), http.StatusNotFound, r.Code)
	assert.Nil(suite.T(), err)
	assert.Nil(suite.T(), jsonerr)
	assert.Equal(suite.T(), model.Customer{}, actual)
}

func (suite *CustomerControllerTestSuite) TestGetByCredentialsCustomerApi_Success() {
	customer := dummyCustomerList[0]
	suite.CustomerUsecaseMock.On("GetByCredentials", customer.Username, customer.Password).Return(customer, nil)

	controller.NewLoginController(suite.CustomerUsecaseMock, suite.routerMock)

	r := httptest.NewRecorder()

	credentials := model.Credential{
		Username: customer.Username,
		Password: customer.Password,
	}

	reqBody, _ := json.Marshal(credentials)
	request, err := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(reqBody))
	suite.routerMock.ServeHTTP(r, request)

	assert.Equal(suite.T(), http.StatusOK, r.Code)
	assert.Nil(suite.T(), err)
}

func (suite *CustomerControllerTestSuite) TestGetByCredentialsCustomerApi_Failed() {
	customer := dummyCustomerList[0]
	suite.CustomerUsecaseMock.On("GetByCredentials", customer.Username, customer.Password).Return(model.Customer{}, errors.New("failed"))

	controller.NewLoginController(suite.CustomerUsecaseMock, suite.routerMock)

	r := httptest.NewRecorder()

	credentials := model.Credential{
		Username: customer.Username,
		Password: customer.Password,
	}

	reqBody, _ := json.Marshal(credentials)
	request, err := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(reqBody))
	suite.routerMock.ServeHTTP(r, request)

	var actual model.Customer
	response := r.Body.String()
	jsonerr := json.Unmarshal([]byte(response), &actual)

	assert.Equal(suite.T(), http.StatusNotFound, r.Code)
	assert.Nil(suite.T(), err)
	assert.Nil(suite.T(), jsonerr)
}

func (suite *CustomerControllerTestSuite) TestCreateNewCustomerApi_Success() {
	customer := model.Customer{
		Id:       utils.GenerateId(),
		Name:     "new cust",
		Username: "new",
		Password: "new password",
		Balance:  123,
	}

	suite.CustomerUsecaseMock.On("GetByUsername", customer.Username).Return(model.Customer{})
	suite.CustomerUsecaseMock.On("Insert", &customer).Return(customer, nil)

	controller.NewCustomerController(suite.CustomerUsecaseMock, suite.routerMock)

	r := httptest.NewRecorder()

	reqBody, _ := json.Marshal(customer)
	request, err := http.NewRequest(http.MethodPost, "/customer", bytes.NewBuffer(reqBody))
	suite.routerMock.ServeHTTP(r, request)

	var actual model.Customer
	response := r.Body.String()
	jsonerr := json.Unmarshal([]byte(response), &actual)

	assert.Equal(suite.T(), http.StatusOK, r.Code)
	assert.Nil(suite.T(), err)
	assert.Nil(suite.T(), jsonerr)
	assert.Equal(suite.T(), customer, actual)
}

func (suite *CustomerControllerTestSuite) TestCreateNewCustomerApi_FailedTakenUsername() {
	customer := model.Customer{
		Id:       utils.GenerateId(),
		Name:     "new cust",
		Username: "new",
		Password: "new password",
		Balance:  123,
	}

	suite.CustomerUsecaseMock.On("GetByUsername", customer.Username).Return(customer)
	suite.CustomerUsecaseMock.On("Insert", &customer).Return(model.Customer{}, nil)

	controller.NewCustomerController(suite.CustomerUsecaseMock, suite.routerMock)

	r := httptest.NewRecorder()

	reqBody, _ := json.Marshal(customer)
	request, err := http.NewRequest(http.MethodPost, "/customer", bytes.NewBuffer(reqBody))
	suite.routerMock.ServeHTTP(r, request)

	var actual model.Customer
	response := r.Body.String()
	jsonerr := json.Unmarshal([]byte(response), &actual)

	assert.Equal(suite.T(), http.StatusBadRequest, r.Code)
	assert.Nil(suite.T(), err)
	assert.Nil(suite.T(), jsonerr)
}

func (suite *CustomerControllerTestSuite) TestCreateNewCustomerApi_FailedEmptyStruct() {
	customer := model.Customer{}

	suite.CustomerUsecaseMock.On("GetByUsername", customer.Username).Return(model.Customer{})
	suite.CustomerUsecaseMock.On("Insert", &customer).Return(model.Customer{}, nil)

	controller.NewCustomerController(suite.CustomerUsecaseMock, suite.routerMock)

	r := httptest.NewRecorder()

	reqBody, _ := json.Marshal(customer)
	request, err := http.NewRequest(http.MethodPost, "/customer", bytes.NewBuffer(reqBody))
	suite.routerMock.ServeHTTP(r, request)

	var actual model.Customer
	response := r.Body.String()
	jsonerr := json.Unmarshal([]byte(response), &actual)

	assert.Equal(suite.T(), http.StatusBadRequest, r.Code)
	assert.Nil(suite.T(), err)
	assert.Nil(suite.T(), jsonerr)
}

func TestCustomerControllerTestSuite(t *testing.T) {
	suite.Run(t, new(CustomerControllerTestSuite))
}
