// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package authenticated_origin_pulls_certificate

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

var _ datasource.DataSourceWithConfigValidators = (*AuthenticatedOriginPullsCertificatesDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"zone_id": schema.StringAttribute{
				Description: "Identifier",
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
				CustomType:  customfield.NewNestedObjectListType[AuthenticatedOriginPullsCertificatesResultDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Identifier",
							Computed:    true,
						},
						"certificate": schema.StringAttribute{
							Description: "The zone's leaf certificate.",
							Computed:    true,
						},
						"expires_on": schema.StringAttribute{
							Description: "When the certificate from the authority expires.",
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
							Description: "Status of the certificate activation.\nAvailable values: \"initializing\", \"pending_deployment\", \"pending_deletion\", \"active\", \"deleted\", \"deployment_timed_out\", \"deletion_timed_out\".",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"initializing",
									"pending_deployment",
									"pending_deletion",
									"active",
									"deleted",
									"deployment_timed_out",
									"deletion_timed_out",
								),
							},
						},
						"uploaded_on": schema.StringAttribute{
							Description: "This is the time the certificate was uploaded.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"enabled": schema.BoolAttribute{
							Description: "Indicates whether zone-level authenticated origin pulls is enabled.",
							Computed:    true,
						},
						"private_key": schema.StringAttribute{
							Description: "The zone's private key.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func (d *AuthenticatedOriginPullsCertificatesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *AuthenticatedOriginPullsCertificatesDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
