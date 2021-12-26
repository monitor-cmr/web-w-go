package router

import (
    "database/sql"

    "github.com/gin-gonic/gin"
    "github.com/monitor-cmr/web-w-go/handler"
    "github.com/monitor-cmr/web-w-go/service"
    "github.com/monitor-cmr/web-w-go/storage/mysql"
)

// Router ...
// func Router() *gin.Engine {
func Router(db *sql.DB) *gin.Engine {
    router := gin.Default()

    // Albums - Start

    // // Create Memory storage
    // store := memory.NewAlbumMemory()

    // // Create MySQL storage
    store := mysql.NewAlbumMySQL(db)

    // Create AlbumService
    svcAlbum := service.NewAlbumService(store)

    // Create a new AlbumHandler
    al := handler.NewAlbumHandler(svcAlbum)

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

