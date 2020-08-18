# Onboarding

provider "influxdbv2-onboarding" {
    url = "http://localhost:9999"
}

resource "influxdbv2-onboarding_setup" "setup" {
    username = "joe"
    password = "changeme"
    bucket = "defaultbucket"
    org = "company"
    retention_period = 4
}

# Influxdbv2 provider use

provider "influxdbv2" {
    url = "http://localhost:9999/"
    token = influxdbv2-onboarding_setup.setup.token
}

data "influxdbv2_ready" "status" {}

output "influxdbv2_is_ready" {
    value = data.influxdbv2_ready.status.output["url"]
}

resource "influxdbv2_bucket" "temp" {
    description = "Temperature sensors dara."
    name = "temp"
    org_id = influxdbv2-onboarding_setup.setup.org_id
    retention_rules {
        every_seconds = 3600*24*30
    }
}

resource "influxdbv2_authorization" "api" {
    org_id = influxdbv2_bucket.temp.org_id
    description = "Influxdb token to give api access to the temp bucket."
    status = "active"
    permissions {
        action = "read"
        resource {
            id = influxdbv2_bucket.temp.id
            org_id = influxdbv2_bucket.temp.org_id
            type = "buckets"
        }
    }
    permissions {
        action = "write"
        resource {
            id = influxdbv2_bucket.temp.id
            org_id = influxdbv2_bucket.temp.org_id
            type = "buckets"
        }
    }
}

output "api_token" {
    value = influxdbv2_authorization.api.token
}
