DB_HOST = db
DB_PORT = 5432
DB_USER = user
DB_PASSWORD = password
DB_NAME = mydb

.PHONY: generate start clean user-generator

generate:
	docker-compose run --rm oapi-codegen oapi-codegen -generate types,server -package api -o /app/api/gen.go /app/spec.yaml

start:
	docker-compose up -d db
	docker-compose up http-service

flyway:
	docker-compose up -d flyway

clean:
	docker-compose down
	rm -rf ./client

user-generator:
	docker-compose up user-generator

clean-generator:
	docker stop catersease-user-generator-1 || true
	docker rm catersease-user-generator-1 || true
	docker rmi catersease-user-generator || true
