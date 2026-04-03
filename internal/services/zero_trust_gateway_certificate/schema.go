// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_gateway_certificate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*ZeroTrustGatewayCertificateResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Identify the certificate with a UUID.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"account_id": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"validity_period_days": schema.Int64Attribute{
				Description:   "Sets the certificate validity period in days (range: 1-10,950 days / ~30 years). Defaults to 1,825 days (5 years). **Important**: This field is only settable during the certificate creation.  Certificates becomes immutable after creation - use the `/activate` and `/deactivate` endpoints to manage certificate lifecycle.",
				Optional:      true,
				PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
			},
			"activate": schema.BoolAttribute{
				Description: "Whether to activate the certificate on Cloudflare's edge. When true, the certificate will be activated. When false, the certificate will be deactivated at the edge. This is a Terraform-only field and does not appear in the API response. Monitor `binding_status` for the activation status. Once a certificate is activated, you may use the certificate to intercept traffic",
				Optional:    true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"binding_status": schema.StringAttribute{
				Description: "Indicate the read-only deployment status of the certificate on Cloudflare's edge. Gateway TLS interception can use certificates in the 'available' (previously called 'active') state.\nAvailable values: \"pending_deployment\", \"available\", \"pending_deletion\", \"inactive\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"pending_deployment",
						"available",
						"pending_deletion",
						"inactive",
					),
				},
			},
			"certificate": schema.StringAttribute{
				Description: "Provide the CA certificate (read-only).",
				Computed:    true,
			},
			"created_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"expires_on": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"fingerprint": schema.StringAttribute{
				Description: "Provide the SHA256 fingerprint of the certificate (read-only).",
				Computed:    true,
			},
			"in_use": schema.BoolAttribute{
				Description: "Indicate whether Gateway TLS interception uses this certificate (read-only). You cannot set this value directly. To configure interception, use the Gateway configuration setting named `certificate` (read-only).",
				Computed:    true,
			},
			"issuer_org": schema.StringAttribute{
				Description: "Indicate the organization that issued the certificate (read-only).",
				Computed:    true,
			},
			"issuer_raw": schema.StringAttribute{
				Description: "Provide the entire issuer field of the certificate (read-only).",
				Computed:    true,
			},
			"type": schema.StringAttribute{
				Description: "Indicate the read-only certificate type, BYO-PKI (custom) or Gateway-managed.\nAvailable values: \"custom\", \"gateway_managed\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("custom", "gateway_managed"),
				},
			},
			"updated_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"uploaded_on": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
		},
	}
}

func (r *ZeroTrustGatewayCertificateResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *ZeroTrustGatewayCertificateResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
