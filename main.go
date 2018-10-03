package main

import (
	"net/http"
	"strconv"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	//Set router as default one shipped with gin
	router := gin.Default()

	//Serve frontend static files
	router.Use(static.Serve("/", static.LocalFile("./views", true)))

	//Setup route group for the API
	api := router.Group("api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}
	//Api has two routes for now:
	//Bingo returns the bingo card of serial killer traits
	//Bingo/mark/:bingoID captures when a particular trait is marked
	api.GET("/bingo", BingoHandler)
	api.POST("/bingo/mark/:bingoID", MarkBingo)

	router.Run(":3000")
}

// BingoHandler returns the bingo card of serial killer traits
func BingoHandler(c *gin.Context) {
	c.Header("Content-type", "application/json")
	c.JSON(http.StatusOK, Bingos)
}

// MarkBingo marks the trait of a serial killer
func MarkBingo(c *gin.Context) {
	// Confirms bingo id is valid
	if bingoid, err := strconv.Atoi(c.Param("bingoID")); err == nil {
		for i := 0; i < len(Bingos); i++ {
			if Bingos[i].ID == bingoid {
				if Bingos[i].Marked == false {
					Bingos[i].Marked = true
				} else {
					Bingos[i].Marked = false
				}
			}
		}
		c.JSON(http.StatusOK, &Bingos)
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}
