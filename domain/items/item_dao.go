package items

import (
	"github.com/Bookstore-GolangMS/bookstore_items-api/client/elasticsearch"
	"github.com/Bookstore-GolangMS/bookstore_utils-go/errors"
)

const (
	indexItems = "items"
)

func (i *Item) Save() errors.RestErr {
	result, err := elasticsearch.Client.Index(indexItems, i)

	if err != nil {
		return errors.NewInternalServerError("error trying to save item", errors.NewError("database error"))
	}

	i.Id = result.Id
	return nil
}
