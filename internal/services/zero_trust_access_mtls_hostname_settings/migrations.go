// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_mtls_hostname_settings

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ resource.ResourceWithUpgradeState = (*ZeroTrustAccessMTLSHostnameSettingsResource)(nil)

// zeroTrustAccessMTLSHostnameSettingsResourceSchemaV0 defines the v0 schema (v4 provider format)
var zeroTrustAccessMTLSHostnameSettingsResourceSchemaV0 = schema.Schema{
	Attributes: map[string]schema.Attribute{
		"account_id": schema.StringAttribute{
			Optional: true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"zone_id": schema.StringAttribute{
			Optional: true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
	},
	Blocks: map[string]schema.Block{
		"settings": schema.ListNestedBlock{
			NestedObject: schema.NestedBlockObject{
				Attributes: map[string]schema.Attribute{
					"hostname": schema.StringAttribute{
						Required: true,
					},
					"china_network": schema.BoolAttribute{
						Optional: true,
					},
					"client_certificate_forwarding": schema.BoolAttribute{
						Optional: true,
					},
				},
			},
		},
	},
}

func (r *ZeroTrustAccessMTLSHostnameSettingsResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		0: {
			PriorSchema:   &zeroTrustAccessMTLSHostnameSettingsResourceSchemaV0,
			StateUpgrader: upgradeZeroTrustAccessMTLSHostnameSettingsStateV0toV1,
		},
	}
}

// upgradeZeroTrustAccessMTLSHostnameSettingsStateV0toV1 migrates from v4 provider state format to v5
func upgradeZeroTrustAccessMTLSHostnameSettingsStateV0toV1(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	// Parse the old state using the raw state data
	var oldState map[string]interface{}
	if req.RawState == nil || len(req.RawState.JSON) == 0 {
		return
	}

	// Extract raw state attributes from JSON
	err := json.Unmarshal(req.RawState.JSON, &oldState)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failed to parse raw state",
			fmt.Sprintf("Could not parse raw state during migration: %s", err),
		)
		return
	}

	// Create new state structure
	var newState ZeroTrustAccessMTLSHostnameSettingsModel

	// Migrate basic attributes
	if accountID, ok := oldState["account_id"].(string); ok && accountID != "" {
		newState.AccountID = types.StringValue(accountID)
	}
	if zoneID, ok := oldState["zone_id"].(string); ok && zoneID != "" {
		newState.ZoneID = types.StringValue(zoneID)
	}

	// Migrate settings from v4 block format to v5 list format
	if settingsData, ok := oldState["settings"].([]interface{}); ok && len(settingsData) > 0 {
		var settings []*ZeroTrustAccessMTLSHostnameSettingsSettingsModel

		for _, settingItem := range settingsData {
			if settingMap, ok := settingItem.(map[string]interface{}); ok {
				var setting ZeroTrustAccessMTLSHostnameSettingsSettingsModel

				// Migrate hostname (required)
				if hostname, ok := settingMap["hostname"].(string); ok {
					setting.Hostname = types.StringValue(hostname)
				}

				// Migrate china_network with default to false if not present
				if chinaNetwork, ok := settingMap["china_network"].(bool); ok {
					setting.ChinaNetwork = types.BoolValue(chinaNetwork)
				} else {
					// Default to false if not present in v4 state
					setting.ChinaNetwork = types.BoolValue(false)
				}

				// Migrate client_certificate_forwarding with default to false if not present
				if clientCertForwarding, ok := settingMap["client_certificate_forwarding"].(bool); ok {
					setting.ClientCertificateForwarding = types.BoolValue(clientCertForwarding)
				} else {
					// Default to false if not present in v4 state
					setting.ClientCertificateForwarding = types.BoolValue(false)
				}

				settings = append(settings, &setting)
			}
		}

		if len(settings) > 0 {
			newState.Settings = &settings
		}
	}

	// Set the upgraded state
	resp.Diagnostics.Append(resp.State.Set(ctx, &newState)...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, fmt.Sprintf("Failed to set new state: %v", resp.Diagnostics.Errors()))
		return
	}
}
