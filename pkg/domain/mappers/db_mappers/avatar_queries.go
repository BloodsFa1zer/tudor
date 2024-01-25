package dbmappers

import (
	"strings"
	"study_marketplace/database/queries"
	entities "study_marketplace/pkg/domain/models/entities"
)

func AvatarToCreateAvatar(avatar *entities.Avatar) queries.CreateAvatarParams {
	return queries.CreateAvatarParams{
		Filename:   avatar.Filename,
		Fileadress: avatar.FileAddr,
		Data:       avatar.Data,
		ProviderID: avatar.Provider.ID,
	}
}

func DBAvatarToAvatar(dbavatar *queries.Avatar) *entities.Avatar {
	return &entities.Avatar{
		ID:       dbavatar.ID,
		Filename: dbavatar.Filename,
		Format:   strings.Split(dbavatar.Filename, ".")[1],
		Data:     dbavatar.Data,
		FileAddr: dbavatar.Fileadress,
		Provider: &entities.User{
			ID: dbavatar.ProviderID,
		},
	}
}

func AvatarToUpdateAvatar(avatar *entities.Avatar) queries.UpdateAvatarByProviderIDParams {
	return queries.UpdateAvatarByProviderIDParams{
		Column1:    avatar.Filename,
		Column2:    avatar.FileAddr,
		Column3:    avatar.Data,
		ProviderID: avatar.Provider.ID,
	}
}
