---
layout: "influxdbv2"
page_title: "InfluxDB V2: influxdbv2_ready"
sidebar_current: "docs-influxdbv2-datasource-ready"
description: |-
  The influxdbv2_ready data source allows an InfluxDB to return status informations of it.
---

# influxdbv2\_ready

The influxdbv2_ready data source allows to get the information status of the InfluxDB V2.

## Example Usage

```hcl
data "influxdbv2_ready" "test" {}

output "influxdbv2_ready" {
   value = data.influxdbv2_ready.test.output["status"]
}

output "ready_started"  {
    value = data.influxdbv2_ready.test.output["started"]
}

output "ready_up"  {
    value = data.influxdbv2_ready.test.output["up"]
}

```