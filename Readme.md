
# Majoo Backend Test

Simple Transaction Application

## Technologies

The various packages used are as follows:

- [Gin Gonic](https://github.com/gin-gonic/gin) - HTTP handler framework
- [SQL](https://pkg.go.dev/database/sql) - MySQL database driver
- [CORS](https://github.com/gin-contrib/cors) - Gin Gonic CORS middleware
- [Go Dot ENV](https://github.com/joho/godotenv) - Dot ENV vars loader
- [JWT](https://github.com/dgrijalva/jwt-go) - JWT generator & parser
- [Bcrypt](https://pkg.go.dev/golang.org/x/crypto) - Tools for generate password & check password

## Installation

This project requires [Golang](https://golang.org/) v1.17+ to run.

Clone the repository then change your current directory to the repository.

Install Required Package
```sh
go get
```

Copy the example env file and make the required configuration changes in the .env file

```sh
cp .env.example .env
```

Start the REST server.

```sh
make run
```

## Server

Open the development server by navigating to your server address in your preferred browser.

http://localhost:8080

## Database Migrations & Seeding

Run database migrations up.

```sh
make migrate-up
```

Run database migrations down.

```sh
make migrate-down
```

## API Collections

Import postman collection file on docs folder, to access all the endpoints in this REST API.