package influxdbv2

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// TestAccReadOrganization tests the read organization data source
func TestAccReadOrganization(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testDataSourceOrganizationConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.influxdb-v2_organization.by_name", "name", "AcctestName"),
					resource.TestCheckResourceAttr("data.influxdb-v2_organization.by_name", "description", "Desc Acctest"),
				),
			},
		},
	})
}

func testDataSourceOrganizationConfig() string {
	return `resource "influxdb-v2_organization" "org" {
			name = "AcctestName"
			description = "Desc Acctest"
		}
		data "influxdb-v2_organization" "by_name" {
			name = influxdb-v2_organization.org.name
			//requirement of terraform v0.13
			depends_on = [influxdb-v2_organization.org]
		}
`
}
