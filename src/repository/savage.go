package repository

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"src/src/model"
)

type Savage interface {
	GetSavage() (*model.Savage, error)
}

type savage struct {
	client interface {
		Do(req *http.Request) (*http.Response, error)
	}
}

func NewSavage(ctx context.Context) *savage {
	return &savage{}

}

func (sa *savage) GetSavage() (*model.Savage, error) {
	url := "https://jsonplaceholder.typicode.com/todos/1"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.New("http format error")
	}

	req.Close = true
	resp, err := sa.client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		return nil, errors.New("error request to savage")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	savageResposne := &model.Savage{}
	json.Unmarshal(body, savageResposne)
	return savageResposne, nil
}
