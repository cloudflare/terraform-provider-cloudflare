package v500

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// V5SchemaV0 is set by parent package init() to avoid circular imports.
// It returns the v5 schema with Version=0 for decoding early published v5 state.
var V5SchemaV0 func(context.Context) *schema.Schema

// UpgradeFromV0 handles schema version 0 states, which may be either:
// - v4 SDKv2 format (list-based nested blocks)
// - early v5 format (object-based nested blocks)
func UpgradeFromV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	if req.RawState == nil {
		resp.Diagnostics.AddError("Missing raw state", "RawState was nil for schema version 0 migration")
		return
	}

	isV4, detectErr := detectV4State(req)
	if detectErr != nil {
		resp.Diagnostics.AddError("Failed to detect state format", fmt.Sprintf("Could not determine v4 vs v5 state: %s", detectErr))
		return
	}

	if isV4 {
		tflog.Info(ctx, "Detected v4 SDKv2 state at version=0, performing v4->v5 transform")
		var v4State SourceV4TunnelConfigModel
		v4Diags := unmarshalV4State(ctx, req.RawState, &v4State)
		resp.Diagnostics.Append(v4Diags...)
		if resp.Diagnostics.HasError() {
			return
		}

		v5State, diags := Transform(ctx, v4State)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

		resp.Diagnostics.Append(resp.State.Set(ctx, v5State)...)
		tflog.Info(ctx, "State upgrade from v4 to v5 completed successfully")
		return
	}

	tflog.Info(ctx, "Detected early v5 state at version=0, performing version bump")
	state, v5Diags := unmarshalV5StateV0(ctx, req.RawState)
	resp.Diagnostics.Append(v5Diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	normalizeV5State(state)
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
	tflog.Info(ctx, "State version bump from 0 to 500 completed")
}

// UpgradeFromV4 handles state upgrades from v4 SDKv2 provider (schema_version=0) to v5 (version=500).
//
// This performs a full transformation from v4 → v5 format including:
// - config: array[0] → pointer (TypeList MaxItems:1 → SingleNestedAttribute)
// - ingress_rule → ingress (rename)
// - warp_routing → dropped
// - origin_request: array[0] → pointer (at both config and ingress level)
// - Duration fields: string ("30s") → Int64 (30 seconds)
// - access: array[0] → pointer (or nil if aud_tag / team_name missing)
// - aud_tag: TypeSet → *[]types.String
// - Dropped: bastion_mode, proxy_address, proxy_port, ip_rules
func UpgradeFromV4(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading zero_trust_tunnel_cloudflared_config state from v4 SDKv2 provider (schema_version=0)")

	// Parse v4 state using v4 model
	var v4State SourceV4TunnelConfigModel
	resp.Diagnostics.Append(req.State.Get(ctx, &v4State)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Transform v4 → v5
	v5State, diags := Transform(ctx, v4State)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Write transformed state
	resp.Diagnostics.Append(resp.State.Set(ctx, v5State)...)
	tflog.Info(ctx, "State upgrade from v4 to v5 completed successfully")
}

// UpgradeFromV5 handles state upgrades from v5 Plugin Framework provider (version=1) to v5 (version=500).
//
// This is mostly a no-op upgrade — the schema is structurally compatible — but normalizes
// origin_request fields that the v4 Plugin Framework provider (v4.x) may have stored as a
// non-null all-null/zero object (e.g., origin_request = {}) rather than null. This happens when
// the v4 provider's Read reads an API-returned origin_request with no meaningful values and
// stores it as a non-null object. Without normalization, terraform plan would detect
// drift: state has non-null all-null object, plan expects null.
func UpgradeFromV5(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading zero_trust_tunnel_cloudflared_config state from version=1 to version=500")

	// Read prior state using v5 schema
	var state TargetV5TunnelConfigModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	normalizeV5State(&state)

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	tflog.Info(ctx, "State version bump from 1 to 500 completed")
}

func normalizeV5State(state *TargetV5TunnelConfigModel) {
	if state.Config == nil {
		return
	}

	if state.Config.OriginRequest != nil && isV5OriginRequestAllNull(state.Config.OriginRequest) {
		state.Config.OriginRequest = nil
	}
	if state.Config.Ingress != nil {
		for _, ingress := range *state.Config.Ingress {
			if ingress != nil && ingress.OriginRequest != nil && isV5IngressOriginRequestAllNull(ingress.OriginRequest) {
				ingress.OriginRequest = nil
			}
		}
	}
}

func detectV4State(req resource.UpgradeStateRequest) (bool, error) {
	if req.RawState != nil && len(req.RawState.JSON) > 0 {
		var rawJSON map[string]interface{}
		if err := json.Unmarshal(req.RawState.JSON, &rawJSON); err == nil {
			if configRaw, ok := rawJSON["config"]; ok {
				switch configRaw.(type) {
				case []interface{}:
					return true, nil
				case map[string]interface{}:
					return false, nil
				}
			}
		}
	}

	return false, fmt.Errorf("unable to determine state format from raw state")
}

func unmarshalV4State(ctx context.Context, rawState *tfprotov6.RawState, target *SourceV4TunnelConfigModel) diag.Diagnostics {
	var diags diag.Diagnostics

	sourceSchema := SourceV4TunnelConfigSchema()
	sourceType := sourceSchema.Type().TerraformType(ctx)

	rawValue, err := rawState.Unmarshal(sourceType)
	if err != nil {
		diags.AddError("Failed to unmarshal v4 state", "Could not parse raw state as v4 format: "+err.Error())
		return diags
	}

	state := tfsdk.State{Raw: rawValue, Schema: sourceSchema}
	diags.Append(state.Get(ctx, target)...)
	return diags
}

func unmarshalV5StateV0(ctx context.Context, rawState *tfprotov6.RawState) (*TargetV5TunnelConfigModel, diag.Diagnostics) {
	var diags diag.Diagnostics
	if V5SchemaV0 == nil {
		diags.AddError("V5 schema initializer missing", "V5SchemaV0 was not configured")
		return nil, diags
	}

	v5Schema := V5SchemaV0(ctx)
	v5Type := v5Schema.Type().TerraformType(ctx)

	rawValue, err := rawState.Unmarshal(v5Type)
	if err != nil {
		diags.AddError("Failed to unmarshal v5 state", "Could not parse raw state as v5 format: "+err.Error())
		return nil, diags
	}

	state := tfsdk.State{Raw: rawValue, Schema: *v5Schema}
	var out TargetV5TunnelConfigModel
	diags.Append(state.Get(ctx, &out)...)
	if diags.HasError() {
		return nil, diags
	}

	return &out, diags
}

// isV5OriginRequestAllNull reports whether a v5 config-level origin_request object has all-null fields.
// Used to normalize v4 PF state that stored origin_request = {} instead of null.
func isV5OriginRequestAllNull(or *TargetV5OriginRequestModel) bool {
	return or.Access == nil &&
		(or.CAPool.IsNull() || or.CAPool.ValueString() == "") &&
		or.ConnectTimeout.IsNull() &&
		or.DisableChunkedEncoding.IsNull() &&
		or.HTTP2Origin.IsNull() &&
		(or.HTTPHostHeader.IsNull() || or.HTTPHostHeader.ValueString() == "") &&
		or.KeepAliveConnections.IsNull() &&
		or.KeepAliveTimeout.IsNull() &&
		or.MatchSnItoHost.IsNull() &&
		or.NoHappyEyeballs.IsNull() &&
		or.NoTLSVerify.IsNull() &&
		(or.OriginServerName.IsNull() || or.OriginServerName.ValueString() == "") &&
		(or.ProxyType.IsNull() || or.ProxyType.ValueString() == "") &&
		or.TCPKeepAlive.IsNull() &&
		or.TLSTimeout.IsNull()
}

// isV5IngressOriginRequestAllNull reports whether a v5 ingress-level origin_request object has all-null fields.
func isV5IngressOriginRequestAllNull(or *TargetV5IngressOriginRequestModel) bool {
	return or.Access == nil &&
		(or.CAPool.IsNull() || or.CAPool.ValueString() == "") &&
		or.ConnectTimeout.IsNull() &&
		or.DisableChunkedEncoding.IsNull() &&
		or.HTTP2Origin.IsNull() &&
		(or.HTTPHostHeader.IsNull() || or.HTTPHostHeader.ValueString() == "") &&
		or.KeepAliveConnections.IsNull() &&
		or.KeepAliveTimeout.IsNull() &&
		or.MatchSnItoHost.IsNull() &&
		or.NoHappyEyeballs.IsNull() &&
		or.NoTLSVerify.IsNull() &&
		(or.OriginServerName.IsNull() || or.OriginServerName.ValueString() == "") &&
		(or.ProxyType.IsNull() || or.ProxyType.ValueString() == "") &&
		or.TCPKeepAlive.IsNull() &&
		or.TLSTimeout.IsNull()
}
