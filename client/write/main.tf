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
  variable "iaas_type" {
     description="Type of iaas provider"
  }
  variable "iaas_url" {
     description="Url of iaas provider"
  }
  resource "orch_iaas_resource" "my_iaas" {
  name=var.iaas_name
  type=var.iaas_type
  url=var.iaas_url
 }
 resource "orch_iaas_resource" "my_iaas1" {
  name=var.iaas_name
  type=var.iaas_type
  url=var.iaas_url
 }
 output "my_iaas_output" {
  value=[orch_iaas_resource.my_iaas.guid_id,orch_iaas_resource.my_iaas.name]

}





