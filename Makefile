docker-build:
	docker build -t place-chat/account-service .

docker-up:
	docker compose up --build