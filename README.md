# Go Socket Chat

Socket chat app built with Golang

Framework: Gin <br />
Architecture: Hexagonal <br />
Database: Postgres

## How to run?

1. Start the docker compose file

  ```sh
  docker compose up -d
  ```

2. Build the go application

  ```sh
  go build cmd/main.go -o bin/server
  ```

3. Run the go application

  ```sh
  ./bin/server
  ```

Or you could also utilize the `Makefile` file.
