package api

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go-pos/delivery/apprequest"
	"go-pos/model"
	"net/http"
	"net/http/httptest"
	"testing"
)

var dummyCashiers = []model.Cashier{
	{
		Name: "Dummy Cashier 1",
		Passcode: "passworddummy"
	}, {
		Name: "Dummy Cashier 2",
		Passcode: "asal123"
	},
}

type cashierUseCaseMock struct {
	mock.Mock
}

func (c *cashierUseCaseMock) ListCashier(limit, skip int) []model.Cashier {
	//TODO implement me
	panic("implement me")
}

func (c *cashierUseCaseMock) DetailCashier(id int) model.Cashier {
	//TODO implement me
	panic("implement me")
}

func (c *cashierUseCaseMock) CreateCashier(cashier apprequest.Cashier) (model.Cashier, error) {
	args := c.Called(cashier)
	var dummyCashier *model.Cashier
	if args.Get(0) != nil {
		dummyCashier = args.Get(0).(*model.Cashier)
	}
	return dummyCashier, args.Error(1)
}

func (c *cashierUseCaseMock) UpdateCashier(cashier apprequest.Cashier, id int) error {
	//TODO implement me
	panic("implement me")
}

func (c *cashierUseCaseMock) DeleteCashier(id int) error {
	//TODO implement me
	panic("implement me")
}

type CashierApiTestSuite struct {
	suite.Suite
	useCaseTest	*cashierUseCaseMock
	routerTest	*gin.Engine
	routerGroupTest	*gin.RouterGroup
}

func (suite *CashierApiTestSuite) SetupTest() {
	suite.useCaseTest = new(cashierUseCaseMock)
	suite.routerTest = gin.Default()
	suite.routerGroupTest = suite.routerTest.Group("/cashiers")
}

type MockResponse struct {
	Message model.Cashier
}

func (suite *CashierApiTestSuite) Test_CreateCashier_Success() {
	dummyCashier := dummyCashiers[1]
	suite.useCaseTest.On("NewRegistration", dummyCashier).Return(&dummyCashier, nil)
	cashierApi, _ := NewCashierApi(suite.routerGroupTest, suite.useCaseTest)
	handler := cashierApi.CreateCashier
	suite.routerTest.POST("",handler)

	rr:= httptest.NewRecorder()
	reqBody, _ := json.Marshal(dummyCashier)
	request, _ := http.NewRequest(http.MethodPost,"/cashiers", bytes.NewBuffer(reqBody))
	suite.routerTest.ServeHTTP(rr, request)
	assert.Equal(suite.T(), rr.Code, http.StatusOK)

	a := rr.Body.String()
	actualCashier := new(MockResponse)
	json.Unmarshal([]byte(a), actualCashier)
	assert.Equal(suite.T(), dummyCashier.Name, actualCashier.Message.Name)
}

func (suite *CashierApiTestSuite) Test_CreateStudent_FailedBinding() {
	dummyCashierWithoutName := model.Cashier{
		Name: "",
	}
	cashierApi, _ := NewCashierApi(suite.routerGroupTest, suite.useCaseTest)
	handler := cashierApi.CreateCashier
	suite.routerTest.POST("", handler)

	rr := httptest.NewRecorder()
	reqBody, _ := json.Marshal(dummyCashierWithoutName)
	request, _ := http.NewRequest(http.MethodPost,"/cashiers", bytes.NewBuffer(reqBody))
	suite.routerTest.ServeHTTP(rr, request)
	assert.Equal(suite.T(), rr.Code, http.StatusBadRequest)
}

func TestCashierApiTestSuite(t *testing.T) {
	suite.Run(t, new(CashierApiTestSuite))
}