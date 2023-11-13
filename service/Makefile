DB_URL=postgresql://root:secret@localhost:5432/toeic?sslmode=disable
# DB on cloud: DB_URL=postgres://root:YCdFlv6VUDoTr1rLotzWzAs7k5kxzzMr@dpg-cl57p9182rpc73fqeld0-a.singapore-postgres.render.com/toeic
postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root toeic

dropdb:
	docker exec -it postgres12 dropdb toeic

migrateup:
	migrate -path db/migrations -database "${DB_URL}" -verbose up

migratedown:
	migrate -path db/migrations -database "${DB_URL}" -verbose down

sqlc:
	sqlc generate

createdata:
	cd scripts && python3 create_db.py

.PHONY: postgres createdb dropdb migrateup migratedown sqlc createdata