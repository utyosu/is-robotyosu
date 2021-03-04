fmt:
	go fmt ./...

tidy:
	go mod tidy

build-local: fmt tidy
	go build -o bin/rfe-local -tags="local"

build-production:
	GOOS=linux GOARCH=amd64 go build -o bin/rfe-production -tags="production"

run-local: build-local
	sleep 0.5
	./bin/rfe-local

reset-db-local:
	sudo mysql -e " \
		DROP DATABASE IF EXISTS ${RFE_DATABASE_NAME_LOCAL}; \
		CREATE DATABASE ${RFE_DATABASE_NAME_LOCAL}; \
	"
	@make migrate-db-up-local

migrate-db-up-local:
	migrate -path db/migrations -database "mysql://${RFE_DATABASE_USER_LOCAL}:@tcp(${RFE_DATABASE_HOST_LOCAL}:${RFE_DATABASE_PORT_LOCAL})/${RFE_DATABASE_NAME_LOCAL}" up
	mysqldump -u ${RFE_DATABASE_USER_LOCAL} -h ${RFE_DATABASE_HOST_LOCAL} -P ${RFE_DATABASE_PORT_LOCAL} ${RFE_DATABASE_NAME_LOCAL} -d --skip-comments --no-tablespaces > db/schema.sql

migrate-db-down-local:
	migrate -path db/migrations -database "mysql://${RFE_DATABASE_USER_LOCAL}:@tcp(${RFE_DATABASE_HOST_LOCAL}:${RFE_DATABASE_PORT_LOCAL})/${RFE_DATABASE_NAME_LOCAL}" down 1
	mysqldump -u ${RFE_DATABASE_USER_LOCAL} -h ${RFE_DATABASE_HOST_LOCAL} -P ${RFE_DATABASE_PORT_LOCAL} ${RFE_DATABASE_NAME_LOCAL} -d --skip-comments --no-tablespaces > db/schema.sql
