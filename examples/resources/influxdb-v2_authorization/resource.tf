locals {
  org_id = "example_org_id"
  bucket_id = "example_bucket_id"
}

resource "influxdb-v2_authorization" "example_authorization" {
  org_id      = local.org_id
  description = "Example description"
  status      = "active"
  permissions {
    action = "read"
    resource {
      id     = local.bucket_id
      org_id = local.org_id
      type   = "buckets"
    }
  }
  permissions {
    action = "write"
    resource {
      id     = local.bucket_id
      org_id = local.org_id
      type   = "buckets"
    }
  }
}