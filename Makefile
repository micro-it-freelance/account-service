SERVICE_NAME := account-service

up:
	docker compose up --build

test:
	docker compose -f ./docker-compose.yaml -f ./docker-compose.test.yaml up $(SERVICE_NAME) --build --exit-code-from $(SERVICE_NAME)