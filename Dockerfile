# Usar a imagem oficial do Go como imagem de construção
FROM golang:1.22.1-alpine as builder

WORKDIR /app
COPY . .
COPY .env .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .

# Copiar o arquivo .env do construtor
COPY --from=builder /app/.env .

CMD ["./main"]
