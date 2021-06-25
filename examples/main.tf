terraform {
  required_providers {
    a = {
      source = "timyates/random-name"
      version = "0.0.7"
    }
  }
}

provider "a" {
  # Configuration options
}

resource "a_random_name" "woo" {}

output "out" {
  value = a_random_name.woo.id
}