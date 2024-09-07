.PHONY: build lint test migrate-up migrate-down

build:
	sqlc generate
	GOOS=linux go build -o arvosana ./cmd

lint:
	golangci-lint run

test:
	go test ./...

migrate-up:
	migrate -path db/migrations -database sqlite3://${DB} up

migrate-down:
	migrate -path db/migrations -database sqlite3://${DB} down
