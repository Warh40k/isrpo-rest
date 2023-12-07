package openapi

type AppEntity struct {
	Id          int32  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	LastUpdated string `json:"last-updated,omitempty"`
}

var curid int32 = 1
var apps []AppEntity
