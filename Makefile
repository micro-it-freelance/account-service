build:
	docker build -t	micro-it-fl/account-service .

up:
	docker compose up --build

test:
	docker compose -f ./docker-compose.yaml -f ./docker-compose.test.yaml up --build