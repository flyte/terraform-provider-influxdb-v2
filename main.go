package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/lancey-energy-storage/terraform-provider-influxdb-v2/influxdb-v2"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: influxdb-v2.Provider})
}
