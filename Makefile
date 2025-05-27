sqlc_gen:	
	sqlc generate --file ./configs/sqlc.yaml

migrate_db: sqlc_gen
	atlas migrate diff $(MIGRATION_NAME) \
		--dir "file://migrations" \
		--to "file://database/schema.sql" \
		--dev-url "docker://postgres?search_path=public"

validate:
	go generate ./...
	go mod tidy
	go build ./cmd/lqd && rm lqd # check if project can build
	swag init --parseDependency --parseInternal --parseFuncBody -d ./internal/routes/api/ -g ./api.go -o ./internal/routes/api/docs -td "[[,]]"
	swag fmt
	go fmt ./...
	go vet ./...
	golangci-lint run --fix
	go build ./cmd/lqd && rm lqd

all: validate
