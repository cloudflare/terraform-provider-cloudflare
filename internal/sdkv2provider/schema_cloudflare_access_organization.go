package sdkv2provider

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareAccessOrganizationSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.AccountIDSchemaKey: {
			Description:   consts.AccountIDSchemaDescription,
			Type:          schema.TypeString,
			Optional:      true,
			Computed:      true,
			ConflictsWith: []string{consts.ZoneIDSchemaKey},
		},
		consts.ZoneIDSchemaKey: {
			Description:   consts.ZoneIDSchemaDescription,
			Type:          schema.TypeString,
			Optional:      true,
			Computed:      true,
			ConflictsWith: []string{consts.AccountIDSchemaKey},
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
		"ui_read_only_toggle_reason": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "A description of the reason why the UI read only field is being toggled.",
		},
		"user_seat_expiration_inactive_time": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The amount of time a user seat is inactive before it expires. When the user seat exceeds the set time of inactivity, the user is removed as an active seat and no longer counts against your Teams seat count. Must be in the format `300ms` or `2h45m`.",
		},
		"auto_redirect_to_identity": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "When set to true, users skip the identity provider selection step during login",
		},
		"session_duration": {
			Type:     schema.TypeString,
			Optional: true,
			ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
				v := val.(string)
				_, err := time.ParseDuration(v)
				if err != nil {
					errs = append(errs, fmt.Errorf(`%q only supports "ns", "us" (or "µs"), "ms", "s", "m", or "h" as valid units`, key))
				}
				return
			},
			Description: "How often a user will be forced to re-authorise. Must be in the format `48h` or `2h45m`",
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
		"custom_pages": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "Custom pages for your Zero Trust organization.",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"identity_denied": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "The id of the identity denied page",
					},
					"forbidden": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "The id of the forbidden page",
					},
				},
			},
		},
	}
}

func convertLoginDesignSchemaToStruct(d *schema.ResourceData) *cloudflare.AccessOrganizationLoginDesign {
	LoginDesign := cloudflare.AccessOrganizationLoginDesign{}

	if _, ok := d.GetOk("login_design.0"); ok {
		LoginDesign.BackgroundColor = d.Get("login_design.0.background_color").(string)
		LoginDesign.TextColor = d.Get("login_design.0.text_color").(string)
		LoginDesign.LogoPath = d.Get("login_design.0.logo_path").(string)
		LoginDesign.HeaderText = d.Get("login_design.0.header_text").(string)
		LoginDesign.FooterText = d.Get("login_design.0.footer_text").(string)
	}

	return &LoginDesign
}

func convertCustomPageSchemaToStruct(d *schema.ResourceData) *cloudflare.AccessOrganizationCustomPages {
	CustomPages := cloudflare.AccessOrganizationCustomPages{}

	if _, ok := d.GetOk("custom_pages.0"); ok {
		identityDenied := d.Get("custom_pages.0.identity_denied").(string)
		if identityDenied != "" {
			CustomPages.IdentityDenied = cloudflare.AccessCustomPageType(identityDenied)
		}
		forbidden := d.Get("custom_pages.0.forbidden").(string)
		if forbidden != "" {
			CustomPages.Forbidden = cloudflare.AccessCustomPageType(forbidden)
		}
	}

	return &CustomPages
}

func convertCustomPageStructToSchema(ctx context.Context, d *schema.ResourceData, customPages *cloudflare.AccessOrganizationCustomPages) []interface{} {
	var onImport bool
	var ok bool
	if onImport, ok = ctx.Value(orgAccessImportCtxKey).(bool); !ok {
		onImport = false
	}

	if _, ok := d.GetOk("custom_pages"); !ok && !onImport {
		return []interface{}{}
	}

	m := map[string]interface{}{
		"identity_denied": string(customPages.IdentityDenied),
		"forbidden":       string(customPages.Forbidden),
	}

	return []interface{}{m}
}

func convertLoginDesignStructToSchema(ctx context.Context, d *schema.ResourceData, loginDesign *cloudflare.AccessOrganizationLoginDesign) []interface{} {
	var onImport bool
	var ok bool
	if onImport, ok = ctx.Value(orgAccessImportCtxKey).(bool); !ok {
		onImport = false
	}

	if _, ok := d.GetOk("login_design"); !ok && !onImport {
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
