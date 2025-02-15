postgres:
	docker run --name db -p 5432:5432 -e POSTGRES_PASSWORD=postgres -d postgres:12-alpine
createdb:
	docker exec -it db createdb --username=postgres --owner=postgres simple_bank
dropdb:
	docker exec -it db dropdb -U postgres simple_bank
migrateup:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/simple_bank?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/simple_bank?sslmode=disable" -verbose down
sqlc:
	sqlc generate
.PHONY: postgres createdb dropdb migrateup migratedown sqlc