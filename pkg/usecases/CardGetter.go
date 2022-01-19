package usecases

import (
	"github.com/SeoSoojin/TKG_Kards/pkg/domain/models"
	"github.com/SeoSoojin/TKG_Kards/pkg/domain/services"
)

type cardGetter struct {
	CardGetterService *services.CardGetterService
}

func NewUCCardGetter(cardGetterService *services.CardGetterService) *cardGetter {

	return &cardGetter{
		CardGetterService: cardGetterService,
	}

}

func (cg *cardGetter) GetCardsByUserId(userId int) ([]models.Card, error) {

	return cg.CardGetterService.GetCardsByUserId(userId)

}

func (cg *cardGetter) GetCardsByCollectionId(collectionId int) ([]models.Card, error) {

	return cg.CardGetterService.GetCardsByCollectionId(collectionId)

}

func (cg *cardGetter) GetCardsByIdolId(idolId int) ([]models.Card, error) {

	return cg.CardGetterService.GetCardsByIdolId(idolId)

}

func (cg *cardGetter) GetCardsByGroupId(groupId int) ([]models.Card, error) {

	return cg.CardGetterService.GetCardsByGroupId(groupId)

}

func (cg *cardGetter) GetCardsByAlbumId(albumId int) ([]models.Card, error) {

	return cg.CardGetterService.GetCardsByAlbumId(albumId)

}

func (cg *cardGetter) GetCards() ([]models.Card, error) {

	return cg.CardGetterService.GetCards()

}
