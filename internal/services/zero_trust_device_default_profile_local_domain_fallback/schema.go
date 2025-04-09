// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_device_default_profile_local_domain_fallback

import (
  "context"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
  "github.com/hashicorp/terraform-plugin-framework/resource"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*ZeroTrustDeviceDefaultProfileLocalDomainFallbackResource)(nil)

func ResourceSchema(ctx context.Context) (schema.Schema) {
  return schema.Schema{
    Attributes: map[string]schema.Attribute{
      "account_id": schema.StringAttribute{
        Required: true,
        PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
      },
      "domains": schema.ListNestedAttribute{
        Required: true,
        NestedObject: schema.NestedAttributeObject{
          Attributes: map[string]schema.Attribute{
            "suffix": schema.StringAttribute{
              Description: "The domain suffix to match when resolving locally.",
              Required: true,
            },
            "description": schema.StringAttribute{
              Description: "A description of the fallback domain, displayed in the client UI.",
              Optional: true,
            },
            "dns_server": schema.ListAttribute{
              Description: "A list of IP addresses to handle domain resolution.",
              Optional: true,
              ElementType: types.StringType,
            },
          },
        },
      },
      "description": schema.StringAttribute{
        Description: "A description of the fallback domain, displayed in the client UI.",
        Computed: true,
      },
      "suffix": schema.StringAttribute{
        Description: "The domain suffix to match when resolving locally.",
        Computed: true,
      },
      "dns_server": schema.ListAttribute{
        Description: "A list of IP addresses to handle domain resolution.",
        Computed: true,
        CustomType: customfield.NewListType[types.String](ctx),
        ElementType: types.StringType,
      },
    },
  }
}

func (r *ZeroTrustDeviceDefaultProfileLocalDomainFallbackResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
  resp.Schema = ResourceSchema(ctx)
}

func (r *ZeroTrustDeviceDefaultProfileLocalDomainFallbackResource) ConfigValidators(_ context.Context) ([]resource.ConfigValidator) {
  return []resource.ConfigValidator{
  }
}
