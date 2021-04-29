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
