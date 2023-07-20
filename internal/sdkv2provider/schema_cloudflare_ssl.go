package sdkv2provider

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func sslValidationErrorsSchema() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Schema: map[string]*schema.Schema{
			"message": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}
func sslValidationRecordsSchema() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Schema: map[string]*schema.Schema{
			"cname_target": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"cname_name": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"txt_name": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"txt_value": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"http_url": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"http_body": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"emails": {
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}
