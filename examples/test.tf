provider "influxdbv2-onboarding" {
    url = "http://localhost:9999"
}

resource "influxdbv2-onboarding_setup" "setup" {
    username = "test"
    password = "test1234"
    bucket = "test-bucket"
    org = "test-org"
    retention_period = 4
}

provider "influxdbv2" {
    url = "http://localhost:9999/"
    token = influxdbv2-onboarding_setup.setup.token
}

data "influxdbv2_ready" "test" {}

output "influxdbv2_ready" {
    value = data.influxdbv2_ready.test.output["status"]
}

output "ready_started"  {
    value = data.influxdbv2_ready.test.output["started"]
}

output "ready_up"  {
    value = data.influxdbv2_ready.test.output["up"]
}

resource "influxdbv2_bucket" "test" {
    description = ""
    name = "le bucket de test terraform"
    org_id = influxdbv2-onboarding_setup.setup.org_id
    retention_rules {
        every_seconds = 40
        type = "expire"
    }
    rp = ""
}

output "created_at" {
    value = influxdbv2_bucket.test.created_at
}
output "updated_at" {
    value = influxdbv2_bucket.test.updated_at
}
output "type" {
    value = influxdbv2_bucket.test.type
}
output "retention_rules" {
    value = influxdbv2_bucket.test.retention_rules
}
output "rp" {
    value = influxdbv2_bucket.test.rp
}