terraform {
  required_providers {
    opsgenie = {
      source = "opsgenie/opsgenie"
      version = "0.6.10"
    }
  }
}

provider "opsgenie" {
  # Configuration options
}