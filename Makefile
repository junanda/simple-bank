include .$(PWD)/.env

postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=$(POSTGRES_USER) -e POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) -d postgres:15.2-alpine

createdb:
	docker exec -it postgres15.2 createdb --username=$(POSTGRES_USER) --owner=$(POSTGRES_OWNER) $(PSQL_DB)

dropdb:
	docker exec -it postgres15.2 dropdb $(PSQL_DB)

migrateup:
	migrate -path db/migration -database "postgresql://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(HOST_PSQL)/$(PSQL_DB)?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(HOST_PSQL)/$(PSQL_DB)?sslmode=disable" -verbose down

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown test