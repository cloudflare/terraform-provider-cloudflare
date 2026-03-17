package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// SourceCloudflareAccessOrganizationSchema returns the legacy schema (schema_version=0).
// This is used by MoveState and UpgradeFromLegacyV0 to parse state from the legacy SDKv2 provider.
//
// Handles BOTH v4 resource names (they share identical schemas):
// - cloudflare_access_organization
// - cloudflare_zero_trust_access_organization
//
// Reference: https://github.com/cloudflare/terraform-provider-cloudflare/blob/v4/internal/sdkv2provider/schema_cloudflare_access_organization.go
func SourceCloudflareAccessOrganizationSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"account_id": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"zone_id": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"auth_domain": schema.StringAttribute{
				Required: true,
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			"is_ui_read_only": schema.BoolAttribute{
				Optional: true,
			},
			"ui_read_only_toggle_reason": schema.StringAttribute{
				Optional: true,
			},
			"user_seat_expiration_inactive_time": schema.StringAttribute{
				Optional: true,
			},
			"auto_redirect_to_identity": schema.BoolAttribute{
				Optional: true,
			},
			"session_duration": schema.StringAttribute{
				Optional: true,
			},
			"allow_authenticate_via_warp": schema.BoolAttribute{
				Optional: true,
			},
			"warp_auth_session_duration": schema.StringAttribute{
				Optional: true,
			},
		},
		Blocks: map[string]schema.Block{
			// Source login_design is a list block with MaxItems: 1 (SDK v2 pattern)
			// Target login_design is a SingleNestedAttribute
			"login_design": schema.ListNestedBlock{
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"background_color": schema.StringAttribute{
							Optional: true,
						},
						"text_color": schema.StringAttribute{
							Optional: true,
						},
						"logo_path": schema.StringAttribute{
							Optional: true,
						},
						"header_text": schema.StringAttribute{
							Optional: true,
						},
						"footer_text": schema.StringAttribute{
							Optional: true,
						},
					},
				},
			},
			// Source custom_pages is a list block with MaxItems: 1 (SDK v2 pattern)
			// Target custom_pages is a SingleNestedAttribute
			"custom_pages": schema.ListNestedBlock{
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"identity_denied": schema.StringAttribute{
							Optional: true,
						},
						"forbidden": schema.StringAttribute{
							Optional: true,
						},
					},
				},
			},
		},
	}
}
