package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// UpgradeFromV0 handles v4 SDKv2 schema version 0 states.
//
// NOTE: early v5 version=0 state is handled by the parent package dispatcher
// (internal/services/zero_trust_tunnel_cloudflared_config/migrations.go) to avoid
// init()-based callback wiring and keep migration flow explicit.
func UpgradeFromV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	if req.RawState == nil {
		resp.Diagnostics.AddError("Missing raw state", "RawState was nil for schema version 0 migration")
		return
	}

	tflog.Info(ctx, "Upgrading zero_trust_tunnel_cloudflared_config state from v4 SDKv2 provider (schema_version=0)")
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
