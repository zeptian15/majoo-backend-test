# Include ENV
include .env

# Install All Required Package
install:
	go get
	
# Run Locally
run:
	go run main.go

# Create New Migration
migration:
	migrate create -ext sql -dir db/migrations -seq $(name)

# Up Migration
migrate-up:
	migrate -path db/migrations -database "mysql://${MYSQL_USERNAME}:${MYSQL_PASSWORD}@tcp(${MYSQL_HOST}:${MYSQL_PORT})/${MYSQL_DBNAME}" -verbose up

# Down Migration
migrate-down:
	migrate -path db/migrations -database "mysql://${MYSQL_USERNAME}:${MYSQL_PASSWORD}@tcp(${MYSQL_HOST}:${MYSQL_PORT})/${MYSQL_DBNAME}" -verbose down