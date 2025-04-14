package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type NoteController struct {}

func (c *NoteController) InitNoteControllerRoutes(router *gin.Engine) {
	notes := router.Group("/product")
	notes.POST("/", c.CreateProduct)
	notes.GET("/", c.GetAllProduct)
	notes.GET("/:id/:category", c.SingleProduct)
}

func (c *NoteController) GetAllProduct(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get All Product",
	})
}

func (c *NoteController) SingleProduct(ctx *gin.Context) {
	Productid := ctx.Param("id")
	category := ctx.Param("category")
	ctx.JSON(http.StatusOK, gin.H{
		"productid": Productid,
		"category":  category,
	})
}
func (c *NoteController) CreateProduct(ctx *gin.Context) {
	type ProductList struct {
		ProductName string `json:"productname" binding:"required"`
		Category    string `json:"category"`
		Image       string `json:"image"`
		Rating      string `json:"rating"`
	}
	var newRequest ProductList
	if err := ctx.BindJSON(&newRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"ProductName": newRequest.ProductName,
		"category":    newRequest.Category,
		"image":       newRequest.Image,
		"rating":      newRequest.Rating,
		"message": gin.H{
			"success": "Product added successfully",
		},
	})
}