// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_security_allow_policy

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/schemata"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*EmailSecurityAllowPolicyDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: schemata.Description{
			Scopes: []string{
				"Cloud Email Security: Read",
				"Cloud Email Security: Write",
			},
		}.String(),
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Allow policy identifier.",
				Computed:    true,
			},
			"policy_id": schema.StringAttribute{
				Description: "Allow policy identifier.",
				Optional:    true,
			},
			"account_id": schema.StringAttribute{
				Description: "Identifier.",
				Required:    true,
			},
			"comments": schema.StringAttribute{
				Computed: true,
			},
			"created_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"is_acceptable_sender": schema.BoolAttribute{
				Description: "Exempts messages from this sender from Spam, Spoof and Bulk dispositions only; Malicious and Suspicious dispositions still apply.",
				Computed:    true,
			},
			"is_exempt_recipient": schema.BoolAttribute{
				Description: "Bypasses all detections for messages to this recipient.",
				Computed:    true,
			},
			"is_recipient": schema.BoolAttribute{
				Description:        "Deprecated as of July 1, 2025. Use `is_exempt_recipient` instead. End of life: July 1, 2026.",
				Computed:           true,
				DeprecationMessage: "Use `is_exempt_recipient` instead.",
			},
			"is_regex": schema.BoolAttribute{
				Computed: true,
			},
			"is_sender": schema.BoolAttribute{
				Description:        "Deprecated as of July 1, 2025. Use `is_trusted_sender` instead. End of life: July 1, 2026.",
				Computed:           true,
				DeprecationMessage: "Use `is_trusted_sender` instead.",
			},
			"is_spoof": schema.BoolAttribute{
				Description:        "Deprecated as of July 1, 2025. Use `is_acceptable_sender` instead. End of life: July 1, 2026.",
				Computed:           true,
				DeprecationMessage: "Use `is_acceptable_sender` instead.",
			},
			"is_trusted_sender": schema.BoolAttribute{
				Description: "Bypasses all detections and link following for messages from this sender.",
				Computed:    true,
			},
			"last_modified": schema.StringAttribute{
				Description:        "Deprecated, use `modified_at` instead. End of life: November 1, 2026.",
				Computed:           true,
				DeprecationMessage: "Use `modified_at` instead.",
				CustomType:         timetypes.RFC3339Type{},
			},
			"modified_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"pattern": schema.StringAttribute{
				Description: "The pattern value to match. The format depends on `pattern_type`: a valid email address for EMAIL (e.g. `user@example.com`), a valid domain name for DOMAIN (e.g. `example.com`), or a plain IPv4 or IPv6 address or CIDR block for IP (e.g. `1.2.3.4`, `1.2.3.0/24`, `2606:4700:4700::1111`, or `2606:4700:4700::/48`); the API rejects private or unique-local, loopback, link-local, unspecified, and IPv4 broadcast addresses, including their IPv4-mapped IPv6 equivalents.",
				Computed:    true,
			},
			"pattern_type": schema.StringAttribute{
				Description: "Type of pattern matching.\n- EMAIL: matches a full email address (e.g. `user@example.com`)\n- DOMAIN: matches a domain name (e.g. `example.com`)\n- IP: matches a plain IPv4 or IPv6 address (e.g. `1.2.3.4` or `2606:4700:4700::1111`) or CIDR block (e.g. `1.2.3.0/24` or `2606:4700:4700::/48`). The API rejects private or unique-local, loopback, link-local, unspecified, and IPv4 broadcast addresses, including their IPv4-mapped IPv6 equivalents.\n- UNKNOWN: deprecated; you cannot use this when creating or updating policies, but it may appear on existing entries.\nAvailable values: \"EMAIL\", \"DOMAIN\", \"IP\", \"UNKNOWN\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"EMAIL",
						"DOMAIN",
						"IP",
						"UNKNOWN",
					),
				},
			},
			"verify_sender": schema.BoolAttribute{
				Description: "Enforce DMARC, SPF or DKIM authentication. When on, Email Security only honors policies that pass authentication.",
				Computed:    true,
			},
			"filter": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"direction": schema.StringAttribute{
						Description: "The sorting direction.\nAvailable values: \"asc\", \"desc\".",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("asc", "desc"),
						},
					},
					"is_acceptable_sender": schema.BoolAttribute{
						Description: "Filter to show only policies where messages from the sender are exempted from Spam, Spoof, and Bulk dispositions (not Malicious or Suspicious).",
						Optional:    true,
					},
					"is_exempt_recipient": schema.BoolAttribute{
						Description: "Filter to show only policies where messages to the recipient bypass all detections.",
						Optional:    true,
					},
					"is_trusted_sender": schema.BoolAttribute{
						Description: "Filter to show only policies where messages from the sender bypass all detections and link following.",
						Optional:    true,
					},
					"order": schema.StringAttribute{
						Description: "Field to sort by.\nAvailable values: \"pattern\", \"created_at\".",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("pattern", "created_at"),
						},
					},
					"pattern": schema.StringAttribute{
						Optional: true,
					},
					"pattern_type": schema.StringAttribute{
						Description: "Type of pattern matching.\n- EMAIL: matches a full email address (e.g. `user@example.com`)\n- DOMAIN: matches a domain name (e.g. `example.com`)\n- IP: matches a plain IPv4 or IPv6 address (e.g. `1.2.3.4` or `2606:4700:4700::1111`) or CIDR block (e.g. `1.2.3.0/24` or `2606:4700:4700::/48`). The API rejects private or unique-local, loopback, link-local, unspecified, and IPv4 broadcast addresses, including their IPv4-mapped IPv6 equivalents.\n- UNKNOWN: deprecated; you cannot use this when creating or updating policies, but it may appear on existing entries.\nAvailable values: \"EMAIL\", \"DOMAIN\", \"IP\", \"UNKNOWN\".",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"EMAIL",
								"DOMAIN",
								"IP",
								"UNKNOWN",
							),
						},
					},
					"search": schema.StringAttribute{
						Description: "Search term for filtering records. Behavior may change.",
						Optional:    true,
					},
					"verify_sender": schema.BoolAttribute{
						Description: "Filter to show only policies that enforce DMARC, SPF, or DKIM authentication.",
						Optional:    true,
					},
				},
			},
		},
	}
}

func (d *EmailSecurityAllowPolicyDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *EmailSecurityAllowPolicyDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("policy_id"), path.MatchRoot("filter")),
	}
}
