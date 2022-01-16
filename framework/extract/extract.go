package main

import (
	"log"

	usecase "github.com/cledupe/cross-commerce-sort-number/application/use_case"
	"github.com/cledupe/cross-commerce-sort-number/framework/extract/adapter"
	"github.com/cledupe/cross-commerce-sort-number/framework/infra"
)

func main() {
	var link string = "http://challenge.dienekes.com.br/api/numbers"
	restConsumer, err := adapter.NewRestConsumer(link)
	if err != nil {
		log.Fatal(err)
	}
	var jsonDb infra.JsonDb
	usecase.InsertNumbers(restConsumer, jsonDb)
}
