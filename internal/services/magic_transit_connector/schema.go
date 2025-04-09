// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit_connector

import (
  "context"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
  "github.com/hashicorp/terraform-plugin-framework/resource"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

var _ resource.ResourceWithConfigValidators = (*MagicTransitConnectorResource)(nil)

func ResourceSchema(ctx context.Context) (schema.Schema) {
  return schema.Schema{
    Attributes: map[string]schema.Attribute{
      "id": schema.StringAttribute{
        Computed: true,
        PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
      },
      "account_id": schema.StringAttribute{
        Required: true,
        PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
      },
      "connector_id": schema.StringAttribute{
        Required: true,
        PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
      },
      "activated": schema.BoolAttribute{
        Optional: true,
      },
      "interrupt_window_duration_hours": schema.Float64Attribute{
        Optional: true,
      },
      "interrupt_window_hour_of_day": schema.Float64Attribute{
        Optional: true,
      },
      "notes": schema.StringAttribute{
        Optional: true,
      },
      "timezone": schema.StringAttribute{
        Optional: true,
      },
      "last_heartbeat": schema.StringAttribute{
        Computed: true,
      },
      "last_seen_version": schema.StringAttribute{
        Computed: true,
      },
      "last_updated": schema.StringAttribute{
        Computed: true,
      },
      "device": schema.SingleNestedAttribute{
        Computed: true,
        CustomType: customfield.NewNestedObjectType[MagicTransitConnectorDeviceModel](ctx),
        Attributes: map[string]schema.Attribute{
          "id": schema.StringAttribute{
            Computed: true,
          },
          "serial_number": schema.StringAttribute{
            Computed: true,
          },
        },
      },
    },
  }
}

func (r *MagicTransitConnectorResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
  resp.Schema = ResourceSchema(ctx)
}

func (r *MagicTransitConnectorResource) ConfigValidators(_ context.Context) ([]resource.ConfigValidator) {
  return []resource.ConfigValidator{
  }
}
