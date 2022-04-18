package elasticsearch_test

import (
	"context"
	"errors"
	"fmt"
	"github.com/olivere/elastic"
	"hello/elasticsearch/model"
	"testing"
)

const (
	indexName = "twitter"

	mapping = `
{
	"settings":{
		"number_of_shards": 1,
		"number_of_replicas": 0
	},
	"mappings":{
		"tweet":{
			"properties":{
				"user":{
					"type":"keyword"
				},
				"message":{
					"type":"text",
					"store": true,
					"fielddata": true
				},
				"image":{
					"type":"keyword"
				},
				"created":{
					"type":"date"
				},
				"tags":{
					"type":"keyword"
				},
				"location":{
					"type":"geo_point"
				},
				"suggest_field":{
					"type":"completion"
				}
			}
		}
	}
}
`
)

func TestCreateIndex(t *testing.T) {
	client, err := elastic.NewClient()
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()

	exist, err := client.IndexExists(indexName).Do(ctx)
	if err != nil {
		t.Fatal(err)
	}

	if exist {
		t.Logf("%s index has existed, exec delete index ...", indexName)
		client.DeleteIndex(indexName)
	}

	resp, err := client.CreateIndex(indexName).
		Body(mapping).
		Do(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if !resp.Acknowledged {
		t.Fatal("Not acknowledged")
	}
}

func TestCrudDocument(t *testing.T) {
	type testcase struct {
		indexName    string
		documentType string
		documentId   string
		document     model.Tweet
	}

	tests := []testcase{
		{
			indexName:    indexName,
			documentType: "doc",
			documentId:   "1",
			document: model.Tweet{
				User:     "olivere",
				Message:  "Take Five",
				Retweets: 0,
			},
		},
		{
			indexName:    indexName,
			documentType: "doc",
			documentId:   "2",
			document: model.Tweet{
				User:    "olivere",
				Message: "It's a Raggy Waltz",
			},
		},
	}

	client, err := elastic.NewClient()
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()

	for _, tt := range tests {
		// index
		putResponse, err := client.Index().
			Index(tt.indexName).
			Type(tt.documentType).
			Id(tt.documentId).
			BodyJson(tt.document).
			Do(ctx)
		if err != nil {
			t.Fatal(err)
		}

		checkPutResponse := func(except *testcase, actual *elastic.IndexResponse) error {
			if putResponse.Id != tt.documentId {
				return errors.New(fmt.Sprintf("put document. except document id: %s, but got %s", putResponse.Id, tt.documentId))
			}
			if putResponse.Type != tt.documentType {
				return errors.New(fmt.Sprintf("put document. except index type: %s, but got %s", putResponse.Id, tt.documentId))
			}
			if putResponse.Index != tt.indexName {
				return errors.New(fmt.Sprintf("put document. except index name: %s, but got %s", putResponse.Id, tt.documentId))
			}
			return nil
		}
		if err := checkPutResponse(&tt, putResponse); err != nil {
			t.Fatal(err)
		}

		// get
		getResponse, err := client.Get().
			Index(tt.indexName).
			Type(tt.documentType).
			Id(tt.documentId).
			Do(ctx)
		if err != nil {
			t.Fatal(err)
		}
		checkGetResult := func(except *testcase, actual *elastic.GetResult) error {
			if putResponse.Id != tt.documentId {
				return errors.New(fmt.Sprintf("get document. except document id: %s, but got %s", putResponse.Id, tt.documentId))
			}
			if putResponse.Type != tt.documentType {
				return errors.New(fmt.Sprintf("get document. except index type: %s, but got %s", putResponse.Id, tt.documentId))
			}
			if putResponse.Index != tt.indexName {
				return errors.New(fmt.Sprintf("get document. except index name: %s, but got %s", putResponse.Id, tt.documentId))
			}
			return nil
		}
		if err := checkGetResult(&tt, getResponse); err != nil {
			t.Fatal(err)
		}

		// delete
		if _, err = client.Delete().
			Index(tt.indexName).
			Type(tt.documentType).
			Id(tt.documentId).
			Do(ctx); err != nil {
			t.Fatal(err)
		}
	}
}
