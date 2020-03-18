package influxdb

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceInfluxdb() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ready": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: false,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"url": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
		},
	}
}
