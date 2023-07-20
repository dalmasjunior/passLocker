package repository

import (
	"github.com/dalmasjunior/passLocker/internal/model"
	"github.com/google/uuid"
)

type CardsData struct {
	data map[string]*model.Card
}

func NewCardsData() *CardsData {
	return &CardsData{
		data: make(map[string]*model.Card),
	}
}

func (cd *CardsData) GetCardByID(id string) (*model.Card, bool) {
	card, exists := cd.data[id]
	return card, exists
}

func (cd *CardsData) GetAllCards() []*model.Card {
	cards := make([]*model.Card, 0, len(cd.data))
	for _, v := range cd.data {
		cards = append(cards, v)
	}
	return cards
}

func (cd *CardsData) SaveCard(card *model.Card) string {
	card.ID = uuid.New().String()
	cd.data[card.ID] = card
	return card.ID
}

func (cd *CardsData) DeleteCard(id string) bool {
	_, exists := cd.data[id]
	if exists {
		delete(cd.data, id)
	}
	return exists
}
