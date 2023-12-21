package services

import (
	"fmt"
	"time"

	"study_marketplace/internal/database/queries"
	"study_marketplace/internal/database/repositories"
	"study_marketplace/models"

	"github.com/gin-gonic/gin"
)

type AdvertisementService interface {
	AdvCreate(ctx *gin.Context, inputModel models.AdvertisementInput, userID int64, u UserService, c CategoriesService) (queries.Advertisement, error)
	AdvPatch(ctx *gin.Context, patch models.AdvertisementUpdate, userID int64, c *CategoriesService) (queries.Advertisement, error)
	AdvDelete(ctx *gin.Context, advId int64, userId int64) error
	AdvGetAll(ctx *gin.Context) ([]queries.Advertisement, error)
	AdvGetByID(ctx *gin.Context, id int64, c *CategoriesService) (queries.Advertisement, error)
	AdvGetFiltered(ctx *gin.Context, filter models.AdvertisementFilter) ([]queries.FilterAdvertisementsRow, error)
	AdvGetMy(ctx *gin.Context, userID int64) ([]queries.Advertisement, error)
}

type advertisementService struct {
	db repositories.AdvertisementsRepository
}

func NewAdvertisementService(db repositories.AdvertisementsRepository) AdvertisementService {
	return &advertisementService{db}
}

func (t *advertisementService) AdvCreate(ctx *gin.Context, inputModel models.AdvertisementInput, userID int64, u UserService, c CategoriesService) (queries.Advertisement, error) {
	user, err := u.UserInfo(ctx, userID)
	if err != nil {
		return queries.Advertisement{}, err
	}

	category, err := c.CatGetByName(ctx, inputModel.Category)

	if err != nil || category.ParentID.Int32 == 0 {
		return queries.Advertisement{}, fmt.Errorf("failed to get category")
	}

	args := &queries.CreateAdvertisementParams{
		Title:       inputModel.Title,
		Provider:    user.Name,
		ProviderID:  user.ID,
		Attachment:  inputModel.Attachment,
		Experience:  inputModel.Experience,
		Category:    category.Name,
		Time:        inputModel.Time,
		Price:       inputModel.Price,
		Format:      inputModel.Format,
		Language:    inputModel.Language,
		Description: inputModel.Description,
		MobilePhone: inputModel.MobilePhone,
		Email:       user.Email,
		Telegram:    inputModel.Telegram,
		CreatedAt:   time.Now(),
	}

	advertisement, err := t.db.CreateAdvertisement(ctx, *args)
	if err != nil {
		return queries.Advertisement{}, err
	}

	return advertisement, nil
}

func (t *advertisementService) AdvPatch(ctx *gin.Context, patch models.AdvertisementUpdate, userID int64, c *CategoriesService) (queries.Advertisement, error) {
	adv, err := t.db.GetAdvertisementByID(ctx, patch.ID)

	if err != nil {
		return queries.Advertisement{}, err
	}

	cat, err := c.CatGetByName(ctx, patch.Category)

	if err != nil || cat.ParentID.Int32 == 0 {
		return queries.Advertisement{}, fmt.Errorf("failed to get category")
	}

	if adv.ProviderID != userID {
		return queries.Advertisement{}, fmt.Errorf("Unauthorized")
	}

	advertisementTmp := &queries.UpdateAdvertisementParams{
		ID:          patch.ID,
		Title:       patch.Title,
		CreatedAt:   time.Now(),
		Attachment:  patch.Attachment,
		Experience:  patch.Experience,
		Category:    cat.Name,
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

func (t *advertisementService) AdvGetByID(ctx *gin.Context, id int64, c *CategoriesService) (queries.Advertisement, error) {
	advertisement, err := t.db.GetAdvertisementByID(ctx, id)

	if err != nil {
		return queries.Advertisement{}, err
	}

	catFull, err := c.CatGetFullName(ctx, advertisement.Category)

	if err != nil {
		return queries.Advertisement{}, err
	}
	advertisement.Category = fmt.Sprintf("%s: %s", catFull.ParentName.String, catFull.CategoryName)
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
