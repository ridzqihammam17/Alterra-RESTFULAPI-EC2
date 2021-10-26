package controllers

import (
	"altastore/models"
	"net/http"
	"strconv"

	echo "github.com/labstack/echo/v4"
)

type ProductController struct {
	productModel models.ProductModel
}

func NewProductController(productModel models.ProductModel) *ProductController {
	return &ProductController{
		productModel,
	}
}

func (controller *ProductController) GetAllProductController(c echo.Context) error {
	product, err := controller.productModel.GetAll()
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
		"message": "Success Get All Product",
		"data":    product,
	})
}

func (controller *ProductController) GetProductController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"code":    400,
			"message": "Bad Request",
		})
	}

	product, err := controller.productModel.Get(id)
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
		"message": "Success Get All Product",
		"data":    product,
	})
}

func (controller *ProductController) PostProductController(c echo.Context) error {
	// bind request value
	var productRequest models.Product
	if err := c.Bind(&productRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"code":    400,
			"message": "Bad Request",
		})
	}

	product := models.Product{
		Name:       productRequest.Name,
		Price:      productRequest.Price,
		Stock:      productRequest.Stock,
		CategoryID: productRequest.CategoryID,
	}
	_, err := controller.productModel.Insert(product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"code":    500,
			"message": "Internal Server Error",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"code":    200,
		"message": "Success Add Product",
	})
}

func (controller *ProductController) UpdateProductController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"code":    400,
			"message": "Bad Request",
		})
	}

	// bind request value
	var productRequest models.Product
	if err := c.Bind(&productRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"code":    400,
			"message": "Bad Request",
		})
	}

	product := models.Product{
		Name:       productRequest.Name,
		Price:      productRequest.Price,
		Stock:      productRequest.Stock,
		CategoryID: productRequest.CategoryID,
	}

	if _, err := controller.productModel.Edit(product, id); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"code":    400,
			"message": "Bad Request",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"code":    200,
		"message": "Success Edit Category",
	})
}

func (controller *ProductController) DeleteProductController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"code":    400,
			"message": "Bad Request",
		})
	}

	if _, err := controller.productModel.Delete(id); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"code":    400,
			"message": "Bad Request",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"code":    200,
		"message": "Success Delete Product",
	})
}
