APP_NAME = .\cmd\server\main.go

run:
	set PROFILE=local&&go run ${APP_NAME}

up-local:
	set PROFILE=local&& docker-compose up -d mydbdev redis-stack&& go run ${APP_NAME}

up-dev:
	set PROFILE=dev&& docker-compose up -d mydbdev redis-stack&& go run ${APP_NAME}

up-prod:
	docker-compose --env-file .env_prod up -d --build

down:
	docker-compose down