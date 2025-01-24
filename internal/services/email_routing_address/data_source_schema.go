// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_routing_address

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

var _ datasource.DataSourceWithConfigValidators = (*EmailRoutingAddressDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Destination address identifier.",
				Computed:    true,
			},
			"destination_address_identifier": schema.StringAttribute{
				Description: "Destination address identifier.",
				Optional:    true,
			},
			"account_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
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
			"filter": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"direction": schema.StringAttribute{
						Description: "Sorts results in an ascending or descending order.",
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
				},
			},
		},
	}
}

func (d *EmailRoutingAddressDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *EmailRoutingAddressDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("destination_address_identifier"), path.MatchRoot("filter")),
	}
}
