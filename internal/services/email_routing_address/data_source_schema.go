// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_routing_address

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = &EmailRoutingAddressDataSource{}
var _ datasource.DataSourceWithValidateConfig = &EmailRoutingAddressDataSource{}

func (r EmailRoutingAddressDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_identifier": schema.StringAttribute{
				Description: "Identifier",
				Optional:    true,
			},
			"destination_address_identifier": schema.StringAttribute{
				Description: "Destination address identifier.",
				Optional:    true,
			},
			"id": schema.StringAttribute{
				Description: "Destination address identifier.",
				Optional:    true,
			},
			"created": schema.StringAttribute{
				Description: "The date and time the destination address has been created.",
				Optional:    true,
			},
			"email": schema.StringAttribute{
				Description: "The contact email address of the user.",
				Optional:    true,
			},
			"modified": schema.StringAttribute{
				Description: "The date and time the destination address was last modified.",
				Optional:    true,
			},
			"tag": schema.StringAttribute{
				Description: "Destination address tag. (Deprecated, replaced by destination address identifier)",
				Optional:    true,
			},
			"verified": schema.StringAttribute{
				Description: "The date and time the destination address has been verified. Null means not verified yet.",
				Optional:    true,
			},
			"find_one_by": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"account_identifier": schema.StringAttribute{
						Description: "Identifier",
						Required:    true,
					},
					"direction": schema.StringAttribute{
						Description: "Sorts results in an ascending or descending order.",
						Computed:    true,
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("asc", "desc"),
						},
					},
					"page": schema.Float64Attribute{
						Description: "Page number of paginated results.",
						Computed:    true,
						Optional:    true,
					},
					"per_page": schema.Float64Attribute{
						Description: "Maximum number of results per page.",
						Computed:    true,
						Optional:    true,
					},
					"verified": schema.BoolAttribute{
						Description: "Filter by verified destination addresses.",
						Computed:    true,
						Optional:    true,
					},
				},
			},
		},
	}
}

func (r *EmailRoutingAddressDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

func (r *EmailRoutingAddressDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
}
