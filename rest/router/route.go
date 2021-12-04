package router

import (
    "github.com/gin-gonic/gin"
    "github.com/monitor-cmr/web-w-go/handler"
)

// Router ...
func Router() *gin.Engine {
    router := gin.Default()

    // Albums - Start

    // Create a new AlbumHandler
    al := handler.NewAlbumHandler()

    // curl -XGET http://127.0.0.1:8080/albums
    router.GET("/albums", al.GetAlbums)

    // curl -XPOST -d @createAlbum.json http://127.0.0.1:8080/albums
    router.POST("/albums", al.PostAlbums)

    // curl -XGET http://127.0.0.1:8080/albums/[1-n] - method GET
    router.GET("/albums/:id", al.GetAlbumsByID)

    // curl -XPUT -d @updateAlbum.json http://127.0.0.1:8080/albums/[1-n] - method PUT
    router.PUT("/albums/:id", al.PutAlbums)

    // curl -XDELETE http://127.0.0.1:8080/albums/[1-n] - method DELETE
    router.DELETE("/albums/:id", al.DeleteAlbums)

    // Albums - End

    return router
}

