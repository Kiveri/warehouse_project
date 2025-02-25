# Этап 1: Сборка приложения (builder)


# Используется официальный образ Go версии 1.24 на базе Alpine Linux (минималистичный образ),
# AS builder указывает, что это первый этап сборки, и результаты этого этапа будут использоваться в следующем.
FROM golang:1.24-alpine AS builder

# Устанавливает утилиту git в контейнер. Она необходима для загрузки зависимостей Go, которые могут быть указаны
# через Git-репозитории. Флаг --no-cache предотвращает кэширование пакетов, чтобы уменьшить размер конечного образа.
RUN apk add --no-cache git

# Устанавливает рабочую директорию внутри контейнера на /app. Все последующие команды будут выполняться в этой директории.
WORKDIR /app

# Копирует файлы go.mod и go.sum из локальной директории в рабочую директорию контейнера (/app).
# Эти файлы содержат информацию о зависимостях проекта.
COPY go.mod go.sum ./

# Скачивает все зависимости, указанные в go.mod и go.sum, используя команду
RUN go mod download

# Устанавливает утилиту goose (для миграций) в систему. Она будет доступна в /go/bin/goose
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

# Копирует все файлы из локальной директории в рабочую директорию контейнера (/app).
COPY . .

# Переходит в директорию cmd. Собирает бинарный файл приложения с именем main и сохраняет его в /app/main.
RUN cd cmd && go build -o /app/main .


# Этап 2: Финальный образ (запуск приложения)


# Используется минимальный образ Alpine Linux для финального контейнера. Это уменьшает размер итогового образа.
FROM alpine:latest

# Устанавливает bash и postgresql-client (клиент PostgreSQL) в финальный образ.
# bash используется для выполнения скриптов, а postgresql-client содержит утилиту pg_isready,
# которая проверяет доступность PostgreSQL.
RUN apk add --no-cache bash postgresql-client

# Устанавливает рабочую директорию внутри контейнера на /root/.
WORKDIR /root/

# Копирует собранный бинарный файл main из этапа builder в текущую рабочую директорию (/root/).
COPY --from=builder /app/main .

# Копирует утилиту goose из этапа builder в /usr/local/bin/goose, чтобы она была доступна в финальном контейнере.
COPY --from=builder /go/bin/goose /usr/local/bin/goose

# Копирует папку с миграциями (db/migrations) из локальной директории в /root/db/migrations внутри контейнера.
COPY db/migrations /root/db/migrations

# Копирует файл .env (с переменными окружения) из локальной директории в /root/.env внутри контейнера.
COPY .env /root/.env

# Указывает, что приложение будет использовать порт 8080. Это не открывает порт автоматически,
# но документирует его использование.
EXPOSE 8080

# Запускает приложение с помощью Bash.

# set -o allexport; source /root/.env; set +o allexport:
    # Экспортирует все переменные из файла .env в окружение.

# until pg_isready -h \"$DB_HOST\" -p \"$DB_PORT\"; do echo 'Waiting for PostgreSQL...'; sleep 2; done:
    # Ожидает, пока PostgreSQL станет доступным, проверяя его с помощью pg_isready.
    # Если база данных недоступна, выводит сообщение и ждёт 2 секунды перед повторной проверкой.

# goose -dir ./db/migrations postgres \"host=$DB_HOST user=$POSTGRES_USER password=$POSTGRES_PASSWORD dbname=$POSTGRES_DB sslmode=$POSTGRES_SSL_MODE\" up:
    # Запускает миграции с помощью goose, используя параметры подключения к PostgreSQL из переменных окружения.

# ./main:

# Запускает собранное приложение.
CMD ["bash", "-c", "et -o allexport; source /root/.env; set +o allexport; until pg_isready -h \"$DB_HOST\" -p \"$DB_PORT\"; do echo 'Waiting for PostgreSQL...'; sleep 2; done && goose -dir ./db/migrations postgres \"host=$DB_HOST user=$POSTGRES_USER password=$POSTGRES_PASSWORD dbname=$POSTGRES_DB sslmode=$POSTGRES_SSL_MODE\" up && ./main"]