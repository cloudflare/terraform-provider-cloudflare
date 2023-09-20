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
			Description:   consts.AccountIDSchemaDescription,
			Type:          schema.TypeString,
			Optional:      true,
			ForceNew:      true,
			ConflictsWith: []string{consts.ZoneIDSchemaKey},
		},
		consts.ZoneIDSchemaKey: {
			Description:   consts.ZoneIDSchemaDescription,
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
			ValidateFunc: validation.StringInSlice([]string{"azureAD", "centrify", "facebook", "github", "google", "google-apps", "linkedin", "oidc", "okta", "onelogin", "onetimepin", "pingone", "saml", "yandex"}, false),
			Description:  fmt.Sprintf("The provider type to use. %s", renderAvailableDocumentationValuesStringSlice([]string{"azureAD", "centrify", "facebook", "github", "google", "google-apps", "linkedin", "oidc", "okta", "onelogin", "onetimepin", "pingone", "saml", "yandex"})),
		},
		"config": {
			Type:        schema.TypeList,
			Optional:    true,
			Computed:    true,
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
						Computed: true,
					},
					"auth_url": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"authorization_server_id": {
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
					"claims": {
						Type:     schema.TypeList,
						Optional: true,
						Elem: &schema.Schema{
							Type: schema.TypeString,
						},
						Computed: true,
					},
					"email_claim_name": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"scopes": {
						Type:     schema.TypeList,
						Optional: true,
						Elem: &schema.Schema{
							Type: schema.TypeString,
						},
						Computed: true,
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
					"ping_env_id": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"redirect_url": {
						Type:     schema.TypeString,
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
					"conditional_access_enabled": {
						Type:     schema.TypeBool,
						Optional: true,
					},
				},
			},
		},
		"scim_config": {
			Type:        schema.TypeList,
			Optional:    true,
			Computed:    true,
			Description: "Configuration for SCIM settings for a given IDP",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"enabled": {
						Type:     schema.TypeBool,
						Optional: true,
					},
					"secret": {
						Type:      schema.TypeString,
						Optional:  true,
						Computed:  true,
						Sensitive: true,
					},
					"user_deprovision": {
						Type:     schema.TypeBool,
						Optional: true,
					},
					"seat_deprovision": {
						Type:     schema.TypeBool,
						Optional: true,
					},
					"group_member_deprovision": {
						Type:     schema.TypeBool,
						Optional: true,
					},
				},
			},
		},
	}
}
