FROM golang:1.18.1-alpine AS build
RUN mkdir -p /app
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN GOOS=linux go build -o main.out ./src/main.go

FROM alpine:3.15.4 AS production
RUN mkdir -p /app
WORKDIR /app
COPY --from=build /app/main.out .
EXPOSE 8080
CMD /bin/sh -c ./main.out

FROM golang:1.18.1-alpine AS test
RUN apk add build-base
RUN mkdir -p /app
WORKDIR /app/src
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
CMD go test ./src/tests

FROM golang:1.18.1-alpine AS dev
RUN curl -sSfL \
    https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | \
    sh -s -- -b $(go env GOPATH)/bin
RUN mkdir -p /app
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
CMD air
