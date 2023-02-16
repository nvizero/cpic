DB_URL=postgresql://root:root@localhost:5432/sex51?sslmode=disable

postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d postgres:14-alpine

mysql:
	docker run --name mysql57 -p 3319:3306 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d mysql:5.7

createdb:
	docker exec -it postgres createdb --username=root --owner=root sex51

dropdb:
	docker exec -it postgres dropdb  --owner=root sex51

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination mockdb/store.go sex51/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server mock
