# InfluxDB V2 Provider

The InfluxDB V2 provider allows Terraform to manage
[InfluxDB v2](https://www.influxdata.com/products/influxdb-overview/).

The provider configuration block accepts the following arguments:

* ``url``
    * (Optional) 
    * The root URL of a InfluxDB V2 server. May alternatively be set via the `INFLUXDB_V2_URL` environment variable.
    * Defaults to `http://localhost:9999/`.

## Build

```bash
go build -o terraform-provider-influxdbv2
```

Don't forget to copy `terraform-provider-influxdbv2` to your terraform plugin directory (eg. `~/.terraform.d/plugins/linux_amd64` on linux).

## Test

```bash
go test ./influxdbv2
```

## How to use

### Initialize the provider
```hcl
provider "influxdbv2" {
  url = "http://influxdb.example.com:9999"
  username = "influxdbUsername"
  password = "influxdbPassword"
  token = "influxdbToken"
}
 ```

###Available functionalities

* **ready** status of the influxdbv2, documentation [here](website/docs/r/ready.html.md)

* **bucket** to create, update and delete bucket resources, documentation [here](website/docs/r/bucket.html.md)


###Examples file
Find more examples in `examples/`. To run them:
```bash
terraform init
terraform apply
```

## Dev

In case you need to update the influx client, run `go get github.com/lancey-energy-storage/influxdb-client-go@<commit sha>`.  
Also don't forget to run `go mod tidy` from time to time to remove useless dependencies.