package router

import (
	"altastore/api/controllers"
	"altastore/constants"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Route(e *echo.Echo,
	customerController *controllers.CustomerController,
	productController *controllers.ProductController,
	categoryController *controllers.CategoryController,
	cartController *controllers.CartController,
	cartDetailController *controllers.CartDetailsController,
	// checkoutController *controllers.CheckoutController,
) {
	// ------------------------------------------------------------------
	// Login & Register
	// ------------------------------------------------------------------
	e.POST("/api/register", customerController.RegisterCustomerController)
	e.POST("/api/login", customerController.LoginCustomerController)

	// Auth JWT
	eAuth := e.Group("")
	eAuth.Use(middleware.JWT([]byte(constants.SECRET_JWT)))

	// Customers
	eAuth.GET("/api/customers", customerController.GetAllCustomerController)

	// ------------------------------------------------------------------
	// CRUD Product
	// ------------------------------------------------------------------
	e.GET("/api/products", productController.GetAllProductController)
	eAuth.POST("/api/products", productController.PostProductController)
	e.GET("/api/products/:id", productController.GetProductController)
	eAuth.PUT("/api/products/:id", productController.UpdateProductController)
	eAuth.DELETE("/api/products/:id", productController.DeleteProductController)

	// e.GET("/api/product/category/", productController.GetProductCategoryController)
	// e.GET("/api/product/category", productController.GetProductCategoryController)

	// Category
	// eAuth.GET("/api/category/", categoryController.GetAllCategoryController)
	// eAuth.GET("/api/category", categoryController.GetAllCategoryController)
	// eAuth.POST("/api/category/", categoryController.PostCategoryController)
	// eAuth.POST("/api/category", categoryController.PostCategoryController)
	// eAuth.GET("/api/category/:id/", categoryController.GetCategoryController)
	// eAuth.GET("/api/category/:id", categoryController.GetCategoryController)
	// eAuth.PUT("/api/category/:id/", categoryController.UpdateCategoryController)
	// eAuth.PUT("/api/category/:id", categoryController.UpdateCategoryController)
	// eAuth.DELETE("/api/category/:id/", categoryController.DeleteCategoryController)
	// eAuth.DELETE("/api/category/:id", categoryController.DeleteCategoryController)

	// Checkout
	// eAuth.POST("/api/checkout", checkoutController.PostCheckoutController)

	// Transaction
	// eAuth.GET("/api/transaction/", transactionController.GetAllTransactionController)
	// eAuth.GET("/api/transaction", transactionController.GetAllTransactionController)
	// eAuth.GET("/api/transaction/:id", transactionController.GetAllTransactionController)
	// eAuth.GET("/api/transaction/:id", transactionController.GetAllTransactionController)

	// Cart
	//carts
	eAuth.POST("/carts/:productId/:qty", cartController.CreateCartController)                                 // create new shopping cart
	eAuth.POST("/carts/:cartId/details", cartDetailController.AddToCartController)                            //add product to cart
	eAuth.GET("/carts/:id", cartController.GetCartController)                                                 //get all product on a cart
	eAuth.DELETE("/carts/:id", cartController.DeleteCartController)                                           //delete cart and all products included on it
	eAuth.DELETE("/cartDetails/:carts_id/:products_id", cartDetailController.DeleteProductFromCartController) //delete product from cart

	// ------------------------------------------------------------------
	// CRUD Categories
	// ------------------------------------------------------------------
	eAuth.GET("/api/categories", categoryController.GetAllCategoryController)
	eAuth.GET("/api/categories/:id", categoryController.GetCategoryController)
	eAuth.POST("/api/categories", categoryController.AddCategoryController)
	eAuth.PUT("/api/categories/:id", categoryController.EditCategoryController)
	eAuth.DELETE("/api/categories/:id", categoryController.DeleteCategoryController)
	// ------------------------------------------------------------------
	// Payment Midtrans
	// ------------------------------------------------------------------
	e.GET("/api/payments/request/:id", controllers.RequestPayment)
	e.GET("/api/payments/status/:id", controllers.StatusPayment)

}
