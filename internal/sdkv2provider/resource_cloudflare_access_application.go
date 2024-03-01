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

func resourceCloudflareAccessApplication() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareAccessApplicationSchema(),
		CreateContext: resourceCloudflareAccessApplicationCreate,
		ReadContext:   resourceCloudflareAccessApplicationRead,
		UpdateContext: resourceCloudflareAccessApplicationUpdate,
		DeleteContext: resourceCloudflareAccessApplicationDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceCloudflareAccessApplicationImport,
		},
		Description: heredoc.Doc(`
			Provides a Cloudflare Access Application resource. Access
			Applications are used to restrict access to a whole application using an
			authorisation gateway managed by Cloudflare.
		`),
	}
}

func resourceCloudflareAccessApplicationCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	appType := d.Get("type").(string)

	newAccessApplication := cloudflare.CreateAccessApplicationParams{
		Name:                     d.Get("name").(string),
		Domain:                   d.Get("domain").(string),
		Type:                     cloudflare.AccessApplicationType(appType),
		SessionDuration:          d.Get("session_duration").(string),
		AutoRedirectToIdentity:   cloudflare.BoolPtr(d.Get("auto_redirect_to_identity").(bool)),
		EnableBindingCookie:      cloudflare.BoolPtr(d.Get("enable_binding_cookie").(bool)),
		CustomDenyMessage:        d.Get("custom_deny_message").(string),
		CustomDenyURL:            d.Get("custom_deny_url").(string),
		CustomNonIdentityDenyURL: d.Get("custom_non_identity_deny_url").(string),
		HttpOnlyCookieAttribute:  cloudflare.BoolPtr(d.Get("http_only_cookie_attribute").(bool)),
		SameSiteCookieAttribute:  d.Get("same_site_cookie_attribute").(string),
		LogoURL:                  d.Get("logo_url").(string),
		SkipInterstitial:         cloudflare.BoolPtr(d.Get("skip_interstitial").(bool)),
		AppLauncherVisible:       cloudflare.BoolPtr(d.Get("app_launcher_visible").(bool)),
		ServiceAuth401Redirect:   cloudflare.BoolPtr(d.Get("service_auth_401_redirect").(bool)),
	}

	if _, ok := d.GetOk("allow_authenticate_via_warp"); ok {
		newAccessApplication.AllowAuthenticateViaWarp = cloudflare.BoolPtr(d.Get("allow_authenticate_via_warp").(bool))
	}

	if value, ok := d.GetOk("allowed_idps"); ok {
		newAccessApplication.AllowedIdps = expandInterfaceToStringList(value.(*schema.Set).List())
	}

	if value, ok := d.GetOk("custom_pages"); ok {
		newAccessApplication.CustomPages = expandInterfaceToStringList(value.(*schema.Set).List())
	}

	if value, ok := d.GetOk("self_hosted_domains"); ok {
		newAccessApplication.SelfHostedDomains = expandInterfaceToStringList(value.(*schema.Set).List())
	}

	if _, ok := d.GetOk("cors_headers"); ok {
		CORSConfig, err := convertCORSSchemaToStruct(d)
		if err != nil {
			return diag.FromErr(err)
		}
		newAccessApplication.CorsHeaders = CORSConfig
	}

	if _, ok := d.GetOk("saas_app"); ok {
		newAccessApplication.SaasApplication = convertSaasSchemaToStruct(d)
	}

	if value, ok := d.GetOk("tags"); ok {
		newAccessApplication.Tags = expandInterfaceToStringList(value.(*schema.Set).List())
	}
	if appType == "app_launcher" {
		newAccessApplication.AccessAppLauncherCustomization = cloudflare.AccessAppLauncherCustomization{
			LogoURL:               d.Get("app_launcher_logo_url").(string),
			BackgroundColor:       d.Get("bg_color").(string),
			HeaderBackgroundColor: d.Get("header_bg_color").(string),
		}

		if _, ok := d.GetOk("landing_page_design"); ok {
			landingPageDesign := convertLandingPageDesignSchemaToStruct(d)
			newAccessApplication.AccessAppLauncherCustomization.LandingPageDesign = *landingPageDesign
		}

		if _, ok := d.GetOk("footer_links"); ok {
			footerLinks := convertFooterLinksSchemaToStruct(d)
			newAccessApplication.AccessAppLauncherCustomization.FooterLinks = footerLinks
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("Creating Cloudflare Access Application from struct: %+v", newAccessApplication))

	identifier, err := initIdentifier(d)
	if err != nil {
		return diag.FromErr(err)
	}
	accessApplication, err := client.CreateAccessApplication(ctx, identifier, newAccessApplication)

	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating Access Application for %s %q: %w", identifier.Level, identifier.Identifier, err))
	}

	d.SetId(accessApplication.ID)

	readApplication := resourceCloudflareAccessApplicationRead(ctx, d, meta)

	// client secret is only returned from the create request and should be stored in state
	if accessApplication.SaasApplication != nil && accessApplication.SaasApplication.ClientSecret != "" {
		rawSaasApp, ok := d.GetOk("saas_app")
		if ok {
			saasApp, ok := rawSaasApp.([]interface{})
			if ok {
				saasAppMap, ok := saasApp[0].(map[string]interface{})
				if ok {
					saasAppMap["client_secret"] = accessApplication.SaasApplication.ClientSecret
					d.Set("saas_app", []interface{}{saasAppMap})
				}
			}
		}
	}

	return readApplication
}

func resourceCloudflareAccessApplicationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	identifier, err := initIdentifier(d)
	if err != nil {
		return diag.FromErr(err)
	}

	accessApplication, err := client.GetAccessApplication(ctx, identifier, d.Id())

	if err != nil {
		var notFoundError *cloudflare.NotFoundError
		if errors.As(err, &notFoundError) {
			tflog.Info(ctx, fmt.Sprintf("Access Application %s no longer exists", d.Id()))
			d.SetId("")
			return nil
		}
		return diag.FromErr(fmt.Errorf("error finding Access Application %q: %w", d.Id(), err))
	}

	d.Set("name", accessApplication.Name)
	d.Set("aud", accessApplication.AUD)
	d.Set("session_duration", accessApplication.SessionDuration)
	d.Set("domain", accessApplication.Domain)
	d.Set("type", accessApplication.Type)
	d.Set("auto_redirect_to_identity", accessApplication.AutoRedirectToIdentity)
	d.Set("enable_binding_cookie", accessApplication.EnableBindingCookie)
	d.Set("custom_deny_message", accessApplication.CustomDenyMessage)
	d.Set("custom_deny_url", accessApplication.CustomDenyURL)
	d.Set("custom_non_identity_deny_url", accessApplication.CustomNonIdentityDenyURL)
	d.Set("allowed_idps", accessApplication.AllowedIdps)
	d.Set("http_only_cookie_attribute", cloudflare.Bool(accessApplication.HttpOnlyCookieAttribute))
	d.Set("same_site_cookie_attribute", accessApplication.SameSiteCookieAttribute)
	d.Set("skip_interstitial", accessApplication.SkipInterstitial)
	d.Set("logo_url", accessApplication.LogoURL)
	d.Set("app_launcher_visible", accessApplication.AppLauncherVisible)
	d.Set("service_auth_401_redirect", accessApplication.ServiceAuth401Redirect)
	d.Set("custom_pages", accessApplication.CustomPages)
	d.Set("tags", accessApplication.Tags)
	d.Set("bg_color", accessApplication.AccessAppLauncherCustomization.BackgroundColor)
	d.Set("header_bg_color", accessApplication.AccessAppLauncherCustomization.HeaderBackgroundColor)
	d.Set("app_launcher_logo_url", accessApplication.AccessAppLauncherCustomization.LogoURL)
	d.Set("allow_authenticate_via_warp", accessApplication.AllowAuthenticateViaWarp)

	if _, ok := d.GetOk("footer_links"); ok {
		footerLinks := convertFooterLinksStructToSchema(d, accessApplication.AccessAppLauncherCustomization.FooterLinks)
		if footerLinksErr := d.Set("footer_links", footerLinks); footerLinksErr != nil {
			return diag.FromErr(fmt.Errorf("error setting Access Application footer links: %w", footerLinksErr))
		}
	}

	if _, ok := d.GetOk("landing_page_design"); ok {
		landingPageDesign := convertLandingPageDesignStructToSchema(d, &accessApplication.AccessAppLauncherCustomization.LandingPageDesign)
		if landingPageDesignErr := d.Set("landing_page_design", landingPageDesign); landingPageDesignErr != nil {
			return diag.FromErr(fmt.Errorf("error setting Access Application landing page design: %w", landingPageDesignErr))
		}
	}

	corsConfig := convertCORSStructToSchema(d, accessApplication.CorsHeaders)
	if corsConfigErr := d.Set("cors_headers", corsConfig); corsConfigErr != nil {
		return diag.FromErr(fmt.Errorf("error setting Access Application CORS header configuration: %w", corsConfigErr))
	}

	saasConfig := convertSaasStructToSchema(d, accessApplication.SaasApplication)
	if saasConfigErr := d.Set("saas_app", saasConfig); saasConfigErr != nil {
		return diag.FromErr(fmt.Errorf("error setting Access Application SaaS app configuration: %w", saasConfigErr))
	}

	if _, ok := d.GetOk("self_hosted_domains"); ok {
		d.Set("self_hosted_domains", accessApplication.SelfHostedDomains)
	}

	return nil
}

func resourceCloudflareAccessApplicationUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	appType := d.Get("type").(string)

	updatedAccessApplication := cloudflare.UpdateAccessApplicationParams{
		ID:                       d.Id(),
		Name:                     d.Get("name").(string),
		Domain:                   d.Get("domain").(string),
		Type:                     cloudflare.AccessApplicationType(appType),
		SessionDuration:          d.Get("session_duration").(string),
		AutoRedirectToIdentity:   cloudflare.BoolPtr(d.Get("auto_redirect_to_identity").(bool)),
		EnableBindingCookie:      cloudflare.BoolPtr(d.Get("enable_binding_cookie").(bool)),
		CustomDenyMessage:        d.Get("custom_deny_message").(string),
		CustomDenyURL:            d.Get("custom_deny_url").(string),
		CustomNonIdentityDenyURL: d.Get("custom_non_identity_deny_url").(string),
		HttpOnlyCookieAttribute:  cloudflare.BoolPtr(d.Get("http_only_cookie_attribute").(bool)),
		SameSiteCookieAttribute:  d.Get("same_site_cookie_attribute").(string),
		LogoURL:                  d.Get("logo_url").(string),
		SkipInterstitial:         cloudflare.BoolPtr(d.Get("skip_interstitial").(bool)),
		AppLauncherVisible:       cloudflare.BoolPtr(d.Get("app_launcher_visible").(bool)),
		ServiceAuth401Redirect:   cloudflare.BoolPtr(d.Get("service_auth_401_redirect").(bool)),
	}

	if _, ok := d.GetOk("allow_authenticate_via_warp"); ok {
		updatedAccessApplication.AllowAuthenticateViaWarp = cloudflare.BoolPtr(d.Get("allow_authenticate_via_warp").(bool))
	}

	if appType != "saas" {
		updatedAccessApplication.Domain = d.Get("domain").(string)
	}

	if value, ok := d.GetOk("allowed_idps"); ok {
		updatedAccessApplication.AllowedIdps = expandInterfaceToStringList(value.(*schema.Set).List())
	}

	if value, ok := d.GetOk("custom_pages"); ok {
		updatedAccessApplication.CustomPages = expandInterfaceToStringList(value.(*schema.Set).List())
	}

	if value, ok := d.GetOk("self_hosted_domains"); ok {
		updatedAccessApplication.SelfHostedDomains = expandInterfaceToStringList(value.(*schema.Set).List())
	}

	if _, ok := d.GetOk("cors_headers"); ok {
		CORSConfig, err := convertCORSSchemaToStruct(d)
		if err != nil {
			return diag.FromErr(err)
		}
		updatedAccessApplication.CorsHeaders = CORSConfig
	}

	if _, ok := d.GetOk("saas_app"); ok {
		saasConfig := convertSaasSchemaToStruct(d)
		updatedAccessApplication.SaasApplication = saasConfig
	}

	if value, ok := d.GetOk("tags"); ok {
		updatedAccessApplication.Tags = expandInterfaceToStringList(value.(*schema.Set).List())
	}
	if appType == "app_launcher" {
		updatedAccessApplication.AccessAppLauncherCustomization = cloudflare.AccessAppLauncherCustomization{
			LogoURL:               d.Get("app_launcher_logo_url").(string),
			BackgroundColor:       d.Get("bg_color").(string),
			HeaderBackgroundColor: d.Get("header_bg_color").(string),
		}

		if _, ok := d.GetOk("landing_page_design"); ok {
			landingPageDesign := convertLandingPageDesignSchemaToStruct(d)
			updatedAccessApplication.AccessAppLauncherCustomization.LandingPageDesign = *landingPageDesign
		}

		if _, ok := d.GetOk("footer_links"); ok {
			footerLinks := convertFooterLinksSchemaToStruct(d)
			updatedAccessApplication.AccessAppLauncherCustomization.FooterLinks = footerLinks
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("Updating Cloudflare Access Application from struct: %+v", updatedAccessApplication))

	identifier, err := initIdentifier(d)
	if err != nil {
		return diag.FromErr(err)
	}

	accessApplication, err := client.UpdateAccessApplication(ctx, identifier, updatedAccessApplication)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error updating Access Application for %s %q: %w", identifier.Level, identifier.Identifier, err))
	}

	if accessApplication.ID == "" {
		return diag.FromErr(fmt.Errorf("failed to find Access Application ID in update response; resource was empty"))
	}

	return resourceCloudflareAccessApplicationRead(ctx, d, meta)
}

func resourceCloudflareAccessApplicationDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	appID := d.Id()

	tflog.Debug(ctx, fmt.Sprintf("Deleting Cloudflare Access Application using ID: %s", appID))

	identifier, err := initIdentifier(d)
	if err != nil {
		return diag.FromErr(err)
	}

	err = client.DeleteAccessApplication(ctx, identifier, appID)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error deleting Access Application for %s %q: %w", identifier.Level, identifier.Identifier, err))
	}

	readErr := resourceCloudflareAccessApplicationRead(ctx, d, meta)
	if readErr != nil {
		return readErr
	}

	return nil
}

func resourceCloudflareAccessApplicationImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	attributes := strings.SplitN(d.Id(), "/", 2)

	if len(attributes) != 2 {
		return nil, fmt.Errorf("invalid id (\"%s\") specified, should be in format \"accountID/accessApplicationID\"", d.Id())
	}

	accountID, accessApplicationID := attributes[0], attributes[1]

	tflog.Debug(ctx, fmt.Sprintf("Importing Cloudflare Access Application: id %s for account %s", accessApplicationID, accountID))

	d.Set(consts.AccountIDSchemaKey, accountID)
	d.SetId(accessApplicationID)

	resourceCloudflareAccessApplicationRead(ctx, d, meta)

	return []*schema.ResourceData{d}, nil
}
