package controllers

import (
	"altastore/config"
	"altastore/models"
	"altastore/util"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	echo "github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetAllCategoryController(t *testing.T) {
	// create database connection and create controller
	setup()
	config := config.GetConfig()
	db := util.MysqlDatabaseConnection(config)
	categoryModel := models.NewCategoryModel(db)
	categoryController := NewCategoryController(categoryModel)
	customerModel := models.NewCustomerModel(db)
	customerController := NewCustomerController(customerModel)

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
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath("/api/categories")

	if err := categoryController.GetAllCategoryController(context); err != nil {
		t.Errorf("Should'nt get error, get error: %s", err)
	}

	type Response struct {
		Code    int               `json:"code"`
		Message string            `json:"message"`
		Data    []models.Category `json:"data"`
	}
	var response Response
	json.Unmarshal(res.Body.Bytes(), &response)
	// var productList []models.Category
	// json.Unmarshal(res.Body.Bytes(), &productList)
	// fmt.Println(productList)

	assert.Equal(t, 200, res.Code)
	assert.Equal(t, "Success Get All Category", response.Message)
	assert.NotEmpty(t, response.Data)
	assert.Equal(t, 1, len(response.Data))
	assert.Equal(t, "Category A", response.Data[0].Name)
}

func TestGetCategoryController(t *testing.T) {
	// create database connection and create controller
	config := config.GetConfig()
	db := util.MysqlDatabaseConnection(config)
	categoryModel := models.NewCategoryModel(db)
	categoryController := NewCategoryController(categoryModel)
	customerModel := models.NewCustomerModel(db)
	customerController := NewCustomerController(customerModel)

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
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath("/api/categories/:id")
	context.SetParamNames("id")
	context.SetParamValues("1")

	if err := categoryController.GetCategoryController(context); err != nil {
		t.Errorf("Should'nt get error, get error: %s", err)
	}

	type Response struct {
		Code    int             `json:"code"`
		Message string          `json:"message"`
		Data    models.Category `json:"data"`
	}
	var response Response
	json.Unmarshal(res.Body.Bytes(), &response)
	// var productList []models.Category
	// json.Unmarshal(res.Body.Bytes(), &productList)
	// fmt.Println(productList)

	assert.Equal(t, 200, res.Code)
	assert.Equal(t, "Success Get Category", response.Message)
	assert.NotEmpty(t, response.Data)
	assert.Equal(t, "Category A", response.Data.Name)

}

func TestPostCategoryController(t *testing.T) {
	config := config.GetConfig()
	db := util.MysqlDatabaseConnection(config)
	categoryModel := models.NewCategoryModel(db)
	categoryController := NewCategoryController(categoryModel)
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
	reqBodyPost, _ := json.Marshal(map[string]string{
		"name": "Category B",
	})

	// setting controller
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBodyPost))
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath("/api/categories")

	if err := categoryController.AddCategoryController(context); err != nil {
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
	assert.Equal(t, "Success Add Category", response.Message)
}

func TestUpdateCategoryController(t *testing.T) {
	// create database connection and create controller
	config := config.GetConfig()
	db := util.MysqlDatabaseConnection(config)
	userModel := models.NewCustomerModel(db)
	customerController := NewCustomerController(userModel)
	categoryModel := models.NewCategoryModel(db)
	categoryController := NewCategoryController(categoryModel)

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
	reqBodyPost, _ := json.Marshal(map[string]string{
		"name": "Category B Update",
	})

	// setting controller
	req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(reqBodyPost))
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath("/categories/:id")
	context.SetParamNames("id")
	context.SetParamValues("2")

	categoryController.EditCategoryController(context)

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

func TestDeleteCategoryController(t *testing.T) {
	// create database connection and create controller
	config := config.GetConfig()
	db := util.MysqlDatabaseConnection(config)
	userModel := models.NewCustomerModel(db)
	customerController := NewCustomerController(userModel)
	categoryModel := models.NewCategoryModel(db)
	categoryController := NewCategoryController(categoryModel)

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
	context.SetPath("/categories/:id")
	context.SetParamNames("id")
	context.SetParamValues("2")

	categoryController.DeleteCategoryController(context)

	type Response struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}
	var response Response

	json.Unmarshal(res.Body.Bytes(), &response)

	// testing stuff
	assert.Equal(t, 200, res.Code)
	assert.Equal(t, "Success Delete Category", response.Message)
}
