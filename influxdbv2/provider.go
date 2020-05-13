package influxdbv2

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/lancey-energy-storage/influxdb-client-go"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		DataSourcesMap: map[string]*schema.Resource{
			"influxdbv2_ready": DataReady(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"influxdbv2_bucket": ResourceBucket(),
		},
		Schema: map[string]*schema.Schema{
			"url": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("INFLUXDB_V2_URL", "http://localhost:9999"),
			},
			"token": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("INFLUXDB_V2_TOKEN", ""),
			},
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	influx, error := influxdb.New(d.Get("url").(string), d.Get("token").(string))
	if error != nil {
		return nil, fmt.Errorf("invalid InfluxDBv2 URL: %s", error)
	}

	err := influx.Ping(context.Background())
	if err != nil {
		return nil, fmt.Errorf("error pinging server: %s", err)
	}

	return influx, nil
}