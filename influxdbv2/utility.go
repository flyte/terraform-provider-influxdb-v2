package influxdbv2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/lancey-energy-storage/influxdb-client-go"
	"log"
)

func SetRetentionRules(input interface{}) ([]influxdb.RetentionRules, error) {
	result := []influxdb.RetentionRules{}
	retentionRulesSet := input.(*schema.Set).List()
	for _, retentionRule := range retentionRulesSet {
		rr, ok := retentionRule.(map[string]interface{})
		if ok {
			each := influxdb.RetentionRules{Type: rr["type"].(string), EverySeconds: rr["every_seconds"].(int)}
			log.Printf("everySeconds: %v, type: %v \n", rr["every_seconds"], rr["type"])
			result = append(result, each)
		}
	}
	return result, nil
}
