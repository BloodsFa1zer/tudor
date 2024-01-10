// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: advertisement.sql

package queries

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

const createAdvertisement = `-- name: CreateAdvertisement :one
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
        RETURNING id, title, provider_id, attachment, experience, category_id, time, price, format, language, description, mobile_phone, email, telegram, created_at, updated_at
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
WHERE categories.parent_id IS NOT NULL
`

type CreateAdvertisementParams struct {
	Title       string      `json:"title"`
	ProviderID  int64       `json:"provider_id"`
	Attachment  string      `json:"attachment"`
	Experience  int32       `json:"experience"`
	Name        string      `json:"name"`
	Time        int32       `json:"time"`
	Price       int32       `json:"price"`
	Format      string      `json:"format"`
	Language    string      `json:"language"`
	Description string      `json:"description"`
	MobilePhone pgtype.Text `json:"mobile_phone"`
	Email       pgtype.Text `json:"email"`
	Telegram    pgtype.Text `json:"telegram"`
}

type CreateAdvertisementRow struct {
	ID                 int64       `json:"id"`
	Title              string      `json:"title"`
	Attachment         string      `json:"attachment"`
	Experience         int32       `json:"experience"`
	Time               int32       `json:"time"`
	Price              int32       `json:"price"`
	Format             string      `json:"format"`
	Language           string      `json:"language"`
	Description        string      `json:"description"`
	MobilePhone        pgtype.Text `json:"mobile_phone"`
	Email              pgtype.Text `json:"email"`
	Telegram           pgtype.Text `json:"telegram"`
	CreatedAt          time.Time   `json:"created_at"`
	UpdatedAt          time.Time   `json:"updated_at"`
	ProviderID         int64       `json:"provider_id"`
	ProviderName       pgtype.Text `json:"provider_name"`
	ProviderEmail      string      `json:"provider_email"`
	ProviderPhoto      pgtype.Text `json:"provider_photo"`
	ProviderVerified   bool        `json:"provider_verified"`
	ProviderRole       string      `json:"provider_role"`
	ProviderCreatedAt  time.Time   `json:"provider_created_at"`
	ProviderUpdatedAt  time.Time   `json:"provider_updated_at"`
	CategoryID         int64       `json:"category_id"`
	CategoryName       string      `json:"category_name"`
	ParentCategoryName string      `json:"parent_category_name"`
}

func (q *Queries) CreateAdvertisement(ctx context.Context, arg CreateAdvertisementParams) (CreateAdvertisementRow, error) {
	row := q.db.QueryRow(ctx, createAdvertisement,
		arg.Title,
		arg.ProviderID,
		arg.Attachment,
		arg.Experience,
		arg.Name,
		arg.Time,
		arg.Price,
		arg.Format,
		arg.Language,
		arg.Description,
		arg.MobilePhone,
		arg.Email,
		arg.Telegram,
	)
	var i CreateAdvertisementRow
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Attachment,
		&i.Experience,
		&i.Time,
		&i.Price,
		&i.Format,
		&i.Language,
		&i.Description,
		&i.MobilePhone,
		&i.Email,
		&i.Telegram,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ProviderID,
		&i.ProviderName,
		&i.ProviderEmail,
		&i.ProviderPhoto,
		&i.ProviderVerified,
		&i.ProviderRole,
		&i.ProviderCreatedAt,
		&i.ProviderUpdatedAt,
		&i.CategoryID,
		&i.CategoryName,
		&i.ParentCategoryName,
	)
	return i, err
}

const deleteAdvertisementByID = `-- name: DeleteAdvertisementByID :exec
DELETE FROM advertisements WHERE id = $1 AND provider_id = $2
`

type DeleteAdvertisementByIDParams struct {
	ID         int64 `json:"id"`
	ProviderID int64 `json:"provider_id"`
}

func (q *Queries) DeleteAdvertisementByID(ctx context.Context, arg DeleteAdvertisementByIDParams) error {
	_, err := q.db.Exec(ctx, deleteAdvertisementByID, arg.ID, arg.ProviderID)
	return err
}

const deleteAdvertisementByUserID = `-- name: DeleteAdvertisementByUserID :exec
DELETE FROM advertisements
WHERE provider_id = $1
`

func (q *Queries) DeleteAdvertisementByUserID(ctx context.Context, providerID int64) error {
	_, err := q.db.Exec(ctx, deleteAdvertisementByUserID, providerID)
	return err
}

const filterAdvertisements = `-- name: FilterAdvertisements :many
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
    AND (NULLIF($1::text, '')::text IS NULL OR categories.name = $1::text)
    AND (NULLIF($2::int, 0) IS NULL OR time <= $2::int)
    AND (NULLIF($3::text, '') IS NULL OR format = $3::text)
    AND ((NULLIF($4::int, 0) IS NULL AND NULLIF($5::int, 0) IS NULL) OR (experience >= $4::int AND experience <= $5::int))
    AND ((NULLIF($6::int, 0) IS NULL AND NULLIF($7::int, 0) IS NULL) OR (price >= $6::int AND price <= $7::int))
    AND (NULLIF($8::text, '') IS NULL OR language = $8::text)
    AND (NULLIF($9::text, '') IS NULL OR title ILIKE '%' || $9::text || '%')
ORDER BY
  ( CASE
    WHEN $10::text = 'price' AND $11::text = 'desc' THEN CAST(price AS TEXT)
    WHEN $10::text = 'experience' AND $11::text = 'desc' THEN CAST(experience AS TEXT)
    WHEN $10::text = 'date' AND $11::text = 'desc' THEN CAST(advertisements.created_at AS TEXT) END) DESC,
  ( CASE
    WHEN $10::text = 'price' THEN CAST(price AS TEXT)
    WHEN $10::text = 'experience' THEN CAST(experience AS TEXT)  
    ELSE CAST(advertisements.created_at AS TEXT) END) ASC                                     
LIMIT $13::integer    
OFFSET $12::integer
`

type FilterAdvertisementsParams struct {
	Advcategory  string `json:"advcategory"`
	Timelength   int32  `json:"timelength"`
	Advformat    string `json:"advformat"`
	Minexp       int32  `json:"minexp"`
	Maxexp       int32  `json:"maxexp"`
	Minprice     int32  `json:"minprice"`
	Maxprice     int32  `json:"maxprice"`
	Advlanguage  string `json:"advlanguage"`
	Titlekeyword string `json:"titlekeyword"`
	Orderby      string `json:"orderby"`
	Sortorder    string `json:"sortorder"`
	Offsetadv    int32  `json:"offsetadv"`
	Limitadv     int32  `json:"limitadv"`
}

type FilterAdvertisementsRow struct {
	ID                 int64       `json:"id"`
	Title              string      `json:"title"`
	Attachment         string      `json:"attachment"`
	Experience         int32       `json:"experience"`
	Time               int32       `json:"time"`
	Price              int32       `json:"price"`
	Format             string      `json:"format"`
	Language           string      `json:"language"`
	Description        string      `json:"description"`
	MobilePhone        pgtype.Text `json:"mobile_phone"`
	Email              pgtype.Text `json:"email"`
	Telegram           pgtype.Text `json:"telegram"`
	CreatedAt          time.Time   `json:"created_at"`
	UpdatedAt          time.Time   `json:"updated_at"`
	ProviderID         int64       `json:"provider_id"`
	ProviderName       pgtype.Text `json:"provider_name"`
	ProviderEmail      string      `json:"provider_email"`
	ProviderPhoto      pgtype.Text `json:"provider_photo"`
	ProviderVerified   bool        `json:"provider_verified"`
	ProviderRole       string      `json:"provider_role"`
	ProviderCreatedAt  time.Time   `json:"provider_created_at"`
	ProviderUpdatedAt  time.Time   `json:"provider_updated_at"`
	CategoryID         int64       `json:"category_id"`
	CategoryName       string      `json:"category_name"`
	ParentCategoryName string      `json:"parent_category_name"`
	TotalItems         int64       `json:"total_items"`
}

func (q *Queries) FilterAdvertisements(ctx context.Context, arg FilterAdvertisementsParams) ([]FilterAdvertisementsRow, error) {
	rows, err := q.db.Query(ctx, filterAdvertisements,
		arg.Advcategory,
		arg.Timelength,
		arg.Advformat,
		arg.Minexp,
		arg.Maxexp,
		arg.Minprice,
		arg.Maxprice,
		arg.Advlanguage,
		arg.Titlekeyword,
		arg.Orderby,
		arg.Sortorder,
		arg.Offsetadv,
		arg.Limitadv,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FilterAdvertisementsRow
	for rows.Next() {
		var i FilterAdvertisementsRow
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Attachment,
			&i.Experience,
			&i.Time,
			&i.Price,
			&i.Format,
			&i.Language,
			&i.Description,
			&i.MobilePhone,
			&i.Email,
			&i.Telegram,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.ProviderID,
			&i.ProviderName,
			&i.ProviderEmail,
			&i.ProviderPhoto,
			&i.ProviderVerified,
			&i.ProviderRole,
			&i.ProviderCreatedAt,
			&i.ProviderUpdatedAt,
			&i.CategoryID,
			&i.CategoryName,
			&i.ParentCategoryName,
			&i.TotalItems,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAdvertisementAll = `-- name: GetAdvertisementAll :many
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
ORDER BY advertisements.created_at DESC LIMIT 10
`

type GetAdvertisementAllRow struct {
	ID                 int64       `json:"id"`
	Title              string      `json:"title"`
	Attachment         string      `json:"attachment"`
	Experience         int32       `json:"experience"`
	Time               int32       `json:"time"`
	Price              int32       `json:"price"`
	Format             string      `json:"format"`
	Language           string      `json:"language"`
	Description        string      `json:"description"`
	MobilePhone        pgtype.Text `json:"mobile_phone"`
	Email              pgtype.Text `json:"email"`
	Telegram           pgtype.Text `json:"telegram"`
	CreatedAt          time.Time   `json:"created_at"`
	UpdatedAt          time.Time   `json:"updated_at"`
	ProviderID         int64       `json:"provider_id"`
	ProviderName       pgtype.Text `json:"provider_name"`
	ProviderEmail      string      `json:"provider_email"`
	ProviderPhoto      pgtype.Text `json:"provider_photo"`
	ProviderVerified   bool        `json:"provider_verified"`
	ProviderRole       string      `json:"provider_role"`
	ProviderCreatedAt  time.Time   `json:"provider_created_at"`
	ProviderUpdatedAt  time.Time   `json:"provider_updated_at"`
	CategoryID         int64       `json:"category_id"`
	CategoryName       string      `json:"category_name"`
	ParentCategoryName string      `json:"parent_category_name"`
}

func (q *Queries) GetAdvertisementAll(ctx context.Context) ([]GetAdvertisementAllRow, error) {
	rows, err := q.db.Query(ctx, getAdvertisementAll)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAdvertisementAllRow
	for rows.Next() {
		var i GetAdvertisementAllRow
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Attachment,
			&i.Experience,
			&i.Time,
			&i.Price,
			&i.Format,
			&i.Language,
			&i.Description,
			&i.MobilePhone,
			&i.Email,
			&i.Telegram,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.ProviderID,
			&i.ProviderName,
			&i.ProviderEmail,
			&i.ProviderPhoto,
			&i.ProviderVerified,
			&i.ProviderRole,
			&i.ProviderCreatedAt,
			&i.ProviderUpdatedAt,
			&i.CategoryID,
			&i.CategoryName,
			&i.ParentCategoryName,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAdvertisementByCategory = `-- name: GetAdvertisementByCategory :many
WITH id AS (SELECT id FROM categories WHERE name = $1)
SELECT id, title, provider_id, attachment, experience, category_id, time, price, format, language, description, mobile_phone, email, telegram, created_at, updated_at FROM advertisements
WHERE category_id = id
`

func (q *Queries) GetAdvertisementByCategory(ctx context.Context, name string) ([]Advertisement, error) {
	rows, err := q.db.Query(ctx, getAdvertisementByCategory, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Advertisement
	for rows.Next() {
		var i Advertisement
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.ProviderID,
			&i.Attachment,
			&i.Experience,
			&i.CategoryID,
			&i.Time,
			&i.Price,
			&i.Format,
			&i.Language,
			&i.Description,
			&i.MobilePhone,
			&i.Email,
			&i.Telegram,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAdvertisementByExperience = `-- name: GetAdvertisementByExperience :many
SELECT id, title, provider_id, attachment, experience, category_id, time, price, format, language, description, mobile_phone, email, telegram, created_at, updated_at FROM advertisements
WHERE experience >= $1
AND experience <= $2
`

type GetAdvertisementByExperienceParams struct {
	Experience   int32 `json:"experience"`
	Experience_2 int32 `json:"experience_2"`
}

func (q *Queries) GetAdvertisementByExperience(ctx context.Context, arg GetAdvertisementByExperienceParams) ([]Advertisement, error) {
	rows, err := q.db.Query(ctx, getAdvertisementByExperience, arg.Experience, arg.Experience_2)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Advertisement
	for rows.Next() {
		var i Advertisement
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.ProviderID,
			&i.Attachment,
			&i.Experience,
			&i.CategoryID,
			&i.Time,
			&i.Price,
			&i.Format,
			&i.Language,
			&i.Description,
			&i.MobilePhone,
			&i.Email,
			&i.Telegram,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAdvertisementByFormat = `-- name: GetAdvertisementByFormat :many
SELECT id, title, provider_id, attachment, experience, category_id, time, price, format, language, description, mobile_phone, email, telegram, created_at, updated_at FROM advertisements
WHERE format = $1
`

func (q *Queries) GetAdvertisementByFormat(ctx context.Context, format string) ([]Advertisement, error) {
	rows, err := q.db.Query(ctx, getAdvertisementByFormat, format)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Advertisement
	for rows.Next() {
		var i Advertisement
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.ProviderID,
			&i.Attachment,
			&i.Experience,
			&i.CategoryID,
			&i.Time,
			&i.Price,
			&i.Format,
			&i.Language,
			&i.Description,
			&i.MobilePhone,
			&i.Email,
			&i.Telegram,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAdvertisementByID = `-- name: GetAdvertisementByID :one
SELECT id, title, provider_id, attachment, experience, category_id, time, price, format, language, description, mobile_phone, email, telegram, created_at, updated_at FROM advertisements
WHERE id = $1
`

func (q *Queries) GetAdvertisementByID(ctx context.Context, id int64) (Advertisement, error) {
	row := q.db.QueryRow(ctx, getAdvertisementByID, id)
	var i Advertisement
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.ProviderID,
		&i.Attachment,
		&i.Experience,
		&i.CategoryID,
		&i.Time,
		&i.Price,
		&i.Format,
		&i.Language,
		&i.Description,
		&i.MobilePhone,
		&i.Email,
		&i.Telegram,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getAdvertisementByLanguage = `-- name: GetAdvertisementByLanguage :many
SELECT id, title, provider_id, attachment, experience, category_id, time, price, format, language, description, mobile_phone, email, telegram, created_at, updated_at FROM advertisements
WHERE language = $1
`

func (q *Queries) GetAdvertisementByLanguage(ctx context.Context, language string) ([]Advertisement, error) {
	rows, err := q.db.Query(ctx, getAdvertisementByLanguage, language)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Advertisement
	for rows.Next() {
		var i Advertisement
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.ProviderID,
			&i.Attachment,
			&i.Experience,
			&i.CategoryID,
			&i.Time,
			&i.Price,
			&i.Format,
			&i.Language,
			&i.Description,
			&i.MobilePhone,
			&i.Email,
			&i.Telegram,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAdvertisementByTime = `-- name: GetAdvertisementByTime :many
SELECT id, title, provider_id, attachment, experience, category_id, time, price, format, language, description, mobile_phone, email, telegram, created_at, updated_at FROM advertisements
WHERE time <= $1
`

func (q *Queries) GetAdvertisementByTime(ctx context.Context, argTime int32) ([]Advertisement, error) {
	rows, err := q.db.Query(ctx, getAdvertisementByTime, argTime)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Advertisement
	for rows.Next() {
		var i Advertisement
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.ProviderID,
			&i.Attachment,
			&i.Experience,
			&i.CategoryID,
			&i.Time,
			&i.Price,
			&i.Format,
			&i.Language,
			&i.Description,
			&i.MobilePhone,
			&i.Email,
			&i.Telegram,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAdvertisementByUserID = `-- name: GetAdvertisementByUserID :many
SELECT id, title, provider_id, attachment, experience, category_id, time, price, format, language, description, mobile_phone, email, telegram, created_at, updated_at FROM advertisements
WHERE provider_id = $1
`

func (q *Queries) GetAdvertisementByUserID(ctx context.Context, providerID int64) ([]Advertisement, error) {
	rows, err := q.db.Query(ctx, getAdvertisementByUserID, providerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Advertisement
	for rows.Next() {
		var i Advertisement
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.ProviderID,
			&i.Attachment,
			&i.Experience,
			&i.CategoryID,
			&i.Time,
			&i.Price,
			&i.Format,
			&i.Language,
			&i.Description,
			&i.MobilePhone,
			&i.Email,
			&i.Telegram,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAdvertisementCategoryAndUserByID = `-- name: GetAdvertisementCategoryAndUserByID :one
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
WHERE advertisements.id = $1
`

type GetAdvertisementCategoryAndUserByIDRow struct {
	ID                 int64       `json:"id"`
	Title              string      `json:"title"`
	Attachment         string      `json:"attachment"`
	Experience         int32       `json:"experience"`
	Time               int32       `json:"time"`
	Price              int32       `json:"price"`
	Format             string      `json:"format"`
	Language           string      `json:"language"`
	Description        string      `json:"description"`
	MobilePhone        pgtype.Text `json:"mobile_phone"`
	Email              pgtype.Text `json:"email"`
	Telegram           pgtype.Text `json:"telegram"`
	CreatedAt          time.Time   `json:"created_at"`
	UpdatedAt          time.Time   `json:"updated_at"`
	ProviderID         int64       `json:"provider_id"`
	ProviderName       pgtype.Text `json:"provider_name"`
	ProviderEmail      string      `json:"provider_email"`
	ProviderPhoto      pgtype.Text `json:"provider_photo"`
	ProviderVerified   bool        `json:"provider_verified"`
	ProviderRole       string      `json:"provider_role"`
	ProviderCreatedAt  time.Time   `json:"provider_created_at"`
	ProviderUpdatedAt  time.Time   `json:"provider_updated_at"`
	CategoryID         int64       `json:"category_id"`
	CategoryName       string      `json:"category_name"`
	ParentCategoryName string      `json:"parent_category_name"`
}

func (q *Queries) GetAdvertisementCategoryAndUserByID(ctx context.Context, id int64) (GetAdvertisementCategoryAndUserByIDRow, error) {
	row := q.db.QueryRow(ctx, getAdvertisementCategoryAndUserByID, id)
	var i GetAdvertisementCategoryAndUserByIDRow
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Attachment,
		&i.Experience,
		&i.Time,
		&i.Price,
		&i.Format,
		&i.Language,
		&i.Description,
		&i.MobilePhone,
		&i.Email,
		&i.Telegram,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ProviderID,
		&i.ProviderName,
		&i.ProviderEmail,
		&i.ProviderPhoto,
		&i.ProviderVerified,
		&i.ProviderRole,
		&i.ProviderCreatedAt,
		&i.ProviderUpdatedAt,
		&i.CategoryID,
		&i.CategoryName,
		&i.ParentCategoryName,
	)
	return i, err
}

const getMyAdvertisement = `-- name: GetMyAdvertisement :many
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
WHERE advertisements.provider_id = $1 AND categories.parent_id IS NOT NULL
`

type GetMyAdvertisementRow struct {
	ID                 int64       `json:"id"`
	Title              string      `json:"title"`
	Attachment         string      `json:"attachment"`
	Experience         int32       `json:"experience"`
	Time               int32       `json:"time"`
	Price              int32       `json:"price"`
	Format             string      `json:"format"`
	Language           string      `json:"language"`
	Description        string      `json:"description"`
	MobilePhone        pgtype.Text `json:"mobile_phone"`
	Email              pgtype.Text `json:"email"`
	Telegram           pgtype.Text `json:"telegram"`
	CreatedAt          time.Time   `json:"created_at"`
	UpdatedAt          time.Time   `json:"updated_at"`
	ProviderID         int64       `json:"provider_id"`
	ProviderName       pgtype.Text `json:"provider_name"`
	ProviderEmail      string      `json:"provider_email"`
	ProviderPhoto      pgtype.Text `json:"provider_photo"`
	ProviderVerified   bool        `json:"provider_verified"`
	ProviderRole       string      `json:"provider_role"`
	ProviderCreatedAt  time.Time   `json:"provider_created_at"`
	ProviderUpdatedAt  time.Time   `json:"provider_updated_at"`
	CategoryID         int64       `json:"category_id"`
	CategoryName       string      `json:"category_name"`
	ParentCategoryName string      `json:"parent_category_name"`
}

func (q *Queries) GetMyAdvertisement(ctx context.Context, providerID int64) ([]GetMyAdvertisementRow, error) {
	rows, err := q.db.Query(ctx, getMyAdvertisement, providerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetMyAdvertisementRow
	for rows.Next() {
		var i GetMyAdvertisementRow
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Attachment,
			&i.Experience,
			&i.Time,
			&i.Price,
			&i.Format,
			&i.Language,
			&i.Description,
			&i.MobilePhone,
			&i.Email,
			&i.Telegram,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.ProviderID,
			&i.ProviderName,
			&i.ProviderEmail,
			&i.ProviderPhoto,
			&i.ProviderVerified,
			&i.ProviderRole,
			&i.ProviderCreatedAt,
			&i.ProviderUpdatedAt,
			&i.CategoryID,
			&i.CategoryName,
			&i.ParentCategoryName,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateAdvertisement = `-- name: UpdateAdvertisement :one
UPDATE advertisements
SET
  title = COALESCE($3, title),
  attachment = COALESCE($4, attachment),
  experience = COALESCE($5, experience),
  category_id = COALESCE((SELECT id FROM categories WHERE name = $6), category_id),
  time = COALESCE($7, time),
  price = COALESCE($8, price),
  format = COALESCE($9, format),
  language = COALESCE($10, language),
  description = COALESCE($11, description),
  mobile_phone = COALESCE($12, mobile_phone),
  email = COALESCE($13, email),
  telegram = COALESCE($14, telegram)
WHERE advertisements.id = $1 AND advertisements.provider_id = $2
RETURNING advertisements.id
`

type UpdateAdvertisementParams struct {
	ID          int64       `json:"id"`
	ProviderID  int64       `json:"provider_id"`
	Title       pgtype.Text `json:"title"`
	Attachment  pgtype.Text `json:"attachment"`
	Experience  pgtype.Int4 `json:"experience"`
	Name        pgtype.Text `json:"name"`
	Time        pgtype.Int4 `json:"time"`
	Price       pgtype.Int4 `json:"price"`
	Format      pgtype.Text `json:"format"`
	Language    pgtype.Text `json:"language"`
	Description pgtype.Text `json:"description"`
	MobilePhone pgtype.Text `json:"mobile_phone"`
	Email       pgtype.Text `json:"email"`
	Telegram    pgtype.Text `json:"telegram"`
}

func (q *Queries) UpdateAdvertisement(ctx context.Context, arg UpdateAdvertisementParams) (int64, error) {
	row := q.db.QueryRow(ctx, updateAdvertisement,
		arg.ID,
		arg.ProviderID,
		arg.Title,
		arg.Attachment,
		arg.Experience,
		arg.Name,
		arg.Time,
		arg.Price,
		arg.Format,
		arg.Language,
		arg.Description,
		arg.MobilePhone,
		arg.Email,
		arg.Telegram,
	)
	var id int64
	err := row.Scan(&id)
	return id, err
}
