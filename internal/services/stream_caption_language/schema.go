// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stream_caption_language

import (
  "context"

  "github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
  "github.com/hashicorp/terraform-plugin-framework/resource"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
  "github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*StreamCaptionLanguageResource)(nil)

func ResourceSchema(ctx context.Context) (schema.Schema) {
  return schema.Schema{
    Attributes: map[string]schema.Attribute{
      "account_id": schema.StringAttribute{
        Description: "Identifier",
        Required: true,
        PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
      },
      "identifier": schema.StringAttribute{
        Description: "A Cloudflare-generated unique identifier for a media item.",
        Required: true,
        PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
      },
      "language": schema.StringAttribute{
        Description: "The language tag in BCP 47 format.",
        Required: true,
        PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
      },
      "file": schema.StringAttribute{
        Description: "The WebVTT file containing the caption or subtitle content.",
        Optional: true,
      },
      "generated": schema.BoolAttribute{
        Description: "Whether the caption was generated via AI.",
        Computed: true,
      },
      "label": schema.StringAttribute{
        Description: "The language label displayed in the native language to users.",
        Computed: true,
      },
      "status": schema.StringAttribute{
        Description: "The status of a generated caption.\nAvailable values: \"ready\", \"inprogress\", \"error\".",
        Computed: true,
        Validators: []validator.String{
        stringvalidator.OneOfCaseInsensitive(
          "ready",
          "inprogress",
          "error",
        ),
        },
      },
    },
  }
}

func (r *StreamCaptionLanguageResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
  resp.Schema = ResourceSchema(ctx)
}

func (r *StreamCaptionLanguageResource) ConfigValidators(_ context.Context) ([]resource.ConfigValidator) {
  return []resource.ConfigValidator{
  }
}
