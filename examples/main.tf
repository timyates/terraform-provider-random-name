terraform {
  required_providers {
    a = {
      source = "timyates/random-name"
      version = "0.0.8"
    }
  }
}

provider "a" {
  # Configuration options
}

resource "a_random_name" "woo" {
//  seed = timestamp()
}

output "out" {
  value = a_random_name.woo.id
}