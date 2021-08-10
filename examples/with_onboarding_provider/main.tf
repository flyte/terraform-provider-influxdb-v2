terraform {
  required_providers {
    influxdb-v2-onboarding = {
      source = "lancey-energy-storage/influxdb-v2-onboarding"
    }
    influxdb-v2 = {
      source = "lancey-energy-storage/influxdb-v2"
    }
  }
}

# Onboarding

provider "influxdb-v2-onboarding" {
    url = "http://localhost:8086"
}

resource "influxdb-v2-onboarding_setup" "setup" {
    username = "joe"
    password = "changeme"
    bucket = "defaultbucket"
    org = "company"
    retention_period = 4
}

output "onboarding-token" {
    value = influxdb-v2-onboarding_setup.setup.token
}

# influxdb-v2 provider

provider "influxdb-v2" {
    url = "http://localhost:8086/"
    token = influxdb-v2-onboarding_setup.setup.token
}

data "influxdb-v2_ready" "status" {}

output "influxdb-v2_is_ready" {
    value = data.influxdb-v2_ready.status.output["url"]
}

resource "influxdb-v2_bucket" "temp" {
    description = "Temperature sensors data"
    name = "temp"
    org_id = influxdb-v2-onboarding_setup.setup.org_id
    retention_rules {
        every_seconds = 3600*24*30
    }
}

resource "influxdb-v2_authorization" "api" {
    org_id = influxdb-v2-onboarding_setup.setup.org_id
    description = "api token"
    status = "active"
    permissions {
        action = "read"
        resource {
            id = influxdb-v2_bucket.temp.id
            org_id = influxdb-v2-onboarding_setup.setup.org_id
            type = "buckets"
        }
    }
    permissions {
        action = "write"
        resource {
            id = influxdb-v2_bucket.temp.id
            org_id = influxdb-v2-onboarding_setup.setup.org_id
            type = "buckets"
        }
    }
}

output "api_token" {
    value = influxdb-v2_authorization.api.token
}
