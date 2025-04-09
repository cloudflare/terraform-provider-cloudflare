// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudforce_one_request

import (
  "context"

  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
  "github.com/hashicorp/terraform-plugin-framework/resource"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
  "github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*CloudforceOneRequestResource)(nil)

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
      "content": schema.StringAttribute{
        Description: "Request content",
        Optional: true,
      },
      "priority": schema.StringAttribute{
        Description: "Priority for analyzing the request",
        Optional: true,
      },
      "request_type": schema.StringAttribute{
        Description: "Requested information from request",
        Optional: true,
      },
      "summary": schema.StringAttribute{
        Description: "Brief description of the request",
        Optional: true,
      },
      "tlp": schema.StringAttribute{
        Description: "The CISA defined Traffic Light Protocol (TLP)\nAvailable values: \"clear\", \"amber\", \"amber-strict\", \"green\", \"red\".",
        Optional: true,
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
      "completed": schema.StringAttribute{
        Computed: true,
        CustomType: timetypes.RFC3339Type{

        },
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

func (r *CloudforceOneRequestResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
  resp.Schema = ResourceSchema(ctx)
}

func (r *CloudforceOneRequestResource) ConfigValidators(_ context.Context) ([]resource.ConfigValidator) {
  return []resource.ConfigValidator{
  }
}
