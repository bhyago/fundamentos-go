createmigration:
	migrate create -ext=sql -dir=sql/migrations -seq init
migrateup:
	migrate -path sql/migrations -database "postgres://user:password@localhost:5432/dbname?sslmode=disable" up
migratedown:
	migrate -path sql/migrations -database "postgres://user:password@localhost:5432/dbname?sslmode=disable" down
.PHONY: migrate