---
layout: "influxdbv2"
page_title: "InfluxDB V2: influxdbv2_bucket"
sidebar_current: "docs-influxdbv2-resource-bucket"
description: |-
  The influxdbv2_authorization resource manages influxdb v2 buckets.
---

## Example Usage

```hcl
resource "influxdbv2_bucket" "sensor_data" {
    name = "Sensor Data"
    description = "A bucket for all my sensors"
    org_id = "94d518926178fea7"
    retention_rules {
        every_seconds = 45
    }
}
```

## Argument Reference

The following arguments are supported: 

* ``name`` (Required) The name of the bucket.
* ``org_id`` (Required) The organization id to which the bucket is linked.
* ``retention_rules`` (Required) Retention rules that affect the bucket.
    * ``every_seconds`` (Required) How many seconds the rule should be applied.
* ``description`` (Optional) The description of the bucket.
* ``rp`` (Optional) As of now, the influxdb documentation doesn't say what this paramenter is for.

## Attributes Reference

In addition to the above arguments, the following attributes are exported:

* created_at - The date the bucket has been created.
* updated_at - The date the bucket has been updated.
* type - The type of bucket.
