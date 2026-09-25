// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_sending_subdomain

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*EmailSendingSubdomainsDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"zone_id": schema.StringAttribute{
				Description: "Identifier.",
				Required:    true,
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
				CustomType:  customfield.NewNestedObjectListType[EmailSendingSubdomainsResultDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Sending subdomain identifier.",
							Computed:    true,
						},
						"enabled": schema.BoolAttribute{
							Description: "Whether Email Sending is enabled on this subdomain.",
							Computed:    true,
						},
						"name": schema.StringAttribute{
							Description: "The exact domain name or a leftmost wildcard such as `*.example.com`.",
							Computed:    true,
						},
						"tag": schema.StringAttribute{
							Description: "Sending subdomain identifier.",
							Computed:    true,
						},
						"created": schema.StringAttribute{
							Description: "The date and time the destination address has been created.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"dkim_selector": schema.StringAttribute{
							Description: "The DKIM selector used for email signing. Wildcard rows publish the selector and sign with `d=<base>`.",
							Computed:    true,
						},
						"drop_suppressed_recipients": schema.BoolAttribute{
							Description: "Whether a send request that includes a recipient suppressed on\nthis subdomain drops that recipient and still delivers to the\nrest, instead of failing the entire request.",
							Computed:    true,
						},
						"modified": schema.StringAttribute{
							Description: "The date and time the destination address was last modified.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"preview_enabled": schema.BoolAttribute{
							Description: "Whether sent messages from this subdomain can be previewed in the activity log.",
							Computed:    true,
						},
						"return_path_domain": schema.StringAttribute{
							Description: "The return-path domain used for bounce handling. Wildcard rows use `cf-bounce.<base>`.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func (d *EmailSendingSubdomainsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *EmailSendingSubdomainsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
