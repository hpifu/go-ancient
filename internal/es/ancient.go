package es

import (
	"context"
	"github.com/olivere/elastic/v7"
	"reflect"
	"strings"
)

type Ancient struct {
	ID      int    `json:"id"`
	Title   string `json:"title,omitempty"`
	Author  string `json:"author,omitempty"`
	Dynasty string `json:"dynasty,omitempty"`
	Content string `json:"content,omitempty"`
}

func split(s string) []string {
	f := func(r rune) bool {
		for _, s := range []rune("，。、？！； ,.") {
			if r == s {
				return true
			}
		}
		return false
	}
	return strings.FieldsFunc(s, f)

}

func (e *ES) SearchAncient(value string, offset int, limit int) ([]*Ancient, error) {
	query := elastic.NewBoolQuery()
	for _, val := range split(value) {
		q := elastic.NewBoolQuery()
		q.Should(elastic.NewTermQuery("title", val))
		q.Should(elastic.NewTermQuery("author", val))
		q.Should(elastic.NewTermQuery("dynasty", val))
		q.Should(elastic.NewTermQuery("content", val))
		query.Must(q)
	}

	res, err := e.es.Search().
		Index("ancient").
		Query(query).
		From(offset).Size(limit).
		Do(context.Background())

	if err != nil {
		return nil, err
	}

	if res == nil {
		return nil, nil
	}

	var ancient Ancient
	var ancients []*Ancient
	for _, item := range res.Each(reflect.TypeOf(ancient)) {
		if t, ok := item.(Ancient); ok {
			ancients = append(ancients, &t)
		}
	}

	return ancients, err
}
