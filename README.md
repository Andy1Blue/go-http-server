# GO HTTP Server

## Description

GO HTTP server

## How to start

```bash
$ go mod init go-http-server # generate mod file
$ go build # generate build
$ go run main.go # run app
```

## Docker

```bash
$ docker build -t go-http-server .
$ docker run -p 8080:8080 -d go-http-server
```
