package reqm

import (
	"fmt"
	entities "study_marketplace/pkg/domen/models/entities"
	reqmodels "study_marketplace/pkg/domen/models/request_models"
	respmodels "study_marketplace/pkg/domen/models/response_models"
)

func CreateUpdateAdvRequestToAdvertisement(req *reqmodels.CreateUpdateAdvertisementRequest, userId int64) *entities.Advertisement {
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
