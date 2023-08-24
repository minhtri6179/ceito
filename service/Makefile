DB_URL=postgresql://root:secret@localhost:5432/toeic?sslmode=disable

postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root toeic

dropdb:
	docker exec -it postgres12 dropdb toeic

migrateup:
	migrate -path db/migration -database "${DB_URL}" -verbose up

migratedown:
	migrate -path db/migration -database "${DB_URL}" -verbose down

sqlc:
	sqlc generate

.PHONY: postgres createdb dropdb migrateup migratedown sqlc