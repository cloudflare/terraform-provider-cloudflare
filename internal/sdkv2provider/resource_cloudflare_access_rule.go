package sdkv2provider

import (
	"context"
	"errors"
	"fmt"
	"net"
	"strings"

	"github.com/MakeNowJust/heredoc/v2"
	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareAccessRule() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareAccessRuleSchema(),
		CreateContext: resourceCloudflareAccessRuleCreate,
		ReadContext:   resourceCloudflareAccessRuleRead,
		UpdateContext: resourceCloudflareAccessRuleUpdate,
		DeleteContext: resourceCloudflareAccessRuleDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceCloudflareAccessRuleImport,
		},

		SchemaVersion: 1,

		StateUpgraders: []schema.StateUpgrader{
			{
				Type:    resourceCloudflareAccessRuleV0().CoreConfigSchema().ImpliedType(),
				Upgrade: resourceCloudflareAccessRuleStateUpgradeV1,
				Version: 0,
			},
		},
		Description: heredoc.Doc(`
			Provides a Cloudflare IP Firewall Access Rule resource. Access
			control can be applied on basis of IP addresses, IP ranges, AS
			numbers or countries.
		`),
	}
}

func resourceCloudflareAccessRuleCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)

	newRule := cloudflare.AccessRule{
		Notes: d.Get("notes").(string),
		Mode:  d.Get("mode").(string),
	}

	if configuration, configurationOk := d.GetOk("configuration"); configurationOk {
		for _, config := range configuration.([]interface{}) {
			newRule.Configuration = cloudflare.AccessRuleConfiguration{
				Target: config.(map[string]interface{})["target"].(string),
				Value:  config.(map[string]interface{})["value"].(string),
			}
		}
	}

	var r *cloudflare.AccessRuleResponse
	var err error

	if accountID != "" {
		r, err = client.CreateAccountAccessRule(ctx, accountID, newRule)
	} else {
		r, err = client.CreateZoneAccessRule(ctx, zoneID, newRule)
	}

	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to create access rule: %w", err))
	}

	if r.Result.ID == "" {
		return diag.FromErr(fmt.Errorf("Failed to find access rule in Create response; ID was empty"))
	}

	d.SetId(r.Result.ID)

	return resourceCloudflareAccessRuleRead(ctx, d, meta)
}

func resourceCloudflareAccessRuleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)

	var accessRuleResponse *cloudflare.AccessRuleResponse
	var err error

	if accountID != "" {
		accessRuleResponse, err = client.AccountAccessRule(ctx, accountID, d.Id())
	} else {
		accessRuleResponse, err = client.ZoneAccessRule(ctx, zoneID, d.Id())
	}

	tflog.Debug(ctx, fmt.Sprintf("accessRuleResponse: %#v", accessRuleResponse))
	tflog.Debug(ctx, fmt.Sprintf("accessRuleResponse error: %#v", err))

	if err != nil {
		var notFoundError *cloudflare.NotFoundError
		if errors.As(err, &notFoundError) {
			tflog.Info(ctx, fmt.Sprintf("Access Rule %s no longer exists", d.Id()))
			d.SetId("")
			return nil
		}
		return diag.FromErr(fmt.Errorf("error finding access rule %q: %w", d.Id(), err))
	}

	tflog.Debug(ctx, fmt.Sprintf("Cloudflare Access Rule read configuration: %#v", accessRuleResponse))

	d.Set(consts.ZoneIDSchemaKey, zoneID)
	d.Set("mode", accessRuleResponse.Result.Mode)
	d.Set("notes", accessRuleResponse.Result.Notes)
	tflog.Debug(ctx, fmt.Sprintf("read configuration: %#v", d.Get("configuration")))

	configuration := []map[string]interface{}{}
	configuration = append(configuration, map[string]interface{}{
		"target": accessRuleResponse.Result.Configuration.Target,
		"value":  accessRuleResponse.Result.Configuration.Value,
	})

	d.Set("configuration", configuration)

	return nil
}

func resourceCloudflareAccessRuleUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)

	updatedRule := cloudflare.AccessRule{
		Notes: d.Get("notes").(string),
		Mode:  d.Get("mode").(string),
	}

	if configuration, configurationOk := d.GetOk("configuration"); configurationOk {
		for _, config := range configuration.([]interface{}) {
			updatedRule.Configuration = cloudflare.AccessRuleConfiguration{
				Target: config.(map[string]interface{})["target"].(string),
				Value:  config.(map[string]interface{})["value"].(string),
			}
		}
	}

	var err error

	if accountID != "" {
		_, err = client.UpdateAccountAccessRule(ctx, accountID, d.Id(), updatedRule)
	} else {
		_, err = client.UpdateZoneAccessRule(ctx, zoneID, d.Id(), updatedRule)
	}

	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to update Access Rule: %w", err))
	}

	return resourceCloudflareAccessRuleRead(ctx, d, meta)
}

func resourceCloudflareAccessRuleDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)

	tflog.Info(ctx, fmt.Sprintf("Deleting Cloudflare Access Rule: id %s", d.Id()))

	var err error

	if accountID != "" {
		_, err = client.DeleteAccountAccessRule(ctx, accountID, d.Id())
	} else {
		_, err = client.DeleteZoneAccessRule(ctx, zoneID, d.Id())
	}

	if err != nil {
		return diag.FromErr(fmt.Errorf("error deleting Cloudflare Access Rule: %w", err))
	}

	return nil
}

func resourceCloudflareAccessRuleImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	attributes := strings.Split(d.Id(), "/")

	var (
		accessRuleType           string
		accessRuleTypeIdentifier string
		accessRuleID             string
	)

	if len(attributes) != 3 {
		return nil, fmt.Errorf("invalid id (\"%s\") specified, should be in format \"accessRuleType/accessRuleTypeIdentifier/identiferValue\"", d.Id())
	}

	accessRuleType, accessRuleTypeIdentifier, accessRuleID = attributes[0], attributes[1], attributes[2]

	d.SetId(accessRuleID)

	switch accessRuleType {
	case "account":
		d.Set(consts.AccountIDSchemaKey, accessRuleTypeIdentifier)
	case "zone":
		d.Set(consts.ZoneIDSchemaKey, accessRuleTypeIdentifier)
	}

	resourceCloudflareAccessRuleRead(ctx, d, meta)

	return []*schema.ResourceData{d}, nil
}

func configurationDiffSuppress(k, old, new string, d *schema.ResourceData) bool {
	switch {
	case d.Get("configuration.0.target") == "ip6" && k == "configuration.0.value":
		existingIP := net.ParseIP(old)
		incomingIP := net.ParseIP(new)
		return existingIP.Equal(incomingIP)
	case d.Get("configuration.0.target") == "country" && k == "configuration.0.value":
		return strings.ToUpper(old) == strings.ToUpper(new)
	case d.Get("configuration.0.target") == "asn" && k == "configuration.0.value":
		if !strings.HasPrefix(strings.ToUpper(new), "AS") {
			new = "AS" + strings.ToUpper(new)
		}

		return strings.ToUpper(old) == strings.ToUpper(new)
	}

	return false
}

func validateAccessRuleConfiguration(v interface{}, k string) (warnings []string, errors []error) {
	config := v.(map[string]interface{})

	target := config["target"].(string)
	value := config["value"].(string)

	switch target {
	case "ip_range":
		return validateAccessRuleConfigurationIPRange(value)
	default:
	}

	return warnings, errors
}

func validateAccessRuleConfigurationIPRange(v string) (warnings []string, errors []error) {
	ip, ipNet, err := net.ParseCIDR(v)
	if err != nil {
		errors = append(errors, fmt.Errorf("failed to parse value as CIDR: %w", err))
		return warnings, errors
	}

	if ipNet == nil {
		errors = append(errors, fmt.Errorf("ip_range must hold a range"))
		return warnings, errors
	}

	if ip.To4() != nil {
		ones, _ := ipNet.Mask.Size()
		if ones != 16 && ones != 24 {
			errors = append(errors, fmt.Errorf("ip_range with ipv4 address must be a /16 or /24, got a /%d", ones))
			return warnings, errors
		}
	} else {
		ones, _ := ipNet.Mask.Size()
		if ones != 32 && ones != 48 && ones != 64 {
			errors = append(errors, fmt.Errorf("ip_range with ipv6 address must be in (/32, /48, /64), instead got a /%d", ones))
			return warnings, errors
		}
	}

	return warnings, errors
}
