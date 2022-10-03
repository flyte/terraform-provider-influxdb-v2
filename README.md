# InfluxDB v2 Provider

The InfluxDB v2 provider allows Terraform to manage
[InfluxDB v2](https://www.influxdata.com/products/influxdb-overview/).

## How to use


#### Terraform 0.13.x

Add this snippet to your code:

```hcl
terraform {
  required_providers {
    influxdb-v2 = {
      source = "hasanhakkaev/influxdb-v2"
      version = "0.4.4"
    }
  }
}
```

### Initialize the provider

```hcl
provider "influxdb-v2" {
  url = "http://influxdb.example.com:8086"
  token = "influxdbToken"
}
```

The provider configuration block accepts the following arguments:

* ``url`` (Optional) The root URL of a InfluxDB V2 server. May alternatively be set via the `INFLUXDB_V2_URL` environment variable. Defaults to `http://localhost:8086/`.

* ``token`` (Optional) The token that gives access to the influxdb instance. May alternatively be set via the `INFLUXDB_V2_TOKEN` environment variable.

A token can be acquired by executing the *onboarding* process, which is possible using:

* influx GUI, API or command line (manually)

### Available functionalities

Documentation is available in [website/docs/](website/docs/).
Influxdb v2 api documentation is available [here](https://v2.docs.influxdata.com/v2.0/api/).

#### Data sources

* ready (status of the influxdb-v2 instance)
* organization (get an organization by name)

#### Resources

* bucket
* authorization (tokens)
* organization

### Examples

Find examples in `examples/`. To run them:



## Development
### Requirements
    
    TaskFile (https://taskfile.dev/)
    Precommit framework (https://pre-commit.com/)
    Docker (https://www.docker.com/)
    Go 1.16^ (https://golang.org/doc/install)
    Terraform 0.13^ (https://www.terraform.io/downloads.html)
    InfluxDB 2.0^ (https://docs.influxdata.com/influxdb/v2.0/install/)

```bash
# Run Task init to download dependencies
task init
```

### Test

First execute this command to lint and format the code:

```bash
task lint
```

To run acceptance tests, execute these commands (requires `docker` and `jq`): 

```bash
task start-influx
task test
task stop-influx
```

### Build

```bash
task build
```
