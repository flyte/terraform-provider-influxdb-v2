locals {
  org_id = "example_org_id"
}

resource "influxdb-v2_bucket" "example_bucket" {
  name        = "example_bucket_name"
  description = "Example bucket description"
  org_id      = local.org_id
  retention_rules {
    every_seconds = 0
  }
}