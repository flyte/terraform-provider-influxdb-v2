---
layout: "influxdbv2"
page_title: "InfluxDB V2: influxdbv2_ready"
sidebar_current: "docs-influxdbv2-datasource-ready"
description: |-
  The influxdbv2_ready data source allows an InfluxDB to return status informations of it.
---

# influxdbv2\_ready

The influxdbv2_ready data source allows to get the information status of the InfluxDB V2. If the endpoint server is
online, the function will output its URL.

## Example Usage

```hcl
data "influxdbv2_ready" "test" {}

output "influxdbv2_ready" {
   value = data.influxdbv2_ready.test.output["url"]
}

```