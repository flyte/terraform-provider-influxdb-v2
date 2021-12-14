package influxdbv2

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func TestAccCreateBucket(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccBucketDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccCreateBucket(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"influxdb-v2_bucket.acctest",
						"org_id",
						os.Getenv("INFLUXDB_V2_ORG_ID"),
					),
					resource.TestCheckResourceAttr(
						"influxdb-v2_bucket.acctest",
						"description",
						"Acceptance test bucket",
					),
					resource.TestCheckResourceAttr(
						"influxdb-v2_bucket.acctest",
						"name",
						"acctest",
					),
					resource.TestCheckResourceAttr(
						"influxdb-v2_bucket.acctest",
						"rp",
						"",
					),
					resource.TestCheckResourceAttr(
						"influxdb-v2_bucket.acctest",
						"retention_rules.0.every_seconds",
						"3640",
					),
					resource.TestCheckResourceAttrSet(
						"influxdb-v2_bucket.acctest",
						"created_at",
					),
					resource.TestCheckResourceAttrSet(
						"influxdb-v2_bucket.acctest",
						"updated_at",
					),
					testAccCheckUpdate("influxdb-v2_bucket.acctest"),
					resource.TestCheckResourceAttrSet(
						"influxdb-v2_bucket.acctest",
						"type",
					),
				),
			},
			{
				Config: testAccUpdateBucket(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"influxdb-v2_bucket.acctest",
						"org_id",
						os.Getenv("INFLUXDB_V2_ORG_ID"),
					),
					resource.TestCheckResourceAttr(
						"influxdb-v2_bucket.acctest",
						"description",
						"Acceptance test bucket 2",
					),
					resource.TestCheckResourceAttr(
						"influxdb-v2_bucket.acctest",
						"name",
						"acctest",
					),
					resource.TestCheckResourceAttr(
						"influxdb-v2_bucket.acctest",
						"rp",
						"",
					),
					resource.TestCheckResourceAttr(
						"influxdb-v2_bucket.acctest",
						"retention_rules.0.every_seconds",
						"3630",
					),
					resource.TestCheckResourceAttrSet(
						"influxdb-v2_bucket.acctest",
						"created_at",
					),
					resource.TestCheckResourceAttrSet(
						"influxdb-v2_bucket.acctest",
						"updated_at",
					),
					testAccCheckUpdate("influxdb-v2_bucket.acctest"),
					resource.TestCheckResourceAttrSet(
						"influxdb-v2_bucket.acctest",
						"type",
					),
				),
			},
		},
	})
}

var lastUpdate = ""

func testAccCheckUpdate(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Resource %s doesn't exist", n)
		}

		if lastUpdate == rs.Primary.Attributes["updated_at"] {
			return fmt.Errorf("updated_at has not changed since last execution")
		}
		lastUpdate = rs.Primary.Attributes["updated_at"]

		return nil
	}
}

func testAccCreateBucket() string {
	return `
resource "influxdb-v2_bucket" "acctest" {
    description = "Acceptance test bucket" 
    name = "acctest" 
    org_id = "` + os.Getenv("INFLUXDB_V2_ORG_ID") + `"
    retention_rules {
        every_seconds = "3640"
    }
}
`
}

func testAccUpdateBucket() string {
	return `
resource "influxdb-v2_bucket" "acctest" {
    description = "Acceptance test bucket 2" 
    name = "acctest" 
    org_id = "` + os.Getenv("INFLUXDB_V2_ORG_ID") + `"
    retention_rules {
        every_seconds = "3630"
    }
}
`
}

func testAccBucketDestroyed(s *terraform.State) error {
	influx := influxdb2.NewClient(
		os.Getenv("INFLUXDB_V2_URL"),
		os.Getenv("INFLUXDB_V2_TOKEN"),
	)
	result, err := influx.BucketsAPI().GetBuckets(context.Background())
	// The only buckets are from the onboarding, plus _monitoring and _tasks
	if len(*result) != 3 {
		return fmt.Errorf(
			"There should be only one remaining bucket but there are: %d",
			len(*result),
		)
	}
	if err != nil {
		return fmt.Errorf("Cannot read bucket list")
	}

	return nil
}
