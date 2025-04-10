package main

import (
	"fmt"
	"net/http"
	"os"
	"project/constants"
	"project/login"
	"project/products"

	"github.com/gin-gonic/gin"
)

func main() {
	login.Instantiate()
	r := gin.Default()

	r.GET("/api", func(c *gin.Context) {
		resp, _ := products.GetProducts()
		c.JSON(http.StatusOK, resp)
	})

	r.GET("/api/staples", func(c *gin.Context) {
		resp, _ := products.GetStaples()
		c.JSON(http.StatusOK, resp)
	})

	r.GET("/api/milk-products", func(c *gin.Context) {
		resp, _ := products.GetProducts(constants.MILK_PRODUCTS)
		c.JSON(http.StatusOK, resp)
	})

	r.GET("/api/snacks", func(c *gin.Context) {
		resp, _ := products.GetProducts(constants.SNACKS)
		c.JSON(http.StatusOK, resp)
	})

	r.GET("/api/grocery", func(c *gin.Context) {
		resp, _ := products.GetProducts(constants.GROCERIES)
		c.JSON(http.StatusOK, resp)
	})

	r.POST("/login", func(c *gin.Context) {

		var req constants.LoginRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		// fmt.Println(email, password)

		exists, val := login.Verify(req.Email, req.Password)
		if exists {
			c.JSON(http.StatusOK, val)
		} else {
			c.JSON(http.StatusBadRequest, "No")
		}
	})

	r.POST("/signup", func(c *gin.Context) {

		var req constants.User
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}
		if req.Email == "" || req.Password == "" || req.FirstName == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}
		login.CreateUser(req)
		c.JSON(http.StatusOK, req)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "10000"
	}
	// r.Run(":8080")

	err := r.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		panic(err)
	}
}
