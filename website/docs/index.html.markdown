---
layout: "influxdbv2"
page_title: "Provider: InfluxDB V2"
sidebar_current: "docs-influxdbv2-index"
description: |-
  The InfluxDB V2 provider configures onboarding, buckets, tokens, etc on an InfluxDB V2 server.
---

# InfluxDB V2 Provider

The InfluxDB V2 provider allows Terraform to manage
[InfluxDB v2](https://www.influxdata.com/products/influxdb-overview/).

The provider configuration block accepts the following arguments:

* ``url``
    * (Optional) 
    * The root URL of a InfluxDB V2 server. May alternatively be set via the `INFLUXDB_V2_URL` environment variable.
    * Defaults to `http://localhost:9999/`.

## Example Usage

```hcl
provider "influxdbv2" {
  url = "http://influxdb.example.com:9999"
}
 ```
 