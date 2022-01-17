package adapter

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const ATTEMPTS = 3

type RestConsumer struct {
	url string
}

type List struct {
	List []float64 `json:"numbers"`
}

func NewRestConsumer(link string) (*RestConsumer, error) {
	validUrl, err := url.ParseRequestURI(link)
	if err != nil {
		return nil, err
	}
	return &RestConsumer{validUrl.String()}, nil
}

func (restConsumer RestConsumer) GetData(page int, numbers chan<- []float64) {
	stringPage := strconv.Itoa(page)
	numbers <- restConsumer.getData(stringPage)
}

func (restConsumer RestConsumer) getData(page string) []float64 {
	var link string = restConsumer.url
	var result []float64
	log.Printf("Page number %v\n", page)

	if page != "0" {
		link = link + "?page=" + page
	}

	for retry := 0; retry < ATTEMPTS; retry++ {
		result = restConsumer.callApi(link)
		if len(result) > 0 {
			break
		}
		time.Sleep(time.Millisecond * 100)
	}
	return result

}

func (restConsumer RestConsumer) callApi(link string) []float64 {
	resp, err := http.Get(link)
	if err != nil {
		log.Print(err)
		return []float64{}
	}
	defer resp.Body.Close()
	stream, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Print(err)
		return []float64{}
	}

	template := List{}
	err = json.Unmarshal(stream, &template)
	if err != nil {
		log.Print(err)
		return []float64{}
	}
	return template.List
}
