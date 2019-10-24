package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"src/src/model"
)

type Savage interface {
	GetSavage(number string) (*model.Savage, error)
}

type savage struct {
	client interface {
		Do(req *http.Request) (*http.Response, error)
	}
}

func NewSavage(ctx context.Context, client http.Client) *savage {
	return &savage{
		client: &client,
	}

}

func (sa *savage) GetSavage(number string) (*model.Savage, error) {
	fmt.Printf("GetSavage started")
	url := "https://jsonplaceholder.typicode.com/todos/" + number
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
	fmt.Printf("response %v", body)
	savageResposne := &model.Savage{}
	json.Unmarshal(body, savageResposne)

	return savageResposne, nil
}
