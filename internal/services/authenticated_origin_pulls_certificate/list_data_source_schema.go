// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package authenticated_origin_pulls_certificate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = &AuthenticatedOriginPullsCertificatesDataSource{}
var _ datasource.DataSourceWithValidateConfig = &AuthenticatedOriginPullsCertificatesDataSource{}

func (r AuthenticatedOriginPullsCertificatesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"zone_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"max_items": schema.Int64Attribute{
				Description: "Max items to fetch, default: 1000",
				Optional:    true,
			},
			"result": schema.ListNestedAttribute{
				Description: "The items returned by the data source",
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Identifier",
							Computed:    true,
							Optional:    true,
						},
						"certificate": schema.StringAttribute{
							Description: "The zone's leaf certificate.",
							Computed:    true,
							Optional:    true,
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
							Description: "Status of the certificate activation.",
							Computed:    true,
							Optional:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("initializing", "pending_deployment", "pending_deletion", "active", "deleted", "deployment_timed_out", "deletion_timed_out"),
							},
						},
						"uploaded_on": schema.StringAttribute{
							Description: "This is the time the certificate was uploaded.",
							Computed:    true,
							Optional:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
					},
				},
			},
		},
	}
}

func (r *AuthenticatedOriginPullsCertificatesDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

func (r *AuthenticatedOriginPullsCertificatesDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
}
