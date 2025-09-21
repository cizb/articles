# 📘 Articles

REST API сервис для заметок на Go (Gin + GORM + PostgreSQL). Поддерживает создание и получение статей.

---

### Вариант A — Docker Compose

```bash
docker compose up -d --build
```

После запуска:

- Сервис доступен на `http://localhost:3000`
- PostgreSQL доступен на порту `5432`

### Вариант B — Локальный запуск

1. Поднять PostgreSQL (через docker или локально).
2. Создать файл `.env` в корне проекта:

```env
APP_ENVIRONMENT=dev

APP_HTTP_SERVER_PORT=3000
APP_HTTP_SERVER_BASE_URL=/api/v1

APP_DB_HOST=localhost
APP_DB_PORT=5432
APP_DB_USER=postgres
APP_DB_PASSWORD=postgres
APP_DB_DATABASE=postgres
```

3. Запустить приложение:

```bash
go run cmd/main.go
```

---

## 🔗 API

### Получить статью по ID

**GET** `/api/v1/article/{id}`

Пример:

```http
GET http://localhost:3000/api/v1/article/1
```

Пример ответа:

```json
{
	"id": 1,
	"title": "Оптимизация сервиса статей",
	"tags": ["backend", "gorm", "postgres"],
	"content": "Внедрил AutoMigrate для таблицы articles. Настроил docker-compose с postgres и сервисом, проверил подключение.",
	"created_at": "2025-09-21T18:50:10Z",
	"updated_at": "2025-09-21T18:50:10Z"
}
```

### Создать статью

**POST** `/api/v1/article`

Пример запроса:

```http
POST http://localhost:3000/api/v1/article
Content-Type: application/json

{
  "title": "Оптимизация сервиса статей",
  "tags": ["backend", "gorm", "postgres"],
  "content": "Внедрил AutoMigrate для таблицы articles. Настроил docker-compose с postgres и сервисом, проверил подключение."
}
```

Пример ответа:

```json
{
	"id": 1,
	"title": "Оптимизация сервиса статей",
	"tags": ["backend", "gorm", "postgres"],
	"content": "Внедрил AutoMigrate для таблицы articles. Настроил docker-compose с postgres и сервисом, проверил подключение.",
	"created_at": "2025-09-21T18:50:10Z",
	"updated_at": "2025-09-21T18:50:10Z"
}
```

---

## 🛠️ Технологии

- Gin — HTTP роутер
- GORM — ORM для PostgreSQL
- Clean Architecture (handler / usecase / repository / schema)
- Docker & Docker Compose
- zerolog — логирование

---

## 📂 Структура проекта (пример)

```
.
├── cmd/                  # entrypoint (main.go)
├── internal/
│   ├── handler/          # HTTP handlers (Gin)
│   ├── usecase/          # логика
│   ├── storage/          # repository (GORM)
│   └── schema/           # схемы (GORM)
├── pkg/                  # вспомогательные пакеты
├── Dockerfile
├── docker-compose.yml
├── go.mod
└── README.md
```

---

## ⚙️ Заметки

- На старте выполняется `AutoMigrate(&schema.Article{})` для автосоздания/обновления таблицы.
