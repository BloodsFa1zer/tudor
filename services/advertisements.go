package services

import (
	"fmt"
	"time"

	"study_marketplace/database/queries"
	"study_marketplace/database/repositories"
	"study_marketplace/domen/models"

	"github.com/gin-gonic/gin"
)

type AdvertisementService interface {
	AdvCreate(ctx *gin.Context, inputModel models.AdvertisementInput, userID int64) (queries.Advertisement, error)
	AdvPatch(ctx *gin.Context, patch models.AdvertisementUpdate, userID int64) (queries.Advertisement, error)
	AdvDelete(ctx *gin.Context, advId int64, userId int64) error
	AdvGetAll(ctx *gin.Context) ([]queries.Advertisement, error)
	AdvGetByID(ctx *gin.Context, id int64) (queries.Advertisement, error)
	AdvGetFiltered(ctx *gin.Context, filter models.AdvertisementFilter) ([]queries.FilterAdvertisementsRow, error)
	AdvGetMy(ctx *gin.Context, userID int64) ([]queries.Advertisement, error)
}

type advertisementService struct {
	db repositories.AdvertisementsRepository
}

func NewAdvertisementService(db repositories.AdvertisementsRepository) AdvertisementService {
	return &advertisementService{db}
}

func (t *advertisementService) AdvCreate(ctx *gin.Context, inputModel models.AdvertisementInput, userID int64) (queries.Advertisement, error) {
	// user, err := u.UserInfo(ctx, userID)
	// if err != nil {
	// 	return queries.Advertisement{}, err
	// }

	args := &queries.CreateAdvertisementParams{
		Title: inputModel.Title,
		// Provider:    user.Name,
		// ProviderID:  user.ID,
		Attachment:  inputModel.Attachment,
		Experience:  inputModel.Experience,
		Name:        inputModel.Category,
		Time:        inputModel.Time,
		Price:       inputModel.Price,
		Format:      inputModel.Format,
		Language:    inputModel.Language,
		Description: inputModel.Description,
		MobilePhone: inputModel.MobilePhone,
		// Email:       user.Email,
		Telegram:  inputModel.Telegram,
		CreatedAt: time.Now(),
	}

	advertisement, err := t.db.CreateAdvertisement(ctx, *args)
	if err != nil {
		return queries.Advertisement{}, err
	}

	return advertisement, nil
}

func (t *advertisementService) AdvPatch(ctx *gin.Context, patch models.AdvertisementUpdate, userID int64) (queries.Advertisement, error) {
	advertisementTmp := &queries.UpdateAdvertisementParams{
		ID:          patch.ID,
		Title:       patch.Title,
		Attachment:  patch.Attachment,
		Experience:  patch.Experience,
		Name:        patch.Category,
		Time:        patch.Time,
		Price:       patch.Price,
		Format:      patch.Format,
		Language:    patch.Language,
		Description: patch.Description,
		MobilePhone: patch.MobilePhone,
		Telegram:    patch.Telegram,
	}

	result, err := t.db.UpdateAdvertisement(ctx, *advertisementTmp)

	if err != nil {
		return queries.Advertisement{}, err
	}

	return result, nil
}

func (t *advertisementService) AdvDelete(ctx *gin.Context, advId int64, userId int64) error {
	advertisement, err := t.db.GetAdvertisementByID(ctx, advId)

	if err != nil {
		return fmt.Errorf("Advertisement not found.")
	}

	if advertisement.ProviderID != userId {
		return fmt.Errorf("Don't have permissions to delete advertisement.")
	}

	err = t.db.DeleteAdvertisementByID(ctx, advId)
	if err != nil {
		return err
	}

	return nil
}

func (t *advertisementService) AdvGetAll(ctx *gin.Context) ([]queries.Advertisement, error) {

	advertisements, err := t.db.GetAdvertisementAll(ctx)

	if err != nil {
		return nil, err
	}

	return advertisements, nil
}

func (t *advertisementService) AdvGetByID(ctx *gin.Context, id int64) (queries.Advertisement, error) {
	adverAndCat, err := t.db.GetAdvertisementAndCategoryByID(ctx, id)

	if err != nil {
		return queries.Advertisement{}, err
	}
	advertisement :=
		queries.Advertisement{
			ID:         adverAndCat.ID,
			Title:      adverAndCat.Title,
			Provider:   adverAndCat.Provider,
			ProviderID: adverAndCat.ProviderID,
			Attachment: adverAndCat.Attachment,
			Experience: adverAndCat.Experience,
			CategoryID: adverAndCat.CategoryID,
			Time:       adverAndCat.Time,
			Price:      adverAndCat.Price,
			Format:     adverAndCat.Format,
			Language:   adverAndCat.Language,
			CreatedAt:  adverAndCat.CreatedAt,
		}
	// advertisement.Category = fmt.Sprintf("%s: %s", catFull.ParentName.String, catFull.CategoryName)
	return advertisement, nil
}

func (t *advertisementService) AdvGetFiltered(ctx *gin.Context, filter models.AdvertisementFilter) ([]queries.FilterAdvertisementsRow, error) {
	argFilter := queries.FilterAdvertisementsParams{
		Orderby:      filter.Orderby,
		Sortorder:    filter.Sortorder,
		Offsetadv:    filter.Offsetadv,
		Limitadv:     filter.Limitadv,
		Advcategory:  filter.Category,
		Timelength:   filter.Timelength,
		Advformat:    filter.Format,
		Minexp:       filter.Minexp,
		Maxexp:       filter.Maxexp,
		Minprice:     filter.Minprice,
		Maxprice:     filter.Maxprice,
		Advlanguage:  filter.Language,
		Titlekeyword: filter.Titlekeyword,
	}

	advertisements, err := t.db.FilterAdvertisements(ctx, argFilter)
	if err != nil {
		return nil, err
	}

	return advertisements, nil
}

func (t *advertisementService) AdvGetMy(ctx *gin.Context, userID int64) ([]queries.Advertisement, error) {
	advertisements, err := t.db.GetAdvertisementMy(ctx, userID)

	if err != nil {
		return nil, err
	}

	return advertisements, nil
}
