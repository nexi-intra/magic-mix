package officegraph

import "time"

type SharePointListInfo []struct {
	OdataEtag string `json:"@odata.etag"`
	CreatedBy struct {
		User struct {
			DisplayName string `json:"displayName"`
		} `json:"user"`
	} `json:"createdBy"`
	CreatedDateTime      time.Time `json:"createdDateTime"`
	Description          string    `json:"description"`
	DisplayName          string    `json:"displayName"`
	ETag                 string    `json:"eTag"`
	ID                   string    `json:"id"`
	LastModifiedDateTime time.Time `json:"lastModifiedDateTime"`
	List                 struct {
		ContentTypesEnabled bool   `json:"contentTypesEnabled"`
		Hidden              bool   `json:"hidden"`
		Template            string `json:"template"`
	} `json:"list"`
	Name            string `json:"name"`
	ParentReference struct {
		SiteID string `json:"siteId"`
	} `json:"parentReference"`
	WebURL         string `json:"webUrl"`
	LastModifiedBy struct {
		User struct {
			DisplayName string `json:"displayName"`
			Email       string `json:"email"`
			ID          string `json:"id"`
		} `json:"user"`
	} `json:"lastModifiedBy,omitempty"`
}
