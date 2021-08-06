package influxdbv2

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"os"
	"testing"
)

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

var testAccProviders = map[string]*schema.Provider{
	"influxdb-v2": Provider(),
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("INFLUXDB_V2_TOKEN"); v == "" {
		t.Fatal("INFLUXDB_V2_TOKEN must be set for acceptance tests")
	}
	if v := os.Getenv("INFLUXDB_V2_BUCKET_ID"); v == "" {
		t.Fatal("INFLUXDB_V2_BUCKET_ID must be set for acceptance tests")
	}
	if v := os.Getenv("INFLUXDB_V2_ORG_ID"); v == "" {
		t.Fatal("INFLUXDB_V2_ORG_ID must be set for acceptance tests")
	}
}
