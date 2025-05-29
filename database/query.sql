-- query.sql

-- Добавление нового человека (соответствует REST POST /api/v1/User)
-- name: CreateUser :one
INSERT INTO users (
    name, 
    surname, 
    patronymic,
    age,
    gender,
    nationality
) VALUES (
    $1, $2, $3, $4, $5, $6
) RETURNING *;

-- Получение данных с фильтрами и пагинацией (соответствует REST GET /api/v1/User)
-- name: GetUsers :many
SELECT * FROM users
WHERE 
    ($1 = '' OR $1 IS NULL OR name ILIKE '%' || $1 || '%') AND
    ($2 = '' OR $2 IS NULL OR surname ILIKE '%' || $2 || '%') AND
    ($3 = 0 OR $3 IS NULL OR age >= $3) AND
    ($4 = 0 OR $4 IS NULL OR age <= $4) AND
    ($5 = '' OR $5 IS NULL OR gender = $5) AND
    ($6 = '' OR $6 IS NULL OR nationality = $6)
ORDER BY id
LIMIT $7 OFFSET $8;

-- Удаление по идентификатору (соответствует REST DELETE /api/v1/User/{id})
-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;

-- Обновление сущности (соответствует REST PUT /api/v1/User/{id})
-- name: UpdateUser :one
UPDATE users SET 
    name = COALESCE($2, name),
    surname = COALESCE($3, surname),
    patronymic = COALESCE($4, patronymic),
    age = COALESCE($5, age),
    gender = COALESCE($6, gender),
    nationality = COALESCE($7, nationality),
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING *;

-- Получение по ID (для вспомогательных операций)
-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1;