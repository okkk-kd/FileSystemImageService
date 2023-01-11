package service

import "tagesTestTask/internal/catalog"

type Controllers struct {
	Catalog interface {
		catalog.Registry
	}
}
