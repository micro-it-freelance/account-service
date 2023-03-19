build:
	docker build -t place-chat/account-service .

up:
	docker compose up --build

test:
	docker compose -f ./docker-compose.yaml -f ./docker-compose.test.yaml up --build