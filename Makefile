run:
	go run ./cmd/api

pg_start:
	brew services start postgresql

pg_stop: 
	brew services stop postgresql
