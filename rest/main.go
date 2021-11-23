package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    // Albums - Start
    // curl -XGET http://127.0.0.1:8080/albums
    router.GET("/albums", getAlbums)

    // curl -XPOST http://127.0.0.1:8080/albums
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
    c.IndentedJSON(http.StatusOK, gin.H{"data": "Get all albums"})
}

// postAlbums ...
func postAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, gin.H{"data": "Create albums"})
}

// getAlbumsByID ...
func getAlbumsByID(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, gin.H{"data": "Get an albums"})
}

// putAlbums ...
func putAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, gin.H{"data": "Update an albums"})
}

// deleteAlbums ...
func deleteAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, gin.H{"data": "Delete an albums"})
}

