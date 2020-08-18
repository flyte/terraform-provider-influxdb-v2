package influxdbv2

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/influxdata/influxdb-client-go"
	"github.com/influxdata/influxdb-client-go/domain"
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
				Required: true,
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
			"created_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"updated_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceBucketCreate(d *schema.ResourceData, meta interface{}) error {
	influx := meta.(influxdb2.Client)
	retentionRules, err := getRetentionRules(d.Get("retention_rules"))
	if err != nil {
		return fmt.Errorf("error getting retention rules: %v", err)
	}
	desc := d.Get("description").(string)
	orgid := d.Get("org_id").(string)
	retpe := d.Get("rp").(string)
	newBucket := &domain.Bucket{
		Description:    &desc,
		Name:           d.Get("name").(string),
		OrgID:          &orgid,
		RetentionRules: retentionRules,
		Rp:             &retpe,
	}
	result, err := influx.BucketsAPI().CreateBucket(context.Background(), newBucket)
	if err != nil {
		return fmt.Errorf("error creating bucket: %v", err)
	}
	d.SetId(*result.Id)
	return resourceBucketRead(d, meta)
}

func resourceBucketDelete(d *schema.ResourceData, meta interface{}) error {
	influx := meta.(influxdb2.Client)
	err := influx.BucketsAPI().DeleteBucketWithID(context.Background(), d.Id())
	if err != nil {
		return fmt.Errorf("error deleting bucket: %v", err)
	}
	d.SetId("")
	return nil
}

func resourceBucketRead(d *schema.ResourceData, meta interface{}) error {
	influx := meta.(influxdb2.Client)
	result, err := influx.BucketsAPI().FindBucketByID(context.Background(), d.Id())
	if err != nil {
		return fmt.Errorf("error getting bucket: %v", err)
	}
	d.Set("name", result.Name)
	d.Set("description", result.Description)
	d.Set("org_id", result.OrgID)
	d.Set("retention_rules", result.RetentionRules)
	d.Set("rp", result.Rp)
	d.Set("created_at", result.CreatedAt.String())
	d.Set("updated_at", result.UpdatedAt.String())
	d.Set("type", result.Type)

	return nil
}

func resourceBucketUpdate(d *schema.ResourceData, meta interface{}) error {
	influx := meta.(influxdb2.Client)
	retentionRules, err := getRetentionRules(d.Get("retention_rules"))
	if err != nil {
		return fmt.Errorf("error getting retention rules: %v", err)
	}

	id := d.Id()
	desc := d.Get("description").(string)
	orgid := d.Get("org_id").(string)
	retpe := d.Get("rp").(string)
	updateBucket := &domain.Bucket{
		Id:             &id,
		Description:    &desc,
		Name:           d.Get("name").(string),
		OrgID:          &orgid,
		RetentionRules: retentionRules,
		Rp:             &retpe,
	}
	_, err = influx.BucketsAPI().UpdateBucket(context.Background(), updateBucket)

	if err != nil {
		return fmt.Errorf("error updating bucket: %v", err)
	}

	return resourceBucketRead(d, meta)
}

func getRetentionRules(input interface{}) (domain.RetentionRules, error) {
	result := domain.RetentionRules{}
	retentionRulesSet := input.(*schema.Set).List()
	for _, retentionRule := range retentionRulesSet {
		rr, ok := retentionRule.(map[string]interface{})
		if ok {
			each := domain.RetentionRule{EverySeconds: rr["every_seconds"].(int)}
			result = append(result, each)
		}
	}
	return result, nil
}
