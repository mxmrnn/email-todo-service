# ====================== Builder Stage ======================
FROM golang:1.26-alpine AS builder

WORKDIR /app

# Копируем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем весь исходный код
COPY . .

# Собираем приложение
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /email-service ./cmd/api

# ====================== Runtime Stage ======================
FROM alpine:latest

WORKDIR /app

# Копируем бинарник из builder
COPY --from=builder /email-service .

# Делаем исполняемым
RUN chmod +x ./email-service

# Запускаем
CMD ["./email-service"]