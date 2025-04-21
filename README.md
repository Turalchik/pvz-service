## Описание проекта

Сервис для сотрудников ПВЗ, позволяющий:

- Регистрацию и авторизацию пользователей (клиент, модератор, сотрудник ПВЗ)
- Заведение пунктов выдачи заказов (ПВЗ) модератором
- Открытие/закрытие приёмок товаров сотрудником ПВЗ
- Добавление и удаление товаров в рамках приёмки (LIFO)
- Получение списка ПВЗ с полной информацией по приёмкам по дате, с пагинацией

### Используемые технологии

- Go (grpc, sqlx, squirrel)
- gRPC API (protobuf)
- PostgreSQL

## Структура проекта

```plain
pvz-service/
├── cmd/main          # точка входа приложения
├── internal/         # логика приложения, репозиторий, сервисы
├── migrations/       # SQL-файлы миграций
├── pkg/pvz_service   # gRPC описание (protobuf)
└── integration/      # интеграционные тесты
```

## Prerequisites

- Go 1.20+
- PostgreSQL 12+

## Настройка окружения


Создайте `.env`, указав параметры:
   ```dotenv
   DB_USER=postgres
   DB_PASSWORD=secret
   DB_HOST=localhost
   DB_PORT=5432
   DB_NAME=pvz_database
   SECRET_KEY_FOR_JWT=<ваш base64-ключ>
   ```

## Запуск базы и миграции

### Локально

1. Создайте БД:
   ```sql
   CREATE DATABASE pvz_database;
   ```
2. Примените миграции:
   ```bash
   psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -f migrations/0001_create_pvzs.sql
   psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -f migrations/0001_create_receptions.sql
   psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -f migrations/0001_create_products.sql
   psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -f migrations/0001_create_users.sql
   ```

## Запуск сервера

```bash
go run cmd/main/main.go
```

Сервер будет слушать порт `:8080`.

## gRPC API

После старта сервис доступен по gRPC на `localhost:8080`. Ниже примеры вызовов через `grpcurl`.

- **Регистрация**:
  ```bash
  grpcurl -plaintext -d '{"login":"<LOGIN>","password":"<PASSWORD>","role":"<ROLE>"}' localhost:8080 pvz_service.PVZService/Register
  ```

- **Логин**:
  ```bash
   grpcurl -plaintext -d '{"login":"<LOGIN>","password":"<PASSWORD>"}' localhost:8080 pvz_service.PVZService/Login
   ```

- **Создать ПВЗ** (только модератор):
  ```bash
  grpcurl -plaintext -d '{"token":"<TOKEN>","city":"<CITY>"}' localhost:8080 pvz_service.PVZService/CreatePVZ
  ```

- **Открыть приёмку**:
  ```bash
   grpcurl -plaintext -d '{"token":"<TOKEN>","id":"<PVZ_ID>"}' localhost:8080 pvz_service.PVZService/OpenReception
   ```
  
- **Добавить товар**:
  ```bash
  grpcurl -plaintext -d '{"token":"<TOKEN>","id":"<PVZ_ID>","type":"<TYPE>"}' localhost:8080 pvz_service.PVZService/AddProduct
  ```

- **Удалить товар**:
  ```bash
   grpcurl -plaintext -d '{"token":"<TOKEN>","id":"<PVZ_ID>"}' localhost:8080 pvz_service.PVZService/RemoveProduct
   ```

- **Закрыть приёмку**:
  ```bash
  grpcurl -plaintext -d '{"token":"<TOKEN>","id":"<PVZ_ID>"}' localhost:8080 pvz_service.PVZService/CloseReception
  ```

- **Получить отфильтрованные ПВЗ**:
  ```bash
   grpcurl -plaintext -d '{"token":"<TOKEN>","start":"2024-04-22T00:00:00Z","finish":"2024-04-23T00:00:00Z","limit":<LIMIT>,"offset":<OFFSET>}' localhost:8080 pvz_service.PVZService/GetFilteredPVZs
   ```

## Тестирование

### Юнит-тесты

```bash
go test ./internal/app/repo/repo_test.go
go test ./internal/service/pvz_service/api_test.go
```

### Интеграционные тесты

```bash
go test ./integration/integration_test.go
```

---

*Автор: Девришев Турал, 2025*

