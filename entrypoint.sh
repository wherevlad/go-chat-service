#!/bin/bash
source /root/.env

# wait for the postgres to be ready
while ! nc -z postgres ${POSTGRES_PORT_EXTERNAL}; do
  >&2 echo "postgres is unavailable, wait..."
  sleep 2
done

# Run goose and check if the command was successful
if ! goose -dir "${MIGRATION_DIR}" postgres "${PG_DSN}" up -v; then
    echo "Migration failed"
    exit 1  # Exit with a non-zero status to indicate error
fi

# run server
./chat_server