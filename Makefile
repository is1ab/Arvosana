.PHONY: build migrate-up migrate-down

build:
	GOOS=linux go build -o arvosana ./cmd

migrate-up:
	migrate -path db/migrations -database sqlite3://${DB} up

migrate-down:
	migrate -path db/migrations -database sqlite3://${DB} down
