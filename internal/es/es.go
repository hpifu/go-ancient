package es

import (
	"context"
	"github.com/olivere/elastic/v7"
	"time"
)

type ES struct {
	es *elastic.Client
}

func NewES(uri string) (*ES, error) {
	client, err := elastic.NewClient(
		elastic.SetURL(uri),
		elastic.SetSniff(false),
	)

	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3000*time.Millisecond)
	defer cancel()
	_, _, err = client.Ping(uri).Do(ctx)
	if err != nil {
		return nil, err
	}

	return &ES{
		es: client,
	}, nil
}
