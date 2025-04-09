// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package total_tls

import (
  "context"

  "github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
  "github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
  "github.com/hashicorp/terraform-plugin-framework/resource"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
  "github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*TotalTLSResource)(nil)

func ResourceSchema(ctx context.Context) (schema.Schema) {
  return schema.Schema{
    Attributes: map[string]schema.Attribute{
      "id": schema.StringAttribute{
        Description: "Identifier",
        Computed: true,
        PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
      },
      "zone_id": schema.StringAttribute{
        Description: "Identifier",
        Required: true,
        PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
      },
      "enabled": schema.BoolAttribute{
        Description: "If enabled, Total TLS will order a hostname specific TLS certificate for any proxied A, AAAA, or CNAME record in your zone.",
        Required: true,
        PlanModifiers: []planmodifier.Bool{boolplanmodifier.RequiresReplace()},
      },
      "certificate_authority": schema.StringAttribute{
        Description: "The Certificate Authority that Total TLS certificates will be issued through.\nAvailable values: \"google\", \"lets_encrypt\", \"ssl_com\".",
        Optional: true,
        Validators: []validator.String{
        stringvalidator.OneOfCaseInsensitive(
          "google",
          "lets_encrypt",
          "ssl_com",
        ),
        },
        PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
      },
      "validity_period": schema.Int64Attribute{
        Description: "The validity period in days for the certificates ordered via Total TLS.\nAvailable values: 90.",
        Computed: true,
        Validators: []validator.Int64{
        int64validator.OneOf(90),
        },
      },
    },
  }
}

func (r *TotalTLSResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
  resp.Schema = ResourceSchema(ctx)
}

func (r *TotalTLSResource) ConfigValidators(_ context.Context) ([]resource.ConfigValidator) {
  return []resource.ConfigValidator{
  }
}
