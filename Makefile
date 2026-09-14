include .env
export

service-deploy:
	@docker compose up -d --build

service-undeploy:
	@docker compose down 

db-up:
	@docker compose up -d postgres

run:
	@go run main.go