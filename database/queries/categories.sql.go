// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: categories.sql

package queries

import (
	"context"
	"database/sql"
)

const getCategoriesWithChildren = `-- name: GetCategoriesWithChildren :many
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
GROUP BY p.id, p.name
`

type GetCategoriesWithChildrenRow struct {
	ParentID   int64       `json:"parent_id"`
	ParentName string      `json:"parent_name"`
	Children   interface{} `json:"children"`
}

func (q *Queries) GetCategoriesWithChildren(ctx context.Context) ([]GetCategoriesWithChildrenRow, error) {
	rows, err := q.db.Query(ctx, getCategoriesWithChildren)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetCategoriesWithChildrenRow
	for rows.Next() {
		var i GetCategoriesWithChildrenRow
		if err := rows.Scan(&i.ParentID, &i.ParentName, &i.Children); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCategoryAndParent = `-- name: GetCategoryAndParent :one
SELECT
    c.name AS category_name,
    p.name AS parent_name
FROM
    categories c
LEFT JOIN
    categories p ON c.parent_id = p.id
WHERE 
    c.name = $1
`

type GetCategoryAndParentRow struct {
	CategoryName string         `json:"category_name"`
	ParentName   sql.NullString `json:"parent_name"`
}

func (q *Queries) GetCategoryAndParent(ctx context.Context, name string) (GetCategoryAndParentRow, error) {
	row := q.db.QueryRow(ctx, getCategoryAndParent, name)
	var i GetCategoryAndParentRow
	err := row.Scan(&i.CategoryName, &i.ParentName)
	return i, err
}

const getCategoryByID = `-- name: GetCategoryByID :one
SELECT id, name, parent_id FROM categories
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetCategoryByID(ctx context.Context, id int64) (Category, error) {
	row := q.db.QueryRow(ctx, getCategoryByID, id)
	var i Category
	err := row.Scan(&i.ID, &i.Name, &i.ParentID)
	return i, err
}

const getCategoryByName = `-- name: GetCategoryByName :one
SELECT id, name, parent_id FROM categories 
WHERE name = $1 LIMIT 1
`

func (q *Queries) GetCategoryByName(ctx context.Context, name string) (Category, error) {
	row := q.db.QueryRow(ctx, getCategoryByName, name)
	var i Category
	err := row.Scan(&i.ID, &i.Name, &i.ParentID)
	return i, err
}

const getCategoryParents = `-- name: GetCategoryParents :many
SELECT id, name, parent_id FROM categories
WHERE parent_id = NULL
`

func (q *Queries) GetCategoryParents(ctx context.Context) ([]Category, error) {
	rows, err := q.db.Query(ctx, getCategoryParents)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Category
	for rows.Next() {
		var i Category
		if err := rows.Scan(&i.ID, &i.Name, &i.ParentID); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
