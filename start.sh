#!/bin/sh

set -e 

/app/migrate -path /app/migrations -database "$DB_SOURCE" -verbose up

exec "$@"
