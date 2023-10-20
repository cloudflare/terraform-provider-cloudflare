package sdkv2provider

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/MakeNowJust/heredoc/v2"
	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const CONCEALED_STRING = "**********************************"

func resourceCloudflareAccessIdentityProvider() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareAccessIdentityProviderSchema(),
		CreateContext: resourceCloudflareAccessIdentityProviderCreate,
		ReadContext:   resourceCloudflareAccessIdentityProviderRead,
		UpdateContext: resourceCloudflareAccessIdentityProviderUpdate,
		DeleteContext: resourceCloudflareAccessIdentityProviderDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceCloudflareAccessIdentityProviderImport,
		},
		Description: heredoc.Doc(`
			Provides a Cloudflare Access Identity Provider resource. Identity
			Providers are used as an authentication or authorisation source
			within Access.
		`),
	}
}

func resourceCloudflareAccessIdentityProviderRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	identifier, err := initIdentifier(d)
	if err != nil {
		return diag.FromErr(err)
	}

	accessIdentityProvider, err := client.GetAccessIdentityProvider(ctx, identifier, d.Id())
	if err != nil {
		var notFoundError *cloudflare.NotFoundError
		if errors.As(err, &notFoundError) {
			tflog.Info(ctx, fmt.Sprintf("Access Identity Provider %s no longer exists", d.Id()))
			d.SetId("")
			return nil
		}
		return diag.FromErr(fmt.Errorf("unable to find Access Identity Provider %q: %w", d.Id(), err))
	}

	d.SetId(accessIdentityProvider.ID)
	d.Set("name", accessIdentityProvider.Name)
	d.Set("type", accessIdentityProvider.Type)

	config := convertAccessIDPConfigStructToSchema(accessIdentityProvider.Config)
	if configErr := d.Set("config", config); configErr != nil {
		return diag.FromErr(fmt.Errorf("error setting Access Identity Provider configuration: %w", configErr))
	}

	// We need to get the current secret and set that in the state as it is only
	// returned initially when the resource was created. The value read from the
	// server will always be redacted.
	scimConfig := convertAccessIDPScimConfigStructToSchema(d.Get("scim_config.0.secret").(string), accessIdentityProvider.ScimConfig)
	if scimConfigErr := d.Set("scim_config", scimConfig); scimConfigErr != nil {
		return diag.FromErr(fmt.Errorf("error setting Access Identity Provider scim configuration: %w", scimConfigErr))
	}

	return nil
}

func resourceCloudflareAccessIdentityProviderCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	IDPConfig, _ := convertSchemaToStruct(d)
	ScimConfig := convertScimConfigSchemaToStruct(d)

	identityProvider := cloudflare.CreateAccessIdentityProviderParams{
		Name:       d.Get("name").(string),
		Type:       d.Get("type").(string),
		Config:     IDPConfig,
		ScimConfig: ScimConfig,
	}

	tflog.Debug(ctx, fmt.Sprintf("Creating Cloudflare Access Identity Provider from struct: %+v", identityProvider))

	identifier, err := initIdentifier(d)
	if err != nil {
		return diag.FromErr(err)
	}

	accessIdentityProvider, err := client.CreateAccessIdentityProvider(ctx, identifier, identityProvider)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to create Access Identity Provider: %w", err))
	}

	d.SetId(accessIdentityProvider.ID)

	scimConfig := convertAccessIDPScimConfigStructToSchema(accessIdentityProvider.ScimConfig.Secret, accessIdentityProvider.ScimConfig)
	if scimConfigErr := d.Set("scim_config", scimConfig); scimConfigErr != nil {
		return diag.FromErr(fmt.Errorf("error setting Access Identity Provider scim configuration: %w", scimConfigErr))
	}

	return resourceCloudflareAccessIdentityProviderRead(ctx, d, meta)
}

func resourceCloudflareAccessIdentityProviderUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	IDPConfig, conversionErr := convertSchemaToStruct(d)
	if conversionErr != nil {
		return diag.FromErr(fmt.Errorf("failed to convert schema into struct: %w", conversionErr))
	}

	ScimConfig := convertScimConfigSchemaToStruct(d)

	tflog.Debug(ctx, fmt.Sprintf("updatedConfig: %+v", IDPConfig))
	updatedAccessIdentityProvider := cloudflare.UpdateAccessIdentityProviderParams{
		ID:         d.Id(),
		Name:       d.Get("name").(string),
		Type:       d.Get("type").(string),
		Config:     IDPConfig,
		ScimConfig: ScimConfig,
	}

	tflog.Debug(ctx, fmt.Sprintf("Updating Cloudflare Access Identity Provider from struct: %+v", updatedAccessIdentityProvider))

	identifier, err := initIdentifier(d)
	if err != nil {
		return diag.FromErr(err)
	}

	accessIdentityProvider, err := client.UpdateAccessIdentityProvider(ctx, identifier, updatedAccessIdentityProvider)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error updating Access Identity Provider for ID %q: %w", d.Id(), err))
	}

	if accessIdentityProvider.ID == "" {
		return diag.FromErr(fmt.Errorf("failed to find Access Identity Provider ID in update response; resource was empty"))
	}

	scimSecret := d.Get("scim_config.0.secret").(string)
	if accessIdentityProvider.ScimConfig.Secret != "" && !strings.Contains(accessIdentityProvider.ScimConfig.Secret, "*") {
		scimSecret = accessIdentityProvider.ScimConfig.Secret
	}

	scimConfig := convertAccessIDPScimConfigStructToSchema(scimSecret, accessIdentityProvider.ScimConfig)
	if scimConfigErr := d.Set("scim_config", scimConfig); scimConfigErr != nil {
		return diag.FromErr(fmt.Errorf("error setting Access Identity Provider scim configuration: %w", scimConfigErr))
	}

	return resourceCloudflareAccessIdentityProviderRead(ctx, d, meta)
}

func resourceCloudflareAccessIdentityProviderDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	tflog.Debug(ctx, fmt.Sprintf("Deleting Cloudflare Access Identity Provider using ID: %s", d.Id()))

	identifier, err := initIdentifier(d)
	if err != nil {
		return diag.FromErr(err)
	}

	_, err = client.DeleteAccessIdentityProvider(ctx, identifier, d.Id())
	if err != nil {
		return diag.FromErr(fmt.Errorf("error deleting Access Identity Provider for ID %q: %w", d.Id(), err))
	}

	d.SetId("")

	return nil
}

func resourceCloudflareAccessIdentityProviderImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	attributes := strings.SplitN(d.Id(), "/", 2)

	if len(attributes) != 2 {
		return nil, fmt.Errorf("invalid id (\"%s\") specified, should be in format \"accountID/accessIdentityProviderID\"", d.Id())
	}

	accountID, accessIdentityProviderID := attributes[0], attributes[1]

	tflog.Debug(ctx, fmt.Sprintf("Importing Cloudflare Access Identity Provider: accountID=%s accessIdentityProviderID=%s", accountID, accessIdentityProviderID))

	d.Set(consts.AccountIDSchemaKey, accountID)
	d.SetId(accessIdentityProviderID)

	resourceCloudflareAccessIdentityProviderRead(ctx, d, meta)

	return []*schema.ResourceData{d}, nil
}

func convertSchemaToStruct(d *schema.ResourceData) (cloudflare.AccessIdentityProviderConfiguration, error) {
	IDPConfig := cloudflare.AccessIdentityProviderConfiguration{}

	if _, ok := d.GetOk("config"); ok {
		if _, ok := d.GetOk("config.0.attributes"); ok {
			attrData := make([]string, d.Get("config.0.attributes.#").(int))
			for id := range attrData {
				attrData[id] = d.Get(fmt.Sprintf("config.0.attributes.%d", id)).(string)
			}
			IDPConfig.Attributes = attrData
		}

		if _, ok := d.GetOk("config.0.claims"); ok {
			claimData := make([]string, d.Get("config.0.claims.#").(int))
			for id := range claimData {
				claimData[id] = d.Get(fmt.Sprintf("config.0.claims.%d", id)).(string)
			}
			IDPConfig.Claims = claimData
		}

		if _, ok := d.GetOk("config.0.scopes"); ok {
			scopeData := make([]string, d.Get("config.0.scopes.#").(int))
			for id := range scopeData {
				scopeData[id] = d.Get(fmt.Sprintf("config.0.scopes.%d", id)).(string)
			}
			IDPConfig.Scopes = scopeData
		}

		IDPConfig.APIToken = d.Get("config.0.api_token").(string)
		IDPConfig.AppsDomain = d.Get("config.0.apps_domain").(string)
		IDPConfig.AuthURL = d.Get("config.0.auth_url").(string)
		IDPConfig.CentrifyAccount = d.Get("config.0.centrify_account").(string)
		IDPConfig.CentrifyAppID = d.Get("config.0.centrify_app_id").(string)
		IDPConfig.CertsURL = d.Get("config.0.certs_url").(string)
		IDPConfig.ClientID = d.Get("config.0.client_id").(string)
		IDPConfig.ClientSecret = d.Get("config.0.client_secret").(string)
		IDPConfig.DirectoryID = d.Get("config.0.directory_id").(string)
		IDPConfig.EmailAttributeName = d.Get("config.0.email_attribute_name").(string)
		IDPConfig.EmailClaimName = d.Get("config.0.email_claim_name").(string)
		IDPConfig.IdpPublicCert = d.Get("config.0.idp_public_cert").(string)
		IDPConfig.IssuerURL = d.Get("config.0.issuer_url").(string)
		IDPConfig.OktaAccount = d.Get("config.0.okta_account").(string)
		IDPConfig.OktaAuthorizationServerID = d.Get("config.0.authorization_server_id").(string)
		IDPConfig.OneloginAccount = d.Get("config.0.onelogin_account").(string)
		IDPConfig.PingEnvID = d.Get("config.0.ping_env_id").(string)
		IDPConfig.RedirectURL = d.Get("config.0.redirect_url").(string)
		IDPConfig.SignRequest = d.Get("config.0.sign_request").(bool)
		IDPConfig.SsoTargetURL = d.Get("config.0.sso_target_url").(string)
		IDPConfig.SupportGroups = d.Get("config.0.support_groups").(bool)
		IDPConfig.TokenURL = d.Get("config.0.token_url").(string)
		IDPConfig.PKCEEnabled = cloudflare.BoolPtr(d.Get("config.0.pkce_enabled").(bool))
		IDPConfig.ConditionalAccessEnabled = d.Get("config.0.conditional_access_enabled").(bool)
	}

	return IDPConfig, nil
}

func convertScimConfigSchemaToStruct(d *schema.ResourceData) cloudflare.AccessIdentityProviderScimConfiguration {
	ScimConfig := cloudflare.AccessIdentityProviderScimConfiguration{}

	if _, ok := d.GetOk("scim_config"); ok {
		ScimConfig.Enabled = d.Get("scim_config.0.enabled").(bool)
		ScimConfig.Secret = d.Get("scim_config.0.secret").(string)
		ScimConfig.GroupMemberDeprovision = d.Get("scim_config.0.group_member_deprovision").(bool)
		ScimConfig.UserDeprovision = d.Get("scim_config.0.user_deprovision").(bool)
		ScimConfig.SeatDeprovision = d.Get("scim_config.0.seat_deprovision").(bool)
	}

	return ScimConfig
}

func convertAccessIDPConfigStructToSchema(options cloudflare.AccessIdentityProviderConfiguration) []interface{} {
	attributes := make([]string, 0)
	for _, value := range options.Attributes {
		attributes = append(attributes, value)
	}

	m := map[string]interface{}{
		"api_token":                  options.APIToken,
		"apps_domain":                options.AppsDomain,
		"attributes":                 attributes,
		"auth_url":                   options.AuthURL,
		"authorization_server_id":    options.OktaAuthorizationServerID,
		"centrify_account":           options.CentrifyAccount,
		"centrify_app_id":            options.CentrifyAppID,
		"certs_url":                  options.CertsURL,
		"client_id":                  options.ClientID,
		"client_secret":              options.ClientSecret,
		"claims":                     options.Claims,
		"scopes":                     options.Scopes,
		"directory_id":               options.DirectoryID,
		"email_attribute_name":       options.EmailAttributeName,
		"email_claim_name":           options.EmailClaimName,
		"idp_public_cert":            options.IdpPublicCert,
		"issuer_url":                 options.IssuerURL,
		"okta_account":               options.OktaAccount,
		"onelogin_account":           options.OneloginAccount,
		"ping_env_id":                options.PingEnvID,
		"redirect_url":               options.RedirectURL,
		"sign_request":               options.SignRequest,
		"sso_target_url":             options.SsoTargetURL,
		"support_groups":             options.SupportGroups,
		"token_url":                  options.TokenURL,
		"pkce_enabled":               options.PKCEEnabled,
		"conditional_access_enabled": options.ConditionalAccessEnabled,
	}

	return []interface{}{m}
}

func convertAccessIDPScimConfigStructToSchema(secret string, options cloudflare.AccessIdentityProviderScimConfiguration) []interface{} {
	m := map[string]interface{}{
		"enabled":                  options.Enabled,
		"secret":                   secret,
		"user_deprovision":         options.UserDeprovision,
		"seat_deprovision":         options.SeatDeprovision,
		"group_member_deprovision": options.GroupMemberDeprovision,
	}

	return []interface{}{m}
}
