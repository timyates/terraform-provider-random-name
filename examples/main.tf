terraform {
  required_providers {
    random = {
      source = "timyates/random-name"
      version = "0.0.5"
    }
  }
}

provider "random" {}

data "random_name" "name" {
}

output "name" {
  value = data.random_name.name.generated
}