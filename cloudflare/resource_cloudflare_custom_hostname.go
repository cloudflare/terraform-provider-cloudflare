package cloudflare

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceCloudflareCustomHostname() *schema.Resource {
	return &schema.Resource{
		Create: resourceCloudflareCustomHostnameCreate,
		Read:   resourceCloudflareCustomHostnameRead,
		Update: resourceCloudflareCustomHostnameUpdate,
		Delete: resourceCloudflareCustomHostnameDelete,
		Importer: &schema.ResourceImporter{
			State: resourceCloudflareCustomHostnameImport,
		},

		SchemaVersion: 0,
		Schema: map[string]*schema.Schema{
			"hostname": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringLenBetween(0, 255),
			},
			"custom_origin_servver": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssl": {
				Type:     schema.TypeList,
				Required: true,
				MinItems: 1,
				MaxItems: 1,
				Elem: &schema.Resource{
					SchemaVersion: 1,
					Schema: map[string]*schema.Schema{
						"status": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"method": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.StringInSlice([]string{"http", "txt", "email"}, false),
						},
						"type": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.StringInSlice([]string{"dv"}, false),
						},
						"cname_target": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"cname_name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"wildcard": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"custom_certificate": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"custom_key": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"settings": {
							Type:     schema.TypeList,
							Required: true,
							MinItems: 1,
							MaxItems: 1,
							Elem: &schema.Resource{
								SchemaVersion: 1,
								Schema: map[string]*schema.Schema{
									"http2": {
										Type:         schema.TypeString,
										Optional:     true,
										ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
									},
									"tls13": {
										Type:         schema.TypeString,
										Optional:     true,
										ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
									},
									"min_tls_version": {
										Type:         schema.TypeString,
										Optional:     true,
										ValidateFunc: validation.StringInSlice([]string{"1.0", "1.1", "1.2", "1.3"}, false),
									},
									"ciphers": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
					},
				},
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ownership_verification": {
				Type:     schema.TypeList,
				Computed: true,
				MinItems: 1,
				MaxItems: 1,
				Elem: &schema.Resource{
					SchemaVersion: 1,
					Schema: map[string]*schema.Schema{
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"value": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"ownership_verification_http": {
				Type:     schema.TypeList,
				Computed: true,
				MinItems: 1,
				MaxItems: 1,
				Elem: &schema.Resource{
					SchemaVersion: 1,
					Schema: map[string]*schema.Schema{
						"http_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"http_body": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceCloudflareCustomHostnameRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceCloudflareCustomHostnameDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceCloudflareCustomHostnameCreate(d *schema.ResourceData, meta interface{}) error {
	return resourceCloudflareCustomHostnameRead(d, meta)
}

func resourceCloudflareCustomHostnameUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceCloudflareCustomHostnameRead(d, meta)
}

func resourceCloudflareCustomHostnameImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	return []*schema.ResourceData{d}, nil
}
