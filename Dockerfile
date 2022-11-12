# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./
COPY Limiter/ ./Limiter/
COPY omdbquery/  ./omdbquery/
COPY readFile/ ./readFile/
COPY util/ ./util/
COPY az.csv ./

RUN go build -o /docker-gs-ping

EXPOSE 8080

CMD [ "/docker-gs-ping" ]
