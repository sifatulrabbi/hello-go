package lib

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type refreshTokenRes struct {
	StatusCode int               `json:"statusCode"`
	Message    string            `json:"message"`
	Data       map[string]string `json:"data"`
}

type accessTokenRes struct {
	StatusCode int               `json:"statusCode"`
	Message    string            `json:"message"`
	Data       map[string]string `json:"data"`
}

type DelegatedApp struct {
	ClientID       string
	ClientSecret   string
	OrganizationID string
	RefreshToken   string
	AccessToken    string
}

func (d *DelegatedApp) GetRefreshToken() error {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, getAuthURL("/tokens/refresh"), nil)
	if err != nil {
		return err
	}
	req.Header.Add("client-id", d.ClientID)
	req.Header.Add("client-secret", d.ClientSecret)
	req.Header.Add("organization-id", d.OrganizationID)
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	data := refreshTokenRes{}
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&data)
	return err
}

func (d *DelegatedApp) GetAccessToken() error {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, getAuthURL("/tokens/access"), nil)
	if err != nil {
		return err
	}
	req.Header.Add("client-id", d.ClientID)
	req.Header.Add("organization-id", d.OrganizationID)
	req.Header.Add("Authorization", "Bearer "+d.RefreshToken)
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	data := accessTokenRes{}
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&data)
	return err
}

func (d *DelegatedApp) MakeRequest(method string, url string, body io.Reader) (*http.Response, error) {
	client := &http.Client{}
	req, _ := http.NewRequest(method, url, body)
	req.Header.Add("client-id", d.ClientID)
	req.Header.Add("organization-id", d.OrganizationID)
	req.Header.Add("Authorization", "Bearer "+d.AccessToken)
	return client.Do(req)
}

func getAuthURL(path string) string {
	return "http://localhost:8000/auth/v1" + path
}

func getApiURL(path string) string {
	return "http://localhost:8000/delegated/v1" + path
}

func GetOrganizationDetails() {
	godotenv.Load()

	delegatedApp := DelegatedApp{
		ClientID:       os.Getenv("CLIENT_ID"),
		ClientSecret:   os.Getenv("CLIENT_SECRET"),
		OrganizationID: os.Getenv("ORGANIZATION_ID"),
	}

	err := delegatedApp.GetRefreshToken()
	if err != nil {
		log.Fatalln(err)
	}
	err = delegatedApp.GetAccessToken()
	if err != nil {
		log.Fatalln(err)
	}
	res, err := delegatedApp.MakeRequest(http.MethodGet, getApiURL("/organization"), nil)
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()
	b, _ := io.ReadAll(res.Body)
	var data interface{}
	json.Unmarshal(b, &data)
	fmt.Println(data)
}
