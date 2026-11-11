// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_cws_policy_set

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

var _ datasource.DataSourceWithConfigValidators = (*ZeroTrustCWSPolicySetDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "ID of the policy set to retrieve",
				Computed:    true,
			},
			"policy_set_id": schema.StringAttribute{
				Description: "ID of the policy set to retrieve",
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
			"name": schema.StringAttribute{
				Description: "Human-readable name; unique per account",
				Computed:    true,
			},
			"updated_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"exclude_workload_ids": schema.SetAttribute{
				Description: "IDs of workloads that exclude matching requests. Exclusions take precedence.",
				Computed:    true,
				CustomType:  customfield.NewSetType[types.String](ctx),
				ElementType: types.StringType,
			},
			"include_workload_ids": schema.SetAttribute{
				Description: "IDs of workloads that include matching requests. Empty includes all requests before exclusions.",
				Computed:    true,
				CustomType:  customfield.NewSetType[types.String](ctx),
				ElementType: types.StringType,
			},
			"tags": schema.MapAttribute{
				Description: "Tags attached to this policy set",
				Computed:    true,
				CustomType:  customfield.NewMapType[types.String](ctx),
				ElementType: types.StringType,
			},
			"policies": schema.SetNestedAttribute{
				Description: "Member policies and their per-membership enforcement mode",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectSetType[ZeroTrustCWSPolicySetPoliciesDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"enforce": schema.BoolAttribute{
							Description: "Whether this policy should enforce when it matches. False keeps the policy attached for non-enforcing evaluation.",
							Computed:    true,
						},
						"policy_id": schema.StringAttribute{
							Description: "ID of the member policy",
							Computed:    true,
						},
					},
				},
			},
			"filter": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"deployment_id": schema.StringAttribute{
						Description: "Filter to policy sets attached to this deployment",
						Optional:    true,
					},
					"direction": schema.StringAttribute{
						Description: "Sort direction (default: asc)\nAvailable values: \"asc\", \"desc\".",
						Computed:    true,
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("asc", "desc"),
						},
					},
					"name": schema.StringAttribute{
						Description: "Filter by policy set name (case-insensitive substring match)",
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
				},
			},
		},
	}
}

func (d *ZeroTrustCWSPolicySetDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *ZeroTrustCWSPolicySetDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("policy_set_id"), path.MatchRoot("filter")),
	}
}
