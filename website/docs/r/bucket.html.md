---
layout: "influxdbv2"
page_title: "InfluxDB V2: influxdbv2_bucket"
sidebar_current: "docs-influxdbv2-resource-bucket"
description: |-
  The influxdbv2_bucket resource allows an InfluxDB to create, update and delete buckets.
---

## Example Usage

###Create a new bucket
```hcl
resource "influxdbv2_bucket" "initial" {
    name = "New Bucket"
    description = "Description of new bucket"
    org_id = "94d518926178fea7"
    retention_rules {
        every_seconds = 45
        type = "expire"
    }
    rp = ""
}
```

###Output usage
```hcl
output "created_at" {
    value = influxdbv2_bucket.initial.created_at
}
output "updated_at" {
    value = influxdbv2_bucket.initial.updated_at
}
output "type" {
    value = influxdbv2_bucket.initial.type
}
```

##Argument Reference

The following arguments are supported: 

* ``name`` (Required) The name of the bucket.

* ``description`` (Optional) The description of the bucket.

* ``org_id`` (Optional) The organization id the bucket is linked.

* ``retention_rules`` (Required) Retention rules that affects the bucket.
    * ``every_seconds`` (Required) How many seconds the rule should be applied.
    * ``type`` (Optional) The type of the retention. Default is "expire".

##Attributes Reference

In addition to the above arguments, the following attributes are exported:

* created_at - The date the bucket has been created.
* updated_at - The date the bucket has been updated.
* type - The type of bucket.

