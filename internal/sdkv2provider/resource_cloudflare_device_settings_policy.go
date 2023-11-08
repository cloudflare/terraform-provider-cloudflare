package sdkv2provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareDeviceSettingsPolicy() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareDeviceSettingsPolicySchema(),
		CreateContext: resourceCloudflareDeviceSettingsPolicyCreate,
		ReadContext:   resourceCloudflareDeviceSettingsPolicyRead,
		UpdateContext: resourceCloudflareDeviceSettingsPolicyUpdate,
		DeleteContext: resourceCloudflareDeviceSettingsPolicyDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceCloudflareDeviceSettingsPolicyImport,
		},
		Description: "Provides a Cloudflare Device Settings Policy resource. Device policies configure settings applied to WARP devices.",
	}
}

func resourceCloudflareDeviceSettingsPolicyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)
	defaultPolicy := d.Get("default").(bool)

	tflog.Debug(ctx, fmt.Sprintf("Creating Cloudflare device settings policy: accountID=%s default=%t", accountID, defaultPolicy))

	if defaultPolicy {
		d.SetId(accountID)
		return resourceCloudflareDeviceSettingsPolicyUpdate(ctx, d, meta)
	}

	req, err := buildDeviceSettingsPolicyRequest(d)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating Cloudflare device settings policy request: %q: %w", accountID, err))
	}

	policy, err := client.CreateDeviceSettingsPolicy(ctx, cloudflare.AccountIdentifier(accountID), cloudflare.CreateDeviceSettingsPolicyParams{
		DisableAutoFallback: req.DisableAutoFallback,
		CaptivePortal:       req.CaptivePortal,
		AllowModeSwitch:     req.AllowModeSwitch,
		SwitchLocked:        req.SwitchLocked,
		AllowUpdates:        req.AllowUpdates,
		AutoConnect:         req.AutoConnect,
		AllowedToLeave:      req.AllowedToLeave,
		SupportURL:          req.SupportURL,
		ServiceModeV2:       req.ServiceModeV2,
		Precedence:          req.Precedence,
		Name:                req.Name,
		Match:               req.Match,
		Enabled:             req.Enabled,
		ExcludeOfficeIps:    req.ExcludeOfficeIps,
		Description:         req.Description,
	})
	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating Cloudflare device settings policy %q: %w", accountID, err))
	}

	if policy.PolicyID == nil {
		return diag.FromErr(fmt.Errorf("error creating Cloudflare device settings policy: returned policyID was missing after creating policy for account: %q", accountID))
	}
	d.SetId(fmt.Sprintf("%s/%s", accountID, *policy.PolicyID))
	return resourceCloudflareDeviceSettingsPolicyRead(ctx, d, meta)
}

func resourceCloudflareDeviceSettingsPolicyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)
	_, policyID := parseDevicePolicyID(d.Id())

	tflog.Debug(ctx, fmt.Sprintf("Updating Cloudflare device settings policy: accountID=%s policyID=%s", accountID, policyID))

	req, err := buildDeviceSettingsPolicyRequest(d)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating Cloudflare device settings policy request: %q: %w", accountID, err))
	}

	if policyID == "" {
		_, err = client.UpdateDefaultDeviceSettingsPolicy(ctx, cloudflare.AccountIdentifier(accountID), cloudflare.UpdateDefaultDeviceSettingsPolicyParams{
			DisableAutoFallback: req.DisableAutoFallback,
			CaptivePortal:       req.CaptivePortal,
			AllowModeSwitch:     req.AllowModeSwitch,
			SwitchLocked:        req.SwitchLocked,
			AllowUpdates:        req.AllowUpdates,
			AutoConnect:         req.AutoConnect,
			AllowedToLeave:      req.AllowedToLeave,
			SupportURL:          req.SupportURL,
			ServiceModeV2:       req.ServiceModeV2,
			Precedence:          req.Precedence,
			Name:                req.Name,
			Match:               req.Match,
			Enabled:             req.Enabled,
			ExcludeOfficeIps:    req.ExcludeOfficeIps,
			Description:         req.Description,
		})
	} else {
		_, err = client.UpdateDeviceSettingsPolicy(ctx, cloudflare.AccountIdentifier(accountID), cloudflare.UpdateDeviceSettingsPolicyParams{
			PolicyID:            cloudflare.StringPtr(policyID),
			DisableAutoFallback: req.DisableAutoFallback,
			CaptivePortal:       req.CaptivePortal,
			AllowModeSwitch:     req.AllowModeSwitch,
			SwitchLocked:        req.SwitchLocked,
			AllowUpdates:        req.AllowUpdates,
			AutoConnect:         req.AutoConnect,
			AllowedToLeave:      req.AllowedToLeave,
			SupportURL:          req.SupportURL,
			ServiceModeV2:       req.ServiceModeV2,
			Precedence:          req.Precedence,
			Name:                req.Name,
			Match:               req.Match,
			Enabled:             req.Enabled,
			ExcludeOfficeIps:    req.ExcludeOfficeIps,
			Description:         req.Description,
		})
	}
	if err != nil {
		return diag.FromErr(fmt.Errorf("error updating Cloudflare device settings policy %q: %w", accountID, err))
	}

	return resourceCloudflareDeviceSettingsPolicyRead(ctx, d, meta)
}

func resourceCloudflareDeviceSettingsPolicyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)
	_, policyID := parseDevicePolicyID(d.Id())

	var policy cloudflare.DeviceSettingsPolicy
	var err error
	if policyID == "" {
		policy, err = client.GetDefaultDeviceSettingsPolicy(ctx, cloudflare.AccountIdentifier(accountID), cloudflare.GetDefaultDeviceSettingsPolicyParams{})
	} else {
		policy, err = client.GetDeviceSettingsPolicy(ctx, cloudflare.AccountIdentifier(accountID), cloudflare.GetDeviceSettingsPolicyParams{PolicyID: cloudflare.StringPtr(policyID)})
	}

	if err != nil {
		return diag.FromErr(fmt.Errorf("error reading device settings policy %q %s: %w", accountID, policyID, err))
	}

	if err := d.Set("disable_auto_fallback", policy.DisableAutoFallback); err != nil {
		return diag.FromErr(fmt.Errorf("error parsing disable_auto_fallback"))
	}
	if err := d.Set("captive_portal", policy.CaptivePortal); err != nil {
		return diag.FromErr(fmt.Errorf("error parsing captive_portal"))
	}
	if err := d.Set("allow_mode_switch", policy.AllowModeSwitch); err != nil {
		return diag.FromErr(fmt.Errorf("error parsing allow_mode_switch"))
	}
	if err := d.Set("switch_locked", policy.SwitchLocked); err != nil {
		return diag.FromErr(fmt.Errorf("error parsing switch_locked"))
	}
	if err := d.Set("allow_updates", policy.AllowUpdates); err != nil {
		return diag.FromErr(fmt.Errorf("error parsing allow_updates"))
	}
	if err := d.Set("auto_connect", policy.AutoConnect); err != nil {
		return diag.FromErr(fmt.Errorf("error parsing auto_connect"))
	}
	if err := d.Set("allowed_to_leave", policy.AllowedToLeave); err != nil {
		return diag.FromErr(fmt.Errorf("error parsing allowed_to_leave"))
	}
	if err := d.Set("support_url", policy.SupportURL); err != nil {
		return diag.FromErr(fmt.Errorf("error parsing support_url"))
	}
	if err := d.Set("default", policy.Default); err != nil {
		return diag.FromErr(fmt.Errorf("error setting default"))
	}
	if err := d.Set("service_mode_v2_mode", policy.ServiceModeV2.Mode); err != nil {
		return diag.FromErr(fmt.Errorf("error setting service_mode_v2_mode"))
	}
	if err := d.Set("service_mode_v2_port", policy.ServiceModeV2.Port); err != nil {
		return diag.FromErr(fmt.Errorf("error setting service_mode_v2_port"))
	}
	if err := d.Set("exclude_office_ips", policy.ExcludeOfficeIps); err != nil {
		return diag.FromErr(fmt.Errorf("error parsing exclude_office_ips"))
	}
	// ignore setting forbidden fields for default policies
	if policy.Name != nil {
		if err := d.Set("name", policy.Name); err != nil {
			return diag.FromErr(fmt.Errorf("error parsing name"))
		}
	}
	if policy.Description != nil {
		if err := d.Set("description", policy.Description); err != nil {
			return diag.FromErr(fmt.Errorf("error parsing description"))
		}
	}
	if policy.Precedence != nil {
		if err := d.Set("precedence", apiToProviderRulePrecedence(uint64(*policy.Precedence), d.Get("name").(string))); err != nil {
			return diag.FromErr(fmt.Errorf("error parsing precedence"))
		}
	}
	if policy.Match != nil {
		if err := d.Set("match", policy.Match); err != nil {
			return diag.FromErr(fmt.Errorf("error parsing match"))
		}
	}
	if policy.Enabled != nil {
		if err := d.Set("enabled", policy.Enabled); err != nil {
			return diag.FromErr(fmt.Errorf("error parsing enabled"))
		}
	}

	return nil
}

func resourceCloudflareDeviceSettingsPolicyImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	accountID, policyID, err := parseDeviceSettingsIDImport(d.Id())
	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, fmt.Sprintf("Importing Cloudflare device settings policy: id %s for account %s", policyID, accountID))

	d.Set(consts.AccountIDSchemaKey, accountID)
	if policyID == "default" {
		d.SetId(accountID)
	} else {
		d.SetId(fmt.Sprintf("%s/%s", accountID, policyID))
	}

	resourceCloudflareDeviceSettingsPolicyRead(ctx, d, meta)

	return []*schema.ResourceData{d}, nil
}

func resourceCloudflareDeviceSettingsPolicyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	accountID := d.Get(consts.AccountIDSchemaKey).(string)
	_, policyID := parseDevicePolicyID(d.Id())

	client := meta.(*cloudflare.API)
	if policyID == "" {
		d.SetId("")
		return diag.Diagnostics{diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "Cannot delete a default policy, instead simply removing from terraform management without deleting",
		}}
	}

	if _, err := client.DeleteDeviceSettingsPolicy(ctx, cloudflare.AccountIdentifier(accountID), policyID); err != nil {
		return diag.FromErr(fmt.Errorf("error deleting device settings policy %q: %w", accountID, err))
	}

	d.SetId("")

	return nil
}

func buildDeviceSettingsPolicyRequest(d *schema.ResourceData) (cloudflare.DeviceSettingsPolicy, error) {
	defaultPolicy := (d.Get("default").(bool) || d.Id() == d.Get(consts.AccountIDSchemaKey).(string))

	req := cloudflare.DeviceSettingsPolicy{
		DisableAutoFallback: cloudflare.BoolPtr(d.Get("disable_auto_fallback").(bool)),
		CaptivePortal:       cloudflare.IntPtr(d.Get("captive_portal").(int)),
		AllowModeSwitch:     cloudflare.BoolPtr(d.Get("allow_mode_switch").(bool)),
		SwitchLocked:        cloudflare.BoolPtr(d.Get("switch_locked").(bool)),
		AllowUpdates:        cloudflare.BoolPtr(d.Get("allow_updates").(bool)),
		AutoConnect:         cloudflare.IntPtr(d.Get("auto_connect").(int)),
		AllowedToLeave:      cloudflare.BoolPtr(d.Get("allowed_to_leave").(bool)),
		SupportURL:          cloudflare.StringPtr(d.Get("support_url").(string)),
		ServiceModeV2: &cloudflare.ServiceModeV2{
			Mode: cloudflare.ServiceMode(d.Get("service_mode_v2_mode").(string)),
			Port: d.Get("service_mode_v2_port").(int),
		},
		ExcludeOfficeIps: cloudflare.BoolPtr(d.Get("exclude_office_ips").(bool)),
	}

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	enabled := d.Get("enabled").(bool)
	if defaultPolicy && !enabled {
		return req, fmt.Errorf("enabled cannot be false for default policies")
	}
	if !defaultPolicy {
		req.Name = &name
		req.Enabled = &enabled
		req.Description = &description
	}

	match, ok := d.GetOk("match")
	if defaultPolicy && ok {
		return req, fmt.Errorf("match cannot be set for default policies")
	}
	if !defaultPolicy && !ok {
		return req, fmt.Errorf("match must be set for non-default policies")
	}
	if ok {
		matchStr := match.(string)
		req.Match = &matchStr
	}

	precedence, ok := d.GetOk("precedence")
	if defaultPolicy && ok {
		return req, fmt.Errorf("precedence cannot be set for default policies")
	}
	if !defaultPolicy && !ok {
		return req, fmt.Errorf("precedence must be set for non-default policies")
	}
	if ok {
		precedenceVal := int(providerToApiRulePrecedence(int64(precedence.(int)), d.Get("name").(string)))
		req.Precedence = &precedenceVal
	}

	return req, nil
}

func parseDeviceSettingsIDImport(id string) (string, string, error) {
	attributes := strings.SplitN(id, "/", 2)

	if len(attributes) != 2 {
		return "", "", fmt.Errorf("invalid id (\"%s\") specified, should be in format \"accountID/policyID\" or \"accountID/default\" for the default account policy", id)
	}

	return attributes[0], attributes[1], nil
}
