package controllers

import (
	"altastore/midtrans"
	"net/http"

	"github.com/labstack/echo/v4"
)

func RequestPayment(c echo.Context) error {
	ids := c.Param("id")
	// idStr, err := strconv.Atoi(id)
	// if err != nil {
	// 	return echo.NewHTTPError(http.StatusBadRequest, "ID Invalid")
	// }
	// amount, err := fungsi(idStr)
	// if err != nil {
	// 	return echo.NewHTTPError(http.StatusInternalServerError, "Error Bosqu")
	// }
	redirectURL, err := midtrans.RequestPayment(ids, 20000)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error Bosqu")
	}
	return c.JSON(http.StatusOK, redirectURL)
}

func StatusPayment(c echo.Context) error {
	ids := c.Param("id")
	// idStr, err := strconv.Atoi(id)
	// if err != nil {
	// 	return echo.NewHTTPError(http.StatusBadRequest, "ID Invalid")
	// }
	redirectURL, err := midtrans.StatusPayment(ids)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error Bosqu")
	}
	return c.JSON(http.StatusOK, redirectURL)
}
