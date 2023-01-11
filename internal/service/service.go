package service

import (
	"github.com/jmoiron/sqlx"
	"tagesTestTask/config"
	"tagesTestTask/internal/catalog"
	"tagesTestTask/pkg/logger"
)

type service struct {
	srv Server
}

type Service interface {
	InitService(*config.Config, logger.Logger, *sqlx.DB) (Controllers, error)
}

func NewService() (obj Service, err error) {
	if err != nil {
		return
	}
	return &service{}, err
}

func (s *service) InitService(cfg *config.Config, logger logger.Logger, pgDB *sqlx.DB) (_ Controllers, err error) {
	catalogReg, err := catalog.NewCatalogReg()
	return Controllers{
		catalogReg,
	}, err
}
