-- name: CreateAdvertisement :one
WITH cat_id AS (SELECT id FROM categories WHERE categories.name = $5),
     inserted_ad AS (
        INSERT INTO advertisements (
            title,
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
            telegram
        )
        VALUES (
            $1, $2, $3, $4, (SELECT id FROM categories WHERE categories.name = $5), $6, $7, $8, $9, $10, $11, $12, $13
        )
        RETURNING *
     )
SELECT inserted_ad.id, inserted_ad.title, inserted_ad.attachment, inserted_ad.experience, inserted_ad.time, inserted_ad.price, 
  inserted_ad.format, inserted_ad.language, inserted_ad.description, inserted_ad.mobile_phone, inserted_ad.email, 
  inserted_ad.telegram, inserted_ad.created_at, inserted_ad.updated_at, users.id AS provider_id, users.name AS provider_name,
  users.email AS provider_email, users.photo AS provider_photo, users.verified AS provider_verified, users.role AS provider_role,
  users.created_at AS provider_created_at, users.updated_at AS provider_updated_at, categories.id AS category_id, 
  categories.name AS category_name, parent_category.name AS parent_category_name
FROM inserted_ad
JOIN users ON inserted_ad.provider_id = users.id
JOIN categories ON inserted_ad.category_id = categories.id
LEFT JOIN categories AS parent_category ON categories.parent_id = parent_category.id
WHERE categories.parent_id IS NOT NULL;

-- name: UpdateAdvertisement :one
UPDATE advertisements
SET
  title = COALESCE(sqlc.narg('title'), title),
  attachment = COALESCE(sqlc.narg('attachment'), attachment),
  experience = COALESCE(sqlc.narg('experience'), experience),
  category_id = COALESCE((SELECT id FROM categories WHERE name = sqlc.narg('name')), category_id),
  time = COALESCE(sqlc.narg('time'), time),
  price = COALESCE(sqlc.narg('price'), price),
  format = COALESCE(sqlc.narg('format'), format),
  language = COALESCE(sqlc.narg('language'), language),
  description = COALESCE(sqlc.narg('description'), description),
  mobile_phone = COALESCE(sqlc.narg('mobile_phone'), mobile_phone),
  email = COALESCE(sqlc.narg('email'), email),
  telegram = COALESCE(sqlc.narg('telegram'), telegram)
WHERE advertisements.id = $1 AND advertisements.provider_id = $2
RETURNING advertisements.id;

-- name: GetAdvertisementAll :many
SELECT 
  advertisements.id AS id, 
  advertisements.title AS title,
  advertisements.attachment AS attachment,
  advertisements.experience AS experience,
  advertisements.time AS time,
  advertisements.price AS price,
  advertisements.format AS format,
  advertisements.language AS language,
  advertisements.description AS description, 
  advertisements.mobile_phone AS mobile_phone,
  advertisements.email AS email,
  advertisements.telegram AS telegram, 
  advertisements.created_at AS created_at,
  advertisements.updated_at AS updated_at,
  users.id AS provider_id,
  users.name AS provider_name,
  users.email AS provider_email,
  users.photo AS provider_photo,
  users.verified AS provider_verified,
  users.role AS provider_role,
  users.created_at AS provider_created_at,
  users.updated_at AS provider_updated_at,
  categories.id AS category_id,
  categories.name AS category_name, 
  parent_category.name AS parent_category_name
FROM advertisements
JOIN users ON advertisements.provider_id = users.id
JOIN categories ON advertisements.category_id = categories.id
LEFT JOIN categories AS parent_category ON categories.parent_id = parent_category.id
WHERE categories.parent_id IS NOT NULL
ORDER BY advertisements.created_at DESC LIMIT 10;

-- name: GetMyAdvertisement :many
SELECT 
  advertisements.id AS id, advertisements.title AS title, advertisements.attachment AS attachment, 
  advertisements.experience AS experience, advertisements.time AS time, advertisements.price AS price, 
  advertisements.format AS format, advertisements.language AS language, advertisements.description AS description, 
  advertisements.mobile_phone AS mobile_phone, advertisements.email AS email, advertisements.telegram AS telegram, 
  advertisements.created_at AS created_at, advertisements.updated_at AS updated_at, users.id AS provider_id, 
  users.name AS provider_name, users.email AS provider_email, users.photo AS provider_photo,
  users.verified AS provider_verified, users.role AS provider_role, users.created_at AS provider_created_at,
  users.updated_at AS provider_updated_at, categories.id AS category_id, categories.name AS category_name, 
  parent_category.name AS parent_category_name
FROM advertisements 
JOIN users ON advertisements.provider_id = users.id
JOIN categories ON advertisements.category_id = categories.id
LEFT JOIN categories AS parent_category ON categories.parent_id = parent_category.id
WHERE advertisements.provider_id = $1 AND categories.parent_id IS NOT NULL;

-- name: GetAdvertisementCategoryAndUserByID :one
SELECT 
  advertisements.id AS id, advertisements.title AS title, advertisements.attachment AS attachment, advertisements.experience AS experience,
  advertisements.time AS time, advertisements.price AS price, advertisements.format AS format, advertisements.language AS language,
  advertisements.description AS description, advertisements.mobile_phone AS mobile_phone, advertisements.email AS email,
  advertisements.telegram AS telegram, advertisements.created_at AS created_at, advertisements.updated_at AS updated_at,
  users.id AS provider_id, users.name AS provider_name, users.email AS provider_email, users.photo AS provider_photo,
  users.verified AS provider_verified, users.role AS provider_role, users.created_at AS provider_created_at,
  users.updated_at AS provider_updated_at, categories.id AS category_id, categories.name AS category_name, parent_category.name AS parent_category_name
FROM advertisements
JOIN users ON advertisements.provider_id = users.id
JOIN categories ON advertisements.category_id = categories.id
LEFT JOIN categories AS parent_category ON categories.parent_id = parent_category.id
WHERE advertisements.id = $1;

-- name: GetAdvertisementByID :one
SELECT * FROM advertisements
WHERE id = $1;

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

-- name: DeleteAdvertisementByID :exec
DELETE FROM advertisements WHERE id = $1 AND provider_id = $2;

-- name: DeleteAdvertisementByUserID :exec
DELETE FROM advertisements
WHERE provider_id = $1;

-- name: FilterAdvertisements :many
SELECT 
  advertisements.id,
  advertisements.title,
  advertisements.attachment,
  advertisements.experience,
  advertisements.time,
  advertisements.price,
  advertisements.format,
  advertisements.language,
  advertisements.description,
  advertisements.mobile_phone,
  advertisements.email,
  advertisements.telegram,
  advertisements.created_at,
  advertisements.updated_at,
  users.id AS provider_id,
  users.name AS provider_name,
  users.email AS provider_email,
  users.photo AS provider_photo,
  users.verified AS provider_verified,
  users.role AS provider_role,
  users.created_at AS provider_created_at,
  users.updated_at AS provider_updated_at,
  categories.id AS category_id,
  categories.name AS category_name, 
  parent_category.name AS parent_category_name,
  COUNT(*) OVER () AS total_items
FROM advertisements
JOIN users ON advertisements.provider_id = users.id
JOIN categories ON advertisements.category_id = categories.id
LEFT JOIN categories AS parent_category ON categories.parent_id = parent_category.id
WHERE categories.parent_id IS NOT NULL
    AND (NULLIF(sqlc.arg(advCategory)::text, '')::text IS NULL OR categories.name = sqlc.arg(advCategory)::text)
    AND (NULLIF(sqlc.arg(timeLength)::int, 0) IS NULL OR time <= sqlc.arg(timeLength)::int)
    AND (NULLIF(sqlc.arg(advFormat)::text, '') IS NULL OR format = sqlc.arg(advFormat)::text)
    AND ((NULLIF(sqlc.arg(minExp)::int, 0) IS NULL AND NULLIF(sqlc.arg(maxExp)::int, 0) IS NULL) OR (experience >= sqlc.arg(minExp)::int AND experience <= sqlc.arg(maxExp)::int))
    AND ((NULLIF(sqlc.arg(minPrice)::int, 0) IS NULL AND NULLIF(sqlc.arg(maxPrice)::int, 0) IS NULL) OR (price >= sqlc.arg(minPrice)::int AND price <= sqlc.arg(maxPrice)::int))
    AND (NULLIF(sqlc.arg(advLanguage)::text, '') IS NULL OR language = sqlc.arg(advLanguage)::text)
    AND (NULLIF(sqlc.arg(titleKeyword)::text, '') IS NULL OR title ILIKE '%' || sqlc.arg(titleKeyword)::text || '%')
ORDER BY
  ( CASE
    WHEN sqlc.arg(orderBy)::text = 'price' AND sqlc.arg(sortOrder)::text = 'desc' THEN CAST(price AS TEXT)
    WHEN sqlc.arg(orderBy)::text = 'experience' AND sqlc.arg(sortOrder)::text = 'desc' THEN CAST(experience AS TEXT)
    WHEN sqlc.arg(orderBy)::text = 'date' AND sqlc.arg(sortOrder)::text = 'desc' THEN CAST(advertisements.created_at AS TEXT) END) DESC,
  ( CASE
    WHEN sqlc.arg(orderBy)::text = 'price' THEN CAST(price AS TEXT)
    WHEN sqlc.arg(orderBy)::text = 'experience' THEN CAST(experience AS TEXT)  
    ELSE CAST(advertisements.created_at AS TEXT) END) ASC                                     
LIMIT sqlc.arg(limitAdv)::integer    
OFFSET sqlc.arg(offsetAdv)::integer;
