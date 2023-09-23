# build state
FROM golang:1.21.0-alpine3.18 AS builder

WORKDIR /app

COPY . .

RUN go build -o main main.go 

RUN apk add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.16.2/migrate.linux-arm64.tar.gz | tar xvz

# run state
FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/migrate ./migrate
COPY app.env .
COPY start.sh .
COPY wait-for.sh .
COPY db/migrations ./migrations

EXPOSE 8080

CMD [ "/app/main" ]
ENTRYPOINT [ "/app/start.sh" ]
