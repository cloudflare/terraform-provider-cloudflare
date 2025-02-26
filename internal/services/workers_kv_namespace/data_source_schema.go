// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_kv_namespace

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*WorkersKVNamespaceDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Namespace identifier tag.",
				Computed:    true,
			},
			"namespace_id": schema.StringAttribute{
				Description: "Namespace identifier tag.",
				Optional:    true,
			},
			"account_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"supports_url_encoding": schema.BoolAttribute{
				Description: "True if keys written on the URL will be URL-decoded before storing. For example, if set to \"true\", a key written on the URL as \"%3F\" will be stored as \"?\".",
				Computed:    true,
			},
			"title": schema.StringAttribute{
				Description: "A human-readable string name for a Namespace.",
				Computed:    true,
			},
			"filter": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"direction": schema.StringAttribute{
						Description: "Direction to order namespaces.\nAvailable values: \"asc\", \"desc\".",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("asc", "desc"),
						},
					},
					"order": schema.StringAttribute{
						Description: "Field to order results by.\nAvailable values: \"id\", \"title\".",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("id", "title"),
						},
					},
				},
			},
		},
	}
}

func (d *WorkersKVNamespaceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *WorkersKVNamespaceDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("namespace_id"), path.MatchRoot("filter")),
	}
}
