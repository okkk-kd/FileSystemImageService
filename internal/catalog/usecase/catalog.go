package usecase

import (
	"tagesTestTask/internal/catalog/repository"
	"tagesTestTask/internal/models"
	"tagesTestTask/pkg/diskStorage"
)

type catalog struct {
	repo repository.Catalog
}

type Catalog interface {
	UploadFile(params models.UploadFileRequest) error
	GetFileList(params models.GetFileList) (diskStorage.ImageList, error)
	GetFileByName(params models.GetFileByName) ([]byte, error)
}

func NewCatalogUC(repo repository.Catalog) (obj Catalog, err error) {
	return &catalog{
		repo: repo,
	}, err
}

func (c *catalog) UploadFile(params models.UploadFileRequest) error {
	if err := c.repo.UploadFile(params); err != nil {
		return err
	}
	return nil
}

func (c *catalog) GetFileList(params models.GetFileList) (diskStorage.ImageList, error) {
	list, err := c.repo.GetFileList(params)
	if err != nil {
		return diskStorage.ImageList{}, err
	}
	return list, nil
}

func (c *catalog) GetFileByName(params models.GetFileByName) ([]byte, error) {
	img, err := c.repo.GetFileByName(params)
	if err != nil {
		return nil, err
	}
	return img, nil
}
