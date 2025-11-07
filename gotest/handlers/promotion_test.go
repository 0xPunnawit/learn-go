package handlers

import (
	"fmt"
	"io"
	"net/http/httptest"
	"strconv"
	"testing"

	"gotest/services"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestPromotionCalculateDiscount(t *testing.T) {
	t.Run("success", func(t *testing.T) {

		// Arrange
		amount := 100
		expected := 90

		promoService := services.NewPromtionServiceMock()
		promoService.On("CalculateDiscount", amount).Return(expected, nil)

		promoHandler := NewPromotionHandler(promoService)

		app := fiber.New()
		app.Get("/calculate", promoHandler.CalculateDiscount)

		req := httptest.NewRequest("GET", fmt.Sprintf("/calculate?amount=%v", amount), nil)

		// Act
		res, err := app.Test(req)
		assert.NoError(t, err)
		defer res.Body.Close()

		// Assert
		if assert.Equal(t, fiber.StatusOK, res.StatusCode) {
			body, _ := io.ReadAll(res.Body)
			assert.Equal(t, strconv.Itoa(expected), string(body))
		}

	})
}
