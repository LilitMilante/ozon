db:
	docker run --name rates_db -p 8085:5432 -e POSTGRES_PASSWORD=dev -d postgres:15.3-alpine

migrate-new:
	goose -dir ./migrations create $(name) sql

migrate-up:
	goose -dir ./migrations postgres "user=postgres dbname=postgres password=dev host=localhost port=8085 sslmode=disable" up

migrate-down:
	goose -dir ./migrations postgres "user=postgres dbname=postgres password=dev host=localhost port=8085 sslmode=disable" down
