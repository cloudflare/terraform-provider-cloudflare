// GET accounts/:identifier/access/organizations

package cloudflare

import (
	"context"
	"fmt"
	"log"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCloudflareAccessOrganization() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceCloudflareAccessOrganizationRead,

		Schema: map[string]*schema.Schema{
			"account_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"auth_domain": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"login_design": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceCloudflareAccessOrganizationRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	accountID := d.Get("account_id").(string)

	log.Printf("[DEBUG] Reading Access Organization")
	accessOrganization, _, err := client.AccessOrganization(context.Background(), accountID)
	if err != nil {
		return fmt.Errorf("error finding Access Organization: %s", err)
	}

	d.Set("account_id", accountID)
	d.Set("name", accessOrganization.Name)
	d.Set("auth_domain", accessOrganization.AuthDomain)
	d.Set("login_design", map[string]interface{}{
		"background_color": accessOrganization.LoginDesign.BackgroundColor,
		"logo_path":        accessOrganization.LoginDesign.LogoPath,
		"header_text":      accessOrganization.LoginDesign.HeaderText,
		"footer_text":      accessOrganization.LoginDesign.FooterText,
	})

	d.SetId(stringChecksum(accessOrganization.UpdatedAt.String()))

	return nil
}
