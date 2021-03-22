

migrate:
	migrate -path ./migrations -database postgres://local:local@localhost:5432/uber?sslmode=disable up