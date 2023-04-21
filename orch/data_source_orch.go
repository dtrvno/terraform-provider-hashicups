package orch

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceIAAS() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIaaSRead,
		Schema: map[string]*schema.Schema{
			"iaas": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"available_nodes": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"created_date": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"guid_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"password": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"url": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"user": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						
					},
				},
			},
		},
	}
}


func dataSourceIaaSRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := &http.Client{Timeout: 10 * time.Second}

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v1/iaas_providers", "http://localhost:5020"), nil)
	if err != nil {
		return diag.FromErr(err)
	}

	r, err := client.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}
	defer r.Body.Close()

	iaas_providers := make([]map[string]interface{}, 0)
	err = json.NewDecoder(r.Body).Decode(&iaas_providers)
	if err != nil {
		return diag.FromErr(err)
	}
    log.Printf("[DEBUG] %s: *************all providers******", iaas_providers)
	if err := d.Set("iaas", iaas_providers); err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
    iaas := IAAS{
		ID:    "id1",
		Name:   "test1",
		Type:  "type1",
		Url: "url1",
	}
	iaas_marshal, err := json.Marshal(iaas)
	log.Printf("[DEBUG] %s ******  iaas ***",string(iaas_marshal))
	
	return diags
}
