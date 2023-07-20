package router

import (
	passCardsHandler "github.com/dalmasjunior/passLocker/internal/handlers/password-cards"
	"github.com/dalmasjunior/passLocker/internal/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type customRouter struct {
	*fiber.App
}

func NewCustomRouter() *customRouter {
	app := fiber.New()
	app.Use(cors.New())
	return &customRouter{
		App: app,
	}
}

func (cr *customRouter) SetupRoutes() {
	card := cr.Group("password-cards", logger.New())

	cardsData := repository.NewCardsData()

	cardsHandler := passCardsHandler.NewCustomCardHandler(cardsData)

	card.Get("/", cardsHandler.GetAllCards)
	card.Post("/", cardsHandler.CreateCard)
	card.Get("/:cardId", cardsHandler.GetCard)
	card.Put("/:cardId", cardsHandler.UpdateCard)
	card.Delete("/:cardId", cardsHandler.DeleteCard)
}
