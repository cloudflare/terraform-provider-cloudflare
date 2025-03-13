// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudforce_one_request_priority

import (
  "context"

  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
  "github.com/hashicorp/terraform-plugin-framework/resource"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
  "github.com/hashicorp/terraform-plugin-framework/schema/validator"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*CloudforceOneRequestPriorityResource)(nil)

func ResourceSchema(ctx context.Context) (schema.Schema) {
  return schema.Schema{
    Attributes: map[string]schema.Attribute{
      "id": schema.StringAttribute{
        Description: "UUID",
        Computed: true,
        PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
      },
      "account_identifier": schema.StringAttribute{
        Description: "Identifier",
        Required: true,
        PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
      },
      "priority": schema.Int64Attribute{
        Description: "Priority",
        Required: true,
      },
      "requirement": schema.StringAttribute{
        Description: "Requirement",
        Required: true,
      },
      "tlp": schema.StringAttribute{
        Description: "The CISA defined Traffic Light Protocol (TLP)\nAvailable values: \"clear\", \"amber\", \"amber-strict\", \"green\", \"red\".",
        Required: true,
        Validators: []validator.String{
        stringvalidator.OneOfCaseInsensitive(
          "clear",
          "amber",
          "amber-strict",
          "green",
          "red",
        ),
        },
      },
      "labels": schema.ListAttribute{
        Description: "List of labels",
        Required: true,
        ElementType: types.StringType,
      },
      "completed": schema.StringAttribute{
        Computed: true,
        CustomType: timetypes.RFC3339Type{

        },
      },
      "content": schema.StringAttribute{
        Description: "Request content",
        Computed: true,
      },
      "created": schema.StringAttribute{
        Computed: true,
        CustomType: timetypes.RFC3339Type{

        },
      },
      "message_tokens": schema.Int64Attribute{
        Description: "Tokens for the request messages",
        Computed: true,
      },
      "readable_id": schema.StringAttribute{
        Description: "Readable Request ID",
        Computed: true,
      },
      "request": schema.StringAttribute{
        Description: "Requested information from request",
        Computed: true,
      },
      "status": schema.StringAttribute{
        Description: "Request Status\nAvailable values: \"open\", \"accepted\", \"reported\", \"approved\", \"completed\", \"declined\".",
        Computed: true,
        Validators: []validator.String{
        stringvalidator.OneOfCaseInsensitive(
          "open",
          "accepted",
          "reported",
          "approved",
          "completed",
          "declined",
        ),
        },
      },
      "summary": schema.StringAttribute{
        Description: "Brief description of the request",
        Computed: true,
      },
      "tokens": schema.Int64Attribute{
        Description: "Tokens for the request",
        Computed: true,
      },
      "updated": schema.StringAttribute{
        Computed: true,
        CustomType: timetypes.RFC3339Type{

        },
      },
    },
  }
}

func (r *CloudforceOneRequestPriorityResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
  resp.Schema = ResourceSchema(ctx)
}

func (r *CloudforceOneRequestPriorityResource) ConfigValidators(_ context.Context) ([]resource.ConfigValidator) {
  return []resource.ConfigValidator{
  }
}
