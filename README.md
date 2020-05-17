# to-do-app

Back-end of an application that allows you to **add, delete and update items to a to-do list**. This is a simple project that uses Go as a primary language and PostgreSQL as a database.

## How to set up your environment

1. [Install Go](https://golang.org/doc/install) if not already present in your system
2. Install [PostgreSQL](https://www.postgresql.org/)
3. Clone the repository using SSH or HTTPS
4. (Optional) Install [Postman](https://www.postman.com/)

## How to run the application

1. Launch your database server. For instance, run `pg_ctl -D /usr/local/var/postgres start` on MacOS.
2. Create a new database. You can keep the default user `postgres`.
3. Set environment variables: `HOST`, `DB_PORT`, `USERNAME`, `DB`. Usually, you will run the application locally, so the host should be `localhost` and the port should be `5432`.
4. Launch the web server by running `go run main.go` at the project's directory level.

## How to test your set up

1. Open Postman or use another API client like cURL
2. Create a task using the corresponding API endpoint. If you properly set up your environment, you shoud get the task created as a response.
