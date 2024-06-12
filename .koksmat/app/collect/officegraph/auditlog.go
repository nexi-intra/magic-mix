package officegraph

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type AuditLog struct {
	ID                  string    `json:"id"`
	ActivityDateTime    time.Time `json:"activityDateTime"`
	ActivityDisplayName string    `json:"activityDisplayName"`
	// Add other fields as needed
}
type AuditLogQuery struct {
	Type                string   `json:"@odata.type"`
	DisplayName         string   `json:"displayName"`
	FilterStartDateTime string   `json:"filterStartDateTime"`
	FilterEndDateTime   string   `json:"filterEndDateTime"`
	RecordTypeFilters   []string `json:"recordTypeFilters"`
}

type AuditLogQueryResponse struct {
	OdataContext                string    `json:"@odata.context"`
	ID                          string    `json:"id"`
	DisplayName                 string    `json:"displayName"`
	FilterStartDateTime         time.Time `json:"filterStartDateTime"`
	FilterEndDateTime           time.Time `json:"filterEndDateTime"`
	RecordTypeFilters           []string  `json:"recordTypeFilters"`
	KeywordFilter               any       `json:"keywordFilter"`
	ServiceFilters              []any     `json:"serviceFilters"`
	OperationFilters            []any     `json:"operationFilters"`
	UserPrincipalNameFilters    []any     `json:"userPrincipalNameFilters"`
	IPAddressFilters            []any     `json:"ipAddressFilters"`
	ObjectIDFilters             []any     `json:"objectIdFilters"`
	AdministrativeUnitIDFilters []any     `json:"administrativeUnitIdFilters"`
	Status                      string    `json:"status"`
}

func AuditLogQuerySetup(authToken string, startTime, endTime time.Time) (*AuditLogQueryResponse, error) {
	// Define the query parameters

	startDateTime := startTime.Format(time.RFC3339)
	endDateTime := endTime.Format(time.RFC3339)
	query := AuditLogQuery{
		Type:                "#microsoft.graph.security.auditLogQuery",
		DisplayName:         "SharePoint",
		FilterStartDateTime: startDateTime,
		FilterEndDateTime:   endDateTime,
		RecordTypeFilters:   []string{"sharePoint"},
	}
	url := "https://graph.microsoft.com/beta/security/auditLog/queries"
	body, _ := json.Marshal(query)
	req, err := http.NewRequest("POST", url, strings.NewReader(string(body)))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}
	client := &http.Client{}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", authToken))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept", "*/*")
	//	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	log.Println("Requesting audit log query")
	log.Println(url)
	log.Println(string(body))
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}
	defer resp.Body.Close()

	// if resp.StatusCode == 429 {
	// 	retryAfter := resp.Header.Get("Retry-After")
	// 	if retryAfter != "" {
	// 		retrySeconds, err := strconv.Atoi(retryAfter)
	// 		if err == nil {
	// 			time.Sleep(time.Duration(retrySeconds) * time.Second)
	// 			i-- // retry the same page
	// 			continue
	// 		}
	// 	}
	// 	time.Sleep(5 * time.Second)
	// 	i-- // retry the same page
	// 	continue
	// }

	if resp.StatusCode > 299 {
		return nil, fmt.Errorf("request failed with status: %s", resp.Status)
	}

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}
	queryResponse := AuditLogQueryResponse{}
	err = json.Unmarshal(responseBody, &queryResponse)
	log.Println("Response Id: ", queryResponse.ID)
	return &queryResponse, nil

}

func AuditLogQueryRead(authToken string, queryID string) (*AuditLogQueryResponse, error) {
	url := fmt.Sprintf("https://graph.microsoft.com/beta/security/auditLog/queries/%s", queryID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", authToken))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept", "*/*")
	//	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error sending request: %v\n", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		fmt.Printf("Request failed with status: %s\n", resp.Status)
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return nil, err
	}
	queryResponse := AuditLogQueryResponse{}
	err = json.Unmarshal(body, &queryResponse)
	// log.Println("Response Id: ", queryResponse.ID)
	// log.Println("Status: ", queryResponse.Status)
	return &queryResponse, nil
}

func GetAuditLogs(batchID string, startTime time.Time, endTime time.Time) error {
	authToken, err := GetAuthToken()
	if err != nil {
		log.Fatal("Error getting auth token", err)
	}
	log.Println("Getting audit logs", startTime, endTime)
	batchFolder := batchID
	if err := ensureBatchFolderExists(batchFolder); err != nil {
		fmt.Println("Error creating batch folder:", err)
		return err
	}
	parentFilename := filepath.Join(batchFolder, fmt.Sprintf("query-%s-%s.json", startTime, endTime))
	filename := filepath.Join(batchFolder, fmt.Sprintf("records-%s-%s.json", startTime, endTime))

	if fileExists(filename) {
		log.Println("Audit log already downloaded")
		return nil
	}
	var auditLogQuery *AuditLogQueryResponse = &AuditLogQueryResponse{}

	if !fileExists(parentFilename) {
		var err error
		auditLogQuery, err = AuditLogQuerySetup(authToken, startTime, endTime)
		if err != nil {
			log.Println("Error setting up audit log query", err)
			return err
		}
		output, err := json.MarshalIndent(auditLogQuery, "", "  ")
		os.WriteFile(parentFilename, output, 0644)
	} else {
		parent, err := os.ReadFile(parentFilename)
		if err != nil {
			log.Println("Error reading audit log query", err)
			return err
		}
		err = json.Unmarshal(parent, auditLogQuery)
		if err != nil {
			log.Println("Error unmarshalling audit log query", err)
			return err
		}
	}
	//return nil
	for {

		auditLogQueryStatus, err := AuditLogQueryRead(authToken, auditLogQuery.ID)
		if err != nil {
			log.Println("Error reading audit log query", err)
			return err
		}

		if auditLogQueryStatus.Status == "running" {
			break
		}
		if auditLogQueryStatus.Status == "succeeded" {
			break
		}
		log.Println("Status is", auditLogQueryStatus.Status, " ...waiting 5 seconds for audit log query to start")

		time.Sleep(time.Second * 5)
	}

	log.Println("AuditLog query started")
	url := fmt.Sprintf("https://graph.microsoft.com/beta/security/auditLog/queries/%s/records", auditLogQuery.ID)

	data, err := Download(url, authToken, 1000)
	if err != nil {
		log.Println("Error downloading data", err)
		return err
	}

	if err := os.WriteFile(filename, []byte(*data), 0644); err != nil {
		log.Println("Writing parent data", err)
		return err
	}
	return nil
}

func GetAuditLogsForADayByTheHour(batchID string, day time.Time) error {

	for h := 0; h < 24; h++ {
		from := day.Add(time.Duration(h) * time.Hour)
		to := day.Add(time.Duration(h+1)*time.Hour - 1*time.Second)
		err := GetAuditLogs(batchID, from, to)
		if err != nil {
			log.Println("Error getting audit logs", err)
			return err
		}
	}
	return nil
}
