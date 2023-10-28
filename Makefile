build:
	docker compose build

up:
	docker compose up

upd:
	docker compose up -d

updb:
	docker compose up -d --build

down:
	docker compose down

db:
	docker compose exec database psql -h localhost -U postgres postgres

web:
	docker compose exec server bash

lg:
	docker compose logs -f
