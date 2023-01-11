package catalog

import (
	"google.golang.org/grpc"
	"tagesTestTask/internal/catalog/controller"
	pb "tagesTestTask/pb/catalog"
)

type registry struct {
	catalog controller.Catalog
}

type Registry interface {
	InitCatalog(gSrv *grpc.Server) error
}

func NewCatalogReg() (obj Registry, err error) {
	return &registry{}, err
}

func (r *registry) InitCatalog(gSrv *grpc.Server) error {
	test, err := controller.NewCatalogCtrl()
	if err != nil {
		return nil
	}
	pb.RegisterCatalogServiceServer(gSrv, test)
	return err
}
