// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_kv_namespace

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/schemata"
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
		MarkdownDescription: schemata.Description{
			Scopes: []string{
				"Workers KV Storage Read",
				"Workers KV Storage Write",
			},
		}.String(),
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "ID of the Workers KV namespace.",
				Computed:    true,
			},
			"namespace_id": schema.StringAttribute{
				Description: "ID of the Workers KV namespace.",
				Optional:    true,
			},
			"account_id": schema.StringAttribute{
				Description: "ID of the Cloudflare account that owns the Workers KV namespaces.",
				Optional:    true,
			},
			"jurisdiction": schema.StringAttribute{
				Description: "Specify the jurisdiction to restrict the KV namespace to durably store data within. Can only be set at namespace creation time.\nAvailable values: \"eu\", \"fedramp\", \"us\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"eu",
						"fedramp",
						"us",
					),
				},
			},
			"supports_url_encoding": schema.BoolAttribute{
				Description: `True if keys written on the URL will be URL-decoded before storing. For example, if set to "true", a key written on the URL as "%3F" will be stored as "?".`,
				Computed:    true,
			},
			"title": schema.StringAttribute{
				Description: "Human-readable string name for a Workers KV namespace.",
				Computed:    true,
			},
			"filter": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"direction": schema.StringAttribute{
						Description: "Sort namespaces in ascending (`asc`) or descending (`desc`) order.\nAvailable values: \"asc\", \"desc\".",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("asc", "desc"),
						},
					},
					"order": schema.StringAttribute{
						Description: "Namespace field to sort by (`id` or `title`).\nAvailable values: \"id\", \"title\".",
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
