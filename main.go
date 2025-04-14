package main

import (
	"fmt"
	"github.com/Mzammilishtiaq/gotask1/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("start project")
	router := gin.Default()
	routes.Routes(router)
	
}