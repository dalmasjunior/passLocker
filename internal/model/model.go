package model

type Card struct {
	ID       string `json:"id,omitempty"`
	Title    string `json:"title"`
	URL      string `json:"url"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type CardRepository interface {
	GetCardByID(id string) (*Card, bool)
	GetAllCards() []*Card
	SaveCard(card *Card) string
	DeleteCard(id string) bool
}
