// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package secrets_store_secret

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/schemata"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*SecretsStoreSecretsDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: schemata.Description{
			Scopes: []string{
				"Secrets Store Read",
				"Secrets Store Write",
			},
		}.String(),
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Account Identifier",
				Required:    true,
			},
			"store_id": schema.StringAttribute{
				Description: "Store Identifier",
				Required:    true,
			},
			"search": schema.StringAttribute{
				Description: "Search secrets using a filter string, filtering across name and comment",
				Optional:    true,
			},
			"scopes": schema.ListAttribute{
				Description: "Only secrets with the given scopes will be returned",
				Optional:    true,
				ElementType: types.ListType{
					ElemType: types.StringType,
				},
			},
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
			"max_items": schema.Int64Attribute{
				Description: "Max items to fetch, default: 1000",
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
			},
			"result": schema.ListNestedAttribute{
				Description: "The items returned by the data source",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[SecretsStoreSecretsResultDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Secret identifier tag.",
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
						"store_id": schema.StringAttribute{
							Description: "Store Identifier",
							Computed:    true,
						},
						"comment": schema.StringAttribute{
							Description: "Freeform text describing the secret",
							Computed:    true,
						},
						"scopes": schema.ListAttribute{
							Description: "The list of services that can use this secret.",
							Computed:    true,
							CustomType:  customfield.NewListType[types.String](ctx),
							ElementType: types.StringType,
						},
					},
				},
			},
		},
	}
}

func (d *SecretsStoreSecretsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *SecretsStoreSecretsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
