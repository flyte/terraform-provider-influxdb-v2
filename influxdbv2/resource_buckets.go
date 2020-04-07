package influxdbv2

import (
	"errors"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/lancey-energy-storage/influxdb-client-go"
	"github.com/rs/xid"
)

func resourceCreateBucket() *schema.Resource {
	return &schema.Resource{
		Create: resourceCreateBucketCreate,
		Delete: resourceCreateBucketDelete,
		Read:   resourceCreateBucketRead,
		Update: resourceCreateBucketUpdate,
		Schema: map[string]*schema.Schema{
			"description": {
				Type:     schema.TypeString,
				Required: false,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"org_id": {
				Type:     schema.TypeString,
				Required: false,
				Optional: true,
			},
			"retention_rules": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"every_seconds": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"type": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"rp": {
				Type:     schema.TypeString,
				Required: false,
				Optional: true,
			},
		},
	}
}

func resourceCreateBucketCreate(d *schema.ResourceData, meta interface{}) error {
	if d.Get("name") == "" {
		return errors.New("a name is required")
	}

	rr := d.Get("retention_rules")
	retentionRules, err := SetRetentionRules(rr)
	if err != nil {
		return err
	}
	influx := meta.(*influxdb.Client)
	_, err = influx.CreateBucket(d.Get("description").(string), d.Get("name").(string), d.Get("org_id").(string), retentionRules, d.Get("rp").(string))
	if err != nil {
		return err
	}
	id := xid.New().String()
	d.SetId(id)
	return nil
}

func resourceCreateBucketDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceCreateBucketRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceCreateBucketUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}
