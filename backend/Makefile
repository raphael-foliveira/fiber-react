dev:
	air

build:
	GOOS=linux GOARCH=amd64 go build -o ./bin/main ./cmd/app
	GOOS=windows GOARCH=amd64 go build -o ./bin/main.exe ./cmd/app
	GOOS=darwin GOARCH=amd64 go build -o ./bin/main-darwin ./cmd/app

run:
	go run ./cmd/app