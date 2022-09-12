package database

import (
	"log"

	"github.com/olivere/elastic/v7"
)

var (
	EsClient = newClient()
)

// newClient: create new elasticsearch client.
func newClient() *elastic.Client {
	esClient, err := elastic.NewClient(
		elastic.SetURL(env.EsURL),
		elastic.SetSniff(false),
	)

	if err != nil {
		log.Fatalf("Error creating client: %s", err)
	}

	return esClient
}
