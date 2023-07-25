package handler

import (
	"card-transactions/internal/entity"
	usecase "card-transactions/internal/usecase/accounts"
	repository "card-transactions/test/platform"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

var (
	handler        *AccountHandler
	repositoryMock *repository.MockAccounts
	controller     *gomock.Controller
	usecaseMock    *usecase.AccountUsecase
	anyError       = errors.New("Error")
	id             = "edde2aa4-5a60-44f5-bdf6-db1dd972b41c"
	documentNumber = "123456789010"
)

func (mock *MockHTTP) Do(_ *http.Request) (*http.Response, error) {
	return mock.response, mock.err
}

type MockHTTP struct {
	response *http.Response
	err      error
}

func TestAccountHandler_GetByID(t *testing.T) {
	id := "edde2aa4-5a60-44f5-bdf6-db1dd972b41c"
	setHandlerAndLogMock(t)
	req := httptest.NewRequest("GET", "/accounts/"+id, nil)
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/accounts/{accountId}", handler.GetByID).Methods("GET")

	expectedAccount := entity.Account{
		ID:             id,
		DocumentNumber: "123456789010",
	}

	repositoryMock.EXPECT().GetByID(id).Return(expectedAccount, nil)
	router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)

}

func setHandlerAndLogMock(t *testing.T) {
	controller = gomock.NewController(t)
	repositoryMock = repository.NewMockAccounts(controller)
	usecaseMock = usecase.NewAccountsUsecase(repositoryMock)
	handler = NewAccountHandler(usecaseMock)
}
