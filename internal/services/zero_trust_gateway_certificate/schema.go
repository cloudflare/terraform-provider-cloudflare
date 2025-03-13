// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_gateway_certificate

import (
  "context"

  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
  "github.com/hashicorp/terraform-plugin-framework/resource"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
  "github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*ZeroTrustGatewayCertificateResource)(nil)

func ResourceSchema(ctx context.Context) (schema.Schema) {
  return schema.Schema{
    Attributes: map[string]schema.Attribute{
      "id": schema.StringAttribute{
        Description: "Certificate UUID tag.",
        Computed: true,
        PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
      },
      "account_id": schema.StringAttribute{
        Required: true,
        PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
      },
      "validity_period_days": schema.Int64Attribute{
        Description: "Number of days the generated certificate will be valid, minimum 1 day and maximum 30 years. Defaults to 5 years.",
        Optional: true,
        PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
      },
      "binding_status": schema.StringAttribute{
        Description: "The deployment status of the certificate on Cloudflare's edge. Certificates in the 'available' (previously called 'active') state may be used for Gateway TLS interception.\nAvailable values: \"pending_deployment\", \"available\", \"pending_deletion\", \"inactive\".",
        Computed: true,
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
        Description: "The CA certificate",
        Computed: true,
      },
      "created_at": schema.StringAttribute{
        Computed: true,
        CustomType: timetypes.RFC3339Type{

        },
      },
      "expires_on": schema.StringAttribute{
        Computed: true,
        CustomType: timetypes.RFC3339Type{

        },
      },
      "fingerprint": schema.StringAttribute{
        Description: "The SHA256 fingerprint of the certificate.",
        Computed: true,
      },
      "in_use": schema.BoolAttribute{
        Description: "Use this certificate for Gateway TLS interception",
        Computed: true,
      },
      "issuer_org": schema.StringAttribute{
        Description: "The organization that issued the certificate.",
        Computed: true,
      },
      "issuer_raw": schema.StringAttribute{
        Description: "The entire issuer field of the certificate.",
        Computed: true,
      },
      "type": schema.StringAttribute{
        Description: "The type of certificate, either BYO-PKI (custom) or Gateway-managed.\nAvailable values: \"custom\", \"gateway_managed\".",
        Computed: true,
        Validators: []validator.String{
        stringvalidator.OneOfCaseInsensitive("custom", "gateway_managed"),
        },
      },
      "updated_at": schema.StringAttribute{
        Computed: true,
        CustomType: timetypes.RFC3339Type{

        },
      },
      "uploaded_on": schema.StringAttribute{
        Computed: true,
        CustomType: timetypes.RFC3339Type{

        },
      },
    },
  }
}

func (r *ZeroTrustGatewayCertificateResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
  resp.Schema = ResourceSchema(ctx)
}

func (r *ZeroTrustGatewayCertificateResource) ConfigValidators(_ context.Context) ([]resource.ConfigValidator) {
  return []resource.ConfigValidator{
  }
}
