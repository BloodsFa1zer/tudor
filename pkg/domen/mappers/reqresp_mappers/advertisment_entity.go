package reqm

import (
	"fmt"
	entities "study_marketplace/pkg/domen/models/entities"
	reqmodels "study_marketplace/pkg/domen/models/request_models"
	respmodels "study_marketplace/pkg/domen/models/response_models"
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
		Format:      req.Format,
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
		Format:      req.Format,
		Language:    req.Language,
		Description: req.Description,
		MobilePhone: req.MobilePhone,
		Email:       req.Email,
		Telegram:    req.Telegram,
	}
}

func AdvertisementToCreateUpdateAdvertisementResponse(adv *entities.Advertisement) respmodels.AdvertisementResponse {
	return respmodels.AdvertisementResponse{
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
		Format:       adv.Format,
		Language:     adv.Language,
		MobilePhone:  adv.MobilePhone,
		Email:        adv.Email,
		Telegram:     adv.Telegram,
		CreatedAt:    adv.CreatedAt.GoString(),
		UpdatedAt:    adv.UpdatedAt.GoString(),
	}
}

func AdvertisementsToAdvertisementResponses(adv []entities.Advertisement) []respmodels.AdvertisementResponse {
	advResp := make([]respmodels.AdvertisementResponse, len(adv))
	for i := range adv {
		advResp[i] = AdvertisementToCreateUpdateAdvertisementResponse(&adv[i])
	}
	return advResp
}

func AdvertisementPaginationToAdvertisementPaginationResponse(adv *entities.AdvertisementPagination) respmodels.AdvertisementPaginationResponse {
	return respmodels.AdvertisementPaginationResponse{
		Advertisements: AdvertisementsToAdvertisementResponses(adv.Advertisements),
		TotalPages:     adv.PaginationInfo.TotalPages,
		TotalCount:     adv.PaginationInfo.TotalCount,
		Page:           adv.PaginationInfo.Page,
		PerPage:        adv.PaginationInfo.PerPage,
		Offset:         adv.PaginationInfo.Offset,
	}
}
