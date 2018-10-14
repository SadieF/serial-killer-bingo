package main

import (
	"net/http"
	"strconv"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

//Bingo containes information about a single serial killer trait
type Bingo struct {
	ID     int    `json:"id" binding:"required"`
	Marked bool   `json:"marked"`
	Trait  string `json:"trait" binding:"required"`
}

//Bingos is a list of traits
var Bingos = []Bingo{
	Bingo{1, false, "Drives a van"},
	Bingo{2, false, "Arson"},
	Bingo{3, false, "Childhood Head Injury"},
	Bingo{4, false, "Abandonment"},
	Bingo{5, false, "Mother a prostitute"},
	Bingo{6, false, "Bunker"},
	Bingo{7, false, "Someone is in foster care"},
	Bingo{8, false, "Cannibalism"},
	Bingo{9, false, "Caucasian"},
	Bingo{10, false, "Tortured or killed pets"},
	Bingo{11, false, "Weak or absent father"},
	Bingo{12, false, "FREE or Prison Break"},
	Bingo{13, false, "Someone has a missing limb"},
	Bingo{14, false, "Bedwetting"},
	Bingo{15, false, "Sexual Deviance"},
	Bingo{16, false, "Three first names"},
	Bingo{17, false, "Dancing"},
	Bingo{18, false, "Religious fanatics"},
	Bingo{19, false, "Oddly specific preferences"},
	Bingo{20, false, "Broom as a weapon"},
	Bingo{21, false, "Two in one night"},
	Bingo{22, false, "Mommy issues"},
	Bingo{23, false, "High IQ"},
	Bingo{24, false, "Catchy nickname"},
}

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
