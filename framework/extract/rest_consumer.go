package extract

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

type RestConsumer struct {
	url string
}

type List struct {
	List []float32 `json:"numbers"`
}

func NewRestConsumer(link string) (*RestConsumer, error) {
	validUrl, err := url.ParseRequestURI(link)
	if err != nil {
		return nil, err
	}
	return &RestConsumer{validUrl.String()}, nil
}

func (restConsumer RestConsumer) GetData(page int) ([]float32, error) {
	stringPage := strconv.Itoa(page)
	return restConsumer.getData(stringPage)
}

func (restConsumer RestConsumer) getData(page string) ([]float32, error) {
	var link string = restConsumer.url

	if page != "0" {
		link = link + "?page=" + page
	}

	resp, err := http.Get(link)
	if err != nil {
		log.Print(err)
	}
	defer resp.Body.Close()
	stream, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	template := List{}
	err = json.Unmarshal(stream, &template)
	if err != nil {
		return nil, err
	}
	return template.List, nil
}
