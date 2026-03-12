// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package client_certificate

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*ClientCertificateResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Version: 500,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Identifier.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"zone_id": schema.StringAttribute{
				Description:   "Identifier.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"csr": schema.StringAttribute{
				Description:   "The Certificate Signing Request (CSR). Must be newline-encoded.",
				Required:      true,
				PlanModifiers: []planmodifier.String{utils.RequiresReplaceIfNotCSRSemantic()},
			},
			"validity_days": schema.Int64Attribute{
				Description:   "The number of days the Client Certificate will be valid after the issued_on date",
				Required:      true,
				PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
			},
			"reactivate": schema.BoolAttribute{
				Optional: true,
			},
			"certificate": schema.StringAttribute{
				Description: "The Client Certificate PEM",
				Computed:    true,
			},
			"common_name": schema.StringAttribute{
				Description: "Common Name of the Client Certificate",
				Computed:    true,
			},
			"country": schema.StringAttribute{
				Description: "Country, provided by the CSR",
				Computed:    true,
			},
			"expires_on": schema.StringAttribute{
				Description: "Date that the Client Certificate expires",
				Computed:    true,
			},
			"fingerprint_sha256": schema.StringAttribute{
				Description: "Unique identifier of the Client Certificate",
				Computed:    true,
			},
			"issued_on": schema.StringAttribute{
				Description: "Date that the Client Certificate was issued by the Certificate Authority",
				Computed:    true,
			},
			"location": schema.StringAttribute{
				Description: "Location, provided by the CSR",
				Computed:    true,
			},
			"organization": schema.StringAttribute{
				Description: "Organization, provided by the CSR",
				Computed:    true,
			},
			"organizational_unit": schema.StringAttribute{
				Description: "Organizational Unit, provided by the CSR",
				Computed:    true,
			},
			"serial_number": schema.StringAttribute{
				Description: "The serial number on the created Client Certificate.",
				Computed:    true,
			},
			"signature": schema.StringAttribute{
				Description: "The type of hash used for the Client Certificate..",
				Computed:    true,
			},
			"ski": schema.StringAttribute{
				Description: "Subject Key Identifier",
				Computed:    true,
			},
			"state": schema.StringAttribute{
				Description: "State, provided by the CSR",
				Computed:    true,
			},
			"status": schema.StringAttribute{
				Description: "Client Certificates may be active or revoked, and the pending_reactivation or pending_revocation represent in-progress asynchronous transitions\nAvailable values: \"active\", \"pending_reactivation\", \"pending_revocation\", \"revoked\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"active",
						"pending_reactivation",
						"pending_revocation",
						"revoked",
					),
				},
			},
			"certificate_authority": schema.SingleNestedAttribute{
				Description: "Certificate Authority used to issue the Client Certificate",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[ClientCertificateCertificateAuthorityModel](ctx),
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
					},
					"name": schema.StringAttribute{
						Computed: true,
					},
				},
			},
		},
	}
}

func (r *ClientCertificateResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *ClientCertificateResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
