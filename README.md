# InfluxDB v2 Provider

The InfluxDB v2 provider allows Terraform to manage
[InfluxDB v2](https://www.influxdata.com/products/influxdb-overview/).

## How to use

### Download the provider

#### Terraform 0.12.x

Download the release and extract it to (on linux):  
`~/.terraform.d/plugins/linux_amd64/terraform-provider-influxdbv2_v0.3.0`

#### Terraform 0.13.x

Add this snippet to your code:

```hcl
terraform {
  required_providers {
    influxdbv2 = {
      source = "lancey-energy-storage/influxdbv2"
      version = "0.3.0"
    }
  }
}
```

Until the provider is available on registry.terraform.io, you need to manually download the release and extract it to, eg (on linux):   
`~/.terraform.d/plugins/registry.terraform.io/lancey-energy-storage/influxdbv2/0.3.0/linux_amd64/terraform-provider-influxdbv2_v0.3.0`

### Initialize the provider

```hcl
provider "influxdbv2" {
  url = "http://influxdb.example.com:9999"
  token = "influxdbToken"
}
```

The provider configuration block accepts the following arguments:

* ``url`` (Optional) The root URL of a InfluxDB V2 server. May alternatively be set via the `INFLUXDB_V2_URL` environment variable. Defaults to `http://localhost:9999/`.

* ``token`` (Optional) The token that gives access to the influxdb instance. May alternatively be set via the `INFLUXDB_V2_TOKEN` environment variable.

A token can be acquired by executing the *onboarding* process, which is possible using:

* influx GUI, API or command line (manually)
* the dedicated provider (terraform) available [here](https://github.com/lancey-energy-storage/terraform-provider-influxdb-v2-onboarding)

### Available functionalities

Documentation is available in [website/docs/](website/docs/).
Influxdb v2 api documentation is available [here](https://v2.docs.influxdata.com/v2.0/api/).

#### Data sources

* ready (status of the influxdbv2 instance)

#### Resources

* bucket

* authorization (tokens)

### Examples

Find examples in `examples/`. To run them:

```bash
source ./start_influxdb.sh
terraform init
terraform apply
./stop_influxdb.sh
```

## Dev

In case you need to update the influx client, run `go get github.com/influxdata/influxdb-client-go@<commit sha>`.  
Also don't forget to run `go mod tidy` from time to time to remove useless dependencies.

### Test

First execute this command to check fmt requirements:
 
```bash
make fmt
```

Then execute this command to run the provider unit tests:

```bash
make test
```

And finally to run acceptance tests, execute these commands (requires `docker` and `jq`): 

```bash
source ./scripts/setup_influxdb.sh
make testacc
make stop-influx
```

### Build

```bash
go build -o terraform-provider-influxdbv2
```
