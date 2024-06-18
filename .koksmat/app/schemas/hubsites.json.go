package schemas

type Hubsites []struct {
	ChildSites []interface{} `json:"ChildSites"`
	Id         string        `json:"Id"`
	SiteUrl    string        `json:"SiteUrl"`
	Title      string        `json:"Title"`
}
