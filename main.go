package main

import (
	controller "card-transactions/internal/controllers/handlers"
	repository "card-transactions/internal/platform/repositories"
	"card-transactions/internal/usecase/accounts"
	"card-transactions/internal/usecase/operationType"
	"card-transactions/internal/usecase/transaction"
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var (
	handlerAccount     *controller.AccountHandler
	handlerTransaction *controller.TransactionHandler
	urlAccounts        = "../internal/platform/data/accounts.txt"
	urlTransactions    = "../internal/platform/data/transactions.txt"
	urlOperationTypes  = "../internal/platform/data/operation_types.txt"
)

func main() {

	// Operation Type
	operationTypeRepository := repository.NewOperationTypeRepository(urlOperationTypes)
	operationTypeUsecase := operationType.NewOperationTypeUsecase(operationTypeRepository)

	// Account
	accountRepository := repository.NewAccountsRepository(urlAccounts)
	accountUsecase := accounts.NewAccountsUsecase(accountRepository)
	handlerAccount = controller.NewAccountHandler(accountUsecase)

	//Transaction
	transactionRepository := repository.NewTransactionRepository(urlTransactions)
	transactionUsecase := transaction.NewTransactionUsecase(transactionRepository, *operationTypeUsecase, *accountUsecase)
	handlerTransaction = controller.NewTransactionHandler(transactionUsecase)

	StartServer()
}

func StartServer() {
	router := RegisterRoutes(handlerAccount, handlerTransaction)
	headers := handlers.AllowedHeaders([]string{"Content-Type"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})

	fmt.Println("Server listening on http://localhost:8080")
	http.ListenAndServe(":8080", handlers.CORS(headers, methods, origins)(router))
}

func RegisterRoutes(handlerAccount *controller.AccountHandler, handlerTransaction *controller.TransactionHandler) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/accounts/{accountId}", handlerAccount.GetByID)
	router.HandleFunc("/accounts", handlerAccount.Add)
	router.HandleFunc("/transactions", handlerTransaction.Add)

	return router
}
