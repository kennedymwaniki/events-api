APP_NAME= events_api
CMD_DIR := ./cmd/server

## Run the application
.PHONY: run
run:
    @echo "==> Running locally..."
    GO111MODULE=on go run $(CMD_DIR)

# test the application
.PHONY: test
test:
    @echo "==> Running tests..."
    GO111MODULE=on go test -v ./...	

migrate-up:
    @echo "==> Applying new migrations..."
    migrate -path ./migrations -database "postgres://postgres:password@localhost:5432/events_db?sslmode=disable" up

migrate-down:
    @echo "==> Rolling back last migration..."
    migrate -path ./migrations -database "postgres://postgres:password@localhost:5432/events_db?sslmode=disable" down

migrate-force:
    migrate -path ./migrations -database "postgres://postgres:password@localhost:5432/events_db?sslmode=disable" force 7
