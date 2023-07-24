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

// func TestAccountHandler_Add(t *testing.T) {
// 	setHandlerAndLogMock(t)

// 	requestBody := struct {
// 		DocumentNumber string `json:"document_number"`
// 	}{
// 		DocumentNumber: documentNumber,
// 	}

// 	requestBodyBytes, err := json.Marshal(requestBody)
// 	if err != nil {
// 		t.Fatalf("Failed to marshal request body: %v", err)
// 	}

// 	req := httptest.NewRequest("POST", "/accounts", bytes.NewReader(requestBodyBytes))
// 	rr := httptest.NewRecorder()

// 	router := mux.NewRouter()
// 	router.HandleFunc("/accounts", handler.Add).Methods("POST")

// 	requestAccount := entity.Account{
// 		ID:             id,
// 		DocumentNumber: "123456789010",
// 	}

// 	repositoryMock.EXPECT().Save(requestAccount).Return(nil)

// 	router.ServeHTTP(rr, req)

// 	assert.Equal(t, http.StatusOK, rr.Code)

// 	// // Parse the response body and check if it matches the expected data
// 	// var response struct {
// 	// 	AccountID string `json:"account_id"`
// 	// 	Message   string `json:"message"`
// 	// }
// 	// err = json.Unmarshal(rr.Body.Bytes(), &response)
// 	// if err != nil {
// 	// 	t.Errorf("Failed to unmarshal response: %v", err)
// 	// }

// 	// // Validate the response fields
// 	// assert.NotEmpty(t, response.AccountID)
// 	// assert.Equal(t, "Save account with success", response.Message)

// 	// // Assert that the mock AccountUsecase's Save method was called with the correct account data
// 	// mockAccount := mockAccountUsecase.GetLastSavedAccount()
// 	// assert.Equal(t, response.AccountID, mockAccount.ID)
// 	// assert.Equal(t, requestBody.DocumentNumber, mockAccount.DocumentNumber)
// }

func setHandlerAndLogMock(t *testing.T) {
	controller = gomock.NewController(t)
	repositoryMock = repository.NewMockAccounts(controller)
	usecaseMock = usecase.NewAccountsUsecase(repositoryMock)
	handler = NewAccountHandler(usecaseMock)
}
