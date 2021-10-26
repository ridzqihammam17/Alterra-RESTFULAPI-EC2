package controllers

import (
	"altastore/config"
	"altastore/models"
	"altastore/util"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestRegisterCustomerController(t *testing.T) {
	setup()
	config := config.GetConfig()
	db := util.MysqlDatabaseConnection(config)
	customerModel := models.NewCustomerModel(db)
	customerController := NewCustomerController(customerModel)

	e := echo.New()

	reqBodyPost, _ := json.Marshal(map[string]string{
		"name":     "Dummy Customer",
		"email":    "dummy_customer@gmail.com",
		"password": "password123",
	})

	// setting controller
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBodyPost))
	// req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath("/api/register")

	if err := customerController.RegisterCustomerController(context); err != nil {
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
	assert.Equal(t, "Success Register", response.Message)
}

func TestLoginCustomerController(t *testing.T) {
	config := config.GetConfig()
	db := util.MysqlDatabaseConnection(config)
	customerModel := models.NewCustomerModel(db)
	customerController := NewCustomerController(customerModel)

	// // setting controller
	e := echo.New()
	reqBodyLogin, _ := json.Marshal(models.Customer{Email: "dummy_customer@gmail.com", Password: "password123"})
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
	// assert.NotEqual(t, "", c.Token)
}
