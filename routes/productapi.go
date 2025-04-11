package routes

import (
	"github.com/Mzammilishtiaq/gotask1/controller"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	r := gin.Default()
	controller := controller.NoteController{}
	controller.InitNoteControllerRoutes(r)
	
	r.Run(":8080")
}