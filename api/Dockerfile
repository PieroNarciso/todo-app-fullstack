FROM golang:1.18.1 AS build
RUN mkdir -p /app
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o main.out ./src/main.go

FROM ubuntu:20.04 AS production
RUN apt-get update -y
RUN mkdir -p /app
WORKDIR /app
COPY --from=build /app/main.out .
EXPOSE 8080
CMD /bin/bash -c ./main.out

FROM golang:1.18.1 AS test
RUN mkdir -p /app
WORKDIR /app/src
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
CMD go test ./src/tests

FROM golang:1.18.1 AS dev
RUN mkdir -p /app
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN curl -sSfL \
    https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | \
    sh -s -- -b $(go env GOPATH)/bin
CMD air
