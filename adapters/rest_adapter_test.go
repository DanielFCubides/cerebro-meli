package adapters

import (
	"bytes"
	"cerebro/usecase/mocks"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type RestAdapterSuite struct {
	suite.Suite
	usecase *mocks.MutantVerifier
	adapter *RestAdapter
}

func TestRestAdapterInit(t *testing.T) {
	suite.Run(t, new(RestAdapterSuite))
}

func (r *RestAdapterSuite) SetupSuite() {
	gin.SetMode(gin.TestMode)
	r.usecase = new(mocks.MutantVerifier)
	r.adapter = NewRestAdapter(r.usecase)
}

func (r *RestAdapterSuite) TestRestAdapter_GetStats() {
	recoder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recoder)
	c.Request, _ = http.NewRequest(http.MethodPost, "/stats", nil)
	r.usecase.Mock.On("").
		Return(true).
		Once()
}

func (r *RestAdapterSuite) TestRestAdapter_MutantVerifierMutant() {
	recoder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recoder)
	postBody, _ := json.Marshal(map[string][]string{
		"dna": {"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"},
	})
	responseBody := bytes.NewBuffer(postBody)
	c.Request, _ = http.NewRequest(http.MethodPost, "/mutant/", responseBody)
	r.usecase.Mock.On("IsMutant", mock.Anything).
		Return(true).
		Once()
	r.adapter.MutantVerifier(c)
	r.Equal(http.StatusOK, recoder.Code)

}

func (r *RestAdapterSuite) TestRestAdapter_MutantVerifierBadRequest() {
	recoder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recoder)
	postBody, _ := json.Marshal(map[string]string{
		"hello": "world",
	})
	responseBody := bytes.NewBuffer(postBody)
	c.Request, _ = http.NewRequest(http.MethodPost, "/mutant/", responseBody)
	r.usecase.Mock.On("IsMutant", mock.Anything).
		Return(true).
		Once()
	r.adapter.MutantVerifier(c)
	r.Equal(http.StatusBadRequest, recoder.Code)
}

func (r *RestAdapterSuite) TestRestAdapter_MutantVerifierHuman() {
	recoder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recoder)
	postBody, _ := json.Marshal(map[string][]string{
		"dna": {"ATGCGA", "CCGTCC", "TTATGT", "AGAAGG", "ACCCTA", "TCACTG"},
	})
	responseBody := bytes.NewBuffer(postBody)
	c.Request, _ = http.NewRequest(http.MethodPost, "/mutant/", responseBody)
	r.usecase.Mock.On("IsMutant", mock.Anything).
		Return(false).
		Once()
	r.adapter.MutantVerifier(c)
	r.Equal(http.StatusForbidden, recoder.Code)
}
