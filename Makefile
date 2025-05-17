DB_USER=user
DB_PASS=pass
DB_PORT=6432
DB_NAME=uow


export DSN=postgresql://$(DB_USER):$(DB_PASS)@localhost:$(DB_PORT)/$(DB_NAME)?sslmode=disable

run_db:
	docker run --name uow-postgres -e POSTGRES_USER=$(DB_USER) -e POSTGRES_PASSWORD=$(DB_PASS) -e POSTGRES_DB=$(DB_NAME) -p $(DB_PORT):5432 -d postgres

run_migration:
	go run cmd/migration/main.go


run_app:
	go run cmd/app/main.go