package controllers

import (
	"altastore/models"
	"net/http"

	echo "github.com/labstack/echo/v4"
)

type CheckoutController struct {
	checkoutModel models.CheckoutModel
}

func NewCheckoutController(checkoutModel models.CheckoutModel) *CheckoutController {
	return &CheckoutController{
		checkoutModel,
	}
}

func (controller *CheckoutController) PostCheckoutController(c echo.Context) error {
	_, err := controller.checkoutModel.AddCheckoutID()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"code":    400,
			"message": "Bad Request",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"code":    200,
		"message": "Success Checkout",
	})
}
