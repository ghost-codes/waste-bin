
migrateup:
	migrate -path db/migrations -database "postgresql://postgres:password@localhost:5432/to_did_db?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migrations -database "postgresql://postgres:password@localhost:5432/to_did_db?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migrations -database "postgresql://postgres:password@localhost:5432/to_did_db?sslmode=disable" -verbose down
	
migratedown1:
	migrate -path db/migrations -database "postgresql://postgres:password@localhost:5432/to_did_db?sslmode=disable" -verbose down 1

createMigration:
	migrate create -ext sql -dir db/migrations -seq $(name)

test:
	go test -v -cover ./...

server:
	go run main.go

proto:
	rm -f pb/*.go
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
    proto/*.proto

.PHONY: migrateup migratedown migrateup1 migratedown1 sqlc test server proto