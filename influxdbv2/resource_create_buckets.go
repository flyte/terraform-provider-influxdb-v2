package influxdbv2

import (
	"errors"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/lancey-energy-storage/influxdb-client-go"
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
	influx := meta.(*influxdb.Client)

	if d.Get("name") == "" {
		return errors.New("a name is required")
	}

	rr := d.Get("retention_rules")
	retentionRules, err := SetRetentionRules(rr)
	if err != nil {
		return err
	}

	result, err := influx.CreateBucket(d.Get("description").(string), d.Get("name").(string), d.Get("org_id").(string), retentionRules, d.Get("rp").(string))
	if err != nil {
		return err
	}

	bucketCreated := &influxdb.BucketCreate{}
	bucketCreated = result
	d.SetId(bucketCreated.Id)

	return nil
}

func resourceCreateBucketDelete(d *schema.ResourceData, meta interface{}) error {
	influx := meta.(*influxdb.Client)
	err := influx.DeleteABucket(d.Id())
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}

func resourceCreateBucketRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceCreateBucketUpdate(d *schema.ResourceData, meta interface{}) error {
	influx := meta.(*influxdb.Client)

	if d.Get("name") == "" {
		return errors.New("a name is required")
	}

	retentionRules, err := SetRetentionRules(d.Get("retention_rules"))
	if err != nil {
		return err
	}

	labels, err := SetLabels(d.Get("labels"))
	if err != nil {
		return err
	}

	_, err = influx.UpdateABucket(d.Id(), d.Get("description").(string), labels, d.Get("name").(string), d.Get("org_id").(string), retentionRules, d.Get("rp").(string))
	if err != nil {
		return err
	}

	return nil
}
