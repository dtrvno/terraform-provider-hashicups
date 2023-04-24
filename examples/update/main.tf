terraform {
  required_providers {
       orch = {
      version = "1.0.1"
      source  = "localhost/supermicro/orch"
    }
  }
}

provider "orch" {

}
  variable "iaas_name" {
     description="Name of iaas provider"
  }
  
  variable "iaas_url" {
     description="Url of iaas provider"
  }
  variable "iaas_update_url" {
     description="Url of iaas provider"
  }
  variable "iaas_type" {
     description="Url of iaas provider"
  }
  variable "iaas_guid" {
     description="guild of iaas provider"
  }
  resource "orch_iaas_resource" "my_iaas" {
    
        type=var.iaas_type
        name=var.iaas_name
        url=var.iaas_url
      
 }
 output "my_iaas_output" {
  value=[orch_iaas_resource.my_iaas.guid_id,orch_iaas_resource.my_iaas.url]
}






