package sdkv2provider

import (
	"fmt"
	"slices"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

const (
	IdentityUpdateBehaviorNoAction  = "no_action"
	IdentityUpdateBehaviorReauth    = "reauth"
	IdentityUpdateBehaviorAutomatic = "automatic"
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
						Type:        schema.TypeBool,
						Optional:    true,
						Description: "A flag to enable or disable SCIM for the identity provider.",
					},
					"secret": {
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
						Sensitive:   true,
						Description: "A read-only token generated when the SCIM integration is enabled for the first time.  It is redacted on subsequent requests.  If you lose this you will need to refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.",
					},
					"user_deprovision": {
						Type:        schema.TypeBool,
						Optional:    true,
						Description: "A flag to enable revoking a user's session in Access and Gateway when they have been deprovisioned in the Identity Provider.",
					},
					"seat_deprovision": {
						Type:        schema.TypeBool,
						Optional:    true,
						Description: "A flag to remove a user's seat in Zero Trust when they have been deprovisioned in the Identity Provider.  This cannot be enabled unless user_deprovision is also enabled.",
					},
					"group_member_deprovision": {
						Type:        schema.TypeBool,
						Optional:    true,
						Description: "Deprecated. Use `identity_update_behavior`.",
					},
					"identity_update_behavior": {
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
						Description: "Indicates how a SCIM event updates a user identity used for policy evaluation. Use \"automatic\" to automatically update a user's identity and augment it with fields from the SCIM user resource. Use \"reauth\" to force re-authentication on group membership updates, user identity update will only occur after successful re-authentication. With \"reauth\" identities will not contain fields from the SCIM user resource. With \"no_action\" identities will not be changed by SCIM updates in any way and users will not be prompted to reauthenticate.",
						ValidateDiagFunc: func(val interface{}, path cty.Path) diag.Diagnostics {
							s, ok := val.(string)

							if !ok {
								return diag.Errorf("value %s was not a string", val)
							}

							allowedValues := []string{IdentityUpdateBehaviorNoAction, IdentityUpdateBehaviorReauth, IdentityUpdateBehaviorAutomatic}

							isValid := slices.Contains(allowedValues, s)

							if !isValid {
								return diag.Errorf("value %s was not one of %s", val, allowedValues)
							}

							return nil
						},
					},
				},
			},
		},
	}
}
