package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// SourceCloudflareAccountSchema returns the legacy cloudflare_account schema.
// Schema version: 0 (implicit in v4 - no Version field set)
// Resource type: cloudflare_account
//
// This minimal schema is used only for reading v4 state during migration.
// It includes only the properties needed for state parsing (Required, Optional, Computed).
// Validators, PlanModifiers, and Descriptions are intentionally omitted.
//
// Key difference from v5:
// - enforce_twofactor is a top-level BoolAttribute (not nested in settings)
// - settings, unit, managed_by, created_on, abuse_contact_email do NOT exist
//
// Reference: cloudflare-terraform-v4 provider cloudflare_account resource
func SourceCloudflareAccountSchema() schema.Schema {
	return schema.Schema{
		// Version: 0 (implicit - v4 schema had no explicit version)
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			"type": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"enforce_twofactor": schema.BoolAttribute{
				Optional: true,
				Computed: true,
			},
		},
	}
}
