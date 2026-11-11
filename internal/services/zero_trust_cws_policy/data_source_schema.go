// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_cws_policy

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*ZeroTrustCWSPolicyDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "ID of the policy to retrieve",
				Computed:    true,
			},
			"policy_id": schema.StringAttribute{
				Description: "ID of the policy to retrieve",
				Optional:    true,
			},
			"account_id": schema.StringAttribute{
				Description: "Cloudflare account ID (32-character hexadecimal string)",
				Required:    true,
			},
			"created_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"description": schema.StringAttribute{
				Description: "Optional free-form description.",
				Computed:    true,
			},
			"entrypoint": schema.StringAttribute{
				Description: "Non-function rule declared in rego_source. Evaluation uses data.<package_path>.<entrypoint>.",
				Computed:    true,
			},
			"filter": schema.StringAttribute{
				Description: "Filter expression that determines whether this policy applies. Maximum 4096 UTF-8 bytes.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "Human-readable name; unique per account",
				Computed:    true,
			},
			"package_path": schema.StringAttribute{
				Description: "Dotted path derived from rego_source. Begins with policy. and is unique per account.",
				Computed:    true,
			},
			"rego_source": schema.StringAttribute{
				Description: `Inline Rego source with a package beginning with "policy." Maximum 16384 UTF-8 bytes.`,
				Computed:    true,
			},
			"updated_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"tags": schema.MapAttribute{
				Description: "Tags attached to this policy",
				Computed:    true,
				CustomType:  customfield.NewMapType[types.String](ctx),
				ElementType: types.StringType,
			},
			"find_one_by": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"direction": schema.StringAttribute{
						Description: "Sort direction (default: asc)\nAvailable values: \"asc\", \"desc\".",
						Computed:    true,
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("asc", "desc"),
						},
					},
					"name": schema.StringAttribute{
						Description: "Filter by policy name (case-insensitive substring match)",
						Optional:    true,
					},
					"order": schema.StringAttribute{
						Description: "Field to order results by (default: created_at)\nAvailable values: \"name\", \"created_at\", \"updated_at\".",
						Computed:    true,
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"name",
								"created_at",
								"updated_at",
							),
						},
					},
					"package_path": schema.StringAttribute{
						Description: "Filter by package path (case-sensitive literal substring match)",
						Optional:    true,
					},
					"policy_set_id": schema.StringAttribute{
						Description: "Filter to policies attached to (members of) this policy set",
						Optional:    true,
					},
				},
			},
		},
	}
}

func (d *ZeroTrustCWSPolicyDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *ZeroTrustCWSPolicyDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("policy_id"), path.MatchRoot("find_one_by")),
	}
}
