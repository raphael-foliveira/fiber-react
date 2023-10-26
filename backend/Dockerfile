FROM golang:alpine

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

COPY . .

RUN go mod tidy

CMD ["air"]