---
layout: "influxdb-v2"
page_title: "InfluxDB V2: influxdb-organization"
sidebar_current: "docs-influxdb-v2-resource-organization"
description: |-
The influxdb-organization resource manages influxdb v2 organizations.
---

## Example Usage

```hcl
resource "influxdb-organization" "organization" {
    name = "my_organization"
    description = "My organization desciption"
}
```

## Argument Reference

The following arguments are supported:

* ``name`` (Required) The organization name
* ``description`` (Required) A short description for your organization

## Attributes Reference

In addition to the above arguments, the following attributes are exported:

* ``id`` - The organization ID which is created


## Import

Import is supported using the following syntax:

```shell
terraform import influxdb-v2_organization.example_organization <ORG_ID>
```
