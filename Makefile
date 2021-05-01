run:
	go run ./cmd/api -port=3000 -env=development

pg_start:
	brew services start postgresql

pg_stop: 
	brew services stop postgresql
