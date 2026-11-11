// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_cws_deployment

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

var _ datasource.DataSourceWithConfigValidators = (*ZeroTrustCWSDeploymentDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "ID of the deployment to retrieve",
				Computed:    true,
			},
			"deployment_id": schema.StringAttribute{
				Description: "ID of the deployment to retrieve",
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
			"default_mode": schema.StringAttribute{
				Description: "Default traffic decision when no policy set matches\nAvailable values: \"ALLOW\", \"BLOCK\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("ALLOW", "BLOCK"),
				},
			},
			"description": schema.StringAttribute{
				Description: "Optional free-form description.",
				Computed:    true,
			},
			"domain_acl_id": schema.StringAttribute{
				Description: "ID of the attached domain ACL. Omitted when none.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "Human-readable name; unique per account",
				Computed:    true,
			},
			"on_error_action": schema.StringAttribute{
				Description: "Action applied when an enforcing policy has a runtime Rego error\nAvailable values: \"BLOCK\", \"SKIP\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("BLOCK", "SKIP"),
				},
			},
			"rego_raise_error": schema.BoolAttribute{
				Description: "Whether OPA builtin failures are raised as runtime evaluation errors",
				Computed:    true,
			},
			"updated_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"basic_auth_credential_ids": schema.SetAttribute{
				Description: "IDs of basic-auth credentials attached to this deployment",
				Computed:    true,
				CustomType:  customfield.NewSetType[types.String](ctx),
				ElementType: types.StringType,
			},
			"policy_set_ids": schema.SetAttribute{
				Description: "IDs of policy sets attached to this deployment",
				Computed:    true,
				CustomType:  customfield.NewSetType[types.String](ctx),
				ElementType: types.StringType,
			},
			"tags": schema.MapAttribute{
				Description: "Tags attached to this deployment",
				Computed:    true,
				CustomType:  customfield.NewMapType[types.String](ctx),
				ElementType: types.StringType,
			},
			"filter": schema.SingleNestedAttribute{
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
						Description: "Filter by deployment name (case-insensitive substring match)",
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

func (d *ZeroTrustCWSDeploymentDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *ZeroTrustCWSDeploymentDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("deployment_id"), path.MatchRoot("filter")),
	}
}
