// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waiting_room_rules

import (
  "context"

  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
  "github.com/hashicorp/terraform-plugin-framework/resource"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
  "github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*WaitingRoomRulesResource)(nil)

func ResourceSchema(ctx context.Context) (schema.Schema) {
  return schema.Schema{
    Attributes: map[string]schema.Attribute{
      "id": schema.StringAttribute{
        Description: "The ID of the rule.",
        Computed: true,
        PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
      },
      "waiting_room_id": schema.StringAttribute{
        Required: true,
        PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
      },
      "zone_id": schema.StringAttribute{
        Description: "Identifier",
        Required: true,
        PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
      },
      "rules": schema.ListNestedAttribute{
        Required: true,
        NestedObject: schema.NestedAttributeObject{
          Attributes: map[string]schema.Attribute{
            "action": schema.StringAttribute{
              Description: "The action to take when the expression matches.\nAvailable values: \"bypass_waiting_room\".",
              Required: true,
              Validators: []validator.String{
              stringvalidator.OneOfCaseInsensitive("bypass_waiting_room"),
              },
            },
            "expression": schema.StringAttribute{
              Description: "Criteria defining when there is a match for the current rule.",
              Required: true,
            },
            "description": schema.StringAttribute{
              Description: "The description of the rule.",
              Computed: true,
              Optional: true,
              Default: stringdefault.  StaticString(""),
            },
            "enabled": schema.BoolAttribute{
              Description: "When set to true, the rule is enabled.",
              Computed: true,
              Optional: true,
              Default: booldefault.  StaticBool(true),
            },
          },
        },
      },
      "action": schema.StringAttribute{
        Description: "The action to take when the expression matches.\nAvailable values: \"bypass_waiting_room\".",
        Computed: true,
        Validators: []validator.String{
        stringvalidator.OneOfCaseInsensitive("bypass_waiting_room"),
        },
      },
      "description": schema.StringAttribute{
        Description: "The description of the rule.",
        Computed: true,
        Default: stringdefault.  StaticString(""),
      },
      "enabled": schema.BoolAttribute{
        Description: "When set to true, the rule is enabled.",
        Computed: true,
        Default: booldefault.  StaticBool(true),
      },
      "expression": schema.StringAttribute{
        Description: "Criteria defining when there is a match for the current rule.",
        Computed: true,
      },
      "last_updated": schema.StringAttribute{
        Computed: true,
        CustomType: timetypes.RFC3339Type{

        },
      },
      "version": schema.StringAttribute{
        Description: "The version of the rule.",
        Computed: true,
      },
    },
  }
}

func (r *WaitingRoomRulesResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
  resp.Schema = ResourceSchema(ctx)
}

func (r *WaitingRoomRulesResource) ConfigValidators(_ context.Context) ([]resource.ConfigValidator) {
  return []resource.ConfigValidator{
  }
}
