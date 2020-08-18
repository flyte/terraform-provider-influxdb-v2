package influxdbv2

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/influxdata/influxdb-client-go"
	"os"
	"testing"
)

func TestAccAuthorization(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccAuthorizationDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccCreateAuthorization(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("influxdbv2_authorization.acctest", "org_id", os.Getenv("INFLUXDB_V2_ORG_ID")),
					resource.TestCheckResourceAttr("influxdbv2_authorization.acctest", "description", "Acceptance test token"),
					resource.TestCheckResourceAttr("influxdbv2_authorization.acctest", "status", "inactive"),
					resource.TestCheckResourceAttrSet("influxdbv2_authorization.acctest", "user_id"),
					resource.TestCheckResourceAttr("influxdbv2_authorization.acctest", "user_org_id", os.Getenv("INFLUXDB_V2_ORG_ID")),
					resource.TestCheckResourceAttrSet("influxdbv2_authorization.acctest", "token"),
					resource.TestCheckResourceAttr("influxdbv2_authorization.acctest", "org_id", os.Getenv("INFLUXDB_V2_ORG_ID")),
					// permissions is a complex array... we'll just check it has the correct length
					resource.TestCheckResourceAttr("influxdbv2_authorization.acctest", "permissions.#", "2"),
				),
			},
			{
				Config: testAccUpdateAuthorization(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("influxdbv2_authorization.acctest", "org_id", os.Getenv("INFLUXDB_V2_ORG_ID")),
					resource.TestCheckResourceAttr("influxdbv2_authorization.acctest", "description", "Acceptance test token 2"),
					resource.TestCheckResourceAttr("influxdbv2_authorization.acctest", "status", "active"),
					resource.TestCheckResourceAttrSet("influxdbv2_authorization.acctest", "user_id"),
					resource.TestCheckResourceAttr("influxdbv2_authorization.acctest", "user_org_id", os.Getenv("INFLUXDB_V2_ORG_ID")),
					resource.TestCheckResourceAttrSet("influxdbv2_authorization.acctest", "token"),
					resource.TestCheckResourceAttr("influxdbv2_authorization.acctest", "org_id", os.Getenv("INFLUXDB_V2_ORG_ID")),
					// permissions is a complex array... we'll just check it has the correct length
					resource.TestCheckResourceAttr("influxdbv2_authorization.acctest", "permissions.#", "1"),
				),
			},
		},
	})
}

func testAccCreateAuthorization() string {
	return `
resource "influxdbv2_authorization" "acctest" {
	org_id = "` + os.Getenv("INFLUXDB_V2_ORG_ID") + `"
    description = "Acceptance test token"
    status = "inactive"
    permissions {
        action = "read"
        resource {
			id = "` + os.Getenv("INFLUXDB_V2_BUCKET_ID") + `"
            org_id = "` + os.Getenv("INFLUXDB_V2_ORG_ID") + `"
            type = "buckets"
        }
    }
    permissions {
        action = "write"
        resource {
            id = "` + os.Getenv("INFLUXDB_V2_BUCKET_ID") + `"
            org_id = "` + os.Getenv("INFLUXDB_V2_ORG_ID") + `"
            type = "buckets"
        }
    }
}
`
}

func testAccUpdateAuthorization() string {
	return `
resource "influxdbv2_authorization" "acctest" {
	org_id = "` + os.Getenv("INFLUXDB_V2_ORG_ID") + `"
    description = "Acceptance test token 2"
    status = "active"
    permissions {
        action = "read"
        resource {
			id = "` + os.Getenv("INFLUXDB_V2_BUCKET_ID") + `"
            org_id = "` + os.Getenv("INFLUXDB_V2_ORG_ID") + `"
            type = "buckets"
        }
    }
}
`
}

func testAccAuthorizationDestroyed(s *terraform.State) error {
	influx := influxdb2.NewClient(os.Getenv("INFLUXDB_V2_URL"), os.Getenv("INFLUXDB_V2_TOKEN"))
	result, err := influx.AuthorizationsAPI().GetAuthorizations(context.Background())
	// The only auth is from the onboarding
	if len(*result) != 1 {
		return fmt.Errorf("There should be only one remaining authorization but there are: %d", len(*result))
	}
	if err != nil {
		return fmt.Errorf("Cannot read authorization list")
	}

	return nil
}
