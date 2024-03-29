# build state
FROM golang:1.21.0-alpine3.18 AS builder

WORKDIR /app

COPY . .

RUN go build -o main main.go 

# run state
FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .

EXPOSE 8080

CMD [ "/app/main" ]
