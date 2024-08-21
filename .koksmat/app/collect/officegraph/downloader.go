// officegraph.go
package officegraph

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"text/template"
	"time"

	"github.com/magicbutton/magic-mix/collect"
)

// const (
// 	concurrency    = 15 // Number of worker goroutines
// 	throttleDelay  = 5  // seconds to wait if throttled
// 	maxPages       = 1000
// 	tokenValidTime = 30 * time.Minute
// )

var (
	authToken      string
	authTokenLock  sync.Mutex
	throttleDelay  int
	maxPages       int
	tokenValidTime time.Duration
)

const DEBUG = true

type Parent struct {
	CreatedDateTime time.Time `json:"createdDateTime"`
	DisplayName     string    `json:"displayName"`
	ID              string    `json:"id"`
	IsPersonalSite  bool      `json:"isPersonalSite"`
	Name            string    `json:"name"`
	Root            struct {
	} `json:"root"`
	SiteCollection struct {
		Hostname string `json:"hostname"`
	} `json:"siteCollection"`
	WebURL string `json:"webUrl"`
}

type FilterFunc func([]byte) bool

type DownloaderOptions struct {
	Concurrency    int
	ThrottleDelay  int
	MaxPages       int
	TokenValidTime time.Duration
	Filter         FilterFunc
}

type DownloadBatchChild struct {
	Url    string `json:"url"`
	Prefix string `json:"prefix"`
}

type DownloadBatchType struct {
	ParentUrl string               `json:"parentUrl"`
	ChildUrls []DownloadBatchChild `json:"childUrls"`
}

type Details struct {
	ParentId interface{}     `json:"parentId"`
	Details  json.RawMessage `json:"details"`
}

func DownloadBatch(batchID string, batchType DownloadBatchType, options *DownloaderOptions) {
	if len(batchType.ChildUrls) == 0 {
		Downloader(batchID, batchType.ParentUrl, "", "", options)
		return
	}
	for _, details := range batchType.ChildUrls {
		Downloader(batchID, batchType.ParentUrl, details.Url, details.Prefix, options)
		//if !DEBUG {
		collect.CombineJsonFiles(batchID, details.Prefix, true)
		//}

	}

}

func Downloader(batchID string, parentUrl string, childUrl string, childPrefix string, options *DownloaderOptions) {
	if options == nil {
		options = &DownloaderOptions{}
	}
	// Create batch folder if it doesn't exist
	if options.Concurrency == 0 {
		options.Concurrency = 50
	}
	if options.ThrottleDelay == 0 {
		options.ThrottleDelay = 5
	}
	if options.MaxPages == 0 {
		options.MaxPages = 1000
	}
	if options.TokenValidTime == 0 {
		options.TokenValidTime = 30 * time.Minute
	}

	throttleDelay = options.ThrottleDelay
	maxPages = options.MaxPages
	tokenValidTime = options.TokenValidTime

	batchFolder := batchID // fmt.Sprintf("batch-%s", batchID)
	if err := ensureBatchFolderExists(batchFolder); err != nil {
		fmt.Println("Error creating batch folder:", err)
		return
	}

	// Start the initial call to get all sites
	// Assuming you have a function getAllSites() that returns the list of site IDs
	allParents := getParents(batchID, parentUrl, options.Filter)
	if childUrl == "" {
		log.Println(childPrefix, "No details requests, so done here")
		return
	}
	// Start a background goroutine to handle token renewal
	go renewTokenPeriodically()

	// Create a channel to send site IDs to workers
	parentitem := make(chan map[string]interface{})

	// Create a wait group to wait for all workers to finish
	var wg sync.WaitGroup

	// Start worker goroutines
	for i := 0; i < options.Concurrency; i++ {
		wg.Add(1)
		go worker(i, parentitem, childUrl, batchFolder, childPrefix, &wg)
	}

	// Send site IDs to the channel
	go func() {
		for _, parent := range allParents {
			parentitem <- parent
		}
		close(parentitem)
	}()

	// Wait for all workers to finish
	wg.Wait()
	log.Println(childPrefix, "All details downloaded successfully.")
}

func ensureBatchFolderExists(folder string) error {
	// Check if folder exists
	_, err := os.Stat(folder)
	if os.IsNotExist(err) {
		// Folder doesn't exist, create it
		if err := os.Mkdir(folder, 0755); err != nil {
			return err
		}
	}
	return nil
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func worker(workerId int, parentItems <-chan map[string]interface{}, childUrl string, batchFolder string, childPrefix string, wg *sync.WaitGroup) {
	defer wg.Done()

	for parentItem := range parentItems {
		tmpl, err := template.New("url").Parse(childUrl)
		if err != nil {
			panic(err)

		}
		var b bytes.Buffer

		err = tmpl.Execute(&b, parentItem)
		if err != nil {
			panic(err)
		}
		url := b.String()
		fileid := strings.ReplaceAll(url[8:], "/", "-")
		filename := filepath.Join(batchFolder, fmt.Sprintf("%s-%s.json", childPrefix, fileid))
		if fileExists(filename) {
			//fmt.Printf("Worker %d: Permissions already downloaded for site %d\n", id, siteID)
			continue
		}

		// Make API call to get details for the site
		details, err := downloadDetails(url)
		if err != nil {
			log.Printf("Error downloading details for  %s: %s\n", parentItem, err)
			continue
		}

		var temp json.RawMessage
		if err := json.Unmarshal(details, &temp); err != nil {
			log.Printf("Error unmarshalling details for  %s: %s\n", parentItem, err)
			continue
		}
		detauilsToStore := Details{
			ParentId: parentItem,
			Details:  temp,
		}
		// Write permissions to file
		detailBuf, err := json.MarshalIndent(detauilsToStore, "", "  ")
		if err := os.WriteFile(filename, detailBuf, 0644); err != nil {
			log.Printf("Error writing details for  %s to file: %s\n", parentItem, err)
			continue
		}

		log.Printf("Worker %d: Details downloaded for  %s\n", workerId, fileid)
	}
}

func downloadDetails(url string) ([]byte, error) {

	//url := fmt.Sprintf(detailsurl, id)
	body, err := Download(url, authToken, maxPages)

	if err != nil {
		return nil, err
	}

	return []byte(*body), nil

}

func getToken() string {
	authTokenLock.Lock()
	defer authTokenLock.Unlock()

	// If token exists and not expired, return it
	if authToken != "" {
		return authToken
	}
	log.Println("Getting auth token ")
	// Otherwise, fetch a new token
	// Implement your logic here to obtain the authentication token
	// This might involve making a request to a token endpoint with client credentials
	// For simplicity, let's assume authToken is obtained synchronously
	newAuthToken, err := GetAuthToken()
	if err != nil {
		log.Fatal(err)
	}

	claim, err := DecodeClaim(newAuthToken)
	if err != nil {
		log.Fatal("DecodeToken", err)
	}
	log.Println("Scopes: ", claim.Roles)
	authToken = newAuthToken
	return authToken
}

func getRetryAfter(header http.Header) int {
	if value := header.Get("Retry-After"); value != "" {
		delay, _ := strconv.Atoi(value)
		return delay
	}
	return throttleDelay
}

func getParents(batchID string, parentUrl string, filter FilterFunc) []map[string]interface{} {
	// token, err := getToken("Sites.FullControl.All")
	// if err != nil {
	// 	log.Fatal("Getting auth token", err)
	// }
	token := getToken()
	batchFolder := batchID
	filename := filepath.Join(batchFolder, "parents.json")
	// Check if parents data already exists, then it has been downloaded before
	// and we can use it.

	// this trick is primarily made for making the debugging faster

	if DEBUG && fileExists(filename) {

		fileData, err := os.ReadFile(filename)
		if err != nil {
			log.Fatal("Reading parents data from file", err)
		}
		var parents []map[string]interface{}

		marsshallErr := json.Unmarshal(fileData, &parents)
		if marsshallErr != nil {
			log.Fatal("Unmarshalling parents data", marsshallErr)
		}
		return parents

	}
	log.Println("Downloading parents")
	log.Println("Parent URL: ", parentUrl)
	siteData, err := Download(parentUrl, token, maxPages)
	if err != nil {
		log.Fatal("Downloading parents", err)
	}

	data := []byte(*siteData)
	var parents []map[string]interface{}
	marsshallErr := json.Unmarshal(data, &parents)
	if marsshallErr != nil {
		log.Fatal("Unmarshalling parent data", marsshallErr)
	}

	return parents
}

func renewTokenPeriodically() {
	for {
		// Sleep for half of token validity period
		time.Sleep(tokenValidTime / 2)

		// Renew token
		authTokenLock.Lock()
		authToken = "" // Invalidate current token
		authTokenLock.Unlock()
	}
}
