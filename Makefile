db:
	docker run --name sellers_db -p 5151:5432 -e POSTGRES_PASSWORD=dev -d postgres:15.3-alpine

db-down:
	docker rm -f sellers_db

migrate-new:
	goose -dir ./migrations create $(name) sql

migrate-up:
	goose -dir ./migrations postgres "user=postgres dbname=postgres password=dev host=localhost port=5151 sslmode=disable" up

migrate-down:
	goose -dir ./migrations postgres "user=postgres dbname=postgres password=dev host=localhost port=5151 sslmode=disable" down
