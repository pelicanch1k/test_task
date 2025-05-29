#!/bin/bash
wait_for_db() {
  echo "Waiting for the database to start..."
  until pg_isready -h postgres -U $POSTGRES_USERNAME -d $POSTGRES_DB; do
    echo "Database unavailable, wait 5 seconds..."
    sleep 5
  done
  echo "The database is up and running!"
}

wait_for_db

echo "Performing Migrations..."
atlas migrate apply \
  --url "postgresql://$POSTGRES_USERNAME:$POSTGRES_PASSWORD@postgres:5432/$POSTGRES_DB?sslmode=disable"