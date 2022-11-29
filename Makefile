PGX_TEST_CONN_STRING:="host=localhost user=test password=secret dbname=pgx_test sslmode=disable"

test:
	PGX_TEST_CONN_STRING=$(PGX_TEST_CONN_STRING) go test -v -cover -race ./...

test-e2e: db-restart db-probe
	PGX_TEST_CONN_STRING=$(PGX_TEST_CONN_STRING) go test -v ./...

db-stop:
	docker compose stop postgres
	docker compose rm --force postgres

db-start:
	docker compose up --detach postgres

db-restart: db-stop db-start

db-probe:
	docker compose run --rm postgres-probe

example:
	PGX_TEST_CONN_STRING=$(PGX_TEST_CONN_STRING) go run ./example/

.PHONY: \
	test-e2e \
	test \
	up \
	db-start \
	db-restart \
	db-stop \
	db-probe \
	example