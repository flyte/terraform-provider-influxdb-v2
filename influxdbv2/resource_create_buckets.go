package influxdbv2

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/lancey-energy-storage/influxdb-client-go"
)

func ResourceBucket() *schema.Resource {
	return &schema.Resource{
		Create: resourceBucketCreate,
		Delete: resourceBucketDelete,
		Read:   resourceBucketRead,
		Update: resourceBucketUpdate,
		Schema: map[string]*schema.Schema{
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"org_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"retention_rules": {
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"every_seconds": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"type": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "expire",
						},
					},
				},
			},
			"rp": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"created_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"updated_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceBucketCreate(d *schema.ResourceData, meta interface{}) error {
	influx := meta.(*influxdb.Client)
	rr := d.Get("retention_rules")
	retentionRules, err := getRetentionRules(rr)
	if err != nil {
		return fmt.Errorf("error getting retention rules: %v", err)
	}

	result, err := influx.CreateBucket(d.Get("description").(string), d.Get("name").(string), d.Get("org_id").(string), retentionRules, d.Get("rp").(string))
	if err != nil {
		return fmt.Errorf("error creating bucket: %v", err)
	}

	d.SetId(result.Id)
	return resourceBucketRead(d, meta)
}

func resourceBucketDelete(d *schema.ResourceData, meta interface{}) error {
	influx := meta.(*influxdb.Client)
	err := influx.DeleteABucket(d.Id())
	if err != nil {
		return fmt.Errorf("error deleting bucket: %v", err)
	}
	d.SetId("")
	return nil
}

func resourceBucketRead(d *schema.ResourceData, meta interface{}) error {
	influx := meta.(*influxdb.Client)
	result, err := influx.GetBucketByID(d.Id())
	if err != nil {
		return fmt.Errorf("error getting bucket: %v", err)
	}
	d.Set("name", result.Name)
	d.Set("description", result.Description)
	d.Set("org_id", result.OrgID)
	d.Set("rr_every_seconds", result.RetentionRules[0].EverySeconds)
	d.Set("rr_type", result.RetentionRules[0].Type)

	d.Set("type", result.Type)
	d.Set("created_at", result.CreatedAt)
	d.Set("updated_at", result.UpdatedAt)

	return nil
}

func resourceBucketUpdate(d *schema.ResourceData, meta interface{}) error {
	influx := meta.(*influxdb.Client)

	retentionRules, err := getRetentionRules(d.Get("retention_rules"))
	if err != nil {
		return fmt.Errorf("error getting retention rules: %v", err)
	}

	labels := []influxdb.Labels{}

	_, err = influx.UpdateABucket(d.Id(), d.Get("description").(string), labels, d.Get("name").(string), d.Get("org_id").(string), retentionRules, d.Get("rp").(string))
	if err != nil {
		return fmt.Errorf("error updating bucket: %v", err)
	}

	return nil
}

func getRetentionRules(input interface{}) ([]influxdb.RetentionRules, error) {
	result := []influxdb.RetentionRules{}
	retentionRulesSet := input.(*schema.Set).List()
	for _, retentionRule := range retentionRulesSet {
		rr, ok := retentionRule.(map[string]interface{})
		if ok {
			each := influxdb.RetentionRules{Type: rr["type"].(string), EverySeconds: rr["every_seconds"].(int)}
			result = append(result, each)
		}
	}
	return result, nil
}
