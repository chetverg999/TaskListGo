.SILENT:

swagger:
	swag init -g cmd/main.go

start: swagger
	go build -o TestTask cmd/main.go
	go run cmd/main.go