package catalog

import (
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
	"tagesTestTask/internal/catalog/controller"
	"tagesTestTask/internal/catalog/repository"
	"tagesTestTask/internal/catalog/usecase"
	pb "tagesTestTask/pb/catalog"
	"tagesTestTask/pkg/diskStorage"
)

type registry struct {
	catalog controller.Catalog
}

type Registry interface {
	InitCatalog(db *sqlx.DB, storage diskStorage.DiskStorage, gSrv *grpc.Server) error
}

func NewCatalogReg() (obj Registry, err error) {
	return &registry{}, err
}

func (r *registry) InitCatalog(db *sqlx.DB, storage diskStorage.DiskStorage, gSrv *grpc.Server) error {
	repo, err := repository.NewCatalogRepo(db, storage)
	if err != nil {
		return err
	}
	uc, err := usecase.NewCatalogUC(repo)
	if err != nil {
		return err
	}
	ctrl, err := controller.NewCatalogCtrl(uc)
	if err != nil {
		return nil
	}
	pb.RegisterCatalogServiceServer(gSrv, ctrl)
	return err
}
