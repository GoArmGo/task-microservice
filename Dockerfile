# 1. Используем официальный образ Go
FROM golang:1.21-alpine

# 2. Установка системных зависимостей
RUN apk add --no-cache ca-certificates

# 3. Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# 4. Копируем go.mod и go.sum, чтобы сначала скачать зависимости
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# 5. Копируем весь остальной проект
COPY . .

# 6. Сборка бинарника
RUN go build -o app ./cmd/main.go

# 7. Открываем порт, который слушает Gin
EXPOSE 8080

# 8. Запускаем бинарник
CMD ["./app"]