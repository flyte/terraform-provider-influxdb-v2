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

func TestAccCreateOrganization(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccOrganizationDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccCreateOrganization(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"influxdb-v2_organization.acctest",
						"description",
						"Acceptance test organization",
					),
					resource.TestCheckResourceAttr(
						"influxdb-v2_organization.acctest",
						"name",
						"acctest",
					),
					resource.TestCheckResourceAttrSet(
						"influxdb-v2_organization.acctest",
						"id",
					),
					resource.TestCheckResourceAttrSet(
						"influxdb-v2_organization.acctest",
						"created_at",
					),
					resource.TestCheckResourceAttrSet(
						"influxdb-v2_organization.acctest",
						"updated_at",
					),
					testAccCheckOrganizationUpdate("influxdb-v2_organization.acctest"),
				),
			},
			{
				Config: testAccUpdateOrganization(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"influxdb-v2_organization.acctest",
						"description",
						"Acceptance test organization 2",
					),
					resource.TestCheckResourceAttr(
						"influxdb-v2_organization.acctest",
						"name",
						"acctest",
					),
					resource.TestCheckResourceAttrSet(
						"influxdb-v2_organization.acctest",
						"id",
					),
					resource.TestCheckResourceAttrSet(
						"influxdb-v2_organization.acctest",
						"created_at",
					),
					resource.TestCheckResourceAttrSet(
						"influxdb-v2_organization.acctest",
						"updated_at",
					),
					testAccCheckOrganizationUpdate("influxdb-v2_organization.acctest"),
				),
			},
		},
	})
}

var lastOrgUpdate = ""

func testAccCheckOrganizationUpdate(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Resource %s doesn't exist", n)
		}

		if lastOrgUpdate == rs.Primary.Attributes["updated_at"] {
			return fmt.Errorf("updated_at has not changed since last execution")
		}
		lastOrgUpdate = rs.Primary.Attributes["updated_at"]

		return nil
	}
}

func testAccCreateOrganization() string {
	return `
resource "influxdb-v2_organization" "acctest" {
    description = "Acceptance test organization" 
    name = "acctest" 
}
`
}

func testAccUpdateOrganization() string {
	return `
resource "influxdb-v2_organization" "acctest" {
    description = "Acceptance test organization 2" 
    name = "acctest" 
}
`
}

func testAccOrganizationDestroyed(s *terraform.State) error {
	influx := influxdb2.NewClient(
		os.Getenv("INFLUXDB_V2_URL"),
		os.Getenv("INFLUXDB_V2_TOKEN"),
	)
	result, err := influx.OrganizationsAPI().GetOrganizations(context.Background())
	// The only organization is from the onboarding
	if len(*result) != 1 {
		return fmt.Errorf(
			"There should be only one remaining organization but there are: %d",
			len(*result),
		)
	}
	if err != nil {
		return fmt.Errorf("Cannot read organization list")
	}

	return nil
}
