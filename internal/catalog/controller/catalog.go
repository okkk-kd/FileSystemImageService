package controller

import (
	"context"
	"tagesTestTask/internal/catalog/usecase"
	"tagesTestTask/internal/models"
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
}

func NewCatalogCtrl(uc usecase.Catalog) (obj *catalog, err error) {
	return &catalog{
		uc: uc,
	}, err
}

func (c *catalog) UploadFile(ctx context.Context, req *pbCatalog.UploadFileReq) (resp *pbCatalog.UploadFileRes, err error) {
	if err := c.uc.UploadFile(
		models.UploadFileRequest{
			ClientID: int(req.GetClientID()),
			FileName: req.GetFile().GetName(),
			File:     req.GetFile().GetChunk(),
		},
	); err != nil {
		return nil, err
	}
	return &pbCatalog.UploadFileRes{Response: &pbCatalog.Res{
		Info: "All good!)",
	}}, nil
}

func (c *catalog) GetFilesList(ctx context.Context, req *pbCatalog.GetFileListReq) (res *pbCatalog.GetFileListRes, err error) {
	list, err := c.uc.GetFileList(models.GetFileList{
		int(34568),
	})
	if err != nil {
		return nil, err
	}
	var resList []*pbCatalog.GetFileListRes_File
	for _, el := range list.List {
		resList = append(resList, &pbCatalog.GetFileListRes_File{Id: int64(el.ImgID), Name: el.Name})
	}
	res = &pbCatalog.GetFileListRes{
		Files: resList,
	}
	return res, nil
}

func (c *catalog) GetFileByName(ctx context.Context, req *pbCatalog.GetFileByNameReq) (*pbCatalog.GetFileByNameRes, error) {
	img, err := c.uc.GetFileByName(models.GetFileByName{
		int(34568),
		req.GetName(),
	})
	if err != nil {
		return nil, err
	}
	return &pbCatalog.GetFileByNameRes{Chunk: img}, nil
}
