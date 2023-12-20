FROM golang:alpine

WORKDIR /app

RUN apk add make

RUN go install github.com/cosmtrek/air@latest

COPY . .

RUN go get ./...

RUN make build

CMD ["./bin", "main"]