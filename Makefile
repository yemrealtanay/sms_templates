postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=yemretest -e POSTGRES_PASSWORD=123456789 -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=yemretest --owner=yemretest sms_templates

dropdb:
	docker exec -it postgres12 dropdb sms_templates

migrateup:
	migrate -path ./db/migrations -database "postgresql://yemretest:123456789@0.0.0.0:5432/sms_templates?sslmode=disable" -verbose up

migratedown:
	migrate -path ./db/migrations -database "postgresql://yemretest:123456789@0.0.0.0:5432/sms_templates?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test