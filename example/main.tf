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

/* output "testing" {
  value = data.pingaccess_enginelistener.test
} */
/* module "test" {
  source = "./pingaccess_listener"
}
output "test" {
  value = module.test.pingaccess_listener
} */
