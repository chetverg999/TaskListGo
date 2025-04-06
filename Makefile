.SILENT:

swagger:
	swag init -g cmd/main.go

start: swagger
	docker compose up -d
