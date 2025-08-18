FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY . .

WORKDIR /app/services/trip-service

RUN CGOOS=linux CGO_ENABLED=0 go build -o trip-service ./cmd/server/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/services/trip-service/trip-service .

CMD [ "./trip-service" ]
