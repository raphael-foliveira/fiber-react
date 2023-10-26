FROM golang:alpine

WORKDIR /app

RUN apk add make

RUN go install github.com/cosmtrek/air@latest

COPY . .

RUN go get ./...

CMD ["make", "dev"]