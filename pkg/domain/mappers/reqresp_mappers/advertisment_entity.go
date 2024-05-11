package reqm

import (
	"fmt"
	entities "study_marketplace/pkg/domain/models/entities"
	reqmodels "study_marketplace/pkg/domain/models/request_models"
	respmodels "study_marketplace/pkg/domain/models/response_models"
)

func CreateAdvRequestToAdvertisement(req *reqmodels.CreateAdvertisementRequest, userId int64) *entities.Advertisement {
	return &entities.Advertisement{
		Title:       req.Title,
		Provider:    &entities.User{ID: userId},
		Attachment:  req.Attachment,
		Experience:  int(req.Experience),
		Category:    &entities.Category{Name: req.CategoryName},
		Time:        int(req.Time),
		Price:       int(req.Price),
		Currency:    entities.AdvertisementCurrency(req.Currency),
		Format:      entities.AdvertisementFormat(req.Format),
		Language:    req.Language,
		Description: req.Description,
		MobilePhone: req.MobilePhone,
		Email:       req.Email,
		Telegram:    req.Telegram,
	}
}

func UpdateAdvRequestToAdvertisement(req *reqmodels.UpdateAdvertisementRequest, userId int64) *entities.Advertisement {
	return &entities.Advertisement{
		ID:          req.ID,
		Title:       req.Title,
		Provider:    &entities.User{ID: userId},
		Attachment:  req.Attachment,
		Experience:  int(req.Experience),
		Category:    &entities.Category{Name: req.CategoryName},
		Time:        int(req.Time),
		Price:       int(req.Price),
		Currency:    entities.AdvertisementCurrency(req.Currency),
		Format:      entities.AdvertisementFormat(req.Format),
		Language:    req.Language,
		Description: req.Description,
		MobilePhone: req.MobilePhone,
		Email:       req.Email,
		Telegram:    req.Telegram,
	}
}

func AdvertisementToCreateUpdateAdvertisementResponse(adv *entities.Advertisement) respmodels.AdvertisementResponse {
	return respmodels.AdvertisementResponse{
		Advertisement: respmodels.ResponseAdvertisement{
			ID:           adv.ID,
			Title:        adv.Title,
			ProviderID:   adv.Provider.ID,
			ProviderName: adv.Provider.Name,
			Description:  adv.Description,
			Attachment:   adv.Attachment,
			Experience:   int32(adv.Experience),
			CategoryName: fmt.Sprintf("%s: %s", adv.Category.ParentCategory.Name, adv.Category.Name),
			Time:         int32(adv.Time),
			Price:        int32(adv.Price),
			Currency:     entities.AdvertisementCurrency(adv.Currency),
			Format:       entities.AdvertisementFormat(adv.Format),
			Language:     adv.Language,
			MobilePhone:  adv.MobilePhone,
			Email:        adv.Email,
			Telegram:     adv.Telegram,
			CreatedAt:    adv.CreatedAt,
			UpdatedAt:    adv.UpdatedAt,
		},
		Status: "success",
	}
}

func AdvertisementsToAdvertisementsResponses(adv []entities.Advertisement) *respmodels.AdvertisementsResponse {
	advResp := make([]respmodels.ResponseAdvertisement, len(adv))
	for i := range adv {
		advResp[i] = respmodels.ResponseAdvertisement{
			ID:           adv[i].ID,
			Title:        adv[i].Title,
			ProviderID:   adv[i].Provider.ID,
			ProviderName: adv[i].Provider.Name,
			Description:  adv[i].Description,
			Attachment:   adv[i].Attachment,
			Experience:   int32(adv[i].Experience),
			CategoryName: fmt.Sprintf("%s: %s", adv[i].Category.ParentCategory.Name, adv[i].Category.Name),
			Time:         int32(adv[i].Time),
			Price:        int32(adv[i].Price),
			Currency:     entities.AdvertisementCurrency(adv[i].Currency),
			Format:       entities.AdvertisementFormat(adv[i].Format),
			Language:     adv[i].Language,
			MobilePhone:  adv[i].MobilePhone,
			Email:        adv[i].Email,
			Telegram:     adv[i].Telegram,
			CreatedAt:    adv[i].CreatedAt,
			UpdatedAt:    adv[i].UpdatedAt,
		}
	}
	return &respmodels.AdvertisementsResponse{
		Advertisements: advResp,
		Status:         "success",
	}
}

func AdvertisementPaginationToAdvertisementPaginationResponse(adv *entities.AdvertisementPagination) *respmodels.AdvertisementPaginationResponse {
	return &respmodels.AdvertisementPaginationResponse{
		ResponseAdvertisementPagination: respmodels.ResponseAdvertisementPagination{
			Advertisements: AdvertisementsToAdvertisementsResponses(adv.Advertisements).Advertisements,
			PaginationInfo: adv.PaginationInfo,
		},
		Status: "success",
	}
}
