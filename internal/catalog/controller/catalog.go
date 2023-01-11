package controller

import (
	"context"
	"tagesTestTask/internal/catalog/usecase"
	pbCatalog "tagesTestTask/pb/catalog"
)

type catalog struct {
	uc usecase.Catalog
	pbCatalog.UnimplementedCatalogServiceServer
}

type Catalog interface {
	UploadFile(ctx context.Context, req *pbCatalog.UploadFileReq) (*pbCatalog.UploadFileRes, error)
	GetFilesList(ctx context.Context, req *pbCatalog.GetFileListReq) (*pbCatalog.GetFileListRes, error)
	GetFileByName(ctx context.Context, req *pbCatalog.GetFileByNameReq) (*pbCatalog.GetFileByNameRes, error)
	GetFilesByCategory(ctx context.Context, req *pbCatalog.GetFilesByCategoryReq) (*pbCatalog.GetFileByCategoryRes, error)
}

func NewCatalogCtrl() (obj *catalog, err error) {
	return &catalog{}, err
}

func (c *catalog) UploadFile(ctx context.Context, req *pbCatalog.UploadFileReq) (*pbCatalog.UploadFileRes, error) {
	//TODO implement me
	panic("implement me")
}

func (c *catalog) GetFilesList(ctx context.Context, req *pbCatalog.GetFileListReq) (*pbCatalog.GetFileListRes, error) {
	//TODO implement me
	panic("implement me")
}

func (c *catalog) GetFileByName(ctx context.Context, req *pbCatalog.GetFileByNameReq) (*pbCatalog.GetFileByNameRes, error) {
	//TODO implement me
	panic("implement me")
}

func (c *catalog) GetFilesByCategory(ctx context.Context, req *pbCatalog.GetFilesByCategoryReq) (*pbCatalog.GetFileByCategoryRes, error) {
	//TODO implement me
	panic("implement me")
}
