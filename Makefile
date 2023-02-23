.PHONY: start start-with-docker test gen-swagger

start:
	go run main.go

start-with-docker:
	docker-compose up

test:
	go test ./internal/... -p 1 -count=1 -timeout 5s

gen-swagger:
	# 1. Génération des fichiers swagger
	@swag init -o ./deployments/swagger