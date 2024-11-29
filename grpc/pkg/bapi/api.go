package bapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	APP_KEY    = "eddycjy"
	APP_SECRET = "go-programming-tour-book"
)

type AccessToken struct {
	Token string `json:"token"`
}

type API struct {
	URL string
}

func (a *API) getAccessToken() (string, error) {
	body, err := a.httpGet(fmt.Sprintf("%s?app_key=%s&app_secret=%s", "auth", APP_KEY, APP_SECRET))
	if err != nil {
		return "", err
	}
	var accessToken AccessToken
	_ = json.Unmarshal(body, &accessToken)
	return accessToken.Token, nil
}

func (a *API) httpGet(path string) ([]byte, error) {
	resp, err := http.Get(fmt.Sprintf("%s/%s", a.URL, path))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	return body, nil
}

func NewAPI(url string) *API {
	return &API{URL: url}
}

func (a *API) GetTagList(name string) ([]byte, error) {
	token, err := a.getAccessToken()
	fmt.Println(token)
	if err != nil {
		return nil, err
	}
	fmt.Println(token)
	body, err := a.httpGet(fmt.Sprintf("%s?token=%s&name=%s", "api/v1/tags", token, name))
	if err != nil {
		return nil, err
	}
	return body, nil
}
