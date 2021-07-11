FROM golang:1.14.9-alpine
RUN mkdir /build
WORKDIR /build
ADD go.mod main.go /build/
RUN go build
CMD ["./go-http-server"]