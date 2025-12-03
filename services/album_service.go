package services

import (
	"fmt"

	"github.com/SteliosKoulinas/recordStoreApi/db"
	"github.com/SteliosKoulinas/recordStoreApi/models"
	"gorm.io/gorm"
)

type AlbumService struct{}

func (AlbumService) GetAll() ([]models.Album, error) {
	var albums []models.Album
	result := db.DB.Find(&albums)
	return albums, result.Error
}

func (s *AlbumService) GetByID(id string) (models.Album, error) {
	var album models.Album
	result := db.DB.First(&album, id)
	return album, result.Error
}

func (s *AlbumService) Create(album *models.Album) error {
	// Check if album already exists
	var existing models.Album
	result := db.DB.
		Where("artist = ? AND title = ?", album.Artist, album.Title).
		First(&existing)

	if result.Error == nil {
		// record exists
		return fmt.Errorf("album already exists")
	}

	if result.Error != gorm.ErrRecordNotFound {
		// some other db error
		return result.Error
	}
	return db.DB.Create(album).Error
}

func (s *AlbumService) Delete(id string) error {
	return db.DB.Delete(&models.Album{}, id).Error
}
