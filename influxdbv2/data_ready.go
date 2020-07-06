package influxdbv2

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/influxdata/influxdb-client-go"
	"log"
)

func DataReady() *schema.Resource {
	return &schema.Resource{
		Read: DataGetReady,
		Schema: map[string]*schema.Schema{
			"output": {
				Type:     schema.TypeMap,
				Computed: true,
			},
		},
	}
}

func DataGetReady(d *schema.ResourceData, meta interface{}) error {
	influx := meta.(influxdb2.Client)
	bool, error := influx.Ready(context.Background())
	if error != nil {
		return error
	}
	if bool {
		log.Printf("Server is ready !")
	}

	output := map[string]string{
		"url": influx.ServerUrl(),
	}

	id := ""
	id = influx.ServerUrl()
	d.SetId(id)
	d.Set("output", output)

	return nil
}
