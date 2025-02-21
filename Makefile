migrate-up:
	goose -dir ./db/migrations postgres "host=localhost user=postgres dbname=postgres password=10Denis00 sslmode=disable" up

migrate-create:
	goose -dir db/migrations create $(F) sql

lint:
	golangci-lint run ./...