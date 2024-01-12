swag:
	swag init --parseDependency --parseInternal --parseDepth 1 -md ./documentation -o ./docs

db:
	sqlc generate

migrate:
	migrate -verbose -path ./database/migrations -database $(MIGRATION_URL) up

migrate-down:
	migrate -verbose -path ./database/migrations -database $(MIGRATION_URL) down

migrate-create:
	migrate create -ext sql -dir ./database/migrations -seq $(name)

build:
	go build -o ./bin/run ./main.go

run: migrate build
	./bin/run

fly:
	flyctl auth login

dpl:
	flyctl deploy

fenv:
	flyctl secrets set $(env)

test:
	go test -cover ./pkg/controllers | gocol

.PHONY: swag sql migrate migrate-down migrate-create build run fly dpl fenv test
