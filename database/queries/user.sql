-- name: CreateUser :one
INSERT INTO users (
  name,
  email,
  photo,
  verified,
  password,
  role,
  updated_at
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
)
RETURNING *;




-- name: CreateorUpdateUser :one
DO $$
  BEGIN
     PERFORM * FROM users WHERE email = $2;
  IF FOUND THEN
BEGIN
  UPDATE users SET 
    name = COALESCE($1, name),
    photo = COALESCE($3, photo),
    verified = COALESCE($4, verified),
    password = COALESCE($5, password),
    role = COALESCE($6, role),
    updated_at = CURRENT_TIMESTAMP
  WHERE email = $2
  RETURNING users.id, users.name, users.email, users.photo, users.verified, users.password, users.role, users.created_at, users.updated_at;
ELSE
  INSERT INTO users (
    name,
    email,
    photo,
    verified,
    password,
    role
  ) VALUES (
    $1, $2, $3, $4, $5, $6
  )
  RETURNING users.id, users.name, users.email, users.photo, users.verified, users.password, users.role, users.created_at, users.updated_at;
END IF;
END;
$$;

-- name: GetUserById :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: IsUserEmailExist :one
SELECT EXISTS ( SELECT 1 FROM users WHERE email = $1);

-- name: ListUsers :many
SELECT * FROM users
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateUser :one
UPDATE users
set name = COALESCE($2, name),
email = COALESCE($3, email),
photo = COALESCE($4, photo),
verified = COALESCE($5, verified),
password = COALESCE($6, password),
role = COALESCE($7, role),
updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;
