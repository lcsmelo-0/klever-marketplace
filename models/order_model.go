package models

type Order struct {
	Id             string   `json:"id,omitempty" bson:"_id,omitempty"`
	CompanyName    string   `json:"companyName"`
	Description    string   `json:"description"`
	Categories     []string `json:"categories"`
	SocialNetworks []string `json:"socialNetworks"`
	Token          string   `json:"token"`
	Reach          int64    `json:"reach"`
	Deadline       string   `json:"deadline"`
	CampaignStart  string   `json:"campaignStart"`
}
