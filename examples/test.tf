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

output "ready_started" {
    value = data.influxdbv2_ready.test.output["started"]
}

output "ready_up" {
    value = data.influxdbv2_ready.test.output["up"]
}
