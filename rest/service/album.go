package service

import "github.com/monitor-cmr/web-w-go/domain"

// AlbumServiceRepository ....
type AlbumServiceRepository interface {
    FetchAll() ([]domain.Album, error)
    Fetch(int) (*domain.Album, error)
    Save(domain.Album) (*int, error)
    Update(domain.Album) error
    Delete(int) error
}

// AlbumService ...
type AlbumService struct {
    repo domain.AlbumRepository
}

// NewAlbumService ...
func NewAlbumService(repo domain.AlbumRepository) *AlbumService {
    return &AlbumService{
        repo: repo,
    }
}

// FetchAll ...
func (s AlbumService) FetchAll() ([]domain.Album, error) {
    return s.repo.SelectAll()
}

// Fetch ...
func (s AlbumService) Fetch(id int) (*domain.Album, error) {
    return s.repo.Select(id)
}

// Save ...
func (s AlbumService) Save(a domain.Album) (*int, error) {
    return s.repo.Save(a)
}

// Update ...
func (s AlbumService) Update(a domain.Album) error {
    return s.repo.Update(a)
}

// Delete ...
func (s AlbumService) Delete(id int) error {
    return s.repo.Delete(id)
}

