package usecase

import (
	"errors"
	"testing"

	"enigmacamp.com/golatihanlagi/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

var dummyCustomers = []model.Customer{
	{
		Id:      "C001",
		Nama:    "Dummy Name 1",
		Address: "Dummy Address 1",
	},
	{
		Id:      "C002",
		Nama:    "Dummy Name 2",
		Address: "Dummy Address 2",
	},
}

type repoMock struct {
	mock.Mock
}

type CustomerUseCaseTestSuite struct {
	suite.Suite
	repoMock *repoMock
}

func (r *repoMock) Create(newCustomer model.Customer) error {
	args := r.Called(newCustomer)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}

func (r *repoMock) RetrieveAll() ([]model.Customer, error) {
	args := r.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.Customer), nil
}

func (r *repoMock) FindById(id string) (model.Customer, error) {
	args := r.Called(id)
	if args.Get(1) != nil {
		return model.Customer{}, args.Error(1)
	}
	return args.Get(0).(model.Customer), nil
}

func (suite *CustomerUseCaseTestSuite) SetupTest() {
	suite.repoMock = new(repoMock)
}

func (suite *CustomerUseCaseTestSuite) TestCustomerFindById_Success() {
	dummyCustomer := dummyCustomers[0]
	suite.repoMock.On("FindById", dummyCustomer.Id).Return(dummyCustomer, nil)
	// seolah-olah call FindById dengan parameter id dan return model.customer, error

	customerUsecaseTest := NewCustomerUseCase(suite.repoMock)
	customer, err := customerUsecaseTest.FindCustomerById(dummyCustomer.Id)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), dummyCustomer.Id, customer.Id)
}

func (suite *CustomerUseCaseTestSuite) TestCustomerFindById_Failed() {
	dummyCustomers := dummyCustomers[0]
	suite.repoMock.On("FindById", dummyCustomers.Id).Return(model.Customer{}, errors.New("failed"))
	customerUseCaseTest := NewCustomerUseCase(suite.repoMock)
	customer, err := customerUseCaseTest.FindCustomerById(dummyCustomers.Id)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), "failed", err.Error())
	assert.Equal(suite.T(), "", customer.Id)
}

func (suite *CustomerUseCaseTestSuite) TestCustomerRetrieveAll_Success() {
	suite.repoMock.On("RetrieveAll").Return(dummyCustomers, nil)
	customerUsecaseTest := NewCustomerUseCase(suite.repoMock)
	customer, err := customerUsecaseTest.GetAllCustomer()
	assert.Nil(suite.T(), err)
	assert.NotEmpty(suite.T(), customer)
}

// masih 75%
func (suite *CustomerUseCaseTestSuite) TestCustomerRetrieveAll_Failed() {
	suite.repoMock.On("RetrieveAll").Return(nil, errors.New("failed"))
	customerUseCaseTest := NewCustomerUseCase(suite.repoMock)
	customers, err := customerUseCaseTest.GetAllCustomer()
	assert.Error(suite.T(), err)
	assert.NotNil(suite.T(), err)
	assert.Empty(suite.T(), customers)
	assert.Equal(suite.T(), []model.Customer(nil), customers)
}

func (suite *CustomerUseCaseTestSuite) TestCustomerCreate_Success() {
	dummyCustomer := dummyCustomers[0]
	suite.repoMock.On("Create", dummyCustomer).Return(nil)
	customerUsecaseTest := NewCustomerUseCase(suite.repoMock)
	err := customerUsecaseTest.RegisterCustomer(dummyCustomer)
	assert.Nil(suite.T(), err)
}

func (suite *CustomerUseCaseTestSuite) TestCustomerCreate_Failed() {
	dummyCustomers := dummyCustomers[0]
	suite.repoMock.On("Create", dummyCustomers).Return(errors.New("failed"))
	customerUseCaseTest := NewCustomerUseCase(suite.repoMock)
	err := customerUseCaseTest.RegisterCustomer(dummyCustomers)
	assert.Equal(suite.T(), "failed", err.Error())
}

func TestCustomerUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(CustomerUseCaseTestSuite))
}

/*
go test -cover
go test ./... -cover -coverprofile=c.out
go tool cover -html=c.out -o "coverage.html"
*/

// untuk mock -> butuh interface repo -> seperti repo asli. method menerima receiver mock
// testing -> testing unit terkecil, gaperlu konek ke db
// repo konek ke db -> karena repo pake interface -> buat repo tiruan yg dikondisikan/set ekspektasinya -> tidak butuh db
// poin penting interface -> membuat struct apapun selama kontrak dibuat
// implementasi -> dibuat bohongan -> hardcode
