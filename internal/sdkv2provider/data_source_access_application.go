package sdkv2provider

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCloudflareAccessApplication() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			consts.AccountIDSchemaKey: {
				Description:  consts.AccountIDSchemaDescription,
				Type:         schema.TypeString,
				Optional:     true,
				ExactlyOneOf: []string{consts.ZoneIDSchemaKey, consts.AccountIDSchemaKey},
			},
			consts.ZoneIDSchemaKey: {
				Description:  consts.ZoneIDSchemaDescription,
				Type:         schema.TypeString,
				Optional:     true,
				ExactlyOneOf: []string{consts.ZoneIDSchemaKey, consts.AccountIDSchemaKey},
			},
			"name": {
				Description:  "Friendly name of the Access Application.",
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ExactlyOneOf: []string{"name", "domain"},
			},
			"domain": {
				Description:  "The primary hostname and path that Access will secure.",
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ExactlyOneOf: []string{"name", "domain"},
			},
			"aud": {
				Description: "Application Audience (AUD) Tag of the application.",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
		Description: "Use this data source to lookup a single [Access Application](https://developers.cloudflare.com/cloudflare-one/applications/)",
		ReadContext: dataSourceCloudflareAccessApplicationRead,
	}
}

func dataSourceCloudflareAccessApplicationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	identifier, err := initIdentifier(d)
	if err != nil {
		return diag.FromErr(err)
	}
	name := d.Get("name").(string)
	domain := d.Get("domain").(string)

	applications, _, err := client.ListAccessApplications(ctx, identifier, cloudflare.ListAccessApplicationsParams{})
	if err != nil {
		return diag.FromErr(fmt.Errorf("error listing Access Applications: %w", err))
	}
	if len(applications) == 0 {
		return diag.Errorf("no Access Applications found")
	}
	var accessApplication cloudflare.AccessApplication
	for _, application := range applications {
		if application.Name == name || application.Domain == domain {
			accessApplication = application
			break
		}
	}
	if accessApplication.ID == "" {
		return diag.Errorf("no Access Application matching name %q domain %q", name, domain)
	}
	d.SetId(accessApplication.ID)
	d.Set("name", accessApplication.Name)
	d.Set("domain", accessApplication.Domain)
	d.Set("aud", accessApplication.AUD)
	return nil
}
