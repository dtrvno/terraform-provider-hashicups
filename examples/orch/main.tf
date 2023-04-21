terraform {
  required_providers {
    orch = {
      version = "1.0.1"
      source  = "localhost/supermicro/orch"
    }
  }
}



data "orch_datasource" "all" {}

# Returns all iaas
output "all_iaas" {
  value = data.orch_datasource.all.iaas
}



