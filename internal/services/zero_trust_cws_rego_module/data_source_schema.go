// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_cws_rego_module

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*ZeroTrustCWSRegoModuleDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "ID of the Rego module to retrieve",
				Computed:    true,
			},
			"rego_module_id": schema.StringAttribute{
				Description: "ID of the Rego module to retrieve",
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
			"name": schema.StringAttribute{
				Description: "Human-readable name; unique per account",
				Computed:    true,
			},
			"package_path": schema.StringAttribute{
				Description: "Dotted path derived from rego_source. Begins with lib. and is unique per account.",
				Computed:    true,
			},
			"rego_source": schema.StringAttribute{
				Description: `Library Rego source with a package beginning with "lib." Maximum 16384 UTF-8 bytes.`,
				Computed:    true,
			},
			"updated_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
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
						Description: "Filter by Rego module name (case-insensitive substring match)",
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
				},
			},
		},
	}
}

func (d *ZeroTrustCWSRegoModuleDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *ZeroTrustCWSRegoModuleDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("rego_module_id"), path.MatchRoot("filter")),
	}
}
