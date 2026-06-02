// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package secrets_store_secret

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/schemata"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*SecretsStoreSecretDataSource)(nil)

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
				Description: "Secret identifier tag.",
				Computed:    true,
			},
			"secret_id": schema.StringAttribute{
				Description: "Secret identifier tag.",
				Optional:    true,
			},
			"account_id": schema.StringAttribute{
				Description: "Account Identifier",
				Required:    true,
			},
			"store_id": schema.StringAttribute{
				Description: "Store Identifier",
				Required:    true,
			},
			"comment": schema.StringAttribute{
				Description: "Freeform text describing the secret",
				Computed:    true,
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
				Description: "The name of the secret",
				Computed:    true,
			},
			"status": schema.StringAttribute{
				Description: `Available values: "pending", "active", "deleted".`,
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"pending",
						"active",
						"deleted",
					),
				},
			},
			"scopes": schema.ListAttribute{
				Description: "The list of services that can use this secret.",
				Computed:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
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
					"scopes": schema.ListAttribute{
						Description: "Only secrets with the given scopes will be returned",
						Optional:    true,
						ElementType: types.ListType{
							ElemType: types.StringType,
						},
					},
					"search": schema.StringAttribute{
						Description: "Search secrets using a filter string, filtering across name and comment",
						Optional:    true,
					},
				},
			},
		},
	}
}

func (d *SecretsStoreSecretDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *SecretsStoreSecretDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("secret_id"), path.MatchRoot("filter")),
	}
}
