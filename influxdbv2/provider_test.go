package influxdbv2

import (
	"github.com/hashicorp/terraform-plugin-sdk/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"testing"
)

var testProviders = map[string]terraform.ResourceProvider{
	"influxdbv2": Provider(),
}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestMain(m *testing.M) {
	acctest.UseBinaryDriver("influxdbv2", Provider)
	resource.TestMain(m)
}
