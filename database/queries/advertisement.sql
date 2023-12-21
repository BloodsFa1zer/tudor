-- name: CreateAdvertisement :one
WITH category_id AS (SELECT id FROM categories WHERE name = $6)
INSERT INTO advertisements (
  title,
  provider,
  provider_id,
  attachment,
  experience,
  category_id,
  time,
  price,
  format,
  language,
  description,
  mobile_phone,
  email,
  telegram,
  created_at
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15
)
RETURNING *;

-- name: GetAdvertisementAll :many
SELECT * FROM advertisements
ORDER BY created_at DESC
LIMIT 10;

-- name: GetAdvertisementMy :many
SELECT * FROM advertisements 
WHERE provider_id = $1;

-- name: GetAdvertisementByID :one
SELECT * FROM advertisements
WHERE id = $1;

-- name: GetAdvertisementByUsername :many
SELECT * FROM advertisements
WHERE provider = $1;

-- name: GetAdvertisementByUserID :many
SELECT * FROM advertisements
WHERE provider_id = $1;

-- name: GetAdvertisementByCategory :many
WITH id AS (SELECT id FROM categories WHERE name = $1)
SELECT * FROM advertisements
WHERE category_id = id;

-- name: GetAdvertisementByTime :many
SELECT * FROM advertisements
WHERE time <= $1;

-- name: GetAdvertisementByFormat :many
SELECT * FROM advertisements
WHERE format = $1;

-- name: GetAdvertisementByExperience :many
SELECT * FROM advertisements
WHERE experience >= $1
AND experience <= $2;

-- name: GetAdvertisementByLanguage :many
SELECT * FROM advertisements
WHERE language = $1;

-- name: UpdateAdvertisement :one
WITH c_id AS (SELECT id FROM categories WHERE name = $6)
UPDATE advertisements
SET
  title = COALESCE($2, title),
  created_at = COALESCE($3, created_at),
  attachment = COALESCE($4, attachment),
  experience = COALESCE($5, experience),
  category_id = COALESCE($6, c_id),
  time = COALESCE($7, time),
  price = COALESCE($8, price),
  format = COALESCE($9, format),
  language = COALESCE($10, language),
  description = COALESCE($11, description),
  mobile_phone = COALESCE($12, mobile_phone),
  telegram = COALESCE($13, telegram)
WHERE advertisements.id = $1
RETURNING *;

-- name: DeleteAdvertisementByID :exec
DELETE FROM advertisements
WHERE id = $1;

-- name: DeleteAdvertisementByUserID :exec
DELETE FROM advertisements
WHERE provider_id = $1;

-- name: FilterAdvertisements :many
WITH filtered_ads AS (
SELECT * FROM advertisements
  WHERE
        (NULLIF(sqlc.arg(advCategory)::text, '')::text IS NULL OR category = sqlc.arg(advCategory)::text)
        AND (NULLIF(sqlc.arg(timeLength)::int, 0) IS NULL OR time <= sqlc.arg(timeLength)::int)
        AND (NULLIF(sqlc.arg(advFformat)::text, '') IS NULL OR format = sqlc.arg(advFormat)::text)
        AND ((NULLIF(sqlc.arg(minExp)::int, 0) IS NULL AND NULLIF(sqlc.arg(maxExp)::int, 0) IS NULL) OR (experience >= sqlc.arg(minExp)::int AND experience <= sqlc.arg(maxExp)::int))
        AND ((NULLIF(sqlc.arg(minPrice)::int, 0) IS NULL AND NULLIF(sqlc.arg(maxPrice)::int, 0) IS NULL) OR (price >= sqlc.arg(minPrice)::int AND price <= sqlc.arg(maxPrice)::int))
        AND (NULLIF(sqlc.arg(advLanguage)::text, '') IS NULL OR language = sqlc.arg(advLanguage)::text)
        AND (NULLIF(sqlc.arg(titleKeyword)::text, '') IS NULL OR title ILIKE '%' || sqlc.arg(titleKeyword)::text || '%')
)
SELECT *,
    COUNT(*) OVER () AS total_items
FROM filtered_ads
ORDER BY
  ( CASE
    WHEN sqlc.arg(orderBy)::text = 'price' AND sqlc.arg(sortOrder)::text = 'desc' THEN CAST(price AS TEXT)
    WHEN sqlc.arg(orderBy)::text = 'experience' AND sqlc.arg(sortOrder)::text = 'desc' THEN CAST(experience AS TEXT)
    WHEN sqlc.arg(orderBy)::text = 'date' AND sqlc.arg(sortOrder)::text = 'desc' THEN CAST(created_at AS TEXT) END) DESC,
  ( CASE
    WHEN sqlc.arg(orderBy)::text = 'price' THEN CAST(price AS TEXT)
    WHEN sqlc.arg(orderBy)::text = 'experience' THEN CAST(experience AS TEXT)  
    ELSE CAST(created_at AS TEXT) END) ASC                                     
LIMIT sqlc.arg(limitAdv)::integer    
OFFSET sqlc.arg(offsetAdv)::integer; 