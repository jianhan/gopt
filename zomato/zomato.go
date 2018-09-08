package zomato

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type base struct {
	apiBaseURL string
}

type CommonAPI interface {
	Categories() (*Categories, error)
	Cities() ([]*City, error)
}

type commonAPI struct {
	base
}

func NewCommonAPI() CommonAPI {
	return &commonAPI{base: base{apiBaseURL: os.Getenv("ZOMATO_API_URL")}}
}

func (c *commonAPI) Categories() (*Categories, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/categories"), nil)
	req.Header.Add("user-key", os.Getenv("ZOMATO_API_KEY"))
	rsp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}

	categories := new(Categories)
	if err := json.Unmarshal(body, &categories); err != nil {
		return nil, err
	}

	return categories, nil
}

func (c *commonAPI) Cities() ([]*City, error) {
	return nil, nil
}
