package orch
import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strings"
	"net/http"	
	"encoding/json"
	"fmt"
	"time"
	"io/ioutil"
	"log"
)

func resourceIAAS() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIAASCreate,
		ReadContext:   resourceIAASRead,
		UpdateContext: resourceIAASCreateUpdate,
		DeleteContext: resourceIAASDelete,
		Schema: map[string]*schema.Schema{
			"orch_iaas_resource": &schema.Schema{
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
							Type:        schema.TypeString,
							Computed:    true,
							Description: "the guid_id value returned",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "name of iaas provider",
						},
						"password": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "type of iaas provider",
						},
						"url": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "url of iaas provider",
						},
						"user": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

					},
				},
			},
			"guid_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "the guid_id value",
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "name of iaas provider",
			},
			"type": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "type of iaas provider",
			},
			"url": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "the url value",
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceUpdateIAAS() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIAASUpdate,
		ReadContext:   resourceIAASRead,
		UpdateContext: resourceIAASUpdate,
		DeleteContext: resourceIAASDelete,
		Schema: map[string]*schema.Schema{
			"orch_iaas_update_resource": &schema.Schema{
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
							Type:        schema.TypeString,
							Computed:    true,
							Description: "the guid_id value returned",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "name of iaas provider",
						},
						"password": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "type of iaas provider",
						},
						"url": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "url of iaas provider",
						},
						"user": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

					},
				},
			},
			"guid_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "the guid_id value",
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "name of iaas provider",
			},
			"type": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "type of iaas provider",
			},
			"url": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "the url value",
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}


func resourceIAASCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
    client := &http.Client{Timeout: 10 * time.Second}
	iaas_name := d.Get("name").(string)
	interface_password:=d.Get("password")
	interface_user:=d.Get("user")
	iaas_password:="" 
	iaas_user:=""
	if interface_password != nil {
		iaas_password=interface_password.(string)
	}
	if interface_user!=nil {
		iaas_user=interface_user.(string)
	} 
	iaas_type := d.Get("type").(string)
	iaas_url:=d.Get("url").(string)

	
	iaas := IAAS_Input{
		Name:   iaas_name,
		Password: iaas_password,
		Type:  iaas_type,
		Url: iaas_url,
		User: iaas_user,
	}
	iaas_marshal, err := json.Marshal(iaas)
    log.Printf("[DEBUG] %s: *************create provider******", string(iaas_marshal))
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/v1/iaas_providers", "http://localhost:5020"), 
	       strings.NewReader(string(iaas_marshal)))
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		return diag.FromErr(err)
	}
	log.Printf("[DEBUG]: *************before call******")
    resp, err := client.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}
	log.Printf("[DEBUG]: *************after call******")
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var iaas_output IAAS
	err = json.Unmarshal(body, &iaas_output)
	if err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("guid_id", iaas_output.ID); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("name", iaas_output.Name); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("type", iaas_output.Type); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("url", iaas_output.Url); err != nil {
		return diag.FromErr(err)
	}
	d.SetId("iaas_resource_"+iaas_output.ID)
	return diags
}

func resourceIAASCreateUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourceIAASRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}
func resourceIAASUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
    client := &http.Client{Timeout: 10 * time.Second}
	iaas_guid_id:=""
	iaas_name:=""
	iaas_password:="" 
	iaas_type:=""
	iaas_url:=""
	iaas_user:=""
	interface_guid_id:=d.Get("guid_id")
	interface_name:=d.Get("name")
	interface_password:=d.Get("password")
	interface_type:=d.Get("type")
	interface_url:=d.Get("url")
	interface_user:=d.Get("user")
	if interface_guid_id != nil {
		iaas_guid_id=interface_guid_id.(string)
	}
	if interface_name != nil {
		iaas_name=interface_name.(string)
	}
	if interface_password != nil {
		iaas_password=interface_password.(string)
	}
	if interface_type!=nil {
		iaas_type=interface_type.(string)
	}  
	if interface_url!=nil {
		iaas_url=interface_url.(string)
	} 
    if interface_user!=nil {
		iaas_user=interface_user.(string)
	}
	
	iaas := IAAS {
		ID: iaas_guid_id,
		Name:   iaas_name,
		Password: iaas_password,
		Type:  iaas_type,
		Url: iaas_url,
		User: iaas_user,
	}
	iaas_marshal, err := json.Marshal(iaas)
    log.Printf("[DEBUG] %s: *************update provider******", string(iaas_marshal))
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/v1/iaas_providers/%s", "http://localhost:5020",iaas_guid_id), 
	       strings.NewReader(string(iaas_marshal)))
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		return diag.FromErr(err)
	}
    resp, err := client.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	log.Printf("[DEBUG] %s: *************update provider after******", string(body))
	var iaas_output IAAS
	err = json.Unmarshal(body, &iaas_output)
	if err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("guid_id", iaas.ID); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("name", iaas.Name); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("type", iaas.Type); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("url", iaas.Url); err != nil {
		return diag.FromErr(err)
	}
	d.SetId("iaas_resource_update_"+iaas.ID)
	log.Printf("[DEBUG] %s: *************update provider before exit******", string(body))
	return diags
}
func resourceIAASDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}
