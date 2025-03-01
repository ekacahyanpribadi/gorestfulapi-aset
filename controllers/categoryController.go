package controllers

import (
	//"errors"
	//"net/http"
	"errors"
	"net/http"
	"restfulapi/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// function get error message
func GetErrorMsgKategori(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	}
	return "Unknown error"
}

// Get Kategori
func GetKategori(c *gin.Context) {
	reqToken := c.Request.Header.Get("token")
	var token []models.Token_access
	goSql := models.DB.Raw("SELECT * FROM production.token_access WHERE 1=1 AND token='" + reqToken + "'").Find(&token)
	//var tokenDb = token[0].Token
	if goSql.RowsAffected < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token access not found!"})
		return
	}

	//get data from database using model
	var category []models.Kategori_aset
	//models.DB.Find(&category)

	models.DB.Raw("SELECT * FROM kategori_aset WHERE 1=1 ORDER BY ins_date DESC").Find(&category)

	//return json
	c.JSON(200, gin.H{
		"success": true,
		"message": "Lists Data Category",
		"data":    category,
	})
}

// Create Kategori
func CreateKategori(c *gin.Context) {
	reqToken := c.Request.Header.Get("token")
	var token []models.Token_access
	goSql := models.DB.Raw("SELECT * FROM production.token_access WHERE 1=1 AND token='" + reqToken + "'").Find(&token)
	//var tokenDb = token[0].Token
	if goSql.RowsAffected < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token access not found!"})
		return
	}

	//validate input
	var input models.ValidateKategoriInput
	if err := c.ShouldBindJSON(&input); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = ErrorMsg{fe.Field(), GetErrorMsg(fe)}
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
		}
		return
	}

	//create post
	postKategori := models.Kategori_aset{
		Id_kategori:                input.Id_kategori,
		Kategori:                   input.Kategori,
		Sub_kategori:               input.Sub_kategori,
		Keterangan:                 input.Keterangan,
		Jumlah_aset:                input.Jumlah_aset,
		Status_kategori:            input.Status_kategori,
		Masa_manfaat:               input.Masa_manfaat,
		Penyusutan_persen_pertahun: input.Penyusutan_persen_pertahun,
		Ins_user:                   input.Ins_user,
		Ins_date:                   input.Ins_date,
		Upd_user:                   input.Upd_user,
		Upd_date:                   input.Upd_date,
	}

	models.DB.Create(&postKategori)

	//return response json
	c.JSON(201, gin.H{
		"success": true,
		"message": "Id: " + postKategori.Id_kategori + " Created Successfully",
		"data":    postKategori,
	})
}

// get kategori by id_kategori
func GetKategoriId(c *gin.Context) {
	reqToken := c.Request.Header.Get("token")
	var token []models.Token_access
	goSqlToken := models.DB.Raw("SELECT * FROM production.token_access WHERE 1=1 AND token='" + reqToken + "'").Find(&token)
	//var tokenDb = token[0].Token
	if goSqlToken.RowsAffected < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token access not found!"})
		return
	}

	var kategori models.Kategori_aset
	if err := models.DB.Where("id_kategori = ?", c.Param("id_kategori")).First(&kategori).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "Detail Data Kategori By ID : " + c.Param("id_kategori"),
		"data":    kategori,
	})
}

// update kategori start
func UpdateKategori(c *gin.Context) {
	reqToken := c.Request.Header.Get("token")
	var token []models.Token_access
	goSqlToken := models.DB.Raw("SELECT * FROM production.token_access WHERE 1=1 AND token='" + reqToken + "'").Find(&token)
	//var tokenDb = token[0].Token
	if goSqlToken.RowsAffected < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token access not found!"})
		return
	}

	var kategori models.Kategori_aset
	if err := models.DB.Where("id_kategori = ?", c.Param("id_kategori")).First(&kategori).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	//validate input
	var input models.ValidateKategoriInput
	if err := c.ShouldBindJSON(&input); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = ErrorMsg{fe.Field(), GetErrorMsg(fe)}
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
		}
		return
	}

	//update kategori to db
	models.DB.Model(&kategori).Updates(input)

	c.JSON(200, gin.H{
		"success": true,
		"message": "Kategori: " + c.Param("id_kategori") + "Updated Successfully",
		"data":    kategori,
	})
}

// delete kategori
func DeleteKategori(c *gin.Context) {
	reqToken := c.Request.Header.Get("token")
	var token []models.Token_access
	goSqlToken := models.DB.Raw("SELECT * FROM production.token_access WHERE 1=1 AND token='" + reqToken + "'").Find(&token)
	//var tokenDb = token[0].Token
	if goSqlToken.RowsAffected < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token access not found!"})
		return
	}

	var kategori models.Kategori_aset
	if err := models.DB.Where("id_kategori = ?", c.Param("id_kategori")).First(&kategori).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	//delete post
	models.DB.Delete(&kategori)

	c.JSON(200, gin.H{
		"success": true,
		"message": "Kategori: " + c.Param("id_kategori") + " Deleted Successfully",
	})
}
