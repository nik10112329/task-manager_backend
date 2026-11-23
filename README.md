# Task Manager — Backend (Go)

REST API для мобильного таск-менеджера ([frontend на Flutter](https://github.com/nik10112329/task_manager)). Первый пет-проект на Go.

> **Статус:** проект в активной разработке. Реализованы подключение к PostgreSQL, конфигурация, логирование, регистрация пользователей и схема БД (миграции). JWT-авторизация, OAuth-линковка аккаунтов и остальные CRUD-эндпоинты — в процессе.

## Стек

- **Go 1.25**
- **PostgreSQL** (драйвер `pgx/v5`)
- **julienschmidt/httprouter** — роутер
- **golang-jwt/v5** — JWT (в процессе)
- **cleanenv + godotenv** — конфигурация
- **logrus** — логирование
- **google/uuid** — UUID
- **goose** — миграции (CLI)

## Структура

```
cmd/app/          — точка входа
internal/
  auth/           — авторизация (в процессе)
  user/           — пользователи (handler / service / storage)
  config/         — загрузка конфигурации
  apperror/       — обработка ошибок
  transport/rest/ — роутер
pkg/
  client/postgresql/ — клиент PostgreSQL
  logging/        — логирование
migrations/       — SQL-миграции (goose)
```

## Как запустить

1. Установи Go 1.25+.
2. Подними PostgreSQL (локально или, например, [Neon](https://neon.tech)).
3. Скопируй шаблон окружения и заполни его:

```bash
cp .env.example .env
```

```
ENV=local
BIND_IP=127.0.0.1
PORT=10000

POSTGRES_HOST=...
POSTGRES_PORT=5432
POSTGRES_USER=...
POSTGRES_PASSWORD=...
POSTGRES_DB=...
POSTGRES_SSLMODE=require        # require для Neon, disable для локального Postgres
DATABASE_URL=postgres://USER:PASSWORD@HOST:PORT/DB?sslmode=require
JWT_SECRET=...
```

4. Примени миграции:

```bash
goose -dir migrations postgres "$DATABASE_URL" up
```

5. Запусти сервер:

```bash
go run ./cmd/app
```

Сервер поднимется на `http://127.0.0.1:10000`.

## API

**Пользователи**

| Метод | Путь | Описание | Статус |
|---|---|---|---|
| POST | `/users` | Создать (зарегистрировать) пользователя | реализован |
| GET | `/users` | Список пользователей | заглушка |
| GET | `/users/:uuid` | Получить пользователя | заглушка |
| PUT | `/users/:uuid` | Обновить пользователя | заглушка |
| PATCH | `/users/:uuid` | Частично обновить | заглушка |
| DELETE | `/users/:uuid` | Удалить пользователя | заглушка |

**Авторизация** *(в разработке, пока не подключена)*

| Метод | Путь | Описание |
|---|---|---|
| POST | `/auth/register` | Регистрация |
| POST | `/auth/login` | Вход |
| POST | `/auth/refresh` | Обновление токенов |
| POST | `/auth/logout` | Выход |

## Схема БД

Миграции описывают: пользователей, привязку OAuth-провайдеров (`user_linked_providers`), device-токены для push, проекты, задачи, теги, участников проектов, уведомления и refresh-токены.
