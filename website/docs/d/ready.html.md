---
layout: "influxdbv2"
page_title: "InfluxDB V2: influxdbv2_ready"
sidebar_current: "docs-influxdbv2-datasource-ready"
description: |-
  The influxdbv2_ready data source returns influxdb status.
---

# influxdbv2\_ready

The influxdbv2_ready data source retrieves influxdb instance status information.
If the endpoint server is online, the function will output its URL, otherwise, the field will be empty.

## Example Usage

```hcl
data "influxdbv2_ready" "test" {}

output "influxdbv2_ready" {
   value = data.influxdbv2_ready.test.output["url"]
}
```

## Argument Reference

This data source doesn't support arguments.

## Attributes Reference

The following attributes are exported:

* ``url`` - The URL of the influx instance (empty if not ready).
