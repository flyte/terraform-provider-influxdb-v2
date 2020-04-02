package influxdbv2

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/lancey-energy-storage/influxdb-client-go"
	"github.com/rs/xid"
	"log"
)

func dataReady() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceInfluxdbV2Ready,
		Schema: map[string]*schema.Schema{
			"output": {
				Type:     schema.TypeMap,
				Computed: true,
			},
		},
	}
}

func dataSourceInfluxdbV2Ready(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] Get ready status of influxdb")

	influx := meta.(*influxdb.Client)
	result, err := influx.Ready(context.Background())
	if err != nil {
		return err
	}

	output := make(map[string]*string)
	output["status"] = result.Status
	output["started"] = result.Started
	output["up"] = result.Up

	d.Set("output", output)
	id := xid.New().String()
	d.SetId(id)

	return nil
}
