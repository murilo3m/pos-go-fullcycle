DB_NAME = orders
DB_USER = root
DB_PASSWORD = root
DB_HOST = db
DB_PORT = 3306
migrate:
	migrate -path=configs/sql -database "mysql://$(DB_USER):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)" -verbose up
migratedown:
	migrate -path=configs/sql -database "mysql://$(DB_USER):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)" -verbose down
.PHONY: migrate migratedown