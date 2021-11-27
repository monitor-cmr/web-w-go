package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
    ID     int
    Title  string
    Artist string
    Price  float64
}

// albums slice to seed record album data.
var albums = []album{
    {ID: 1, Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: 2, Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: 3, Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
    router := gin.Default()

    // Albums - Start
    // curl -XGET http://127.0.0.1:8080/albums
    router.GET("/albums", getAlbums)

    // curl -XPOST -d @createAlbum.json http://127.0.0.1:8080/albums
    router.POST("/albums", postAlbums)

    // curl -XGET http://127.0.0.1:8080/albums/[1-n] - method GET
    router.GET("/albums/:id", getAlbumsByID)

    // curl -XPUT http://127.0.0.1:8080/albums/[1-n] - method PUT
    router.PUT("/albums/:id", putAlbums)

    // http://127.0.0.1:8080/albums/[1-n] - method DELETE
    router.DELETE("/albums/:id", deleteAlbums)

    // Albums - End

    router.Run(":8080")
}

// getAlbums ...
func getAlbums(c *gin.Context) {
    // c.IndentedJSON(http.StatusOK, gin.H{"data": "Get all albums"})
    c.IndentedJSON(http.StatusOK, gin.H{"data": albums})
}

// postAlbums ...
func postAlbums(c *gin.Context) {
    var newAlbum album

    // Call BindJSON to bind the received JSON to newAlbum.
    if err := c.BindJSON(&newAlbum); err != nil {
        c.IndentedJSON(http.StatusBadRequest, gin.H{"data": "Bad Request"})
        return
    }

    // Get a length of albums
    lenAlbum := len(albums)

    // Get a last album's ID
    // - index start 0
    // - a last album = length - 1 => lenAlbum - 1
    newAlbum.ID = albums[lenAlbum-1].ID + 1

    // Add the new album to the slice.
    albums = append(albums, newAlbum)

    // response to user request
    c.IndentedJSON(http.StatusCreated, gin.H{"data": newAlbum.ID})
    // c.IndentedJSON(http.StatusCreated, gin.H{"data": "Create albums"})
}

// getAlbumsByID ...
func getAlbumsByID(c *gin.Context) {
    // c.IndentedJSON(http.StatusOK, gin.H{"data": "Get an albums"})
}

// putAlbums ...
func putAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, gin.H{"data": "Update an albums"})
}

// deleteAlbums ...
func deleteAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, gin.H{"data": "Delete an albums"})
}

