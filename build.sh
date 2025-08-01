#!/bin/bash
# chmod +x build.sh OR ./build.sh
# To watch logs: tail -f app.log


#!/bin/bash

# Exit immediately if a command exits with a non-zero status
set -e

APP_NAME="ecommerce-ver-2"
CMD_DIR="./cmd/server"
DB_DRIVER="mysql"
DB_DSN="root:1234@tcp(localhost:3306)/go_eco"
MIGRATIONS_DIR="./db/schema"
GOOSE="goose"
TOPIC_NAME="otp-auth-topic"
KAFKA_BROKER="localhost:9192"

echo "🐳 Starting Docker containers..."
docker-compose up -d

echo "🔧 Building application..."
go build -o $APP_NAME $CMD_DIR/main.go

echo "🛠️ Running database migrations..."
$GOOSE -dir $MIGRATIONS_DIR $DB_DRIVER "$DB_DSN" up

echo "🔎 Checking if Kafka topic '$TOPIC_NAME' exists..."
# Check if topic exists using kafka-topics.sh (requires Kafka CLI)
if ! kafka-topics.sh --bootstrap-server $KAFKA_BROKER --list | grep -q "^$TOPIC_NAME$"; then
  echo "📦 Creating Kafka topic '$TOPIC_NAME'..."
  kafka-topics.sh --create \
    --topic $TOPIC_NAME \
    --bootstrap-server $KAFKA_BROKER \
    --partitions 1 \
    --replication-factor 1
else
  echo "✅ Kafka topic '$TOPIC_NAME' already exists."
fi

echo "🚀 Starting application and logging to app.log..."
./$APP_NAME > app.log 2>&1