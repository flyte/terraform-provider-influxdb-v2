package influxdbv2

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// TestAccReadBucket tests the read bucket data source
func TestAccReadBucket(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testDataSourceBucketConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.influxdb-v2_bucket.by_name", "name", "AcctestBucket"),
					resource.TestCheckResourceAttr("data.influxdb-v2_bucket.by_name", "description", "Desc Acctest"),
				),
			},
		},
	})
}

func testDataSourceBucketConfig() string {
	return `resource "influxdb-v2_bucket" "bucket" {
			name = "AcctestBucket"
			description = "Desc Acctest"
			org_id = "` + os.Getenv("INFLUXDB_V2_ORG_ID") + `"
    		retention_rules {
        		every_seconds = "3640"
    			}
			}
			data "influxdb-v2_bucket" "by_name" {
				name = influxdb-v2_bucket.bucket.name
				depends_on = [influxdb-v2_bucket.bucket]
			}
`
}
