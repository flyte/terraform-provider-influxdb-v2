package influxdbv2

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{},

		DataSourcesMap: map[string]*schema.Resource{
			"influxdbv2_ready": dataSourceInfluxdbReady(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{}

	return config.Client()
}
