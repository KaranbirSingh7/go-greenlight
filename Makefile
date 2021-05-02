DATABASE_URL=postgres://greenlight:password@localhost/greenlight?sslmode=disable

run:
	go run ./cmd/api -port=4000 -env=development

migrate.up:
	migrate -path=./migrations -database "$(DATABASE_URL)" up


migrate.down:
	migrate -path=./migrations -database "$(DATABASE_URL)" down
	
db.start:
	brew services start postgresql
	
db.stop:
	brew services stop postgresql
