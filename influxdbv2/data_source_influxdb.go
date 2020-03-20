package influxdbv2

import (
	"encoding/json"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/rs/xid"
	"io/ioutil"
	"log"
	"net/http"
)

func dataSourceInfluxdbReady() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceInfluxdbReadReady,
		Schema: map[string]*schema.Schema{
			"url": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"output": {
				Type:     schema.TypeMap,
				Computed: true,
			},
		},
	}
}

func dataSourceInfluxdbReadReady(d *schema.ResourceData, meta interface{}) error {
	var r Ready
	log.Printf("[DEBUG] Pinging resource ")
	url := d.Get("url").(string)
	req, err := http.Get(url + "/ready")

	if err != nil {
		return err
	}

	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)
	err = json.Unmarshal(body, &r)
	if err != nil {
		return err
	}

	output := make(map[string]string)
	output["status"] = r.Status
	output["started"] = r.Started
	output["up"] = r.Up

	d.Set("output", output)
	id := xid.New().String()
	d.SetId(id)

	return nil
}
