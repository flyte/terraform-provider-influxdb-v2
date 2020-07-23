package influxdbv2

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"os/exec"
	"strings"
	"testing"
)

func TestAccCreateAuthorization(t *testing.T) {
	org, _ := exec.Command("sh", "-c", "terraform output -state=../scripts/terraform.tfstate org_id").CombinedOutput()
	token, _ := exec.Command("sh", "-c", "terraform output -state=../scripts/terraform.tfstate token").CombinedOutput()

	resource.Test(t, resource.TestCase{
		Providers: testProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCreateAuthorization(strings.TrimSuffix(string(token), "\n"), strings.TrimSuffix(string(org), "\n")),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCreateAuthorization("influxdbv2_authorization.fred"),
				),
			},
		},
	})
}
func testAccCheckCreateAuthorization(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found %s", n)
		}
		if rs.Primary.Attributes["status"] != "inactive" {
			return fmt.Errorf("the authorization doesn't have correct status, should be inactive, not: %v", rs.Primary.Attributes["status"])
		}
		if rs.Primary.Attributes["description"] != "test token" {
			return fmt.Errorf("the authorization description is not correct, should be \"test token \", not: %v", rs.Primary.Attributes["description"])
		}
		if len(rs.Primary.Attributes["permissisions"]) != 2 {
			return fmt.Errorf("the authorization permissions doesn't have correct length, should be 2, not: %v", len(rs.Primary.Attributes["permissions"]))
		}
		if rs.Primary.Attributes["token"] == "" {
			return fmt.Errorf("the token is not correct, should not be empty")
		}

		return nil
	}
}

func testAccCreateAuthorization(token string, orgId string) string {
	var testAccConfigBucket = `
provider "influxdbv2" {
	url = "http://localhost:9999/"
    token = "` + token + `"
}
resource "influxdbv2_bucket" "test" {
    description = "Le bucket terraform" 
    name = "le bucket de test terraform" 
    org_id = "` + orgId + `"
    retention_rules {
        every_seconds = 40
    }
    rp = ""
}
resource "influxdbv2_authorization" "fred" {
    org_id = influxdbv2_bucket.test.org_id
    description = "test token"
    status = "inactive"
    permissions {
        action = "read"
        resource {
            id = influxdbv2_bucket.test.id
            org_id = influxdbv2_bucket.test.org_id
            type = "buckets"
        }
    }
    permissions {
        action = "write"
        resource {
            id = influxdbv2_bucket.test.id
            org_id = influxdbv2_bucket.test.org_id
            type = "buckets"
        }
    }
}
`
	return testAccConfigBucket
}
