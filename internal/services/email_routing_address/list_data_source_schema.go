// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_routing_address

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*EmailRoutingAddressesDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"direction": schema.StringAttribute{
				Description: "Sorts results in an ascending or descending order.\nAvailable values: \"asc\", \"desc\".",
				Computed:    true,
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("asc", "desc"),
				},
			},
			"verified": schema.BoolAttribute{
				Description: "Filter by verified destination addresses.",
				Computed:    true,
				Optional:    true,
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
				CustomType:  customfield.NewNestedObjectListType[EmailRoutingAddressesResultDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Destination address identifier.",
							Computed:    true,
						},
						"created": schema.StringAttribute{
							Description: "The date and time the destination address has been created.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"email": schema.StringAttribute{
							Description: "The contact email address of the user.",
							Computed:    true,
						},
						"modified": schema.StringAttribute{
							Description: "The date and time the destination address was last modified.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"tag": schema.StringAttribute{
							Description: "Destination address tag. (Deprecated, replaced by destination address identifier)",
							Computed:    true,
						},
						"verified": schema.StringAttribute{
							Description: "The date and time the destination address has been verified. Null means not verified yet.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
					},
				},
			},
		},
	}
}

func (d *EmailRoutingAddressesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *EmailRoutingAddressesDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
