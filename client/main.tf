terraform {
  required_providers {
    orch = {
      version = "1.0.1"
      source  = "localhost/supermicro/orch"
    }
  }
}

provider "orch" {}

module "my_iaas" {
  source = "./orch"
}

output "iaas_out" {
  value = module.my_iaas
}
