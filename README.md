# card-transactions

### Install mockgen if necessary
```
go install github.com/golang/mock/mockgen@v1.6.0
go get github.com/golang/mock/gomock
go get github.com/stretchr/testify/assert


```
### run project locale

Write the line of command bellow:
```
go run main.go
```
 
### run the project with Docker
```
docker build -t card-transaction .

docker run -p 8080:8080 card-transaction
```


## **Create mock**

### **Platform**
```
~/go/bin/mockgen -source=internal/platform/repositories/account.go -destination=test/platform/account.go -package=repository
~/go/bin/mockgen -source=internal/platform/repositories/operationType.go -destination=test/platform/operationType.go -package=repository
~/go/bin/mockgen -source=internal/platform/repositories/transaction.go -destination=test/platform/transaction.go -package=repository
```


### **Usecase**
```
~/go/bin/mockgen -source=internal/usecase/accounts/accountUsecase.go -destination=test/usecase/accountUsecase.go -package=usecase
~/go/bin/mockgen -source=internal/usecase/operationType/operationTypeUsecase.go -destination=test/usecase/operationTypeUsecase.go -package=usecase
~/go/bin/mockgen -source=internal/usecase/transaction/transactionUsecase.go -destination=test/usecase/transactionUsecase.go -package=usecase
```