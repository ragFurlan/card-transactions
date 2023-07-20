package main

import (
	controller "card-transactions/internal/controllers/handlers"
	repository "card-transactions/internal/platform/repositories"
	"card-transactions/internal/usecase/accounts"
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
)

var (
	accountUseCase *accounts.AccountUseCase
	urlAccounts    = "../internal/platform/data/accounts.txt"
	// urlOperationTypes = "../internal/platform/data/operation_types.txt"
	// urlTransactions   = "../internal/platform/data/transactions.txt"
)

func main() {
	accountRepository := repository.NewAccountsRepository(urlAccounts)
	accountUseCase = accounts.NewAccountsUseCase(accountRepository)
	StartServer()

	fmt.Println("hello word")
}

func StartServer() {
	handler := controller.NewAccountHandler(accountUseCase)
	router := handler.RegisterRoutes()

	headers := handlers.AllowedHeaders([]string{"Content-Type"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})

	fmt.Println("Server listening on http://localhost:8080")
	http.ListenAndServe(":8080", handlers.CORS(headers, methods, origins)(router))
}
