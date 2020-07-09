package influxdbv2

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"os/exec"
	"strings"
	"testing"
)

func TestAccCreateBucket(t *testing.T) {
	org, _ := exec.Command("sh", "-c", "terraform output -state=../scripts/terraform.tfstate org_id").CombinedOutput()
	token, _ := exec.Command("sh", "-c", "terraform output -state=../scripts/terraform.tfstate token").CombinedOutput()

	resource.Test(t, resource.TestCase{
		Providers: testProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCreateBucket(strings.TrimSuffix(string(token), "\n"), strings.TrimSuffix(string(org), "\n")),
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
		if rs.Primary.Attributes["description"] != "Le bucket terraform" {
			return fmt.Errorf("the bucket description is not correct")
		}
		if rs.Primary.Attributes["name"] != "le bucket de test terraform" {
			return fmt.Errorf("the bucket name is not correct")
		}
		if rs.Primary.Attributes["rp"] != "" {
			return fmt.Errorf("the retention policies ar not correct")
		}

		return nil
	}
}

func TestAccUpdateBucket(t *testing.T) {
	org, _ := exec.Command("sh", "-c", "terraform output -state=../scripts/terraform.tfstate org_id").CombinedOutput()
	token, _ := exec.Command("sh", "-c", "terraform output -state=../scripts/terraform.tfstate token").CombinedOutput()

	resource.Test(t, resource.TestCase{
		Providers: testProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdateBucket(strings.TrimSuffix(string(token), "\n"), strings.TrimSuffix(string(org), "\n")),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUpdate("influxdbv2_bucket.initial"),
				),
			},
		},
	})
}
func testAccCheckUpdate(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found %s", n)
		}
		if rs.Primary.Attributes["updated_at"] == "" {
			return fmt.Errorf("the bucket has not been updated")
		}
		if rs.Primary.Attributes["description"] != "Le bucket modifié" {
			return fmt.Errorf("the bucket description is not correct")
		}
		if rs.Primary.Attributes["rp"] != "" {
			return fmt.Errorf("the retention policies ar not correct")
		}

		return nil
	}
}

func testAccCreateBucket(token string, orgId string) string {
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
`
	return testAccConfigBucket
}

func testAccUpdateBucket(token string, orgId string) string {
	var testAccUpdateBucket = `
provider "influxdbv2" {
	url = "http://localhost:9999/"
    token = "` + token + `"
}
resource "influxdbv2_bucket" "initial" {
  description = "Le bucket modifié"
  name = "le bucket de deploiement modifié"
  org_id = "` + orgId + `"
  retention_rules {
    every_seconds = 30
  }
  rp = ""
}
`
	return testAccUpdateBucket
}
