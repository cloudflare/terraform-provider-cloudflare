// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package secrets_store

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

var _ datasource.DataSourceWithConfigValidators = (*SecretsStoreDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: schemata.Description{
			Scopes: []string{
				"Secrets Store Read",
				"Secrets Store Write",
			},
		}.String(),
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Store Identifier",
				Computed:    true,
			},
			"store_id": schema.StringAttribute{
				Description: "Store Identifier",
				Optional:    true,
			},
			"account_id": schema.StringAttribute{
				Description: "Account Identifier",
				Required:    true,
			},
			"created": schema.StringAttribute{
				Description: "Whenthe secret was created.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"modified": schema.StringAttribute{
				Description: "When the secret was modified.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"name": schema.StringAttribute{
				Description: "The name of the store",
				Computed:    true,
			},
			"filter": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"direction": schema.StringAttribute{
						Description: "Direction to sort objects\nAvailable values: \"asc\", \"desc\".",
						Computed:    true,
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("asc", "desc"),
						},
					},
					"order": schema.StringAttribute{
						Description: "Order secrets by values in the given field\nAvailable values: \"name\", \"comment\", \"created\", \"modified\", \"status\".",
						Computed:    true,
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"name",
								"comment",
								"created",
								"modified",
								"status",
							),
						},
					},
				},
			},
		},
	}
}

func (d *SecretsStoreDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *SecretsStoreDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("store_id"), path.MatchRoot("filter")),
	}
}
