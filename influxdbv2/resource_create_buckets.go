package influxdbv2

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/domain"
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
	retentionRules := getRetentionRules(d.Get("retention_rules"))
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

	// Reformat retention rules array
	var rr []map[string]interface{}
	for _, retention_rule := range result.RetentionRules {
		tmp := map[string]interface{}{
			"every_seconds": retention_rule.EverySeconds,
			"type":          retention_rule.Type,
		}
		rr = append(rr, tmp)
	}
	if len(result.RetentionRules) == 0 {
		// If no retention rules, there's a default of 0 expiry
		// but this isn't returned on the API
		tmp := map[string]interface{}{
			"every_seconds": 0,
			"type":          "expire",
		}
		rr = append(rr, tmp)
	}

	err = d.Set("name", result.Name)
	if err != nil {
		return err
	}
	err = d.Set("description", result.Description)
	if err != nil {
		return err
	}
	err = d.Set("org_id", result.OrgID)
	if err != nil {
		return err
	}
	err = d.Set("retention_rules", rr)
	if err != nil {
		return err
	}
	err = d.Set("rp", result.Rp)
	if err != nil {
		return err
	}
	err = d.Set("created_at", result.CreatedAt.String())
	if err != nil {
		return err
	}
	err = d.Set("updated_at", result.UpdatedAt.String())
	if err != nil {
		return err
	}
	err = d.Set("type", result.Type)
	if err != nil {
		return err
	}

	return nil
}

func resourceBucketUpdate(d *schema.ResourceData, meta interface{}) error {
	influx := meta.(influxdb2.Client)
	retentionRules := getRetentionRules(d.Get("retention_rules"))

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
	var err error
	_, err = influx.BucketsAPI().UpdateBucket(context.Background(), updateBucket)

	if err != nil {
		return fmt.Errorf("error updating bucket: %v", err)
	}

	return resourceBucketRead(d, meta)
}

func getRetentionRules(input interface{}) domain.RetentionRules {
	result := domain.RetentionRules{}
	retentionRulesSet := input.(*schema.Set).List()
	for _, retentionRule := range retentionRulesSet {
		rr, ok := retentionRule.(map[string]interface{})
		if ok {
			each := domain.RetentionRule{EverySeconds: int64(rr["every_seconds"].(int))}
			result = append(result, each)
		}
	}
	return result
}
