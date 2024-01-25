-- name: CreateAvatar :one
INSERT INTO avatars (filename, fileadress, data, provider_id) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: UpdateAvatarByProviderID :one
UPDATE avatars SET 
    filename = COALESCE(NULLIF($1, ''), filename),
    fileadress = COALESCE(NULLIF($2, ''), fileadress),
    data = COALESCE(NULLIF($3, ''), data)
WHERE provider_id = $4 RETURNING *;

-- name: DeleteAvatarByProviderID :one
WITH deleted AS (
    DELETE FROM avatars WHERE provider_id = $1 RETURNING *
)
SELECT COUNT(*) FROM deleted;

-- name: GetAvatarByProviderID :one
SELECT * FROM avatars WHERE provider_id = $1;

-- name: GetAvatarByID :one
SELECT * FROM avatars WHERE id = $1;