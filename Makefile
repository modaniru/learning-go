.PHONY: postgres-run
postgres-run:
	docker run -p 5433:5432 -e POSTGRES_PASSWORD=qwerty --name my-learn-postgres -d postgres:latest

.PHONY: postgres-start
postgres-start:
	docker start my-learn-postgres

.PHONY: migrateUp
migrateUp:
	migrate -path db/migrations -database "postgresql://postgres:qwerty@localhost:5433?sslmode=disable" up

.PHONY: migrateDown
migrateDown:
	migrate -path db/migrations -database "postgresql://postgres:qwerty@localhost:5433?sslmode=disable" down

.PHONY: sqlc
sqlc:
	sqlc generate

.PHONY: test
test:
	go test ./... -v -cover