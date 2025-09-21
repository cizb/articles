FROM golang:1.24 AS builder

WORKDIR /opt/app

# Настройка кэша go
RUN go env -w GOCACHE=/go-cache
RUN go env -w GOMODCACHE=/gomod-cache

COPY go.mod go.sum ./

# Кэшируем зависимости
RUN --mount=type=cache,target=/gomod-cache \
    go mod download

# Копируем весь проект
COPY . .

# Собираем бинарь (main.go в cmd/, поправь путь если у тебя другой)
RUN --mount=type=cache,target=/go-cache \
    --mount=type=cache,target=/gomod-cache \
    CGO_ENABLED=0 GOOS=linux go build -o /go/bin/service ./cmd/main.go

# ========== Runtime stage ==========
FROM alpine:3.20

# Устанавливаем только нужное
RUN apk add --no-cache ca-certificates tzdata bash curl

# Создаём пользователя без root
RUN addgroup -S app && adduser -S -g app app
USER app

WORKDIR /app

# Копируем бинарь
COPY --from=builder /go/bin/service .

ENTRYPOINT ["/app/service"]