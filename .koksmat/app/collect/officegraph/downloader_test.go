package officegraph

import (
	"encoding/json"
	"testing"
	"time"
)

type Site struct {
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

func TestDownloadSharePointPermissions(t *testing.T) {
	options := &DownloaderOptions{
		MaxPages: 1,
		Filter: func(item []byte) bool {
			site := &Site{}
			json.Unmarshal(item, site)
			if site.IsPersonalSite {
				return false
			}

			return true
		},
	}
	Downloader("site-permissions-3", "https://graph.microsoft.com/v1.0/sites", "https://graph.microsoft.com/v1.0/sites/%s/permissions", "permissions", options)
}

func TestDownloadOfficeGroupsAndOwners(t *testing.T) {
	options := &DownloaderOptions{
		MaxPages: 1,
	}
	Downloader("groups-1", "https://graph.microsoft.com/v1.0/groups", "https://graph.microsoft.com/v1.0//groups/%s/owners", "owners", options)
	Downloader("groups-1", "https://graph.microsoft.com/v1.0/groups", "https://graph.microsoft.com/v1.0//groups/%s/members", "members", options)
}
func TestDownloadOfficeGroupsAndOwners2(t *testing.T) {
	options := &DownloaderOptions{
		MaxPages: 10000,
	}
	Downloader("groups-2", "https://graph.microsoft.com/v1.0/groups", "https://graph.microsoft.com/v1.0//groups/%s/owners", "owners", options)
	Downloader("groups-2", "https://graph.microsoft.com/v1.0/groups", "https://graph.microsoft.com/v1.0//groups/%s/members", "members", options)
}
