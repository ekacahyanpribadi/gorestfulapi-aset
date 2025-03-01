package main

import (
	"restfulapi/controllers"
	"restfulapi/middlewares"
	"restfulapi/models"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

var logger *log.Logger

func init() {
	logger = logrus.New()
	logger.SetLevel(log.InfoLevel)

}

func main() {
	//gin.SetMode(gin.ReleaseMode)

	// Add logging middleware
	r := gin.Default()
	r.Use(middlewares.RequestLoggingMiddleware(logger))

	//inisialiasai Gin
	router := gin.Default()

	//panggil koneksi database
	models.ConnectDatabase()

	//membuat route dengan method GET
	router.GET("/", func(c *gin.Context) {

		//return response JSON
		c.JSON(200, gin.H{
			"message": "golang server ready!",
			"info":    "restfull api aset",
		})
	})

	//membuat route get all posts
	router.GET("/api/posts", controllers.FindPosts)

	//membuat route store post
	router.POST("/api/posts", controllers.StorePost)

	//membuat route detail post
	router.GET("/api/posts/:id", controllers.FindPostById)

	//membuat route update post
	router.PUT("/api/posts/:id", controllers.UpdatePost)

	//membuat route delete post
	router.DELETE("/api/posts/:id", controllers.DeletePost)

	//membuat route get post test
	router.GET("/api/test", controllers.TestGet)

	//ROUTE CATEGORY
	router.GET("/api/category/", controllers.GetKategori)
	router.POST("/api/category/", controllers.CreateKategori)
	router.GET("/api/category/:id_kategori", controllers.GetKategoriId)
	router.PUT("/api/category/:id_kategori", controllers.UpdateKategori)
	router.DELETE("/api/category/:id_kategori", controllers.DeleteKategori)

	//mulai server dengan port 3000
	router.Run(":3000")
}
