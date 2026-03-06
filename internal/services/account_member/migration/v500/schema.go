package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// SourceV4Schema returns the v4 SDKv2 account_member schema for state parsing.
//
// Reference: cloudflare-terraform-v4/internal/sdkv2provider/schema_cloudflare_account_member.go
//
// Key differences from v5:
//   - email_address (v4) instead of email (v5)
//   - role_ids (v4) instead of roles (v5)
//   - No policies field (v5 only)
//   - No user field (v5 only)
func SourceV4Schema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"account_id": schema.StringAttribute{
				Required: true,
			},
			"email_address": schema.StringAttribute{
				Required: true,
			},
			"role_ids": schema.SetAttribute{
				Required:    true,
				ElementType: types.StringType,
			},
			"status": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
		},
	}
}
