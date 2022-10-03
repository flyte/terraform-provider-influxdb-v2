---
layout: "influxdb-v2"
page_title: "InfluxDB V2: influxdb-v2_organization"
sidebar_current: "docs-influxdb-v2-datasource-organization"
description: |-
The influxdb-v2_organization data source returns influxdb status.
---

# influxdb-v2\_organization (Data Source)

The influxdb-v2_organization data source retrieves influxdb organization information.

## Example Usage

```hcl
data "influxdb-v2_organization" "organization" {
    name = "my-org"
}

output "influxdb-v2_organization_id" {
   value = data.influxdb-v2_organization.organization.id
}
```

## Argument Reference

* ``name`` (Required) The organization name


## Attributes Reference

The following attributes are exported:

* ``id`` - The ID of the Influx organization.
* ``name`` - The name of the Influx organization.
* ``description`` - The description of the Influx organization.