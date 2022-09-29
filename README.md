# InfluxDB v2 Provider

The InfluxDB v2 provider allows Terraform to manage
[InfluxDB v2](https://www.influxdata.com/products/influxdb-overview/).

## How to use

### Download the provider

    cd examples
    mkdir -p terraform.d/plugins/lancey.fr/influx/influxdb-v2/0.3.0/linux_amd64/
    cd terraform.d/plugins/lancey.fr/influx/influxdb-v2/0.3.0/linux_amd64/
    wget https://github.com/hasanhakkaev/terraform-provider-influxdb-v2/releases/download/v0.3.0/terraform-provider-influxdb-v2_v0.3.0-v0.3.0-linux-amd64.tar.gz
    tar xvzf terraform-provider-influxdb-v2_v0.3.0-v0.3.0-linux-amd64.tar.gz && rm -rf terraform-provider-influxdb-v2_v0.3.0-v0.3.0-linux-amd64.tar.gz
    filename=$(echo terraform-provider-influxdb-v2*)
    chmod +x $filename
    mv "$filename" "${filename%.*}"

#### Terraform 0.13.x

Add this snippet to your code:

```hcl
terraform {
  required_providers {
    influxdb-v2 = {
      source = "lancey.fr/influx/influxdb-v2"
      version = "0.3.0"
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
* the dedicated provider (terraform) available [here](https://github.com/hasanhakkaev/terraform-provider-influxdb-v2-onboarding)

### Available functionalities

Documentation is available in [website/docs/](website/docs/).
Influxdb v2 api documentation is available [here](https://v2.docs.influxdata.com/v2.0/api/).

#### Data sources

* ready (status of the influxdb-v2 instance)

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
go build -o terraform-provider-influxdb-v2
```
