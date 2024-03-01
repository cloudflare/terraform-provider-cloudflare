package sdkv2provider

import (
	"context"
	"encoding/json"
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

func resourceCloudflareDevicePostureRule() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareDevicePostureRuleSchema(),
		CreateContext: resourceCloudflareDevicePostureRuleCreate,
		ReadContext:   resourceCloudflareDevicePostureRuleRead,
		UpdateContext: resourceCloudflareDevicePostureRuleUpdate,
		DeleteContext: resourceCloudflareDevicePostureRuleDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceCloudflareDevicePostureRuleImport,
		},
		Description: heredoc.Doc(`
			Provides a Cloudflare Device Posture Rule resource. Device posture rules configure security policies for device posture checks.
		`),
	}
}

func resourceCloudflareDevicePostureRuleCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)

	newDevicePostureRule := cloudflare.DevicePostureRule{
		Name:        d.Get("name").(string),
		Type:        d.Get("type").(string),
		Description: d.Get("description").(string),
		Schedule:    d.Get("schedule").(string),
		Expiration:  d.Get("expiration").(string),
	}

	err := setDevicePostureRuleMatch(&newDevicePostureRule, d)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating Device Posture Rule with provided match input: %w", err))
	}

	setDevicePostureRuleInput(&newDevicePostureRule, d)
	tflog.Debug(ctx, fmt.Sprintf("Creating Cloudflare Device Posture Rule from struct: %+v", newDevicePostureRule))

	rule, err := client.CreateDevicePostureRule(ctx, accountID, newDevicePostureRule)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating Device Posture Rule for account %q: %w", accountID, err))
	}

	d.SetId(rule.ID)

	return resourceCloudflareDevicePostureRuleRead(ctx, d, meta)
}

func resourceCloudflareDevicePostureRuleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)

	devicePostureRule, err := client.DevicePostureRule(ctx, accountID, d.Id())
	if err != nil {
		var notFoundError *cloudflare.NotFoundError
		if errors.As(err, &notFoundError) {
			tflog.Info(ctx, fmt.Sprintf("Device Posture Rule %s no longer exists", d.Id()))
			d.SetId("")
			return nil
		}
		return diag.FromErr(fmt.Errorf("error finding Device Posture Rule %q: %w", d.Id(), err))
	}

	d.Set("name", devicePostureRule.Name)
	d.Set("description", devicePostureRule.Description)
	d.Set("type", devicePostureRule.Type)
	d.Set("schedule", devicePostureRule.Schedule)
	d.Set("expiration", devicePostureRule.Expiration)
	d.Set("match", convertMatchToSchema(devicePostureRule.Match))
	d.Set("input", convertInputToSchema(devicePostureRule.Input))

	return nil
}

func resourceCloudflareDevicePostureRuleUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)

	updatedDevicePostureRule := cloudflare.DevicePostureRule{
		ID:          d.Id(),
		Name:        d.Get("name").(string),
		Type:        d.Get("type").(string),
		Description: d.Get("description").(string),
		Schedule:    d.Get("schedule").(string),
		Expiration:  d.Get("expiration").(string),
	}

	err := setDevicePostureRuleMatch(&updatedDevicePostureRule, d)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating Device Posture Rule with provided match input: %w", err))
	}

	setDevicePostureRuleInput(&updatedDevicePostureRule, d)
	tflog.Debug(ctx, fmt.Sprintf("Updating Cloudflare Device Posture Rule from struct: %+v", updatedDevicePostureRule))

	devicePostureRule, err := client.UpdateDevicePostureRule(ctx, accountID, updatedDevicePostureRule)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error updating Device Posture Rule for account %q: %w", accountID, err))
	}

	if devicePostureRule.ID == "" {
		return diag.FromErr(fmt.Errorf("failed to find Device Posture Rule ID in update response; resource was empty"))
	}

	return resourceCloudflareDevicePostureRuleRead(ctx, d, meta)
}

func resourceCloudflareDevicePostureRuleDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	appID := d.Id()
	accountID := d.Get(consts.AccountIDSchemaKey).(string)

	tflog.Debug(ctx, fmt.Sprintf("Deleting Cloudflare Device Posture Rule using ID: %s", appID))

	err := client.DeleteDevicePostureRule(ctx, accountID, appID)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error deleting Device Posture Rule for account %q: %w", accountID, err))
	}

	resourceCloudflareDevicePostureRuleRead(ctx, d, meta)

	return nil
}

func resourceCloudflareDevicePostureRuleImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	attributes := strings.SplitN(d.Id(), "/", 2)

	if len(attributes) != 2 {
		return nil, fmt.Errorf("invalid id (\"%s\") specified, should be in format \"accountID/devicePostureRuleID\"", d.Id())
	}

	accountID, devicePostureRuleID := attributes[0], attributes[1]

	tflog.Debug(ctx, fmt.Sprintf("Importing Cloudflare Device Posture Rule: id %s for account %s", devicePostureRuleID, accountID))

	d.Set(consts.AccountIDSchemaKey, accountID)
	d.SetId(devicePostureRuleID)

	resourceCloudflareDevicePostureRuleRead(ctx, d, meta)

	return []*schema.ResourceData{d}, nil
}

func setDevicePostureRuleInput(rule *cloudflare.DevicePostureRule, d *schema.ResourceData) {
	if _, ok := d.GetOk("input"); ok {
		input := cloudflare.DevicePostureRuleInput{}
		if inputID, ok := d.GetOk("input.0.id"); ok {
			input.ID = inputID.(string)
		}
		if p, ok := d.GetOk("input.0.path"); ok {
			input.Path = p.(string)
		}
		if exists, ok := d.GetOk("input.0.exists"); ok {
			input.Exists = exists.(bool)
		}
		if tp, ok := d.GetOk("input.0.thumbprint"); ok {
			input.Thumbprint = tp.(string)
		}
		if s, ok := d.GetOk("input.0.sha256"); ok {
			input.Sha256 = s.(string)
		}
		if running, ok := d.GetOk("input.0.running"); ok {
			input.Running = running.(bool)
		}
		if require_all, ok := d.GetOk("input.0.require_all"); ok {
			input.RequireAll = require_all.(bool)
		}
		if check_disks, ok := d.GetOk("input.0.check_disks"); ok {
			values := check_disks.(*schema.Set).List()
			for _, value := range values {
				input.CheckDisks = append(input.CheckDisks, value.(string))
			}
		}
		if enabled, ok := d.GetOk("input.0.enabled"); ok {
			input.Enabled = enabled.(bool)
		}
		if version, ok := d.GetOk("input.0.version"); ok {
			input.Version = version.(string)
		}
		if operator, ok := d.GetOk("input.0.operator"); ok {
			input.Operator = operator.(string)
		}
		if domain, ok := d.GetOk("input.0.domain"); ok {
			input.Domain = domain.(string)
		}
		if complianceStatus, ok := d.GetOk("input.0.compliance_status"); ok {
			input.ComplianceStatus = complianceStatus.(string)
		}
		if connectionID, ok := d.GetOk("input.0.connection_id"); ok {
			input.ConnectionID = connectionID.(string)
		}
		if osDistroName, ok := d.GetOk("input.0.os_distro_name"); ok {
			input.OsDistroName = osDistroName.(string)
		}
		if osDistroRevision, ok := d.GetOk("input.0.os_distro_revision"); ok {
			input.OsDistroRevision = osDistroRevision.(string)
		}
		if os, ok := d.GetOk("input.0.os"); ok {
			input.Os = os.(string)
		}
		if overall, ok := d.GetOk("input.0.overall"); ok {
			input.Overall = overall.(string)
		}
		if sensorConfig, ok := d.GetOk("input.0.sensor_config"); ok {
			input.SensorConfig = sensorConfig.(string)
		}
		if versionOperator, ok := d.GetOk("input.0.version_operator"); ok {
			input.VersionOperator = versionOperator.(string)
		}
		if state, ok := d.GetOk("input.0.state"); ok {
			input.State = state.(string)
		}
		if last_seen, ok := d.GetOk("input.0.last_seen"); ok {
			input.LastSeen = last_seen.(string)
		}
		if countOperator, ok := d.GetOk("input.0.count_operator"); ok {
			input.CountOperator = countOperator.(string)
		}
		if issueCount, ok := d.GetOk("input.0.issue_count"); ok {
			input.IssueCount = issueCount.(string)
		}
		if certificateId, ok := d.GetOk("input.0.certificate_id"); ok {
			input.CertificateID = certificateId.(string)
		}
		if commonName, ok := d.GetOk("input.0.cn"); ok {
			input.CommonName = commonName.(string)
		}
		if activeThreats, ok := d.GetOk("input.0.active_threats"); ok {
			input.ActiveThreats = activeThreats.(int)
		}
		if networkStatus, ok := d.GetOk("input.0.network_status"); ok {
			input.NetworkStatus = networkStatus.(string)
		}
		if infected, ok := d.GetOk("input.0.infected"); ok {
			input.Infected = infected.(bool)
		}
		if isActive, ok := d.GetOk("input.0.is_active"); ok {
			input.IsActive = isActive.(bool)
		}
		if eidLastSeen, ok := d.GetOk("input.0.eid_last_seen"); ok {
			input.EidLastSeen = eidLastSeen.(string)
		}
		if riskLevel, ok := d.GetOk("input.0.risk_level"); ok {
			input.RiskLevel = riskLevel.(string)
		}
		if totalScore, ok := d.GetOk("input.0.total_score"); ok {
			input.TotalScore = totalScore.(int)
		}
		rule.Input = input
	}
}

func setDevicePostureRuleMatch(rule *cloudflare.DevicePostureRule, d *schema.ResourceData) error {
	if _, ok := d.GetOk("match"); ok {
		match := d.Get("match").([]interface{})
		for _, v := range match {
			jsonString, err := json.Marshal(v.(map[string]interface{}))
			if err != nil {
				return err
			}

			var dprMatch cloudflare.DevicePostureRuleMatch
			err = json.Unmarshal(jsonString, &dprMatch)
			if err != nil {
				return err
			}

			rule.Match = append(rule.Match, dprMatch)
		}
	}

	return nil
}

func convertMatchToSchema(matches []cloudflare.DevicePostureRuleMatch) []map[string]interface{} {
	matchSchema := []map[string]interface{}{}
	for _, match := range matches {
		matchSchema = append(matchSchema, map[string]interface{}{"platform": match.Platform})
	}

	return matchSchema
}

func convertInputToSchema(input cloudflare.DevicePostureRuleInput) []map[string]interface{} {
	m := map[string]interface{}{
		"id":                 input.ID,
		"path":               input.Path,
		"exists":             input.Exists,
		"thumbprint":         input.Thumbprint,
		"sha256":             input.Sha256,
		"running":            input.Running,
		"require_all":        input.RequireAll,
		"check_disks":        input.CheckDisks,
		"enabled":            input.Enabled,
		"version":            input.Version,
		"os_distro_name":     input.OsDistroName,
		"os_distro_revision": input.OsDistroRevision,
		"operator":           input.Operator,
		"domain":             input.Domain,
		"compliance_status":  input.ComplianceStatus,
		"connection_id":      input.ConnectionID,
		"os":                 input.Os,
		"overall":            input.Overall,
		"sensor_config":      input.SensorConfig,
		"version_operator":   input.VersionOperator,
		"count_operator":     input.CountOperator,
		"issue_count":        input.IssueCount,
		"certificate_id":     input.CertificateID,
		"cn":                 input.CommonName,
		"active_threats":     input.ActiveThreats,
		"network_status":     input.NetworkStatus,
		"infected":           input.Infected,
		"is_active":          input.IsActive,
		"eid_last_seen":      input.EidLastSeen,
		"risk_level":         input.RiskLevel,
		"total_score":        input.TotalScore,
	}

	return []map[string]interface{}{m}
}
