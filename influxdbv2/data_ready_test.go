package influxdb-v2

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccReady(t *testing.T) {
	currentUrl, exists := os.LookupEnv("INFLUXDB_V2_URL")
	if !exists {
		currentUrl = "http://localhost:9999"
	}

	resource.Test(t, resource.TestCase{
		// no need to precheck the token env var, we don't need it
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccReady(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.influxdb-v2_ready.acctest", "output.url", currentUrl),
				),
			},
		},
	})
}

func testAccReady() string {
	return `
data "influxdb-v2_ready" "acctest" {}`
}
