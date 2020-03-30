---
layout: "influxdbv2"
page_title: "Provider: InfluxDB V2"
sidebar_current: "docs-influxdbv2-index"
description: |-
  The InfluxDB V2 provider configures databases, etc on an InfluxDB server.
---

# InfluxDB V2 Provider

The InfluxDB V2 provider allows Terraform to create Databases in
[InfluxDB](https://influxdb.com/). InfluxDB is a database server optimized
for time-series data.

The provider configuration block accepts the following arguments:

* ``url`` - (Optional) The root URL of a InfluxDB V2 server. May alternatively be
  set via the ``INFLUXDB_URL`` environment variable. Defaults to
  `http://localhost:9999/`.

## Example Usage

```hcl
provider "influxdb" {
  url      = "http://influxdb.example.com:9999"
}