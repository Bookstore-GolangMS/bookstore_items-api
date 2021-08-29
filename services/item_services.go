package services

import (
	"github.com/Bookstore-GolangMS/bookstore_items-api/domain/items"
	"github.com/Bookstore-GolangMS/bookstore_utils-go/errors"
)

var (
	ItemsService ItemsServiceInterface = &itemsService{}
)

type ItemsServiceInterface interface {
	Create(items.Item) (*items.Item, errors.RestErr)
	Get(string) (*items.Item, errors.RestErr)
}

type itemsService struct {
}

func (s *itemsService) Create(itemRequest items.Item) (*items.Item, errors.RestErr) {
	if err := itemRequest.Save(); err != nil {
		return nil, err
	}

	return &itemRequest, nil
}

func (s *itemsService) Get(string) (*items.Item, errors.RestErr) {
	return nil, nil
}
