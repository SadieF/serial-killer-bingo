package main

import (
	"net/http"

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

//Bingo containes information about a single serial killer trait
type Bingo struct {
	ID int `json:"id" binding:"required"`
	Marked bool `json:"marked"`
	Trait string `json:"trait" binding:"required"`
}

//List of traits
var Bingo = []Bingo{
	Bingo{1, false, "Drives a van"},
	Bingo{1, false, "Arson"},
	Bingo{1, false, "Childhood Head Injury"},
	Bingo{1, false, "Abandonment"},
	Bingo{1, false, "Mother a prostitute"},
	Bingo{1, false, "Bunker"},
	Bingo{1, false, "Someone is in foster care"},
	Bingo{1, false, "Cannibalism"},
	Bingo{1, false, "Caucasian"},
	Bingo{1, false, "Tortured or killed pets"
	Bingo{1, false, "Weak or absent father"},
	Bingo{1, false, "FREE or Prison Break"},
	Bingo{1, false, "Someone has a missing limb"},
	Bingo{1, false, "Bedwetting"},
	Bingo{1, false, "Sexual Deviance"},
	Bingo{1, false, "Three first names"},
	Bingo{1, false, "Dancing"},
	Bingo{1, false, "Religious fanatics"},
	Bingo{1, false, "Oddly specific preferences"},
	Bingo{1, false, "Broom as a weapon"},
	Bingo{1, false, "Two in one night"},
	Bingo{1, false, "Mommy issues"},
	Bingo{1, false, "High IQ"},
	Bingo{1, false, "Catchy nickname"},
}

func main() {
	// Leave this untouched
}

// BingoHandler returns the bingo card of serial killer traits
func BingoHandler(c *gin.Context) {
	c.Header("Content-type", "application/json")
	c.JSON(http.StatusOK, gin.H{
		"message": "Bingo handler not implemented yet",
	})
}

// MarkBingo marks the trait of a serial killer
func MarkBingo(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{
		"message": "MarkBingo handler not implemented yet",
	})
}




}