createmigration:
	migrate create -ext=sql -dir=sql/migrations -seq init

migrate:
	migrate -path=sql/migrations -database "mysql://root:@tcp(localhost:3306)/go_courses" -verbose up

migratedown:
	migrate -path=sql/migrations -database "mysql://root:@tcp(localhost:3306)/go_courses" -verbose down

.PHONY: migrate migratedown createmigration
