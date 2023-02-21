package sdkv2provider

import (
	"context"
	"errors"
	"fmt"
	"strings"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareWaitingRoomRules() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCloudflareWaitingRoomRulesCreate,
		ReadContext:   resourceCloudflareWaitingRoomRulesRead,
		UpdateContext: resourceCloudflareWaitingRoomRulesUpdate,
		DeleteContext: resourceCloudflareWaitingRoomRulesDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceCloudflareWaitingRoomRulesImport,
		},

		Schema:      resourceCloudflareWaitingRoomRulesSchema(),
		Description: "Provides a Cloudflare Waiting Room Rules resource.",
	}
}

func expandWaitingRoomRules(d *schema.ResourceData) ([]cloudflare.WaitingRoomRule, error) {
	var waitingRoomRules []cloudflare.WaitingRoomRule
	rules, ok := d.Get("rules").([]interface{})
	if !ok {
		return nil, errors.New("unable to create interface array type assertion")
	}

	for _, v := range rules {
		var rule cloudflare.WaitingRoomRule

		resourceRule, ok := v.(map[string]interface{})
		if !ok {
			return nil, errors.New("unable to create interface map type assertion for rule")
		}

		rule.Action = resourceRule["action"].(string)
		rule.Enabled = statusToAPIEnabledFieldConversion(resourceRule["status"].(string))

		if resourceRule["expression"] != nil {
			rule.Expression = resourceRule["expression"].(string)
		}

		if resourceRule["description"] != nil {
			rule.Description = resourceRule["description"].(string)
		}

		waitingRoomRules = append(waitingRoomRules, rule)
	}

	return waitingRoomRules, nil
}

func flattenWaitingRoomRules(rules []cloudflare.WaitingRoomRule) interface{} {
	var rulesData []map[string]interface{}
	for _, r := range rules {
		rule := map[string]interface{}{
			"id":         r.ID,
			"expression": r.Expression,
			"action":     r.Action,
			"status":     apiEnabledToStatusFieldConversion(r.Enabled),
			"version":    r.Version,
		}

		if r.Description != "" {
			rule["description"] = r.Description
		}

		rulesData = append(rulesData, rule)
	}

	return rulesData
}

func resourceCloudflareWaitingRoomRulesCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	waitingRoomID := d.Get("waiting_room_id").(string)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)
	waitingRoomRules, err := expandWaitingRoomRules(d)

	if err != nil {
		return diag.FromErr(fmt.Errorf("error building waiting room rules %q: %w", waitingRoomID, err))
	}

	_, err = client.ReplaceWaitingRoomRules(ctx, cloudflare.ResourceIdentifier(zoneID), cloudflare.ReplaceWaitingRoomRuleParams{
		WaitingRoomID: waitingRoomID,
		Rules:         waitingRoomRules,
	})

	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating waiting room rules %q: %w", waitingRoomID, err))
	}

	d.SetId(waitingRoomID)

	return resourceCloudflareWaitingRoomRulesRead(ctx, d, meta)
}

func resourceCloudflareWaitingRoomRulesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	waitingRoomID := d.Get("waiting_room_id").(string)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	waitingRoomRules, err := client.ListWaitingRoomRules(ctx, cloudflare.ResourceIdentifier(zoneID), cloudflare.ListWaitingRoomRuleParams{
		WaitingRoomID: waitingRoomID,
	})
	if err != nil {
		var notFoundError *cloudflare.NotFoundError
		if errors.As(err, &notFoundError) {
			tflog.Warn(ctx, fmt.Sprintf("Removing waiting room rules from state because it's not found in API"))
			d.SetId("")
			return nil
		}
		return diag.FromErr(fmt.Errorf("error getting waiting room rules %q: %w", d.Get("name").(string), err))
	}

	if err := d.Set("rules", flattenWaitingRoomRules(waitingRoomRules)); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceCloudflareWaitingRoomRulesUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	waitingRoomID := d.Get("waiting_room_id").(string)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)
	waitingRoomRules, err := expandWaitingRoomRules(d)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error building waiting room rules%q: %w", waitingRoomID, err))
	}

	_, err = client.ReplaceWaitingRoomRules(ctx, cloudflare.ResourceIdentifier(zoneID), cloudflare.ReplaceWaitingRoomRuleParams{
		WaitingRoomID: waitingRoomID,
		Rules:         waitingRoomRules,
	})

	if err != nil {
		return diag.FromErr(fmt.Errorf("error updating waiting room rules %q: %w", waitingRoomID, err))
	}

	return resourceCloudflareWaitingRoomRulesRead(ctx, d, meta)
}

func resourceCloudflareWaitingRoomRulesDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	waitingRoomID := d.Get("waiting_room_id").(string)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	_, err := client.ReplaceWaitingRoomRules(ctx, cloudflare.ResourceIdentifier(zoneID), cloudflare.ReplaceWaitingRoomRuleParams{
		WaitingRoomID: waitingRoomID,
		Rules:         []cloudflare.WaitingRoomRule{},
	})

	if err != nil {
		return diag.FromErr(fmt.Errorf("error deleting waiting room rules %q: %w", waitingRoomID, err))
	}

	return nil
}

func resourceCloudflareWaitingRoomRulesImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	client := meta.(*cloudflare.API)
	idAttr := strings.SplitN(d.Id(), "/", 3)
	var zoneID string
	var waitingRoomID string
	if len(idAttr) == 2 {
		zoneID = idAttr[0]
		waitingRoomID = idAttr[1]
	} else {
		return nil, fmt.Errorf("invalid id (\"%s\") specified, should be in format \"zoneID/waitingRoomID/\" for import", d.Id())
	}

	_, err := client.ListWaitingRoomRules(ctx, cloudflare.ResourceIdentifier(zoneID), cloudflare.ListWaitingRoomRuleParams{
		WaitingRoomID: waitingRoomID,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch Waiting room rules %s", waitingRoomID)
	}

	d.SetId(waitingRoomID)
	d.Set("waiting_room_id", waitingRoomID)
	d.Set(consts.ZoneIDSchemaKey, zoneID)

	resourceCloudflareWaitingRoomRulesRead(ctx, d, meta)

	return []*schema.ResourceData{d}, nil
}
