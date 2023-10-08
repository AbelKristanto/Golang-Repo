package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// [1] CREATE STRUCT FOR DATASET THAT WILL BE USED
// Struct Data Artist
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// Data inside struct
var albums = []album{
	{"1", "Blue Train", "John Coltrane", 56.99},
	{"2", "Jeru", "Gerry Mulligan", 17.99},
	{"3", "Sarah Vaughan", "Sarah Vaughan", 39.99},
}

// [2] HANDLER TO ACCESS ALL ITEMS

// get albums responds as JSON
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// add an album from JSON
func postAlbums(c *gin.Context) {
	var newAlbum album

	// CALL BindJSON to receive JSON
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)

}

// GET Album by ID as PARAMETER
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// loop over the data and looking for
	// an album show ID value matchers the parameter
	for _, name := range albums {
		if name.ID == id {
			c.IndentedJSON(http.StatusOK, name)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.Run("localhost:8080")
}
