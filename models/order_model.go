package models

type Order struct {
	Id          string                 `json:"id,omitempty" bson:"_id,omitempty"`
	Category    string                 `json:"category"`
	Description string                 `json:"description"`
	PreviewDate string                 `json:"previewDate"`
	Contact     map[string]interface{} `json:"contact"`
}
