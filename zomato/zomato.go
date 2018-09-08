package zomato

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"
)

type base struct {
	apiBaseURL string
}

type CommonAPI interface {
	Categories() (*CategoryResponse, error)
	Cities() ([]*City, error)
}

type commonAPI struct {
	base
}

func NewCommonAPI() CommonAPI {
	return &commonAPI{base: base{apiBaseURL: os.Getenv("ZOMATO_API_URL")}}
}

func (c *commonAPI) Categories() (*CategoryResponse, error) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/categories", os.Getenv("ZOMATO_API_URL")), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("user-key", os.Getenv("ZOMATO_API_KEY"))
	rsp, err := client.Do(req)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	defer rsp.Body.Close()

	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}

	categoryResponse := CategoryResponse{}
	if err := json.Unmarshal(body, &categoryResponse); err != nil {
		return nil, err
	}

	return &categoryResponse, nil
}

func (c *commonAPI) Cities() ([]*City, error) {
	return nil, nil
}
