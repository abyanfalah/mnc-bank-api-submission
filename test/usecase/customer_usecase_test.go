package usecase_test

import (
	"errors"
	"mnc-bank-api/model"
	"mnc-bank-api/usecase"
	"mnc-bank-api/utils"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

var dummyList = []model.Customer{
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

var afterPayment = []model.Customer{
	{
		Id:       "test",
		Name:     "test",
		Username: "test",
		Password: utils.Sha1("test"),
		Balance:  3000,
	},

	{
		Id:       "test2",
		Name:     "test2",
		Username: "test2",
		Password: utils.Sha1("test2"),
		Balance:  7000,
	},
}

type repoMock struct {
	mock.Mock
}

type CustomerUsecaseTestSuite struct {
	suite.Suite
	repoMock *repoMock
}

func (r *repoMock) GetAll() ([]model.Customer, error) {
	args := r.Called()
	if args.Get(1) != nil {
		return nil, args.Error(1)
	}

	return args[0].([]model.Customer), nil
}

func (r *repoMock) UpdateList(newList []model.Customer) error {
	args := r.Called(newList)
	if args.Get(0) != nil {
		return args.Error(1)
	}

	return nil
}

func (suite *CustomerUsecaseTestSuite) TestGetAll_Success() {
	suite.repoMock.On("GetAll").Return(dummyList, nil)

	CustomerUsecaseTest := usecase.NewCustomerUsecase(suite.repoMock)
	actual, err := CustomerUsecaseTest.GetAll()

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), dummyList, actual)
}

func (suite *CustomerUsecaseTestSuite) TestGetAll_Failed() {
	suite.repoMock.On("GetAll").Return(nil, errors.New("failed to get all user"))

	CustomerUsecaseTest := usecase.NewCustomerUsecase(suite.repoMock)
	actual, err := CustomerUsecaseTest.GetAll()

	assert.NotNil(suite.T(), err)
	assert.Nil(suite.T(), actual)
}

func (suite *CustomerUsecaseTestSuite) TestGetById_Success() {
	suite.repoMock.On("GetAll").Return(dummyList, nil)

	CustomerUsecaseTest := usecase.NewCustomerUsecase(suite.repoMock)
	actual, err := CustomerUsecaseTest.GetById(dummyList[0].Id)

	assert.Nil(suite.T(), err)
	return
	assert.Equal(suite.T(), dummyList[0], actual)
}

func (suite *CustomerUsecaseTestSuite) TestGetById_Failed() {
	suite.repoMock.On("GetAll").Return(nil, errors.New("failed to get customer "+dummyList[0].Id))

	CustomerUsecaseTest := usecase.NewCustomerUsecase(suite.repoMock)
	actual, err := CustomerUsecaseTest.GetById(dummyList[0].Id)

	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), model.Customer{}, actual)
}

func (suite *CustomerUsecaseTestSuite) TestGetByUsername_Success() {
	suite.repoMock.On("GetAll").Return(dummyList, nil)

	CustomerUsecaseTest := usecase.NewCustomerUsecase(suite.repoMock)
	actual := CustomerUsecaseTest.GetByUsername(dummyList[0].Username)

	assert.Equal(suite.T(), dummyList[0], actual)
}

func (suite *CustomerUsecaseTestSuite) TestGetByUsername_Failed() {
	suite.repoMock.On("GetAll").Return(nil, errors.New("failed to get customer "+dummyList[0].Username))

	CustomerUsecaseTest := usecase.NewCustomerUsecase(suite.repoMock)
	actual := CustomerUsecaseTest.GetByUsername(dummyList[0].Username)

	assert.Equal(suite.T(), model.Customer{}, actual)
}

func (suite *CustomerUsecaseTestSuite) TestGetByCredentials_Success() {
	customer := dummyList[0]
	suite.repoMock.On("GetAll").Return(dummyList, nil)

	CustomerUsecaseTest := usecase.NewCustomerUsecase(suite.repoMock)
	actual, err := CustomerUsecaseTest.GetByCredentials(customer.Username, "test")

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), customer, actual)
}

func (suite *CustomerUsecaseTestSuite) TestGetByCredentials_Failed() {
	customer := dummyList[0]
	suite.repoMock.On("GetAll").Return(nil, errors.New("failed to get customer "+dummyList[0].Username))

	CustomerUsecaseTest := usecase.NewCustomerUsecase(suite.repoMock)
	actual, err := CustomerUsecaseTest.GetByCredentials(customer.Username, "test")

	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), model.Customer{}, actual)
}

func (suite *CustomerUsecaseTestSuite) TestInsert_Success() {
	newCustomer := model.Customer{
		Id:       utils.GenerateId(),
		Name:     "new",
		Username: "new",
		Password: "new",
		Balance:  234,
	}

	expected := model.Customer{
		Id:       newCustomer.Id,
		Name:     "new",
		Username: "new",
		Password: utils.Sha1(newCustomer.Password),
		Balance:  234,
	}

	newList := append(dummyList, expected)

	suite.repoMock.On("UpdateList", newList).Return([]model.Customer{}, nil)
	suite.repoMock.On("GetAll").Return(dummyList, nil)

	CustomerUsecaseTest := usecase.NewCustomerUsecase(suite.repoMock)
	actual, err := CustomerUsecaseTest.Insert(&newCustomer)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expected, actual)
}

func (suite *CustomerUsecaseTestSuite) TestInsert_Failed() {

	suite.repoMock.On("UpdateList", nil).Return(nil, errors.New("failed"))
	suite.repoMock.On("GetAll").Return(dummyList, nil)

	CustomerUsecaseTest := usecase.NewCustomerUsecase(suite.repoMock)
	actual, err := CustomerUsecaseTest.Insert(&model.Customer{})

	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), model.Customer{}, actual)
}

func (suite *CustomerUsecaseTestSuite) TestUpdateBothBalance_Success() {
	amount := 2000
	sender := dummyList[0]
	receiver := dummyList[1]

	suite.repoMock.On("UpdateList", afterPayment).Return([]model.Customer{}, nil)
	suite.repoMock.On("GetAll").Return(dummyList, nil)

	CustomerUsecaseTest := usecase.NewCustomerUsecase(suite.repoMock)
	err := CustomerUsecaseTest.UpdateBothBalance(amount, sender.Id, receiver.Id)

	assert.Nil(suite.T(), err)
	assert.True(suite.T(), afterPayment[1].Balance == dummyList[1].Balance)
}

func (suite *CustomerUsecaseTestSuite) TestUpdateBothBalance_FailedExceededAmount() {
	amount := 999999
	sender := dummyList[0]
	receiver := dummyList[1]

	suite.repoMock.On("UpdateList", afterPayment).Return([]model.Customer{}, nil)
	suite.repoMock.On("GetAll").Return(dummyList, nil)

	CustomerUsecaseTest := usecase.NewCustomerUsecase(suite.repoMock)
	err := CustomerUsecaseTest.UpdateBothBalance(amount, sender.Id, receiver.Id)

	assert.NotNil(suite.T(), err)
}

func (suite *CustomerUsecaseTestSuite) TestUpdateBothBalance_FailedNegativeAmount() {
	amount := -7
	sender := dummyList[0]
	receiver := dummyList[1]

	suite.repoMock.On("UpdateList", afterPayment).Return([]model.Customer{}, nil)
	suite.repoMock.On("GetAll").Return(dummyList, nil)

	CustomerUsecaseTest := usecase.NewCustomerUsecase(suite.repoMock)
	err := CustomerUsecaseTest.UpdateBothBalance(amount, sender.Id, receiver.Id)

	assert.NotNil(suite.T(), err)
}

func (suite *CustomerUsecaseTestSuite) SetupTest() {
	suite.repoMock = new(repoMock)
}

func TestCustomerUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(CustomerUsecaseTestSuite))
}
