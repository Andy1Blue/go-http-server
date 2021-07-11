FROM golang:1.14.9-alpine
RUN mkdir /build
WORKDIR /build
ADD go.mod main.go /build/
RUN go build
EXPOSE 3000
CMD ["./go-http-server"]