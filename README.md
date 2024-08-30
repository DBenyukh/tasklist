# Tasklist API

## Описание

Этот проект представляет собой API для управления задачами, реализованный на Go с использованием PostgreSQL в качестве базы данных.

## Функционал

- Создание задачи
- Получение задачи по ID
- Обновление задачи по ID
- Удаление задачи по ID
- Просмотр всех задач

## Требования

- Go 1.20 или выше
- PostgreSQL 13 или выше
- [Postman](https://www.postman.com/) (для тестирования API)

## Установка и запуск

1. **Клонируйте репозиторий**
   ```
   git clone https://github.com/DBenyukh/tasklist.git
   cd your-repo-name
   ```
2. **Установите зависимости**
   ```
   go mod download
   ```

3. **Создать таблицу**
   ```
   CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    due_date TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
   );
   ```
4. **Создать в основном каталоге файл ".env" и поместить код**
   ```
   DB_URI=postgres://user:password@localhost:5432/yourdatabase
   WEB_ADDR=localhost:8090
   ```
   В строке: ```DB_URI=postgres://user:password@localhost:5432/yourdatabase```
   необходимо заменить **user**, **password**, **yourdatabase** и адрес (**localhost:5432**) на свои.

5. **Запустите сервер**\
   В корневом каталоге проекта выполните команду:
   ```
   go run main.go
   ```
   Сервер запустится и будет слушать на порту, указанном в переменной WEB_ADDR.

## Тестирование
Для тестирования API используйте Postman или любой другой инструмент для выполнения HTTP-запросов.
