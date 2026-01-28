#!/bin/bash

# Get the directory where the script is located
SCRIPT_DIR=$(dirname "$(readlink -f "$0")")
# Go one level up to find the project root
PROJECT_ROOT=$(dirname "$SCRIPT_DIR")

# 1. Load environment variables from the root .env
if [ -f "$PROJECT_ROOT/.env" ]; then
    export $(grep -v '^#' "$PROJECT_ROOT/.env" | xargs)
fi

COMMAND=$1
VALUE=$2

# Use $PROJECT_ROOT/migrations to ensure paths are always correct
MIGRATIONS_PATH="$PROJECT_ROOT/migrations"

case $COMMAND in
    "up")
        migrate -path "$MIGRATIONS_PATH" -database "$DATABASE_URL" up
        ;;
    "down")
        COUNT=${VALUE:-1}
        read -p "Rolling back $COUNT migration(s). Continue? [y/N]: " confirm
        if [[ $confirm == [yY] ]]; then
            migrate -path "$MIGRATIONS_PATH" -database "$DATABASE_URL" down "$COUNT"
        fi
        ;;
    "create")
        migrate create -ext sql -dir "$MIGRATIONS_PATH" -seq "$VALUE"
        ;;
    "force")
        migrate -path "$MIGRATIONS_PATH" -database "$DATABASE_URL" force "$VALUE"
        ;;
    *)
        echo "Usage: ./scripts/migrate.sh {up|down|create|force} [value]"
        ;;
esac