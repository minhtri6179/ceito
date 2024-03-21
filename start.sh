#!/bin/sh

set -e 
echo "Waiting for database to be ready..."
/app/migrate -path /app/migrations -database "$DB_SOURCE" -verbose up

exec "$@"
