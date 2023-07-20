FROM golang:1.20.6
WORKDIR /cmd

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /internal/bin/app ./...

CMD ["app"]