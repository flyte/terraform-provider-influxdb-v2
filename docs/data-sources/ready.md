---
layout: "influxdb-v2"
page_title: "InfluxDB V2: influxdb-v2_ready"
sidebar_current: "docs-influxdb-v2-datasource-ready"
description: |-
The influxdb-v2_ready data source returns influxdb status.
---

# influxdb-v2\_ready (Data Source)

The influxdb-v2_ready data source retrieves influxdb instance status information.
If the endpoint server is online, the function will output its URL, otherwise, the field will be empty.

## Example Usage

```hcl
data "influxdb-v2_ready" "test" {}

output "influxdb-v2_ready" {
   value = data.influxdb-v2_ready.test.output["url"]
}
```

## Argument Reference

This data source doesn't support arguments.

## Attributes Reference

The following attributes are exported:

* ``url`` - The URL of the influx instance (empty if not ready).