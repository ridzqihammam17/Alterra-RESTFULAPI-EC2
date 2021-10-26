package main

import (
	"fmt"

	"altastore/api/controllers"
	"altastore/api/router"
	"altastore/config"
	"altastore/models"
	"altastore/util"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func main() {
	//load config if available or set to default
	config := config.GetConfig()

	//initialize database connection based on given config
	db := util.MysqlDatabaseConnection(config)

	//initiate model
	customerModel := models.NewCustomerModel(db)
	productModel := models.NewProductModel(db)
	categoryModel := models.NewCategoryModel(db)
	cartModel := models.NewCartModel(db)
	cartDetailModel := models.NewCartDetailModel(db)
	// checkoutModel := models.NewCheckoutModel(db)

	//initiate controller
	newCustomerController := controllers.NewCustomerController(customerModel)
	newProductController := controllers.NewProductController(productModel)
	newCategoryController := controllers.NewCategoryController(categoryModel)
	newCartController := controllers.NewCartController(cartModel, cartDetailModel, productModel)
	newCartDetailController := controllers.NewCartDetailController(cartModel, cartDetailModel, productModel)

	// newCheckoutController := controllers.NewCheckoutController(checkoutModel)
	//create echo http
	e := echo.New()

	//register API path and controller
	router.Route(e, newCustomerController, newProductController, newCategoryController, newCartController, newCartDetailController)

	// run server
	address := fmt.Sprintf("localhost:%d", config.Port)

	if err := e.Start(address); err != nil {
		log.Info("shutting down the server")
	}
}
