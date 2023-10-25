migrate:
	migrate -source file:internal/db/migrations -database postgres://postgres:postgres@localhost:5432/emdb?sslmode=disable up	