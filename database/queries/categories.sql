-- name: GetCategoryByName :one
SELECT * FROM categories 
WHERE name = $1 LIMIT 1;

-- name: GetCategoryByID :one
SELECT * FROM categories
WHERE id = $1 LIMIT 1;

-- name: GetCategoryParents :many
SELECT * FROM categories
WHERE parent_id = NULL;

-- name: GetCategoryAndParent :one
SELECT
    c.name AS category_name,
    p.name AS parent_name
FROM
    categories c
LEFT JOIN
    categories p ON c.parent_id = p.id
WHERE 
    c.name = $1;

-- name: GetCategoriesWithChildren :many
WITH RECURSIVE RecursiveChildren AS (
    SELECT
        id AS child_id,
        name AS child_name,
        parent_id
    FROM categories
    WHERE parent_id IS NOT NULL
    UNION
    SELECT
        c.id AS child_id,
        c.name AS child_name,
        c.parent_id
    FROM categories c
    JOIN RecursiveChildren pc ON c.parent_id = pc.child_id
)
SELECT
    p.id AS parent_id,
    p.name AS parent_name,
    COALESCE(json_agg(json_build_object('id', c.child_id, 'name', c.child_name)), '[]'::json) AS children
FROM categories p
LEFT JOIN RecursiveChildren c ON p.id = c.parent_id
WHERE p.parent_id IS NULL
GROUP BY p.id, p.name;