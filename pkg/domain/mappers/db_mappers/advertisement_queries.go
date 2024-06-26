package dbmappers

import (
	"math"
	"study_marketplace/database/queries"
	entities "study_marketplace/pkg/domain/models/entities"
	reqmodels "study_marketplace/pkg/domain/models/request_models"
)

func AdvertisementToCreateAdvertisementParams(adv *entities.Advertisement) queries.CreateAdvertisementParams {
	return queries.CreateAdvertisementParams{
		Title:       adv.Title,
		ProviderID:  adv.Provider.ID,
		Attachment:  adv.Attachment,
		Experience:  int32(adv.Experience),
		Name:        adv.Category.Name,
		Time:        int32(adv.Time),
		Price:       int32(adv.Price),
		Currency:    StrToSqlStr(string(adv.Currency)),
		Format:      string(adv.Format),
		Language:    adv.Language,
		Description: adv.Description,
		MobilePhone: StrToSqlStr(adv.MobilePhone),
		Email:       StrToSqlStr(adv.Email),
		Telegram:    StrToSqlStr(adv.Telegram),
	}
}

func CreateAdvertisementRowToAdvertisement(row queries.CreateAdvertisementRow) *entities.Advertisement {
	return &entities.Advertisement{
		ID:    row.ID,
		Title: row.Title,
		Provider: &entities.User{
			ID:        row.ProviderID,
			Name:      row.ProviderName.String,
			Email:     row.ProviderEmail,
			Photo:     row.ProviderPhoto.String,
			Verified:  row.ProviderVerified,
			Role:      row.ProviderRole,
			CreatedAt: row.ProviderCreatedAt,
			UpdatedAt: row.ProviderUpdatedAt,
		},
		Attachment: row.Attachment,
		Experience: int(row.Experience),
		Category: &entities.Category{
			ID:   row.CategoryID,
			Name: row.CategoryName,
			ParentCategory: &entities.ParentCategory{
				Name: row.ParentCategoryName,
			},
		},
		Time:        int(row.Time),
		Price:       int(row.Price),
		Currency:    entities.AdvertisementCurrency(row.Currency.String),
		Format:      entities.AdvertisementFormat(row.Format),
		Language:    row.Language,
		Description: row.Description,
		MobilePhone: row.MobilePhone.String,
		Email:       row.Email.String,
		Telegram:    row.Telegram.String,
		CreatedAt:   row.CreatedAt,
		UpdatedAt:   row.UpdatedAt,
	}
}

func AdvertisementToUpdateAdvertisementParams(adv *entities.Advertisement) queries.UpdateAdvertisementParams {
	return queries.UpdateAdvertisementParams{
		ID:          adv.ID,
		ProviderID:  adv.Provider.ID,
		Title:       StrToSqlStr(adv.Title),
		Attachment:  StrToSqlStr(adv.Attachment),
		Experience:  IntToPostgresInt4(int32(adv.Experience)),
		Name:        StrToSqlStr(adv.Category.Name),
		Time:        IntToPostgresInt4(int32(adv.Time)),
		Price:       IntToPostgresInt4(int32(adv.Price)),
		Currency:    StrToSqlStr(string(adv.Currency)),
		Format:      StrToSqlStr(string(adv.Format)),
		Language:    StrToSqlStr(adv.Language),
		Description: StrToSqlStr(adv.Description),
		MobilePhone: StrToSqlStr(adv.MobilePhone),
		Email:       StrToSqlStr(adv.Email),
		Telegram:    StrToSqlStr(adv.Telegram),
	}
}

func GetFullAdvToAdvertisement(row queries.GetAdvertisementCategoryAndUserByIDRow) *entities.Advertisement {
	return &entities.Advertisement{
		ID:    row.ID,
		Title: row.Title,
		Provider: &entities.User{
			ID:        row.ProviderID,
			Name:      row.ProviderName.String,
			Email:     row.ProviderEmail,
			Photo:     row.ProviderPhoto.String,
			Verified:  row.ProviderVerified,
			Role:      row.ProviderRole,
			CreatedAt: row.ProviderCreatedAt,
			UpdatedAt: row.ProviderUpdatedAt,
		},
		Attachment: row.Attachment,
		Experience: int(row.Experience),
		Category: &entities.Category{
			ID:   row.CategoryID,
			Name: row.CategoryName,
			ParentCategory: &entities.ParentCategory{
				Name: row.ParentCategoryName,
			},
		},
		Time:        int(row.Time),
		Price:       int(row.Price),
		Currency:    entities.AdvertisementCurrency(row.Currency.String),
		Format:      entities.AdvertisementFormat(row.Format),
		Language:    row.Language,
		Description: row.Description,
		MobilePhone: row.MobilePhone.String,
		Email:       row.Email.String,
		Telegram:    row.Telegram.String,
		CreatedAt:   row.CreatedAt,
		UpdatedAt:   row.UpdatedAt,
	}
}

func GetAdvertisementsAllToAdvertisements(rows []queries.GetAdvertisementAllRow) []entities.Advertisement {
	advertisements := make([]entities.Advertisement, len(rows))
	for i := range rows {
		advertisements[i] = entities.Advertisement{
			ID:    rows[i].ID,
			Title: rows[i].Title,
			Provider: &entities.User{
				ID:        rows[i].ProviderID,
				Name:      rows[i].ProviderName.String,
				Email:     rows[i].ProviderEmail,
				Photo:     rows[i].ProviderPhoto.String,
				Verified:  rows[i].ProviderVerified,
				Role:      rows[i].ProviderRole,
				CreatedAt: rows[i].ProviderCreatedAt,
				UpdatedAt: rows[i].ProviderUpdatedAt,
			},
			Attachment: rows[i].Attachment,
			Experience: int(rows[i].Experience),
			Category: &entities.Category{
				ID:   rows[i].CategoryID,
				Name: rows[i].CategoryName,
				ParentCategory: &entities.ParentCategory{
					Name: rows[i].ParentCategoryName,
				},
			},
			Time:        int(rows[i].Time),
			Price:       int(rows[i].Price),
			Currency:    entities.AdvertisementCurrency(rows[i].Currency.String),
			Format:      entities.AdvertisementFormat(rows[i].Format),
			Language:    rows[i].Language,
			Description: rows[i].Description,
			MobilePhone: rows[i].MobilePhone.String,
			Email:       rows[i].Email.String,
			Telegram:    rows[i].Telegram.String,
			CreatedAt:   rows[i].CreatedAt,
			UpdatedAt:   rows[i].UpdatedAt,
		}
	}
	return advertisements
}

func FilterAdvToAdvPagination(params *queries.FilterAdvertisementsParams, filteredAdvs []queries.FilterAdvertisementsRow,
) *entities.AdvertisementPagination {

	advs := make([]entities.Advertisement, len(filteredAdvs))
	for i := range filteredAdvs {
		advs[i] = entities.Advertisement{
			ID:    filteredAdvs[i].ID,
			Title: filteredAdvs[i].Title,
			Provider: &entities.User{
				ID:        filteredAdvs[i].ProviderID,
				Name:      filteredAdvs[i].ProviderName.String,
				Email:     filteredAdvs[i].ProviderEmail,
				Photo:     filteredAdvs[i].ProviderPhoto.String,
				Verified:  filteredAdvs[i].ProviderVerified,
				Role:      filteredAdvs[i].ProviderRole,
				CreatedAt: filteredAdvs[i].ProviderCreatedAt,
				UpdatedAt: filteredAdvs[i].ProviderUpdatedAt,
			},
			Attachment: filteredAdvs[i].Attachment,
			Experience: int(filteredAdvs[i].Experience),
			Category: &entities.Category{
				ID:   filteredAdvs[i].CategoryID,
				Name: filteredAdvs[i].CategoryName,
				ParentCategory: &entities.ParentCategory{
					Name: filteredAdvs[i].ParentCategoryName,
				},
			},
			Time:        int(filteredAdvs[i].Time),
			Price:       int(filteredAdvs[i].Price),
			Currency:    entities.AdvertisementCurrency(filteredAdvs[i].Currency.String),
			Format:      entities.AdvertisementFormat(filteredAdvs[i].Format),
			Language:    filteredAdvs[i].Language,
			Description: filteredAdvs[i].Description,
			MobilePhone: filteredAdvs[i].MobilePhone.String,
			Email:       filteredAdvs[i].Email.String,
			Telegram:    filteredAdvs[i].Telegram.String,
			CreatedAt:   filteredAdvs[i].CreatedAt,
			UpdatedAt:   filteredAdvs[i].UpdatedAt,
		}
	}
	return &entities.AdvertisementPagination{
		Advertisements: advs,
		PaginationInfo: entities.PaginationInfo{
			TotalPages: int(math.Ceil(float64(filteredAdvs[0].TotalItems) / float64(params.Limitadv))),
			TotalCount: int(filteredAdvs[0].TotalItems),
			Page:       int((params.Offsetadv / params.Limitadv) + 1),
			PerPage:    int(params.Limitadv),
			Offset:     int(params.Offsetadv),
			OrderBy: func() string {
				if params.Orderby == "" {
					return "date"
				}
				return params.Orderby
			}(),
			SortOrder: func() string {
				if params.Sortorder == "" {
					return "asc"
				}
				return params.Sortorder
			}(),
		},
	}
}

func AdvertisementFilterRequestToFilterAdvertisementsParams(filter *reqmodels.AdvertisementFilterRequest) queries.FilterAdvertisementsParams {
	if filter.LimitAdv == 0 {
		filter.LimitAdv = 10
	}

	return queries.FilterAdvertisementsParams{
		Orderby:   filter.OrderBy,
		Sortorder: filter.SortOrder,
		Offsetadv: func() int32 {
			if filter.Page == 0 {
				return 0
			}
			return (filter.Page - 1) * filter.LimitAdv
		}(),
		Limitadv:     filter.LimitAdv,
		Advcategory:  filter.Category,
		Timelength:   filter.TimeLength,
		Currency:     string(entities.AdvertisementFormat(filter.Currency)),
		Advformat:    string(entities.AdvertisementFormat(filter.Format)),
		Minexp:       filter.MinExp,
		Maxexp:       filter.MaxExp,
		Minprice:     filter.MinPrice,
		Maxprice:     filter.MaxPrice,
		Advlanguage:  filter.Language,
		Titlekeyword: filter.TitleKeyword,
	}
}

func GetMyAdvertisementsToAdvertisements(row []queries.GetMyAdvertisementRow) []entities.Advertisement {
	advertisements := make([]entities.Advertisement, len(row))
	for i := range row {
		advertisements[i] = entities.Advertisement{
			ID:    row[i].ID,
			Title: row[i].Title,
			Provider: &entities.User{
				ID:        row[i].ProviderID,
				Name:      row[i].ProviderName.String,
				Email:     row[i].ProviderEmail,
				Photo:     row[i].ProviderPhoto.String,
				Verified:  row[i].ProviderVerified,
				Role:      row[i].ProviderRole,
				CreatedAt: row[i].ProviderCreatedAt,
				UpdatedAt: row[i].ProviderUpdatedAt,
			},
			Attachment: row[i].Attachment,
			Experience: int(row[i].Experience),
			Category: &entities.Category{
				ID:   row[i].CategoryID,
				Name: row[i].CategoryName,
				ParentCategory: &entities.ParentCategory{
					Name: row[i].ParentCategoryName,
				},
			},
			Time:        int(row[i].Time),
			Price:       int(row[i].Price),
			Currency:    entities.AdvertisementCurrency(row[i].Currency.String),
			Format:      entities.AdvertisementFormat(row[i].Format),
			Language:    row[i].Language,
			Description: row[i].Description,
			MobilePhone: row[i].MobilePhone.String,
			Email:       row[i].Email.String,
			Telegram:    row[i].Telegram.String,
			CreatedAt:   row[i].CreatedAt,
			UpdatedAt:   row[i].UpdatedAt,
		}
	}
	return advertisements
}
