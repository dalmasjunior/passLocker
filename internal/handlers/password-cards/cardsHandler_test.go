package passwordCardsHandler_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	passwordCardsHandler "github.com/dalmasjunior/passLocker/internal/handlers/password-cards"
	"github.com/dalmasjunior/passLocker/internal/model"
	"github.com/dalmasjunior/passLocker/internal/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/stretchr/testify/assert"
)

func convertToCardArray(i interface{}) ([]model.Card, error) {
	var cards []model.Card

	// Type assertion
	cardInterface, ok := i.([]interface{})

	if !ok {
		return cards, errors.New("invalid interface type, expected []interface{}")
	}

	// Convert []interface{} to []model.Card
	for _, cardIface := range cardInterface {
		cardMap, ok := cardIface.(map[string]interface{})
		if !ok {
			return cards, errors.New("invalid type within interface slice, expected map[string]interface{}")
		}

		card := model.Card{
			ID:       cardMap["id"].(string),
			Title:    cardMap["title"].(string),
			URL:      cardMap["url"].(string),
			Username: cardMap["username"].(string),
			Password: cardMap["password"].(string),
		}
		cards = append(cards, card)
	}
	return cards, nil
}

func convertToCard(i interface{}) (model.Card, error) {
	var card model.Card

	// Type assertion
	cardMap, ok := i.(map[string]interface{})
	if !ok {
		return card, errors.New("invalid interface type, expected map[string]interface{}")
	}

	// Extract data from the map and populate the card struct
	card.ID, ok = cardMap["id"].(string)
	if !ok {
		return card, errors.New("invalid type for 'id' field, expected string")
	}

	card.Title, ok = cardMap["title"].(string)
	if !ok {
		return card, errors.New("invalid type for 'title' field, expected string")
	}

	card.URL, ok = cardMap["url"].(string)
	if !ok {
		return card, errors.New("invalid type for 'url' field, expected string")
	}

	card.Username, ok = cardMap["username"].(string)
	if !ok {
		return card, errors.New("invalid type for 'username' field, expected string")
	}

	card.Password, ok = cardMap["password"].(string)
	if !ok {
		return card, errors.New("invalid type for 'password' field, expected string")
	}

	return card, nil
}

func Test_CustomCardHandler_GetAllCards(t *testing.T) {
	// Create a new CardsData instance for testing
	cardsData := repository.NewCardsData()

	// Add some test data
	card1 := model.Card{Title: "Card 1", URL: "example.com", Username: "user1", Password: "pass1"}

	card1.ID = cardsData.SaveCard(&card1)

	// Create a new Fiber app and register the cardHandler endpoints
	app := fiber.New()
	card := app.Group("password-cards", logger.New())

	// Create a new customCardHandler instance with the test data
	cardHandler := passwordCardsHandler.NewCustomCardHandler(cardsData)

	card.Get("/", cardHandler.GetAllCards)

	// Perform a GET request to /
	req := httptest.NewRequest("GET", "/password-cards", nil)
	resp, _ := app.Test(req)

	// Assert that the response status code is 200 OK
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %s\n", err)
		return
	}

	var retCard1 passwordCardsHandler.SuccessResponse
	err = json.Unmarshal(body, &retCard1)

	sliceCard, err := convertToCardArray(retCard1.Data)

	// Assert that the response data contains the test cards
	assert.Equal(t, card1, sliceCard[0])
}

func Test_CustomCardHandler_CreateCard(t *testing.T) {
	// Create a new CardsData instance for testing
	cardsData := repository.NewCardsData()
	app := fiber.New()
	card := app.Group("password-cards", logger.New())

	// Create a new customCardHandler instance with the test data
	cardHandler := passwordCardsHandler.NewCustomCardHandler(cardsData)
	card.Post("/", cardHandler.CreateCard)

	card1 := model.Card{Title: "Card 1", URL: "example.com", Username: "user1", Password: "pass1"}

	requestBody, err := json.Marshal(card1)

	if err != nil {
		t.Fatalf("Failed to marshal request body: %v", err)
	}

	//Perform a POST request to /
	req := httptest.NewRequest("POST", "/password-cards", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusCreated, resp.StatusCode)
}

func Test_CustomCardHandler_GetCard(t *testing.T) {
	cardsData := repository.NewCardsData()

	// Add some test data
	card1 := model.Card{Title: "Card 1", URL: "example.com", Username: "user1", Password: "pass1"}
	card1.ID = cardsData.SaveCard(&card1)

	// Create a new Fiber app and register the cardHandler endpoints
	app := fiber.New()
	card := app.Group("password-cards", logger.New()) // Add "/" before "password-cards"

	// Create a new customCardHandler instance with the test data
	cardHandler := passwordCardsHandler.NewCustomCardHandler(cardsData)

	card.Get("/:cardId", cardHandler.GetCard) // Update the route path to include "/:id"

	// Perform a GET request to /password-cards/{ID}
	req := httptest.NewRequest("GET", "/password-cards/"+card1.ID, nil)
	resp, _ := app.Test(req)

	// Assert that the response status code is 200 OK
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %s\n", err)
		return
	}

	var retCard1 passwordCardsHandler.SuccessResponse
	err = json.Unmarshal(body, &retCard1)

	sliceCard, err := convertToCard(retCard1.Data)

	// Assert that the response data contains the test card
	assert.Equal(t, card1, sliceCard)
}

func Test_CustomCardHandler_UpdateCard(t *testing.T) {
	cardsData := repository.NewCardsData()

	// Add some test data
	card1 := model.Card{Title: "Card 1", URL: "example.com", Username: "user1", Password: "pass1"}
	card1.ID = cardsData.SaveCard(&card1)

	cards := cardsData.GetAllCards()
	fmt.Println(cards[0])
	// Create a new Fiber app and register the cardHandler endpoints
	app := fiber.New()
	card := app.Group("password-cards", logger.New())

	// Create a new customCardHandler instance with the test data
	cardHandler := passwordCardsHandler.NewCustomCardHandler(cardsData)

	card.Put("/:cardId", cardHandler.UpdateCard) // Update the route path to include "/:id"

	// Modify card1 with new data
	updatedCard1 := model.Card{
		ID:       card1.ID,
		Title:    "Updated Card 1",
		URL:      "updated-example.com",
		Username: "updated-user1",
		Password: "updated-pass1",
	}

	requestBody, err := json.Marshal(updatedCard1)
	if err != nil {
		t.Fatalf("Failed to marshal request body: %v", err)
	}

	// Perform a PUT request to /password-cards/{ID}
	req := httptest.NewRequest("PUT", "/password-cards/"+card1.ID, bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)

	// Assert that the response status code is 200 OK or 204 No Content
	assert.True(t, resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusNoContent)
}

func Test_CustomCardHandler_DeleteCard(t *testing.T) {
	cardsData := repository.NewCardsData()

	// Add some test data
	card1 := model.Card{Title: "Card 1", URL: "example.com", Username: "user1", Password: "pass1"}
	card1.ID = cardsData.SaveCard(&card1)

	// Create a new Fiber app and register the cardHandler endpoints
	app := fiber.New()
	card := app.Group("/password-cards", logger.New())

	// Create a new customCardHandler instance with the test data
	cardHandler := passwordCardsHandler.NewCustomCardHandler(cardsData)

	card.Delete("/:cardId", cardHandler.DeleteCard) // Update the route path to include "/:id"

	// Perform a DELETE request to /password-cards/{ID}
	req := httptest.NewRequest("DELETE", "/password-cards/"+card1.ID, nil)
	resp, _ := app.Test(req)

	// Assert that the response status code is 204 No Content
	assert.Equal(t, http.StatusOK, resp.StatusCode)

}
