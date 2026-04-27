// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_security_block_sender

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

var _ datasource.DataSourceWithConfigValidators = (*EmailSecurityBlockSenderDataSource)(nil)

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
				Description: "Blocked sender pattern identifier",
				Computed:    true,
			},
			"pattern_id": schema.StringAttribute{
				Description: "Blocked sender pattern identifier",
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
			"is_regex": schema.BoolAttribute{
				Computed: true,
			},
			"last_modified": schema.StringAttribute{
				Description:        "Deprecated, use `modified_at` instead. End of life: November 1, 2026.",
				Computed:           true,
				DeprecationMessage: "This attribute is deprecated.",
				CustomType:         timetypes.RFC3339Type{},
			},
			"modified_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"pattern": schema.StringAttribute{
				Computed: true,
			},
			"pattern_type": schema.StringAttribute{
				Description: "Type of pattern matching.\nNote: UNKNOWN is deprecated and cannot be used when creating or updating policies, but may be returned for existing entries.\nAvailable values: \"EMAIL\", \"DOMAIN\", \"IP\", \"UNKNOWN\".",
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
					"order": schema.StringAttribute{
						Description: "Field to sort by.\nAvailable values: \"pattern\", \"created_at\".",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("pattern", "created_at"),
						},
					},
					"pattern": schema.StringAttribute{
						Description: "Filter by pattern value.",
						Optional:    true,
					},
					"pattern_type": schema.StringAttribute{
						Description: "Filter by pattern type.\nAvailable values: \"EMAIL\", \"DOMAIN\", \"IP\", \"UNKNOWN\".",
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
				},
			},
		},
	}
}

func (d *EmailSecurityBlockSenderDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *EmailSecurityBlockSenderDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("pattern_id"), path.MatchRoot("filter")),
	}
}
