package main

import (
	"net/http"
	"webapp1/handler"
	"webapp1/middleware"

	"github.com/gin-gonic/gin"
)

// buat service yg mengembalikan resep (nama dan bahan2) dan harga total yg dibutuhkan:
// 1. ambil data dari https://api.spoonacular.com/recipes/random
// 2. bikin server
// 3. bikin endpoint dan handler nya
// 4. add middleware api key
// 5. bikin endpoint login untuk validasi user dan return JWT
// 6. bikin middleware untuk validate user token

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
			"data":    nil,
		})
	})
	r.Use(middleware.ValidateAPIKey())
	r.POST("/login", handler.Login)
	r.Use(middleware.ValidateUserToken())
	r.GET("/recipes", handler.HandlerGetRecipe)

	r.Run(":8080")
}
