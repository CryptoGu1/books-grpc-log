FROM golang:1.25-alpine AS builder

WORKDIR /app

# Копируем go.mod и go.sum (они теперь лежат прямо тут)
COPY go.mod go.sum ./
RUN go mod download

# Копируем весь код логгера
COPY . .

# Собираем. Путь ./cmd/main.go (проверь, что main там)
# Называем файл "logger-service"
RUN CGO_ENABLED=0 GOOS=linux go build -o /logger-service ./cmd/main.go

# 2. Финальный образ
FROM alpine:latest

WORKDIR /root/

# Копируем бинарник
COPY --from=builder /logger-service .

# Копируем .env (если он есть в корне логгера)
COPY --from=builder /app/.env .

EXPOSE 9000

CMD ["./logger-service"]