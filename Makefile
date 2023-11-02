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

sh:
	docker compose exec server sh

logs:
	docker compose logs -f

test:
	go test -v -cover ./...

clean: down
	rm -f gotemplate
	docker system prune -f
	docker volume prune -f
