run:
	go run cmd/main.go --profile=dev

build:
	go build -v -o bin/guardian cmd/main.go

build_db:
	sqlite3 guardian.db < ./db/sqlite/create_tables.sql
	sqlite3 guardian.db < ./db/sqlite/insert_data.sql

drop_tables:
	sqlite3 guardian.db < ./db/sqlite/drop_tables.sql