FROM golang:1.24 AS builder

WORKDIR /app

COPY . .

WORKDIR /app/services/api-gateway

RUN CGOOS=linux CGO_ENABLED=0 go build -o api-gateway .


FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/services/api-gateway/api-gateway .

CMD [ "./api-gateway" ]
