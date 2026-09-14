include .env
export

service-deploy:
	@docker compose up -d hr_system

service-undeploy:
	@docker compose down hr_system