// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2_custom_domain

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*R2CustomDomainDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Account ID",
				Required:    true,
			},
			"bucket_name": schema.StringAttribute{
				Description: "Name of the bucket",
				Required:    true,
			},
			"domain": schema.StringAttribute{
				Description: "Name of the custom domain",
				Required:    true,
			},
			"enabled": schema.BoolAttribute{
				Description: "Whether this bucket is publicly accessible at the specified custom domain",
				Computed:    true,
			},
			"min_tls": schema.StringAttribute{
				Description: "Minimum TLS Version the custom domain will accept for incoming connections. If not set, defaults to 1.0.\nAvailable values: \"1.0\", \"1.1\", \"1.2\", \"1.3\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"1.0",
						"1.1",
						"1.2",
						"1.3",
					),
				},
			},
			"zone_id": schema.StringAttribute{
				Description: "Zone ID of the custom domain resides in",
				Computed:    true,
			},
			"zone_name": schema.StringAttribute{
				Description: "Zone that the custom domain resides in",
				Computed:    true,
			},
			"status": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[R2CustomDomainStatusDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"ownership": schema.StringAttribute{
						Description: "Ownership status of the domain\nAvailable values: \"pending\", \"active\", \"deactivated\", \"blocked\", \"error\", \"unknown\".",
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"pending",
								"active",
								"deactivated",
								"blocked",
								"error",
								"unknown",
							),
						},
					},
					"ssl": schema.StringAttribute{
						Description: "SSL certificate status\nAvailable values: \"initializing\", \"pending\", \"active\", \"deactivated\", \"error\", \"unknown\".",
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"initializing",
								"pending",
								"active",
								"deactivated",
								"error",
								"unknown",
							),
						},
					},
				},
			},
		},
	}
}

func (d *R2CustomDomainDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *R2CustomDomainDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
