package controllers

import (
	// "altastore/api/controllers"
	"altastore/config"
	"altastore/models"
	"altastore/util"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	// "github.com/labstack/gommon/bytes"
	"github.com/stretchr/testify/assert"
)

func TestGetAllProductController(t *testing.T) {
	// create database connection and create controller
	setup()
	config := config.GetConfig()
	db := util.MysqlDatabaseConnection(config)
	productModel := models.NewProductModel(db)
	productController := NewProductController(productModel)

	// // setting controller
	e := echo.New()

	// setting controller
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath("/api/products")

	if err := productController.GetAllProductController(context); err != nil {
		t.Errorf("Should'nt get error, get error: %s", err)
	}

	type Response struct {
		Code    int              `json:"code"`
		Message string           `json:"message"`
		Data    []models.Product `json:"data"`
	}
	var response Response
	json.Unmarshal(res.Body.Bytes(), &response)
	// var productList []models.Product
	// json.Unmarshal(res.Body.Bytes(), &productList)
	// fmt.Println(productList)

	assert.Equal(t, 200, res.Code)
	assert.Equal(t, "Success Get All Product", response.Message)
	assert.NotEmpty(t, response.Data)
	assert.Equal(t, 1, len(response.Data))
	assert.Equal(t, "Product A", response.Data[0].Name)
	assert.Equal(t, 10000, response.Data[0].Price)
	assert.Equal(t, 100, response.Data[0].Stock)

}

func TestGetProductController(t *testing.T) {
	// create database connection and create controller
	config := config.GetConfig()
	db := util.MysqlDatabaseConnection(config)
	productModel := models.NewProductModel(db)
	productController := NewProductController(productModel)

	// setting controller
	e := echo.New()

	// setting controller
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath("/api/products/:id")
	context.SetParamNames("id")
	context.SetParamValues("1")

	if err := productController.GetProductController(context); err != nil {
		t.Errorf("Should'nt get error, get error: %s", err)
	}

	type Response struct {
		Code    int            `json:"code"`
		Message string         `json:"message"`
		Data    models.Product `json:"data"`
	}
	var response Response
	json.Unmarshal(res.Body.Bytes(), &response)
	// var productList []models.Product
	// json.Unmarshal(res.Body.Bytes(), &productList)
	// fmt.Println(productList)

	assert.Equal(t, 200, res.Code)
	assert.Equal(t, "Success Get All Product", response.Message)
	assert.NotEmpty(t, response.Data)
	assert.Equal(t, "Product A", response.Data.Name)
	assert.Equal(t, 10000, response.Data.Price)
	assert.Equal(t, 100, response.Data.Stock)

}

func TestPostProductController(t *testing.T) {
	config := config.GetConfig()
	db := util.MysqlDatabaseConnection(config)
	productModel := models.NewProductModel(db)
	productController := NewProductController(productModel)
	customerModel := models.NewCustomerModel(db)
	customerController := NewCustomerController(customerModel)

	// // setting controller
	e := echo.New()

	// login
	reqBodyLogin, _ := json.Marshal(models.Customer{Email: "ilham@gmail.com", Password: "pass123"})
	loginreq := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBodyLogin))
	// loginreq.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
	loginreq.Header.Set("Content-Type", "application/json")
	loginres := httptest.NewRecorder()
	logincontext := e.NewContext(loginreq, loginres)
	logincontext.SetPath("/api/login")

	if err := customerController.LoginCustomerController(logincontext); err != nil {
		t.Errorf("Should'nt get error, get error: %s", err)
	}

	var c models.Customer
	json.Unmarshal(loginres.Body.Bytes(), &c)

	// testing stuff

	assert.Equal(t, 200, loginres.Code)
	assert.NotEqual(t, "", c.Token)

	token := c.Token

	// input controller
	reqBodyPost, _ := json.Marshal(map[string]interface{}{
		"name":        "Product B",
		"price":       5000,
		"stock":       5,
		"category_id": 1,
	})

	// setting controller
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBodyPost))
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath("/api/products")

	if err := productController.PostProductController(context); err != nil {
		t.Errorf("Should'nt get error, get error: %s", err)
	}

	type Response struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}
	var response Response

	json.Unmarshal(res.Body.Bytes(), &response)

	// testing stuff
	assert.Equal(t, 200, res.Code)
	assert.Equal(t, "Success Add Product", response.Message)
}

func TestUpdateProductController(t *testing.T) {
	// create database connection and create controller
	config := config.GetConfig()
	db := util.MysqlDatabaseConnection(config)
	userModel := models.NewCustomerModel(db)
	customerController := NewCustomerController(userModel)
	productModel := models.NewProductModel(db)
	productController := NewProductController(productModel)

	// setting controller
	e := echo.New()

	// login
	reqBodyLogin, _ := json.Marshal(models.Customer{Email: "ilham@gmail.com", Password: "pass123"})
	loginreq := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBodyLogin))
	// loginreq.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
	loginreq.Header.Set("Content-Type", "application/json")
	loginres := httptest.NewRecorder()
	logincontext := e.NewContext(loginreq, loginres)
	logincontext.SetPath("/api/login")

	if err := customerController.LoginCustomerController(logincontext); err != nil {
		t.Errorf("Should'nt get error, get error: %s", err)
	}

	var c models.Customer
	json.Unmarshal(loginres.Body.Bytes(), &c)

	// testing stuff
	assert.Equal(t, 200, loginres.Code)
	assert.NotEqual(t, "", c.Token)

	token := c.Token

	// input controller
	reqBodyPost, _ := json.Marshal(map[string]interface{}{
		"name":        "Product B Update",
		"price":       6000,
		"stock":       6,
		"category_id": 1,
	})

	// setting controller
	req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(reqBodyPost))
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath("/products/:id")
	context.SetParamNames("id")
	context.SetParamValues("2")

	productController.UpdateProductController(context)

	type Response struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}
	var response Response

	json.Unmarshal(res.Body.Bytes(), &response)

	// testing stuff
	assert.Equal(t, 200, res.Code)
	assert.Equal(t, "Success Edit Category", response.Message)
}

func TestDeleteProductController(t *testing.T) {
	// create database connection and create controller
	config := config.GetConfig()
	db := util.MysqlDatabaseConnection(config)
	userModel := models.NewCustomerModel(db)
	customerController := NewCustomerController(userModel)
	productModel := models.NewProductModel(db)
	productController := NewProductController(productModel)

	// setting controller
	e := echo.New()

	// login
	reqBodyLogin, _ := json.Marshal(models.Customer{Email: "ilham@gmail.com", Password: "pass123"})
	loginreq := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBodyLogin))
	// loginreq.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
	loginreq.Header.Set("Content-Type", "application/json")
	loginres := httptest.NewRecorder()
	logincontext := e.NewContext(loginreq, loginres)
	logincontext.SetPath("/api/login")

	if err := customerController.LoginCustomerController(logincontext); err != nil {
		t.Errorf("Should'nt get error, get error: %s", err)
	}

	var c models.Customer
	json.Unmarshal(loginres.Body.Bytes(), &c)

	// testing stuff
	assert.Equal(t, 200, loginres.Code)
	assert.NotEqual(t, "", c.Token)

	token := c.Token

	// setting controller
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath("/products/:id")
	context.SetParamNames("id")
	context.SetParamValues("2")

	productController.DeleteProductController(context)

	type Response struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}
	var response Response

	json.Unmarshal(res.Body.Bytes(), &response)

	// testing stuff
	assert.Equal(t, 200, res.Code)
	assert.Equal(t, "Success Delete Product", response.Message)
}
