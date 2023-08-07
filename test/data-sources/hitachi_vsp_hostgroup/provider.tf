terraform {
  required_providers {
    hitachi = {
      version = "2.0"
      source  = "localhost/hitachi-vantara/hitachi"
    }
  }
}

provider "hitachi" {
  san_storage_system {
    serial        = 40014
    management_ip = "172.25.47.115"
    username      = "ms_vmware"
    password      = "Hitachi1"
  }

}
