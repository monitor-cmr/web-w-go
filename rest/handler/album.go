package handler

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "github.com/monitor-cmr/web-w-go/domain"
)

// AlbumHandler ...
type AlbumHandler struct {
    repo domain.AlbumRepository
}

// NewAlbumHandler ...
func NewAlbumHandler(re domain.AlbumRepository) *AlbumHandler {
    return &AlbumHandler{
        repo: re,
    }
}

// GetAlbums ...
func (al AlbumHandler) GetAlbums(c *gin.Context) {
    // c.IndentedJSON(http.StatusOK, gin.H{"data": "Get all albums"})
    albums, err := al.repo.SelectAll()
    if err != nil {
        c.IndentedJSON(http.StatusInternalServerError, gin.H{"data": "Server Internal Error"})
        return
    }
    c.IndentedJSON(http.StatusOK, gin.H{"data": albums})
}

// PostAlbums ...
func (al AlbumHandler) PostAlbums(c *gin.Context) {
    var newAlbum domain.Album

    // Call BindJSON to bind the received JSON to newAlbum.
    if err := c.BindJSON(&newAlbum); err != nil {
        c.IndentedJSON(http.StatusBadRequest, gin.H{"data": "Bad Request"})
        return
    }

    id, err := al.repo.Save(newAlbum)
    if err != nil {
        c.IndentedJSON(http.StatusInternalServerError, gin.H{"data": "Server Internal Error"})
        return
    }
    // response to user request
    c.IndentedJSON(http.StatusCreated, gin.H{"data": id})
    // c.IndentedJSON(http.StatusCreated, gin.H{"data": "Create albums"})
}

// GetAlbumsByID ...
func (al AlbumHandler) GetAlbumsByID(c *gin.Context) {
    // /albums/:id => c.Param("id") => string value
    tempID := c.Param("id")

    // Convert string to int
    id, err := strconv.Atoi(tempID)

    if err != nil {
        c.IndentedJSON(http.StatusBadRequest, gin.H{"data": "Bad Request"})
        return
    }

    album, err := al.repo.Select(id)
    if err != nil {
        c.IndentedJSON(http.StatusNotFound, gin.H{"data": "album not found"})
        return
    }

    // response to user request
    c.IndentedJSON(http.StatusOK, gin.H{"data": album})
    // c.IndentedJSON(http.StatusOK, gin.H{"data": "Get an albums"})
}

// PutAlbums ...
func (al AlbumHandler) PutAlbums(c *gin.Context) {
    // /albums/:id => c.Param("id") => string value
    tempID := c.Param("id")

    // Convert string to int
    id, err := strconv.Atoi(tempID)
    if err != nil {
        c.IndentedJSON(http.StatusBadRequest, gin.H{"data": "Bad Request"})
        return
    }

    var updateAlbum domain.Album

    // Call BindJSON to bind the received JSON to updateAlbum.
    if err := c.BindJSON(&updateAlbum); err != nil {
        c.IndentedJSON(http.StatusBadRequest, gin.H{"data": "Bad Request"})
        return
    }

    // Assign updateAlbum to id in Request URL
    updateAlbum.ID = id

    err = al.repo.Update(updateAlbum)
    if err != nil {
        c.IndentedJSON(http.StatusNotFound, gin.H{"data": "album not found"})
        return
    }

    // response to user request
    c.IndentedJSON(http.StatusOK, gin.H{"data": "Update Successfull"})
}

// DeleteAlbums ...
func (al AlbumHandler) DeleteAlbums(c *gin.Context) {
    // /albums/:id => c.Param("id") => string value
    tempID := c.Param("id")

    // Convert string to int
    id, err := strconv.Atoi(tempID)
    if err != nil {
        c.IndentedJSON(http.StatusBadRequest, gin.H{"data": "Bad Request"})
        return
    }

    err = al.repo.Delete(id)
    if err != nil {
        c.IndentedJSON(http.StatusNotFound, gin.H{"data": "album not found"})
        return
    }

    // response to user request
    c.IndentedJSON(http.StatusNotFound, gin.H{"data": "Delete successfull"})
    // c.IndentedJSON(http.StatusOK, gin.H{"data": "Delete an albums"})
}

