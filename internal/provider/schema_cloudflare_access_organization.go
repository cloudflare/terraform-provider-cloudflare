package provider

import (
	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareAccessOrganizationSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"account_id": {
			Description:   "The account identifier to target for the resource.",
			Type:          schema.TypeString,
			Optional:      true,
			Computed:      true,
			ConflictsWith: []string{"zone_id"},
		},
		"zone_id": {
			Description:   "The zone identifier to target for the resource.",
			Type:          schema.TypeString,
			Optional:      true,
			Computed:      true,
			ConflictsWith: []string{"account_id"},
		},
		"auth_domain": {
			Description: "The unique subdomain assigned to your Zero Trust organization.",
			Type:        schema.TypeString,
			Required:    true,
		},
		"name": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The name of your Zero Trust organization.",
		},
		"is_ui_read_only": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "When set to true, this will disable all editing of Access resources via the Zero Trust Dashboard",
		},
		"login_design": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"background_color": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "The background color on the login page",
					},
					"text_color": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "The text color on the login page",
					},
					"logo_path": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "The URL of the logo on the login page",
					},
					"header_text": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "The text at the top of the login page",
					},
					"footer_text": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "The text at the bottom of the login page",
					},
				},
			},
		},
	}
}

func convertLoginDesignSchemaToStruct(d *schema.ResourceData) *cloudflare.AccessOrganizationLoginDesign {
	LoginDesign := cloudflare.AccessOrganizationLoginDesign{}

	if _, ok := d.GetOk("login_design"); ok {
		LoginDesign.BackgroundColor = d.Get("login_design.0.background_color").(string)
		LoginDesign.TextColor = d.Get("login_design.0.text_color").(string)
		LoginDesign.LogoPath = d.Get("login_design.0.logo_path").(string)
		LoginDesign.HeaderText = d.Get("login_design.0.header_text").(string)
		LoginDesign.FooterText = d.Get("login_design.0.footer_text").(string)
	}

	return &LoginDesign
}

func convertLoginDesignStructToSchema(d *schema.ResourceData, loginDesign *cloudflare.AccessOrganizationLoginDesign) []interface{} {
	if _, ok := d.GetOk("login_design"); !ok {
		return []interface{}{}
	}

	m := map[string]interface{}{
		"background_color": loginDesign.BackgroundColor,
		"text_color":       loginDesign.TextColor,
		"logo_path":        loginDesign.LogoPath,
		"header_text":      loginDesign.HeaderText,
		"footer_text":      loginDesign.FooterText,
	}

	return []interface{}{m}
}
