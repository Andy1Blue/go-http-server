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
$ docker run -p 8000:8000 -d go-http-server
```

## Heroku

```bash
heroku login
heroku create
git init
git add . && git commit -m "Deploy to Heroku"
heroku stack:set container -a app-name
git push heroku master
```
