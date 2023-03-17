postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=triadpass -d postgres:15.2-alpine

createdb:
	docker exec -it postgres15.2 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres15.2 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:triadpass@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:triadpass@localhost:5432/simple_bank?sslmode=disable" -verbose down

.PHONY: postgres createdb dropdb migrateup migratedown