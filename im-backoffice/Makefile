migrationup:
	migrate -path db/migration -database "postgresql://postgres:password@localhost:5432/im-db?sslmode=disable" -verbose up

migrationdown:
	migrate -path db/migration -database "postgresql://postgres:password@localhost:5432/im-db?sslmode=disable" -verbose down

migrationnew:
	migrate create -ext sql -dir db/migration -seq change_db

sqlc:
	sqlc generate

server:
	go run server.go

swagger:
	swag init -g server.go

test:
	go test -v -cover ./...