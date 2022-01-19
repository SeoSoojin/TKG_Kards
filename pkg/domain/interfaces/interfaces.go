package interfaces

import "github.com/SeoSoojin/TKG_Kards/pkg/domain/models"

type CardGetter interface {
	GetCardsByUserId(userId int) ([]models.Card, error)
	GetCardsByCollectionId(collectionId int) ([]models.Card, error)
	GetCards() ([]models.Card, error)
}
