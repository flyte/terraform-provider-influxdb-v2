provider "influxdbv2" {
    url = "http://influxdbv2.example.com:9999"
    username = "influxdbUsername"
    password = "influxdbPassword"
    token = "influxdbToken"
}

data "influxdbv2_ready" "test" {}

output "ready_status"  {
    value = data.influxdbv2_ready.test.output["status"]
}

output "ready_started"  {
    value = data.influxdbv2_ready.test.output["started"]
}

output "ready_up"  {
    value = data.influxdbv2_ready.test.output["up"]
}