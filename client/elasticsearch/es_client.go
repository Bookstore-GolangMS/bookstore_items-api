package elasticsearch

import (
	"context"
	"fmt"
	"github.com/Bookstore-GolangMS/bookstore_items-api/logger"
	"time"

	"github.com/olivere/elastic"
)

var (
	Client esClientInterface = &esClient{}
)

type esClientInterface interface {
	setClient(c *elastic.Client)
	Index(string, interface{}) (*elastic.IndexResponse, error)
}

type esClient struct {
	client *elastic.Client
}

func (es *esClient) setClient(client *elastic.Client) {
	es.client = client
}

func (es *esClient) Index(index string, doc interface{}) (*elastic.IndexResponse, error) {
	result, err := es.client.Index().
		Index(index).
		BodyJson(doc).
		Type("item").
		Do(context.Background())

	if err != nil {
		logger.Error(fmt.Sprintf("error when trying to index document in index %s", index), err)
		return nil, err
	}

	return result, nil
}

func Init() {
	client, err := elastic.NewClient(
		elastic.SetURL("http://127.0.0.1:9200"),
		elastic.SetHealthcheckInterval(10*time.Second),
		//elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)),
		//elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags))
	)

	if err != nil {
		panic(err)
	}

	Client.setClient(client)
}
