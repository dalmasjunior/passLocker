package passwordCardsHandler

import (
	"net/http"

	"github.com/dalmasjunior/passLocker/internal/model"
	"github.com/gofiber/fiber/v2"
)

type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type SuccessResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type customCardHandler struct {
	cardsData model.CardRepository
}

const (
	ErrorCardIdNotFound = "Card ID not Found"
	ErrorIvalidJsonData = "Invalid JSON data"
)

func NewCustomCardHandler(cardsData model.CardRepository) *customCardHandler {
	return &customCardHandler{
		cardsData: cardsData,
	}
}

func (cd *customCardHandler) GetAllCards(c *fiber.Ctx) error {
	cards := cd.cardsData.GetAllCards()
	return c.Status(http.StatusOK).JSON(SuccessResponse{"success", cards})
}

func (cd *customCardHandler) CreateCard(c *fiber.Ctx) error {
	var card model.Card
	if err := c.BodyParser(&card); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{"error", ErrorIvalidJsonData})
	}

	card.ID = cd.cardsData.SaveCard(&card)
	return c.Status(fiber.StatusCreated).JSON(SuccessResponse{"success", card})
}

func (cd *customCardHandler) GetCard(c *fiber.Ctx) error {
	id := c.Params("cardId")
	card, exists := cd.cardsData.GetCardByID(id)
	if exists {
		return c.Status(fiber.StatusOK).JSON(SuccessResponse{"success", card})
	}

	return c.Status(fiber.StatusNotFound).JSON(ErrorResponse{"error", ErrorCardIdNotFound})
}

func (cd *customCardHandler) UpdateCard(c *fiber.Ctx) error {
	id := c.Params("cardId")

	var card model.Card

	if err := c.BodyParser(&card); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{"error", ErrorIvalidJsonData})
	}

	oldCard, exists := cd.cardsData.GetCardByID(id)
	if exists {
		oldCard.Title = card.Title
		oldCard.URL = card.URL
		oldCard.Username = card.Username
		oldCard.Password = card.Password

		return c.Status(fiber.StatusOK).JSON(SuccessResponse{"success", oldCard})
	}

	return c.Status(fiber.StatusNotFound).JSON(ErrorResponse{"error", ErrorCardIdNotFound})

}

func (cd *customCardHandler) DeleteCard(c *fiber.Ctx) error {
	id := c.Params("cardId")
	exists := cd.cardsData.DeleteCard(id)
	if !exists {
		return c.Status(fiber.StatusNotFound).JSON(ErrorResponse{"error", ErrorCardIdNotFound})
	}

	return c.Status(fiber.StatusOK).JSON(SuccessResponse{"success", nil})
}
