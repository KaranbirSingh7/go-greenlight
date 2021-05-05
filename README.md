# go-greenlight

Repository to keep track of my progress for book https://lets-go-further.alexedwards.net/

### Repository Structure

```bash
`bin` directory will contain our compiled application binaries, ready for deployment to a production server.

`cmd/api` directory will contain the application-specific code for our Greenlight API application. This includes the code for running the server, reading and writing HTTP requests, and managing authentication.

`internal` directory will contain various ancillary packages used by our API. It will contain the code for interacting with our database, doing data validation, sending emails and so on. Basically, any code which isn’t application-specific and can potentially be reused will live in here. Our Go code under cmd/api will import the packages in the internal directory (but never the other way around).

`migrations` directory will contain the SQL migration files for our database.

`remote` directory will contain the configuration files and setup scripts for our production server.

`go.mod` file will declare our project dependencies, versions and module path.

`Makefile` will contain recipes for automating common administrative tasks — like auditing our Go
code, building binaries and executing database migrations.
```

### Things I Learned as I go through the book

1. There are two common approaches to versioning your APIs:
   By prefixing all URLs with the API version, like /v1/healthcheck or /v2/healthcheck.
   By using custom Accept and Content-Type headers on requests and responses to convey the API version, like Accept: application/vnd.greenlight-v1.
   Preffered choice is: myapi/v1/healthcheck

1. Every path in your API need a similar go file that handle those routes. Example:
   /v1/health -> /cmd/api/health.go
   /v1/movies -> /cmd/api/movies.go

1. Any helpers method should not log anything except returning error. All error are logged in specific routes file and there are two loggers. One app.logger for application logging and other is http.error for sending error back to client/user

1. Error logging in Go can be repetitive, it is better to turn common logging into a function into its own file: errors.go

1. Any errors that can be logically produced, you should use panic. Example: json.InvalidUnmarshalError.
   The Go By Example page on panics summarizes all of this quite nicely:
   A panic typically means something went unexpectedly wrong. Mostly we use it to fail fast on errors that shouldn’t occur during normal operation and that we aren’t prepared to handle gracefully.

1. If json data is not proper, send back 422 Unprocessable Entity

1. Want to tune PostgreSQL `https://pgtune.leopard.in.ua/#/`

1. Postgres has hard limit of 100 total connections, can be changes by modifying postgresql.conf

1. For database migrations (schema change, column updates) use `golang-migrate`

1. Use a model/database layer for all operations related to DB. See `./data/models/movies.go`

1. Use of context. Timeout for context starts when it is created. Any time between created and context used will be counted.

### Which code can be reused in other projects?

1. `cmd/api/errors.go`
1. `cmd/api/helpers.go` `readJSON()` function
1.
