package handler

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "github.com/monitor-cmr/web-w-go/domain"
)

// albums slice to seed record album data.
var albums = []domain.Album{
    {ID: 1, Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: 2, Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: 3, Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// AlbumHandler ...
type AlbumHandler struct{}

// NewAlbumHandler ...
func NewAlbumHandler() *AlbumHandler {
    return &AlbumHandler{}
}

// GetAlbums ...
func (al AlbumHandler) GetAlbums(c *gin.Context) {
    // c.IndentedJSON(http.StatusOK, gin.H{"data": "Get all albums"})
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

    // Loop over the list of albums, looking for
    // an album whose ID value matches the parameter.
    for _, a := range albums {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, gin.H{"data": a})
            return
        }
    }

    // response to user request
    c.IndentedJSON(http.StatusNotFound, gin.H{"data": "album not found"})
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

    // Loop over the list of albums, looking for
    // an album whose ID value matches the parameter.
    for i, a := range albums {
        if a.ID == id {
            albums[i] = updateAlbum
            c.IndentedJSON(http.StatusOK, gin.H{"data": updateAlbum.ID})
            return
        }
    }

    // response to user request
    c.IndentedJSON(http.StatusNotFound, gin.H{"data": "album not found"})
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

    // Loop over the list of albums, looking for
    // an album whose ID value matches the parameter.
    for i, a := range albums {
        if a.ID == id {
            // i->length(albums)
            // albums = [1,2,3]
            // index = 0,1,2
            // tempAbum = len(albums) - 1 => 2
            tempAlbum := make([]domain.Album, len(albums)-1)

            // 0->i
            // found an ID at index = 1
            // cp albums[:index] => tempAlbum[:index]
            copy(tempAlbum[:i], albums[:i])

            // Move to index + 1
            // cp albums[index+1:] => tempAlbum[index:]
            copy(tempAlbum[i:], albums[i+1:])

            // assign albums to tempAlbum
            albums = tempAlbum

            // response to user request
            c.IndentedJSON(http.StatusOK, gin.H{"data": "Delete an albums"})
            return
        }
    }

    // response to user request
    c.IndentedJSON(http.StatusNotFound, gin.H{"data": "album not found"})
    // c.IndentedJSON(http.StatusOK, gin.H{"data": "Delete an albums"})
}

