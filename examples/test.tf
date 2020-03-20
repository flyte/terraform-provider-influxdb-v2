data "influxdbv2_ready" "test" {
    url = "http://localhost:9999"
}

output "ready_status" {
    value = data.influxdbv2_ready.test.output["status"]
}
output "ready_started" {
    value = data.influxdbv2_ready.test.output["started"]
}
output "ready_up" {
    value = data.influxdbv2_ready.test.output["up"]
}