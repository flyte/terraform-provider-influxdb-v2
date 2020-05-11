provider "influxdbv2-onboarding" {
    url = "http://localhost:9999"
    username = "test"
    password = "test1234"
}

resource "influxdbv2-onboarding_setup" "setup" {
    bucket = "test-bucket"
    org = "test-org"
    retention_period = 4
}

provider "influxdbv2" {
    url = "http://localhost:9999"
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

resource "influxdbv2_bucket" "initial" {
    description = ""
    name = "le bucket de test terraform"
    org_id = "1de630828fbc3210"
    retention_rules {
        every_seconds = 45
        type = "expire"
    }
    rp = ""
}

output "created_at" {
    value = influxdbv2_bucket.initial.created_at
}
output "updated_at" {
    value = influxdbv2_bucket.initial.updated_at
}
output "type" {
    value = influxdbv2_bucket.initial.type
}
output "retention_rules" {
    value = influxdbv2_bucket.initial.retention_rules
}
output "rp" {
    value = influxdbv2_bucket.initial.rp
}