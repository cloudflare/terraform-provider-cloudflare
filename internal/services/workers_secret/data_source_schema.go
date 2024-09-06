// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_secret

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*WorkersSecretDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Identifier",
				Optional:    true,
			},
			"dispatch_namespace": schema.StringAttribute{
				Description: "Name of the Workers for Platforms dispatch namespace.",
				Optional:    true,
			},
			"script_name": schema.StringAttribute{
				Description: "Name of the script, used in URLs and route configuration.",
				Optional:    true,
			},
			"secret_name": schema.StringAttribute{
				Description: "A JavaScript variable name for the secret binding.",
				Optional:    true,
			},
			"name": schema.StringAttribute{
				Description: "The name of this secret, this is what will be used to access it inside the Worker.",
				Computed:    true,
			},
			"type": schema.StringAttribute{
				Description: "The type of secret.",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("secret_text"),
				},
			},
			"filter": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"account_id": schema.StringAttribute{
						Description: "Identifier",
						Required:    true,
					},
					"dispatch_namespace": schema.StringAttribute{
						Description: "Name of the Workers for Platforms dispatch namespace.",
						Required:    true,
					},
					"script_name": schema.StringAttribute{
						Description: "Name of the script, used in URLs and route configuration.",
						Required:    true,
					},
				},
			},
		},
	}
}

func (d *WorkersSecretDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *WorkersSecretDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.RequiredTogether(
			path.MatchRoot("account_id"),
			path.MatchRoot("dispatch_namespace"),
			path.MatchRoot("script_name"),
			path.MatchRoot("secret_name"),
		),
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("filter"), path.MatchRoot("account_id")),
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("filter"), path.MatchRoot("dispatch_namespace")),
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("filter"), path.MatchRoot("script_name")),
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("filter"), path.MatchRoot("secret_name")),
	}
}
