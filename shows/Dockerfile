# Dockerfile for Account MS for MovieHub
FROM golang:alpine as build

LABEL maintainer=“superDeano”

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

FROM alpine:latest

WORKDIR app/

COPY --from=build /app/main .

RUN chmod +x main

EXPOSE 7000

CMD ["./main"]