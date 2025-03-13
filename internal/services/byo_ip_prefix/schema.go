// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package byo_ip_prefix

import (
  "context"

  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework/resource"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

var _ resource.ResourceWithConfigValidators = (*ByoIPPrefixResource)(nil)

func ResourceSchema(ctx context.Context) (schema.Schema) {
  return schema.Schema{
    Attributes: map[string]schema.Attribute{
      "id": schema.StringAttribute{
        Description: "Identifier of an IP Prefix.",
        Computed: true,
        PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
      },
      "account_id": schema.StringAttribute{
        Description: "Identifier of a Cloudflare account.",
        Required: true,
        PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
      },
      "asn": schema.Int64Attribute{
        Description: "Autonomous System Number (ASN) the prefix will be advertised under.",
        Required: true,
        PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
      },
      "cidr": schema.StringAttribute{
        Description: "IP Prefix in Classless Inter-Domain Routing format.",
        Required: true,
        PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
      },
      "loa_document_id": schema.StringAttribute{
        Description: "Identifier for the uploaded LOA document.",
        Required: true,
        PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
      },
      "description": schema.StringAttribute{
        Description: "Description of the prefix.",
        Optional: true,
      },
      "advertised": schema.BoolAttribute{
        Description: "Prefix advertisement status to the Internet. This field is only not 'null' if on demand is enabled.",
        Computed: true,
      },
      "advertised_modified_at": schema.StringAttribute{
        Description: "Last time the advertisement status was changed. This field is only not 'null' if on demand is enabled.",
        Computed: true,
        CustomType: timetypes.RFC3339Type{

        },
      },
      "approved": schema.StringAttribute{
        Description: "Approval state of the prefix (P = pending, V = active).",
        Computed: true,
      },
      "created_at": schema.StringAttribute{
        Computed: true,
        CustomType: timetypes.RFC3339Type{

        },
      },
      "modified_at": schema.StringAttribute{
        Computed: true,
        CustomType: timetypes.RFC3339Type{

        },
      },
      "on_demand_enabled": schema.BoolAttribute{
        Description: "Whether advertisement of the prefix to the Internet may be dynamically enabled or disabled.",
        Computed: true,
      },
      "on_demand_locked": schema.BoolAttribute{
        Description: "Whether advertisement status of the prefix is locked, meaning it cannot be changed.",
        Computed: true,
      },
    },
  }
}

func (r *ByoIPPrefixResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
  resp.Schema = ResourceSchema(ctx)
}

func (r *ByoIPPrefixResource) ConfigValidators(_ context.Context) ([]resource.ConfigValidator) {
  return []resource.ConfigValidator{
  }
}
