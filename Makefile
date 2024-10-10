build:
	go build bin

run:
	go run main.go

goose-up:
	goose postgres "user=postgres password=postgres dbname=gofiber_boilerplate sslmode=disable" -dir db/migrations up

goose-reset:
	goose postgres "user=postgres password=postgres dbname=gofiber_boilerplate sslmode=disable" -dir db/migrations reset

goose-fix:
	goose postgres "user=postgres password=postgres dbname=gofiber_boilerplate sslmode=disable" -dir db/migrations fix

goose-redo:
	goose postgres "user=postgres password=postgres dbname=gofiber_boilerplate sslmode=disable" -dir db/migrations redo
