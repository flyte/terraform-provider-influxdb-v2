package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/lancey-energy-storage/terraform-provider-influxdb-v2/influxdb"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: influxdb.Provider})
}
