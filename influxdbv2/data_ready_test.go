package influxdbv2

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"testing"
)

func TestAccReady(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccReady(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckReady("data.influxdbv2_ready.test"),
				),
			},
			{
				Config: testAccReady(),
				Check: resource.ComposeTestCheckFunc(
					testAccReadyOutput("data.influxdbv2_ready.test"),
				),
			},
		},
	})
}

func testAccCheckReady(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found %s", n)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("instance is not up")
		}
		return nil
	}
}

func testAccReadyOutput(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found %s", n)
		}
		if rs.Primary.Attributes["output.status"] == "" {
			return fmt.Errorf("instance is not ready")
		}
		if rs.Primary.Attributes["output.started"] == "" {
			return fmt.Errorf("instance is not started")
		}
		if rs.Primary.Attributes["output.up"] == "" {
			return fmt.Errorf("instance is not up")
		}
		return nil
	}
}

func testAccReady() string {
	var ready = `
data "influxdbv2_ready" "test" {}`

	return ready
}
