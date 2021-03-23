

migrate_up:
	migrate -path ./migrations -database postgres://local:local@localhost:5432/uber?sslmode=disable up

migrate_down:
	migrate -path ./migrations -database postgres://local:local@localhost:5432/uber?sslmode=disable down