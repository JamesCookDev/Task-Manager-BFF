package clients

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

func getBaseURL() string {
	return os.Getenv("DJANGO_API_URL")
}

func makeRequest(method, urlPath string, body []byte) ([]byte, int, error) {
	apiToken := os.Getenv("API_TOKEN")

	fullURL := getBaseURL() + urlPath

	req, err := http.NewRequest(method, fullURL, bytes.NewBuffer(body))
	if err != nil {
		return nil, 0, err
	}

	req.Header.Add("Authorization", "Bearer "+apiToken)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, err
	}

	return responseBody, resp.StatusCode, nil
}

func GetProjetos() ([]byte, int, error) {
	return makeRequest("GET", "", nil)
}

func GetProjetoByID(id string) ([]byte, int, error) {
	urlPath := fmt.Sprintf("%s/", id)
	return makeRequest("GET", urlPath, nil)
}

func CreateProjeto(projectData []byte) ([]byte, int, error) {
	return makeRequest("POST", "", projectData)
}

func UpdateProjeto(id string, projectData []byte) ([]byte, int, error) {
	urlPath := fmt.Sprintf("%s/", id)
	return makeRequest("PUT", urlPath, projectData)
}

func DeleteProjeto(id string) (int, error) {
	urlPath := fmt.Sprintf("%s/", id)
	_, statusCode, err := makeRequest("DELETE", urlPath, nil)
	return statusCode, err
}