package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// SourceAccessServiceTokenSchema returns the source schema for legacy resource.
// Schema version: 0 (SDKv2 default)
// Resource type: cloudflare_access_service_token (deprecated name)
//
// This minimal schema is used only for reading v4 state during migration.
// It includes only the properties needed for state parsing (Required, Optional, Computed).
// Validators, PlanModifiers, and Descriptions are intentionally omitted.
func SourceAccessServiceTokenSchema() schema.Schema {
	return schema.Schema{
		Version: 0, // SDKv2 default schema version
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"account_id": schema.StringAttribute{
				Optional: true,
			},
			"zone_id": schema.StringAttribute{
				Optional: true,
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			"client_id": schema.StringAttribute{
				Computed: true,
			},
			"client_secret": schema.StringAttribute{
				Computed: true,
			},
			"expires_at": schema.StringAttribute{
				Computed: true,
			},
			"min_days_for_renewal": schema.Int64Attribute{
				Optional: true,
			},
			"duration": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"client_secret_version": schema.Int64Attribute{
				Optional: true,
				Computed: true,
			},
			"previous_client_secret_expires_at": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
		},
	}
}
