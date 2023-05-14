package influxdbv2

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/domain"
)

func dataSourceBucket() *schema.Resource {
	return &schema.Resource{
		Description: "Lookup a Bucket in InfluxDB2.",
		ReadContext: dataSourceBucketRead,

		Schema: mergeSchemas(map[string]*schema.Schema{
			"id": {
				Description: "Bucket id.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"name": {
				Description: "Bucket name.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"org_id": {
				Description: "ID of organization in which to create a bucket.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"description": {
				Description: "Description of the bucket.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"retention_rules": {
				Description: "Rules to expire or retain data. No rules means data never expires.",
				Type:        schema.TypeSet,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"every_seconds": {
							Description: "Duration in seconds for how long data will be kept in the database. 0 means infinite.",
							Type:        schema.TypeInt,
							Required:    true,
						},
						"shard_group_duration_seconds": {
							Description: "Duration in seconds for how long each shard group will last before it gets rotated out.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"type": {
							Description: "Retention rule type.",
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
			"type": {
				Description: "Bucket type.",
				Type:        schema.TypeString,
				Computed:    true,
			},
		}, createUpdatedSchema("Bucket")),
	}
}

func dataSourceBucketRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	influx := meta.(influxdb2.Client)
	bucketAPI := influx.BucketsAPI()

	// Warning or errors can be collected in a slice type
	var (
		diags  diag.Diagnostics
		bucket *domain.Bucket
		err    error
	)

	if v, ok := d.GetOk("name"); ok {
		bucketName := v.(string)
		if bucket, err = bucketAPI.FindBucketByName(ctx, bucketName); err != nil {
			diags = append(diags, diag.FromErr(err)...)
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  fmt.Sprintf("Can't find Bucket with name: %s", bucketName),
			})
			return diags
		}
	}
	diags = setBucketData(d, bucket)
	d.Set("id", bucket.Id)
	d.SetId(*bucket.Id)
	//id := bucket.Id
	//if id == nil {
	//	diags = append(diags, diag.Diagnostic{
	//		Severity: diag.Error,
	//		Summary:  "Bucket not found",
	//	})
	//	return diags
	//}
	//
	//d.SetId(*id)
	//err = d.Set("id", *id)
	//if err != nil {
	//	return diag.FromErr(err)
	//}
	//if bucket.Description != nil {
	//	err := d.Set("description", *bucket.Description)
	//	if err != nil {
	//		return diag.FromErr(err)
	//	}
	//}
	//err = d.Set("name", bucket.Name)
	//if err != nil {
	//	return diag.FromErr(err)
	//}
	//
	//err = d.Set("type", bucket.Type)
	//if err != nil {
	//	return diag.FromErr(err)
	//}
	//
	//retentionRules := make([]interface{}, len(bucket.RetentionRules))
	//for i, rule := range bucket.RetentionRules {
	//	ruleMap := map[string]interface{}{
	//		"every_seconds": int(rule.EverySeconds),
	//		"type":          rule.Type,
	//	}
	//	if rule.ShardGroupDurationSeconds != nil && *rule.ShardGroupDurationSeconds > rule.EverySeconds {
	//		ruleMap["shard_group_duration_seconds"] = int(*rule.ShardGroupDurationSeconds) // convert *int64 to int
	//	} else {
	//		ruleMap["shard_group_duration_seconds"] = 0
	//	}
	//	retentionRules[i] = ruleMap
	//}
	//if err := d.Set("retention_rules", schema.NewSet(schema.HashResource(dataSourceBucket().Schema["retention_rules"].Elem.(*schema.Resource)), retentionRules)); err != nil {
	//	return diag.FromErr(err)
	//}

	return diags
}

func setBucketData(data *schema.ResourceData, bucket *domain.Bucket) diag.Diagnostics {
	var (
		err error
	)
	err = data.Set("org_id", *bucket.OrgID)
	if err != nil {
		return diag.FromErr(err)
	}
	err = data.Set("name", bucket.Name)
	if err != nil {
		return diag.FromErr(err)
	}
	err = data.Set("description", bucket.Description)
	if err != nil {
		return diag.FromErr(err)
	}
	err = data.Set("created_at", bucket.CreatedAt.Format("2006-01-02T15:04:05.000Z"))
	if err != nil {
		return diag.FromErr(err)
	}
	err = data.Set("updated_at", bucket.UpdatedAt.Format("2006-01-02T15:04:05.000Z"))
	if err != nil {
		return diag.FromErr(err)
	}
	err = data.Set("updated_timestamp", bucket.UpdatedAt.UnixNano()/1000000)
	if err != nil {
		return diag.FromErr(err)
	}
	err = data.Set("created_timestamp", bucket.CreatedAt.UnixNano()/1000000)
	if err != nil {
		return diag.FromErr(err)
	}
	err = data.Set("type", bucket.Type)
	if err != nil {
		return diag.FromErr(err)
	}

	var retentionRules []map[string]interface{}
	for _, rule := range bucket.RetentionRules {
		mapped := map[string]interface{}{
			"every_seconds":                rule.EverySeconds,
			"shard_group_duration_seconds": rule.ShardGroupDurationSeconds,
			"type":                         rule.Type,
		}
		retentionRules = append(retentionRules, mapped)
	}

	err = data.Set("retention_rules", retentionRules)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}
