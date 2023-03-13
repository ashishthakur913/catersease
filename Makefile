.PHONY: generate start clean

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
