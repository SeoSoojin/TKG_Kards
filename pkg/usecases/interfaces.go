package usecases

import "github.com/SeoSoojin/TKG_Kards/pkg/domain/models"

type UCCardGetter interface {
	GetCardsByUserId(userId int) ([]models.Card, error)
	GetCardsByCollectionId(collectionId int) ([]models.Card, error)
	GetCardsByIdolId(idolId int) ([]models.Card, error)
	GetCardsByGroupId(groupId int) ([]models.Card, error)
	GetCardsByAlbumId(albumId int) ([]models.Card, error)
	GetCards() ([]models.Card, error)
}
