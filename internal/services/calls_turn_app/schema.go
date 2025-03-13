// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package calls_turn_app

import (
  "context"

  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework/resource"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

var _ resource.ResourceWithConfigValidators = (*CallsTURNAppResource)(nil)

func ResourceSchema(ctx context.Context) (schema.Schema) {
  return schema.Schema{
    Attributes: map[string]schema.Attribute{
      "account_id": schema.StringAttribute{
        Description: "The account identifier tag.",
        Required: true,
        PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
      },
      "key_id": schema.StringAttribute{
        Description: "A Cloudflare-generated unique identifier for a item.",
        Optional: true,
        PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
      },
      "name": schema.StringAttribute{
        Description: "A short description of a TURN key, not shown to end users.",
        Computed: true,
        Optional: true,
        Default: stringdefault.  StaticString(""),
      },
      "created": schema.StringAttribute{
        Description: "The date and time the item was created.",
        Computed: true,
        CustomType: timetypes.RFC3339Type{

        },
      },
      "key": schema.StringAttribute{
        Description: "Bearer token",
        Computed: true,
        Sensitive: true,
      },
      "modified": schema.StringAttribute{
        Description: "The date and time the item was last modified.",
        Computed: true,
        CustomType: timetypes.RFC3339Type{

        },
      },
      "uid": schema.StringAttribute{
        Description: "A Cloudflare-generated unique identifier for a item.",
        Computed: true,
      },
    },
  }
}

func (r *CallsTURNAppResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
  resp.Schema = ResourceSchema(ctx)
}

func (r *CallsTURNAppResource) ConfigValidators(_ context.Context) ([]resource.ConfigValidator) {
  return []resource.ConfigValidator{
  }
}
