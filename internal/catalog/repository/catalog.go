package repository

import (
	"github.com/jmoiron/sqlx"
	"tagesTestTask/internal/models"
	"tagesTestTask/pkg/diskStorage"
)

type catalog struct {
	db      *sqlx.DB
	storage diskStorage.DiskStorage
}

type Catalog interface {
	UploadFile(params models.UploadFileRequest) error
	GetFileList(params models.GetFileList) (diskStorage.ImageList, error)
	GetFileByName(params models.GetFileByName) ([]byte, error)
}

func NewCatalogRepo(db *sqlx.DB, storage diskStorage.DiskStorage) (obj Catalog, err error) {
	return &catalog{
		db:      db,
		storage: storage,
	}, err
}

func (c *catalog) UploadFile(params models.UploadFileRequest) error {
	if err := c.storage.Upload(params.ClientID, params.FileName, params.File); err != nil {
		return err
	}
	return nil
}

func (c *catalog) GetFileList(params models.GetFileList) (diskStorage.ImageList, error) {
	list, err := c.storage.GetList(params.ClientID)
	if err != nil {
		return diskStorage.ImageList{}, err
	}
	return list, nil
}

func (c *catalog) GetFileByName(params models.GetFileByName) ([]byte, error) {
	img, err := c.storage.Download(params.ClientID, params.Name)
	if err != nil {
		return nil, err
	}
	return img, nil
}
