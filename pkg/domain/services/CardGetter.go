package services

import (
	"github.com/SeoSoojin/TKG_Kards/pkg/domain/interfaces"
	"github.com/SeoSoojin/TKG_Kards/pkg/domain/models"
)

type CardGetterService struct {
	CardGetter interfaces.CardGetter
}

func NewCardGetterService(cardGetter interfaces.CardGetter) *CardGetterService {

	return &CardGetterService{
		CardGetter: cardGetter,
	}

}

func (cgs *CardGetterService) GetCardsByUserId(userId int) ([]models.Card, error) {

	if userId == 0 {

		return nil, models.ErrEmptyUserId

	}

	return cgs.CardGetter.GetCardsByUserId(userId)

}

func (cgs *CardGetterService) GetCardsByCollectionId(collectionId int) ([]models.Card, error) {

	if collectionId == 0 {

		return nil, models.ErrEmptyCollectionId

	}

	return cgs.CardGetter.GetCardsByCollectionId(collectionId)

}

func (cgs *CardGetterService) GetCards() ([]models.Card, error) {

	return cgs.CardGetter.GetCards()

}
