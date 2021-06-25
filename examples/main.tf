terraform {
  required_providers {
    fun = {
      source = "timyates/random-name"
      version = "0.0.6"
    }
  }
}

provider "fun" {}

resource "random_name" "name" {
  seed = "tim"
}

output "name" {
  value = data.random_name.name.generated
}
output "time" {
  value = data.random_name.name.seed
}