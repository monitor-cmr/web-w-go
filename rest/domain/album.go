package domain

// Album represents data about a record album.
type Album struct {
    ID     int
    Title  string
    Artist string
    Price  float64
}

// AlbumRepository ...
type AlbumRepository interface {
    SelectAll() ([]Album, error)
    Select(int) (*Album, error)
    Save(Album) (*int, error)
    Update(Album) error
    Delete(int) error
}

