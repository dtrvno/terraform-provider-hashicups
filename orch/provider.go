package orch

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap:   map[string]*schema.Resource{
			"orch_iaas_resource": resourceIAAS(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"orch_datasource":     dataSourceIAAS(),
		},
	}
}
type IAAS_Input struct {
	Name   string `json:"name,omitempty"`
	Password string `json:"password,omitempty"`
	Type  string `json:"type,omitempty"`
	Url string `json:"url,omitempty"`
	User string `json:"user,omitempty"`
}

type IAAS struct {
	ID     string `json:"guid_id,omitempty"`
	Name   string `json:"name,omitempty"`
	Password string `json:"password,omitempty"`
	Type  string `json:"type,omitempty"`
	Url string `json:"url,omitempty"`
	User string `json:"user,omitempty"`
}
