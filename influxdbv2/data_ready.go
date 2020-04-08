package influxdbv2

import (
	"context"
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/lancey-energy-storage/influxdb-client-go"
	"github.com/rs/xid"
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
	influx := meta.(*influxdb.Client)
	result, err := influx.Ready(context.Background())
	if err != nil {
		return err
	}

	readyResult := &influxdb.ReadyResult{}

	readyResult = result
	var output map[string]interface{}
	marshal, _ := json.Marshal(readyResult)
	err = json.Unmarshal(marshal, &output)
	if err != nil {
		log.Printf("%s", err)
	}

	d.Set("output", output)
	id := xid.New().String()
	d.SetId(id)

	return nil
}
