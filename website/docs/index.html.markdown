---
layout: "influxdb-v2"
page_title: "Provider: InfluxDB V2"
sidebar_current: "docs-influxdb-v2-index"
description: |-
  The InfluxDB V2 provider configures buckets, tokens, etc on an InfluxDB V2 server.
---

# InfluxDB V2 Provider

The InfluxDB V2 provider allows Terraform to manage
[InfluxDB v2](https://v2.docs.influxdata.com/v2.0/get-started/).

The provider configuration block accepts the following arguments:

* ``url``
    * (Optional) 
    * The root URL of a InfluxDB V2 server. May alternatively be set via the `INFLUXDB_V2_URL` environment variable.
    * Defaults to `http://localhost:8086/`.
* ``token``
    * (Optional)
    * The token of the Influwdb V2 account. May alternatively be set via the `INFLUXDB_V2_TOKEN` environment variable.
   
## Example Usage

```hcl
provider "influxdb-v2" {
  url = "http://influxdb.example.com:8086"
  token = "influxdbToken"
}
```

A token can be acquired by executing the *onboarding* process, which is possible using:

* influx GUI, API or command line (manually)
* the dedicated provider (terraform) available [here](https://github.com/lancey-energy-storage/terraform-provider-influxdb-v2-onboarding)
