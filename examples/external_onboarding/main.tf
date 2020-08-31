terraform {
  required_providers {
    influxdbv2 = {
      source = "lancey-energy-storage/influxdbv2"
      version = "0.3.0"
    }
  }
}

provider "influxdbv2" {
    // provider is configured with env vars
}

variable "influx_org_id" {
  description = "Influxdb organization ID defined at the onboarding stage"
}

data "influxdbv2_ready" "status" {}

output "influxdbv2_is_ready" {
    value = data.influxdbv2_ready.status.output["url"]
}

resource "influxdbv2_bucket" "temp" {
    description = "Temperature sensors data"
    name = "temp"
    org_id = var.influx_org_id
    retention_rules {
        every_seconds = 3600*24*30
    }
}

resource "influxdbv2_authorization" "api" {
    org_id = var.influx_org_id
    description = "api token"
    status = "active"
    permissions {
        action = "read"
        resource {
            id = influxdbv2_bucket.temp.id
            org_id = var.influx_org_id
            type = "buckets"
        }
    }
    permissions {
        action = "write"
        resource {
            id = influxdbv2_bucket.temp.id
            org_id = var.influx_org_id
            type = "buckets"
        }
    }
}

output "api_token" {
    value = influxdbv2_authorization.api.token
}
