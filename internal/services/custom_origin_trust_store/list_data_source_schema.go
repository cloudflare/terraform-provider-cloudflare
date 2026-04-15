// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_origin_trust_store

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

var _ datasource.DataSourceWithConfigValidators = (*CustomOriginTrustStoresDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"zone_id": schema.StringAttribute{
				Description: "Identifier.",
				Required:    true,
			},
			"limit": schema.Int64Attribute{
				Description: "Limit to the number of records returned.",
				Optional:    true,
			},
			"offset": schema.Int64Attribute{
				Description: "Offset the results",
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
				CustomType:  customfield.NewNestedObjectListType[CustomOriginTrustStoresResultDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Identifier.",
							Computed:    true,
						},
						"certificate": schema.StringAttribute{
							Description: "The zone's SSL certificate or certificate and the intermediate(s).",
							Computed:    true,
						},
						"expires_on": schema.StringAttribute{
							Description: "When the certificate expires.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"issuer": schema.StringAttribute{
							Description: "The certificate authority that issued the certificate.",
							Computed:    true,
						},
						"signature": schema.StringAttribute{
							Description: "The type of hash used for the certificate.",
							Computed:    true,
						},
						"status": schema.StringAttribute{
							Description: "Status of the zone's custom SSL.\nAvailable values: \"initializing\", \"pending_deployment\", \"active\", \"pending_deletion\", \"deleted\", \"expired\".",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"initializing",
									"pending_deployment",
									"active",
									"pending_deletion",
									"deleted",
									"expired",
								),
							},
						},
						"updated_at": schema.StringAttribute{
							Description: "When the certificate was last modified.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"uploaded_on": schema.StringAttribute{
							Description: "When the certificate was uploaded to Cloudflare.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
					},
				},
			},
		},
	}
}

func (d *CustomOriginTrustStoresDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *CustomOriginTrustStoresDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
