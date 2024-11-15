.PHONY: test migrate migrate-version rollback migration build-run-dependency continue-run-dependency stop-dependency compile run

include .env

URL="postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}"
MODULE="${DB_NAME}"

migrate:
	docker run -v "$(shell pwd)/db/migrations/$(MODULE)":"/migrations" --network host migrate/migrate:v4.18.1 -path /migrations -database "$(URL)?sslmode=disable" -verbose up

migrate-version:
	docker run -v "$(shell pwd)/db/migrations/$(MODULE)":"/migrations" --network host migrate/migrate:v4.18.1 -path /migrations -database "$(URL)?sslmode=disable" -verbose version

rollback:
	docker run -v "$(shell pwd)/db/migrations/$(MODULE)":"/migrations" --network host migrate/migrate:v4.18.1 -path /migrations -database "$(URL)?sslmode=disable" -verbose down 1

test:
	go test -count=1 -race -v ./...

build-run-dependency:
	docker-compose -f deployment/docker-compose.yml up --detach

continue-run-dependency:
	docker-compose -f deployment/docker-compose.yml start

stop-dependency:
	docker-compose -f deployment/docker-compose.yml stop

compile:
	go mod tidy && \
	go mod vendor && \
	go build -o deployment/server cmd/api/main.go

run:
	./deployment/server