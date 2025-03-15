// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package origin_ca_certificate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*OriginCACertificateResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Identifier",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"csr": schema.StringAttribute{
				Description:   "The Certificate Signing Request (CSR). Must be newline-encoded.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"request_type": schema.StringAttribute{
				Description: "Signature type desired on certificate (\"origin-rsa\" (rsa), \"origin-ecc\" (ecdsa), or \"keyless-certificate\" (for Keyless SSL servers).\nAvailable values: \"origin-rsa\", \"origin-ecc\", \"keyless-certificate\".",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"origin-rsa",
						"origin-ecc",
						"keyless-certificate",
					),
				},
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"hostnames": schema.ListAttribute{
				Description:   "Array of hostnames or wildcard names (e.g., *.example.com) bound to the certificate.",
				Optional:      true,
				ElementType:   types.StringType,
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
			},
			"requested_validity": schema.Float64Attribute{
				Description: "The number of days for which the certificate should be valid.\nAvailable values: 7, 30, 90, 365, 730, 1095, 5475.",
				Computed:    true,
				Optional:    true,
				Validators: []validator.Float64{
					float64validator.OneOf(
						7,
						30,
						90,
						365,
						730,
						1095,
						5475,
					),
				},
				PlanModifiers: []planmodifier.Float64{float64planmodifier.RequiresReplace()},
				Default:       float64default.StaticFloat64(5475),
			},
			"certificate": schema.StringAttribute{
				Description: "The Origin CA certificate. Will be newline-encoded.",
				Computed:    true,
			},
			"expires_on": schema.StringAttribute{
				Description: "When the certificate will expire.",
				Computed:    true,
			},
		},
	}
}

func (r *OriginCACertificateResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *OriginCACertificateResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
