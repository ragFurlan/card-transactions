FROM golang:1.20.6

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /app/card-transactions

EXPOSE 8080

CMD ["/app/card-transactions"]