package gateway_app_types

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func (d *CloudflareGatewayAppTypesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Use this data source to retrieve all Gateway application types for an account.",
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Required:    true,
				Description: "The account ID to fetch Gateway App Types from.",
			},
			"app_types": schema.ListNestedAttribute{
				Computed:    true,
				Description: "A list of Gateway App Types.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							Computed:    true,
							Description: "The identifier for this app type. There is only one app type per ID.",
						},
						"application_type_id": schema.Int64Attribute{
							Computed:    true,
							Description: "The identifier for the application type of this app.",
						},
						"name": schema.StringAttribute{
							Computed:    true,
							Description: "The name of the app type.",
						},
						"description": schema.StringAttribute{
							Computed:    true,
							Description: "A short summary of the app type.",
						},
					},
				},
			},
		},
	}
}
