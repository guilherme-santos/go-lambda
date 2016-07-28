# Test go project infrastructe with Golang

### WebServer

To run as webserver:

```
go run main.go -webserver
```

Access:
```
curl -X GET http://localhost:8080/
curl -X GET http://localhost:8080/hello/Guilherme
```

### Command Line

To run as a executable (called inside of lambda function)

```
go run main.go -X GET /
go run main.go -X GET /hello/Guilherme
```
