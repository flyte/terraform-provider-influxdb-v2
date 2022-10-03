data "influxdb-v2_ready" "test" {}

output "influxdb-v2_ready" {
  value = data.influxdb-v2_ready.test.output["url"]
}