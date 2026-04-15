// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_origin_trust_store

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

var _ datasource.DataSourceWithConfigValidators = (*CustomOriginTrustStoreDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Identifier.",
				Computed:    true,
			},
			"custom_origin_trust_store_id": schema.StringAttribute{
				Description: "Identifier.",
				Optional:    true,
			},
			"zone_id": schema.StringAttribute{
				Description: "Identifier.",
				Required:    true,
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
			"filter": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"limit": schema.Int64Attribute{
						Description: "Limit to the number of records returned.",
						Optional:    true,
					},
					"offset": schema.Int64Attribute{
						Description: "Offset the results",
						Optional:    true,
					},
				},
			},
		},
	}
}

func (d *CustomOriginTrustStoreDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *CustomOriginTrustStoreDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("custom_origin_trust_store_id"), path.MatchRoot("filter")),
	}
}
