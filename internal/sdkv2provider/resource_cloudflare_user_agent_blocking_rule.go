package sdkv2provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
)

func resourceCloudflareUserAgentBlockingRules() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareUserAgentBlockingRulesSchema(),
		CreateContext: resourceCloudflareUserAgentBlockingRulesCreate,
		ReadContext:   resourceCloudflareUserAgentBlockingRulesRead,
		UpdateContext: resourceCloudflareUserAgentBlockingRulesUpdate,
		DeleteContext: resourceCloudflareUserAgentBlockingRulesDelete,
		Importer: &schema.ResourceImporter{
			StateContext: nil,
		},
		Description: heredoc.Doc(`
			Provides a resource to manage User Agent Blocking Rules.
		`),
	}
}

func resourceCloudflareUserAgentBlockingRulesCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	newRule := buildUserAgentBlockingRules(d)

	rule, err := client.CreateUserAgentRule(ctx, zoneID, newRule)
	if err != nil {
		return diag.FromErr(errors.Wrap(err, fmt.Sprintf("failed to create User Agent Blocking Rule")))
	}

	d.SetId(rule.Result.ID)

	return resourceCloudflareUserAgentBlockingRulesRead(ctx, d, meta)
}

func resourceCloudflareUserAgentBlockingRulesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	ua, err := client.UserAgentRule(ctx, zoneID, d.Id())
	if err != nil {
		var notFoundError *cloudflare.NotFoundError
		if errors.As(err, &notFoundError) {
			tflog.Info(ctx, fmt.Sprintf("User Agent Blocking Rule %s no longer exists", d.Id()))
			d.SetId("")
			return nil
		}
		return diag.FromErr(fmt.Errorf("error finding User Agent Blocking Rule %q: %w", d.Id(), err))
	}

	d.Set("paused", ua.Result.Paused)
	d.Set("mode", ua.Result.Mode)
	d.Set("description", ua.Result.Description)
	d.Set("configuration", convertConfigurationToSchema(ua.Result.Configuration))

	return nil
}

func resourceCloudflareUserAgentBlockingRulesUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	ua := buildUserAgentBlockingRules(d)

	_, err := client.UpdateUserAgentRule(ctx, zoneID, d.Id(), ua)
	if err != nil {
		return diag.FromErr(errors.Wrap(err, fmt.Sprintf("failed to update User Agent Blocking Rule")))
	}

	return resourceCloudflareUserAgentBlockingRulesRead(ctx, d, meta)
}

func resourceCloudflareUserAgentBlockingRulesDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	_, err := client.DeleteUserAgentRule(ctx, zoneID, d.Id())
	if err != nil {
		return diag.FromErr(errors.Wrap(err, fmt.Sprintf("failed to delete User Agent Blocking Rule")))
	}

	return resourceCloudflareUserAgentBlockingRulesRead(ctx, d, meta)
}

func buildUserAgentBlockingRules(d *schema.ResourceData) cloudflare.UserAgentRule {
	rule := cloudflare.UserAgentRule{
		Description: d.Get("description").(string),
		Paused:      d.Get("paused").(bool),
		Mode:        d.Get("mode").(string),
	}

	if _, ok := d.GetOk("configuration"); ok {
		configuration := cloudflare.UserAgentRuleConfig{}
		if target, ok := d.GetOk("configuration.0.target"); ok {
			configuration.Target = target.(string)
		}
		if value, ok := d.GetOk("configuration.0.value"); ok {
			configuration.Value = value.(string)
		}
		rule.Configuration = configuration
	}

	return rule
}

func convertConfigurationToSchema(configuration cloudflare.UserAgentRuleConfig) []map[string]interface{} {
	m := map[string]interface{}{
		"target": configuration.Target,
		"value":  configuration.Value,
	}

	return []map[string]interface{}{m}
}
