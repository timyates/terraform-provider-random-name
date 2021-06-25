terraform {
  required_providers {
    random = {
      source = "timyates/random-name"
      version = "0.0.4"
    }
  }
}

provider "random" {}

data "random_name" "name" {
  provider = "random"
}

output "name" {
  value = data.random_name.name.name
}