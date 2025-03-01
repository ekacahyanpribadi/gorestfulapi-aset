package controllers

import (
	//"errors"
	//"net/http"
	"restfulapi/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// type validation post input
type ValidateCategoriInput struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

// type error message
type ErrorMsgCategory struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// function get error message
func GetErrorMsgCategory(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	}
	return "Unknown error"
}

var istokenC = "BENAR"

// get all posts
func GetCategory(c *gin.Context) {

	cekToken := c.Request.Header.Get("token")
	if cekToken != istokenC {
		c.JSON(400, gin.H{
			"message": "Header invalid!",
		})
		return
	}

	//get data from database using model
	var category []models.Category
	models.DB.Find(&category)

	//return json
	c.JSON(200, gin.H{
		"success": true,
		"message": "Lists Data Category",
		"data":    category,
	})
}
