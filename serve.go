package main

import (
	"github.com/Mzammilishtiaq/gotask1/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	controller := controller.NoteController{}
	controller.InitNoteControllerRoutes(r)
	
	r.Run(":8080")
}