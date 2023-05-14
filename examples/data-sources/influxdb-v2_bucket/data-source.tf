data "influxdb-v2_bucket" "bucket" {
  name = "testbucket"
}

output "influxdb-v2_bucket" {
  value = data.influxdb-v2_bucket.bucket
}
