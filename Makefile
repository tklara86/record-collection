include .env

server:
	go run ./cmd/web

webpack:
	npm run watch

# migrate create -ext sql -dir ./migrations -seq init_schema
migrateup:
	migrate -path migrations -database "$(DB_TYPE)://$(DB_USER):$(DB_PASSWORD)@tcp($(DB_HOST))/$(DB_NAME)?parseTime=true" up

migratedown:
	migrate -path migrations -database "$(DB_TYPE)://$(DB_USER):$(DB_PASSWORD)@tcp($(DB_HOST))/$(DB_NAME)?parseTime=true" down
