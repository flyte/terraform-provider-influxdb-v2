package influxdbv2

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"testing"
)

func TestCreateBucket(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccConfigBucket,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCreate("influxdbv2_bucket.test"),
				),
			},
		},
	})
}

func testAccCheckCreate(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found %s", n)
		}
		if rs.Primary.Attributes["created_at"] == "" {
			return fmt.Errorf("the bucket has not been created")
		}
		return nil
	}
}

var testAccConfigBucket = `
resource "influxdbv2_bucket" "test" {
    description = "Le bucket terraform"
    name = "le bucket de test terraform"
    org_id = influxdbv2-onboarding_setup.setup.org_id
    retention_rules {
        every_seconds = 40
        type = "expire"
    }
    rp = ""
}
`
