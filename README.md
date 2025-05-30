## Стек технологий
- Язык: Golang
- База данных: PostgreSQL
- Генерация кода: SQLC
- Миграции: Atlas
- Контейнеризация: Docker
- Управление задачами: Makefile

## Запуск приложения
  ```bash
  docker compose up --build
  ```


## Дополнительно
- Сгенерировать код SQLC:
  ```bash
  make sqlc_gen
  ```
- Запустите миграции базы данных:
  ```bash
  make migrate_db MIGRATION_NAME=init
  ```
- Проверить проект:
  ```bash
  make validate
  ```
- Выполнить полную проверку и сборку проекта:
  ```bash
  make all
  ```

## API Документация
API документация доступна по адресу: http://localhost:8080/api/docs/index.html

## База данных
- СУБД: PostgreSQL
- Миграции: Управление миграциями осуществляется с помощью Atlas. Для применения миграций используйте команду:
  ```bash
  make migrate-up
  ```
- Генерация кода: SQLC используется для генерации типизированного кода для работы с базой данных.


## Генерация свагера (swagger)
- Вводим команду для генерации свагера
    ```
    swag init --output internal/routes/api/docs --parseDependency --parseInternal --dir cmd/test_task,internal/routes/api/v1
  ```
