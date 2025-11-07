package services_test

import (
	"errors"
	"gotest/repositories"
	"gotest/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPromotionCalculateDiscount(t *testing.T) {

	type testCase struct {
		name            string
		purchaseMin     int
		discountPercent int
		amount          int
		expected        int
	}

	cases := []testCase{
		{name: "a", purchaseMin: 100, discountPercent: 10, amount: 100, expected: 90}, // ลด 10% เหลือ 90
		{name: "b", purchaseMin: 100, discountPercent: 10, amount: 50, expected: 50},  // ซื้อไม่ถึง 100 ไม่ลด
		{name: "c", purchaseMin: 100, discountPercent: 10, amount: 0, expected: 0},    // ซื้อ 0 เหลือ 0
		{name: "d", purchaseMin: 50, discountPercent: 20, amount: 50, expected: 40},   // ซื้อครบ 50 ลด 20%
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			// Arrange เตรียมค่า
			promoRepo := repositories.NewPromotionRepositoryMock()
			promoRepo.On("GetPromotion").Return(repositories.Promotion{
				ID:              1,
				PurchaseMin:     c.purchaseMin,
				DiscountPercent: c.discountPercent,
			}, nil)

			promoService := services.NewPromotionService(promoRepo)

			// Act จะทำจริงแล้ว
			discount, _ := promoService.CalculateDiscount(c.amount)

			// Assert ตรวจสอบ
			assert.Equal(t, c.expected, discount)
		})
	}

	t.Run("purchase amount zero", func(t *testing.T) {
		// Arrange เตรียมค่า
		promoRepo := repositories.NewPromotionRepositoryMock()
		promoRepo.On("GetPromotion").Return(repositories.Promotion{
			ID:              1,
			PurchaseMin:     100,
			DiscountPercent: 20,
		}, nil)

		promoService := services.NewPromotionService(promoRepo)

		// Act จะทำจริงแล้ว
		_, err := promoService.CalculateDiscount(0)

		// Assert ตรวจสอบ
		assert.ErrorIs(t, err, services.ErrZeroAmount)
		promoRepo.AssertNotCalled(t, "GetPromotion")
	})

	t.Run("repository error", func(t *testing.T) {
		// Arrange
		promoRepo := repositories.NewPromotionRepositoryMock()
		promoRepo.On("GetPromotion").Return(repositories.Promotion{}, errors.New("repository error"))

		promoService := services.NewPromotionService(promoRepo)

		// Act
		_, err := promoService.CalculateDiscount(100)

		// Assert
		assert.ErrorIs(t, err, services.ErrRepository)
	})

}
