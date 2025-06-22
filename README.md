# Auth App (Go + Gin + PostgreSQL + JWT)

Простое API для аутентификации пользователей с JWT-токенами.

## 📌 Возможности

- Регистрация новых пользователей
- Вход с email/паролем
- Обновление JWT-токенов
- Защищенные маршруты
- CORS поддержка
- Работа с PostgreSQL

## 🚀 Быстрый старт

### Требования
- Go 1.18+
- PostgreSQL 12+

### Установка

1. Клонируйте репозиторий:
```bash
git clone https://github.com/lafflame/AuthApp.git
cd auth-app
```

2. Настройте базу данных:
```sql
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    username VARCHAR(20) NOT NULL UNIQUE,
    email VARCHAR(30) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(15) NOT NULL DEFAULT 'user',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
```

3. Настройте переменные окружения в файле main.go:
```ini
DB_HOST=localhost
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=auth_app
DB_PORT=5432
```

4. Запустите приложение:
```bash
go run main.go
```

Или с помощью Makefile:
```bash
make run
```

## 📡 API Endpoints

### Аутентификация

| Метод | Путь | Описание |
|-------|------|----------|
| POST | `/api/auth/register` | Регистрация нового пользователя |
| POST | `/api/auth/login`    | Вход в систему |
| POST | `/api/auth/refresh`  | Обновление токенов |

### Защищенные маршруты

| Метод | Путь | Описание |
|-------|------|----------|
| GET | `/api/protected/test`  | Тестовый защищенный маршрут |
| GET | `/api/protected/check` | Проверка валидности токена  |
| GET | `/api/protected/data`  | Получение защищенных данных |

## 🛠 Технологии

- **Backend**: Go 1.18+
- **Фреймворк**: Gin
- **База данных**: PostgreSQL
- **Аутентификация**: JWT
- **CORS**: gin-contrib/cors

## 📝 Примеры запросов

### Регистрация
```bash
curl -X POST http://localhost:8080/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser","email":"test@example.com","password":"password123"}'
```

### Вход
```bash
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","password":"password123"}'
```

### Защищенный маршрут
```bash
curl -X GET http://localhost:8080/api/protected/test \
  -H "Authorization: Bearer YOUR_ACCESS_TOKEN"
```

## 🏗 Структура проекта

```
auth-app/
├── controllers/       # Контроллеры
├── middleware/        # Мидлвары
├── models/            # Модели данных
├── main.go            # Точка входа
├── go.mod             # Зависимости
└── README.md          # Документация
```
