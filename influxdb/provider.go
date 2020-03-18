package influxdb

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"net/url"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{},

		DataSourcesMap: map[string]*schema.Resource{
			"influxdb": dataSourceInfluxdb(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	url, err := url.Parse(d.Get("url").(string))

	if err != nil {
		return nil, fmt.Errorf("invalid Influxdb URL: %s, %s", err, url)
	}
	return nil, nil
}
