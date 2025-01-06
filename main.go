package main

import (
	"embed"
	"fmt"
	"log"

	"net/http"

	"github.com/gin-gonic/gin"
)

// go:embed public
var f embed.FS

func main() {
	router := gin.Default()

	// router.StaticFile("/", "./public/index.html")

	// router.Static("/public", "./public")

	// router.StaticFS("/fs", http.FileSystem(http.FS(f)))

	router.GET("/employee", func (c *gin.Context) {
		fmt.Println("In employee handler")
		c.File("./public/employee.html")
	})

	router.POST("/employee/submit", func (c *gin.Context) {
		c.String(http.StatusOK, "PTO Submitted Successfully!")
	})

	router.GET("/employee/:username/*rest", func (c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"username": c.Param("username"),
			"rest": c.Param("rest"),
		})
	})

	log.Fatal(router.Run(":5000"))
}