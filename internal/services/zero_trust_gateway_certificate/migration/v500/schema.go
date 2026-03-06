package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// SourceCloudflareZeroTrustGatewayCertificateSchema returns the legacy
// cloudflare_zero_trust_gateway_certificate schema (schema_version=0).
// This is used by UpgradeFromV4 to parse state from the legacy SDKv2 provider
// (implemented as resourceCloudflareTeamsCertificate).
//
// Reference: cloudflare-terraform-v4/internal/sdkv2provider/schema_cloudflare_teams_certificates.go
//
// Only Required, Optional, and Computed properties are included.
// Validators, PlanModifiers, and Descriptions are intentionally omitted —
// this schema is only used for reading existing v4 state, not for planning.
func SourceCloudflareZeroTrustGatewayCertificateSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Required: true,
			},
			// v4-only fields (removed in v5)
			"custom": schema.BoolAttribute{
				Optional: true,
			},
			"gateway_managed": schema.BoolAttribute{
				Optional: true,
			},
			"qs_pack_id": schema.StringAttribute{
				Computed: true,
			},
			// id was Optional+Computed in v4 (Optional for custom certs, Computed for gateway-managed)
			"id": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"validity_period_days": schema.Int64Attribute{
				Optional: true,
				Computed: true,
			},
			"activate": schema.BoolAttribute{
				Optional: true,
				Computed: true,
			},
			// Computed fields
			"in_use": schema.BoolAttribute{
				Computed: true,
			},
			"binding_status": schema.StringAttribute{
				Computed: true,
			},
			"uploaded_on": schema.StringAttribute{
				Computed: true,
			},
			"created_at": schema.StringAttribute{
				Computed: true,
			},
			"expires_on": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}
