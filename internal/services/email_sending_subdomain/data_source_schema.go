// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_sending_subdomain

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = (*EmailSendingSubdomainDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Sending subdomain identifier.",
				Computed:    true,
			},
			"subdomain_id": schema.StringAttribute{
				Description: "Sending subdomain identifier.",
				Required:    true,
			},
			"zone_id": schema.StringAttribute{
				Description: "Identifier.",
				Required:    true,
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
			"enabled": schema.BoolAttribute{
				Description: "Whether Email Sending is enabled on this subdomain.",
				Computed:    true,
			},
			"modified": schema.StringAttribute{
				Description: "The date and time the destination address was last modified.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"name": schema.StringAttribute{
				Description: "The exact domain name or a leftmost wildcard such as `*.example.com`.",
				Computed:    true,
			},
			"preview_enabled": schema.BoolAttribute{
				Description: "Whether sent messages from this subdomain can be previewed in the activity log.",
				Computed:    true,
			},
			"return_path_domain": schema.StringAttribute{
				Description: "The return-path domain used for bounce handling. Wildcard rows use `cf-bounce.<base>`.",
				Computed:    true,
			},
			"tag": schema.StringAttribute{
				Description: "Sending subdomain identifier.",
				Computed:    true,
			},
		},
	}
}

func (d *EmailSendingSubdomainDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *EmailSendingSubdomainDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
