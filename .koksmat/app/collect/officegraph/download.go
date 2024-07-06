package officegraph

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

// GraphResponse represents a typical response from Microsoft Graph API with paging
type GraphResponse struct {
	Value    []interface{} `json:"value"`
	NextLink string        `json:"@odata.nextLink"`
}

// FetchGraphData fetches data from the Microsoft Graph API, handling pagination and throttling
func FetchGraphData(maxPages int, url string, accessToken string) (string, error) {
	var allData []interface{}
	client := &http.Client{}

	for i := 0; i < maxPages; i++ {
		log.Printf("Fetching page %d\n", i+1)
		log.Printf("URL: %s\n", url)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return "", fmt.Errorf("failed to create request: %v", err)
		}
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
		req.Header.Set("Accept", "application/json")
		req.Header.Set("Accept", "*/*")
		resp, err := client.Do(req)
		if err != nil {
			return "", fmt.Errorf("request failed: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode == 429 {
			retryAfter := resp.Header.Get("Retry-After")
			if retryAfter != "" {
				retrySeconds, err := strconv.Atoi(retryAfter)
				if err == nil {
					time.Sleep(time.Duration(retrySeconds) * time.Second)
					i-- // retry the same page
					continue
				}
			}
			time.Sleep(5 * time.Second)
			i-- // retry the same page
			continue
		}

		if resp.StatusCode != http.StatusOK {
			return "", fmt.Errorf("request failed with status: %s", resp.Status)
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return "", fmt.Errorf("failed to read response body: %v", err)
		}

		var result GraphResponse
		err = json.Unmarshal(body, &result)
		if err != nil {

			return "", fmt.Errorf("failed to parse response: %v", err)
		}
		if result.Value == nil {
			var singleResult interface{}
			err = json.Unmarshal(body, &singleResult)
			allData = append(allData, singleResult)
			break

		}
		allData = append(allData, result.Value...)

		if result.NextLink == "" {
			break
		}
		url = result.NextLink
	}

	allDataBytes, err := json.Marshal(allData)
	if err != nil {
		return "", fmt.Errorf("failed to marshal all data: %v", err)
	}

	return string(allDataBytes), nil
}

func Download(url string, accessToken string, maxPages int) (*string, error) {

	data, err := FetchGraphData(maxPages, url, accessToken)
	if err != nil {
		fmt.Printf("Error fetching data: %v\n", err)
		return nil, err
	}
	return &data, nil

}
