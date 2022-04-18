# GoWebApi
This project was developed to be the starting point for the development of a simple bit web api project with go programming.

### Frameworks and libraries used

- GoFiber 2.29.0
- Gorm v1.23.4
- Viper v1.11.0
- MS SqlServer

Build:
```sh
go build main.go
```

Serve on local server:
```sh
go run main.go
```

Environments: (.env file in root folder)

```sh
DB_URL=<database url>
DB_NAME=<database name>
DB_USER=<database user>
DB_PASS=<database password>
```