package controllers

import (
	"errors"
	"math/rand"
	"net/http"
	"restfulapi/models"
	"strings"
	"time"

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

const charset = "012345678901234567890123456789012345678901234567890123456789"
const charset2 = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randomNumber(n int) string {
	sb := strings.Builder{}
	sb.Grow(n)
	for i := 0; i < n; i++ {
		sb.WriteByte(charset[rand.Intn(len(charset))])
	}
	return sb.String()
}

func randomString(n int) string {
	sb := strings.Builder{}
	sb.Grow(n)
	for i := 0; i < n; i++ {
		sb.WriteByte(charset2[rand.Intn(len(charset2))])
	}
	return sb.String()
}

// Get Kategori
func GetKategori(c *gin.Context) {
	//cek token access start
	reqToken := c.Request.Header.Get("token")
	var token []models.Token_access
	goSql := models.DB.Raw("SELECT * FROM production.token_access WHERE 1=1 AND token='" + reqToken + "'").Find(&token)
	//var tokenDb = token[0].Token
	if goSql.RowsAffected < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token access not found!"})
		return
	}
	//cek token access end

	//get data from database using model
	var category []models.Kategori_aset

	//models.DB.Find(&category)

	getSql := models.DB.Raw("SELECT * FROM kategori_aset WHERE 1=1 ORDER BY ins_date DESC").Find(&category)
	if err := getSql.Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

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

	/*c.JSON(http.StatusOK, gin.H{
		"messege": "Test!",
		"data":    "Token access found!",
	})
	return*/

	currentTime := time.Now().UTC().Add(7 * time.Hour)
	const (
		df = "20060102"
	)
	currentTimex := currentTime.Format(df)
	genId := currentTimex + randomNumber(6) + randomString(6)

	//create post
	postKategori := models.Kategori_aset{
		Id_kategori:                genId,
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

	c.JSON(200, gin.H{
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

// Create Kategori
func CreateLog(c *gin.Context) {

	//create post
	postLoghit := models.Kategori_aset{
		Id_kategori:  "input.Id_kategori",
		Kategori:     "input.Kategori",
		Sub_kategori: " input.Sub_kategori",
		Keterangan:   "input.Keterangan",
	}

	models.DB.Create(&postLoghit)

	//return response json
	c.JSON(201, gin.H{
		"success": true,
		"message": "Id: " + postLoghit.Id_kategori + " Created Successfully",
		"data":    postLoghit,
	})
}
