package test

import (
	"github.com/dalmasjunior/passLocker/internal/model"
)

type MockCardRepository struct {
	data map[string]*model.Card
}
