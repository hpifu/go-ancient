package es

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
)

type ES struct {
	es *elastic.Client
}

func NewES(uri string) (*ES, error) {
	client, err := elastic.NewClient(
		elastic.SetURL(uri),
		elastic.SetSniff(false),
	)

	fmt.Println(err, client)

	if err != nil {
		return nil, err
	}

	_, _, err = client.Ping(uri).Do(context.Background())
	if err != nil {
		return nil, err
	}

	return &ES{
		es: client,
	}, nil
}
