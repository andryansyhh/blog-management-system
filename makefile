PROTO_DIR=proto

generate:
	find $(PROTO_DIR) -name "*.proto" -exec protoc \
		--go_out=. \
		--go-grpc_out=. \
		--go_opt=paths=source_relative \
		--go-grpc_opt=paths=source_relative \
		{} \;


# migration database
DB_USER=postgres
DB_PASS=postgres
DB_HOST=localhost
DB_PORT=5432
DB_NAME=tablelink
DB_URL=postgres://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable

DB_URL=postgres://postgres:postgres@localhost:5432/tablelink?sslmode=disable
MIGRATION_PATH=$(shell pwd)/migration

migrate-up:
	migrate -path $(MIGRATION_PATH) -database "$(DB_URL)" up

migrate-down:
	migrate -path $(MIGRATION_PATH) -database "$(DB_URL)" down

migrate-drop:
	migrate -path $(MIGRATION_PATH) -database "$(DB_URL)" drop -f

run:
	go run ./main.go