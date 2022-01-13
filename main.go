package main

import (
	"fmt"
	"log"

	"github.com/cledupe/cross-commerce-sort-number/framework/extract"
)

func main() {
	var link string = "http://challenge.dienekes.com.br/api/numbers"
	restConsumer, err := extract.NewRestConsumer(link)
	if err != nil {
		log.Fatal(err)
	}

	result, err := restConsumer.GetData(1)

	if err != nil {
		log.Fatal(err)
	}

	for _, r := range result {
		fmt.Print(r)
	}

}
