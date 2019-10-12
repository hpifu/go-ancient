package es

import (
	"context"
	"fmt"
	"time"

	"github.com/olivere/elastic/v7"
)

var ancientMapping = `{
    "settings": {
        "analysis": {
            "tokenizer": {
                "ngram_tokenizer": {
                    "type": "nGram",
                    "min_gram": 1,
                    "max_gram": 10,
                    "token_chars": [
                        "letter",
                        "digit"
                    ]
                }
            },
            "analyzer": {
                "ngram_tokenizer_analyzer": {
                    "type": "custom",
                    "tokenizer": "ngram_tokenizer",
                    "filter": [
                        "lowercase"
                    ]
                }
            }
        },
        "max_ngram_diff": "10"
	},
	"mappings": {
		"properties": {
			"id": {
				"type": "long"
			},
			"title": {
				"type": "text",
				"analyzer": "ngram_tokenizer_analyzer",
				"search_analyzer": "standard"
			},
			"author": {
				"type": "text",
				"analyzer": "ngram_tokenizer_analyzer",
				"search_analyzer": "standard"
			},
			"dynasty": {
				"type": "text",
				"analyzer": "ngram_tokenizer_analyzer",
				"search_analyzer": "standard"
			},
			"content": {
				"type": "text",
				"analyzer": "ngram_tokenizer_analyzer",
				"search_analyzer": "standard"
			}
		}
	}
}`

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

	{
		ctx, cancel := context.WithTimeout(context.Background(), 3000*time.Millisecond)
		defer cancel()
		_, _, err = client.Ping(uri).Do(ctx)
		if err != nil {
			return nil, err
		}
	}

	{
		ctx, cancel := context.WithTimeout(context.Background(), 3000*time.Millisecond)
		defer cancel()
		exists, err := client.IndexExists("ancient").Do(ctx)
		if err != nil {
			return nil, err
		}
		if !exists {
			createIndex, err := client.CreateIndex("ancient").Body(ancientMapping).Do(context.Background())
			if err != nil {
				return nil, err
			}
			if !createIndex.Acknowledged {
				return nil, fmt.Errorf("not acknowledged")
			}
		}
	}

	return &ES{
		es: client,
	}, nil
}
