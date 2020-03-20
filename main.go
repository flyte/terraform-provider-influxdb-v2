package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/lancey-energy-storage/terraform-provider-influxdb-v2/influxdbv2"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: influxdbv2.Provider})
}
