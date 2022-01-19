package usecases

import "github.com/SeoSoojin/TKG_Kards/pkg/domain/models"

type UCCardGetter interface {
	GetCardsByUserId(userId int) ([]models.Card, error)
	GetCardsByCollectionId(collectionId int) ([]models.Card, error)
	GetCards() ([]models.Card, error)
}
