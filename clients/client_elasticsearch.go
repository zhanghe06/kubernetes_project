package main

import (
	"log"

	"github.com/elastic/go-elasticsearch"
)

func main() {
	es, _ := elasticsearch.NewDefaultClient()
	log.Println(es.Info())
}
