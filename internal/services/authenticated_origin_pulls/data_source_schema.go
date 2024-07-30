// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package authenticated_origin_pulls

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = &AuthenticatedOriginPullsDataSource{}
var _ datasource.DataSourceWithValidateConfig = &AuthenticatedOriginPullsDataSource{}

func (r AuthenticatedOriginPullsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"zone_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"hostname": schema.StringAttribute{
				Description: "The hostname on the origin for which the client certificate uploaded will be used.",
				Computed:    true,
			},
			"cert_id": schema.StringAttribute{
				Description: "Identifier",
				Optional:    true,
			},
			"cert_status": schema.StringAttribute{
				Description: "Status of the certificate or the association.",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("initializing", "pending_deployment", "pending_deletion", "active", "deleted", "deployment_timed_out", "deletion_timed_out"),
				},
			},
			"cert_updated_at": schema.StringAttribute{
				Description: "The time when the certificate was updated.",
				Optional:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"cert_uploaded_on": schema.StringAttribute{
				Description: "The time when the certificate was uploaded.",
				Optional:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"certificate": schema.StringAttribute{
				Description: "The hostname certificate.",
				Optional:    true,
			},
			"created_at": schema.StringAttribute{
				Description: "The time when the certificate was created.",
				Optional:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"enabled": schema.BoolAttribute{
				Description: "Indicates whether hostname-level authenticated origin pulls is enabled. A null value voids the association.",
				Optional:    true,
			},
			"expires_on": schema.StringAttribute{
				Description: "The date when the certificate expires.",
				Optional:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"issuer": schema.StringAttribute{
				Description: "The certificate authority that issued the certificate.",
				Optional:    true,
			},
			"serial_number": schema.StringAttribute{
				Description: "The serial number on the uploaded certificate.",
				Optional:    true,
			},
			"signature": schema.StringAttribute{
				Description: "The type of hash used for the certificate.",
				Optional:    true,
			},
			"status": schema.StringAttribute{
				Description: "Status of the certificate or the association.",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("initializing", "pending_deployment", "pending_deletion", "active", "deleted", "deployment_timed_out", "deletion_timed_out"),
				},
			},
			"updated_at": schema.StringAttribute{
				Description: "The time when the certificate was updated.",
				Optional:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
		},
	}
}

func (r *AuthenticatedOriginPullsDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

func (r *AuthenticatedOriginPullsDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
}
