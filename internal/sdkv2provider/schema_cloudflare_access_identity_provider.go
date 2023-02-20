package sdkv2provider

import (
	"fmt"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareAccessIdentityProviderSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.AccountIDSchemaKey: {
			Description:   "The account identifier to target for the resource.",
			Type:          schema.TypeString,
			Optional:      true,
			ForceNew:      true,
			ConflictsWith: []string{consts.ZoneIDSchemaKey},
		},
		consts.ZoneIDSchemaKey: {
			Description:   "The zone identifier to target for the resource.",
			Type:          schema.TypeString,
			Optional:      true,
			ForceNew:      true,
			ConflictsWith: []string{consts.AccountIDSchemaKey},
		},
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Friendly name of the Access Identity Provider configuration.",
		},
		"type": {
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validation.StringInSlice([]string{"centrify", "facebook", "google-apps", "oidc", "github", "google", "saml", "linkedin", "azureAD", "okta", "onetimepin", "onelogin", "yandex"}, false),
			Description:  fmt.Sprintf("The provider type to use. %s", renderAvailableDocumentationValuesStringSlice([]string{"centrify", "facebook", "google-apps", "oidc", "github", "google", "saml", "linkedin", "azureAD", "okta", "onetimepin", "onelogin", "yandex"})),
		},
		"config": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "Provider configuration from the [developer documentation](https://developers.cloudflare.com/access/configuring-identity-providers/).",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"api_token": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"apps_domain": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"attributes": {
						Type:     schema.TypeList,
						Optional: true,
						Elem: &schema.Schema{
							Type: schema.TypeString,
						},
					},
					"auth_url": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"centrify_account": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"centrify_app_id": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"certs_url": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"client_id": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"client_secret": {
						Type:     schema.TypeString,
						Optional: true,
						// client_secret is a write only operation from the Cloudflare API
						// and once it's set, it is no longer accessible. To avoid storing
						// it and messing up the state, hardcode in the concealed version.
						StateFunc: func(val interface{}) string {
							return CONCEALED_STRING
						},
					},
					"directory_id": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"email_attribute_name": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"idp_public_cert": {
						Type:     schema.TypeString,
						Optional: true,
						// idp_public_cert is a write only operation from the Cloudflare
						// API and once it's set, it is no longer accessible. To avoid
						// storing it and messing up the state, hardcode in the concealed
						// version.
						StateFunc: func(val interface{}) string {
							return CONCEALED_STRING
						},
					},
					"issuer_url": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"okta_account": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"onelogin_account": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"redirect_url": {
						Type:     schema.TypeString,
						Optional: true,
						Computed: true,
					},
					"sign_request": {
						Type:     schema.TypeBool,
						Optional: true,
					},
					"sso_target_url": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"support_groups": {
						Type:     schema.TypeBool,
						Optional: true,
					},
					"token_url": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"pkce_enabled": {
						Type:     schema.TypeBool,
						Optional: true,
					},
				},
			},
		},
	}
}
