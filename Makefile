.PHONY: start-postgres stop-postgres

start-postgres:
	docker run --name some-postgres -e POSTGRES_PASSWORD=mysecretpassword -p 5432:5432 -d postgres

stop-postgres:
	docker stop some-postgres
	docker rm some-postgres
