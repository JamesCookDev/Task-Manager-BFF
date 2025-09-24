package clients

import (
	"io"
	"net/http"
	"os"
)

func GetProjetos() ([]byte, int, error) {
	djangoAPI := os.Getenv("DJANGO_API_URL")
	apiToken := os.Getenv("API_TOKEN")

	req, err := http.NewRequest("GET", djangoAPI, nil)
	if err != nil {
		return nil, 0, err
	}

	req.Header.Add("Authorization", "Bearer "+apiToken)
	req.Header.Add("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, err
	}

	return body, resp.StatusCode, nil
}