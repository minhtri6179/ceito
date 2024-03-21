DB_URL=postgresql://root:secret@localhost:5432/toeic?sslmode=disable
DB_on_cloud=secret

postgres:
	docker run --name dbceito --network ceito -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it dbceito createdb --username=root --owner=root toeic

dropdb:
	docker exec -it dbceito dropdb toeic

migrateup:
	migrate -path db/migrations -database "${DB_URL}" -verbose up

migrateupcloud:
	migrate -path db/migrations -database "${DB_on_cloud}" -verbose up

migratedown:
	migrate -path db/migrations -database "${DB_URL}" -verbose down

sqlc:
	sqlc generate

createdata:
	cd scripts && python3 create_db.py

.PHONY: postgres createdb dropdb migrateup migrateupcloud migratedown sqlc createdata