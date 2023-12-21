FROM golang:alpine

WORKDIR /app

RUN apk add make

RUN go install github.com/cosmtrek/air@latest

COPY . .

RUN go get ./...

RUN go build -o ./bin/main ./cmd/app

CMD ["go", "run", "./cmd/app"]