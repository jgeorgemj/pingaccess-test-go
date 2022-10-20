
terraform {
  required_providers {
    pingaccess = {
      version = "1.0.0"
      source = "hashicorp.com/edu/pa"
    }
  }
}
/* provider "pingaccess" {
  username = "administrator"
  password = "2FederateM0re"
} */

data "pingaccess_enginelistener" "test" {}

output "test1" {
  value = data.pingaccess_enginelistener.test
}