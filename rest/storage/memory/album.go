package memory

import (
    "errors"

    "github.com/monitor-cmr/web-w-go/domain"
)

// AlbumMemory ...
type AlbumMemory struct {
    albums []domain.Album
}

// NewAlbumMemory ...
func NewAlbumMemory() *AlbumMemory {
    // albums slice to seed record album data.
    var albums = []domain.Album{
        {ID: 1, Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
        {ID: 2, Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
        {ID: 3, Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
    }

    return &AlbumMemory{
        albums: albums,
    }
}

// SelectAll ...
func (a *AlbumMemory) SelectAll() ([]domain.Album, error) {
    return a.albums, nil
}

// Select ...
func (a *AlbumMemory) Select(id int) (*domain.Album, error) {
    // Loop over the list of albums, looking for
    // an album whose ID value matches the parameter.
    for _, v := range a.albums {
        if v.ID == id {
            return &v, nil
        }
    }

    return nil, errors.New("Not Found")
}

// Save ...
func (a *AlbumMemory) Save(album domain.Album) (*int, error) {
    // Get a length of albums
    lenAlbum := len(a.albums)

    // Get a last album's ID
    // - index start 0
    // - a last album = length - 1 => lenAlbum - 1
    album.ID = a.albums[lenAlbum-1].ID + 1
    
    a.albums = append(a.albums, album)
    return &album.ID, nil
}

// Update ...
func (a *AlbumMemory) Update(album domain.Album) error {
    // Loop over the list of albums, looking for
    // an album whose ID value matches the parameter.
    for i, v := range a.albums {
        if v.ID == album.ID {
            a.albums[i] = album
            return nil
        }
    }
    return errors.New("Not Found")

}

// Delete ...
func (a *AlbumMemory) Delete(id int) error {
    // Loop over the list of albums, looking for
    // an album whose ID value matches the parameter.
    for i, v := range a.albums {
        if v.ID == id {
            // i->length(albums)
            // albums = [1,2,3]
            // index = 0,1,2
            // tempAbum = len(albums) - 1 => 2
            tempAlbum := make([]domain.Album, len(a.albums)-1)

            // 0->i
            // found an ID at index = 1
            // cp albums[:index] => tempAlbum[:index]
            copy(tempAlbum[:i], a.albums[:i])

            // Move to index + 1
            // cp albums[index+1:] => tempAlbum[index:]
            copy(tempAlbum[i:], a.albums[i+1:])

            // assign albums to tempAlbum
            a.albums = tempAlbum

            return nil
        }
    }
    return errors.New("Not Found")
}

