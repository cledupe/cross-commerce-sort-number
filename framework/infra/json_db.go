package infra

import (
	"encoding/json"
	"io/ioutil"
)

var FILE_NAME = "numbers.json"
var PATH_FILE = "D:\\go-project\\cross-commerce-sort-number\\"

type JsonDb struct {
	List []float64 `json:"numbers"`
}

func (jsonDb JsonDb) Save(list []float64) {
	jsonDb.List = list
	jsonDb.save(jsonDb.load())
}

func (jsonDb JsonDb) Load() ([]float64, error) {
	plan, _ := ioutil.ReadFile(PATH_FILE + FILE_NAME)
	err := json.Unmarshal(plan, &jsonDb)
	if err != nil {
		return nil, err
	}
	return jsonDb.List, nil

}

func (jsonDb JsonDb) load() []byte {
	file, _ := json.Marshal(jsonDb)
	return file
}

func (jsonDb JsonDb) save(file []byte) {
	_ = ioutil.WriteFile(PATH_FILE+FILE_NAME, file, 0644)
}
