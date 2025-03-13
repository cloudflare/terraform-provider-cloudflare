// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package url_normalization_settings

import (
  "context"

  "github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
  "github.com/hashicorp/terraform-plugin-framework/resource"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
  "github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*URLNormalizationSettingsResource)(nil)

func ResourceSchema(ctx context.Context) (schema.Schema) {
  return schema.Schema{
    Attributes: map[string]schema.Attribute{
      "id": schema.StringAttribute{
        Description: "The unique ID of the zone.",
        Computed: true,
        PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
      },
      "zone_id": schema.StringAttribute{
        Description: "The unique ID of the zone.",
        Required: true,
        PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
      },
      "scope": schema.StringAttribute{
        Description: "The scope of the URL normalization.\nAvailable values: \"incoming\", \"both\".",
        Required: true,
        Validators: []validator.String{
        stringvalidator.OneOfCaseInsensitive("incoming", "both"),
        },
      },
      "type": schema.StringAttribute{
        Description: "The type of URL normalization performed by Cloudflare.\nAvailable values: \"cloudflare\", \"rfc3986\".",
        Required: true,
        Validators: []validator.String{
        stringvalidator.OneOfCaseInsensitive("cloudflare", "rfc3986"),
        },
      },
    },
  }
}

func (r *URLNormalizationSettingsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
  resp.Schema = ResourceSchema(ctx)
}

func (r *URLNormalizationSettingsResource) ConfigValidators(_ context.Context) ([]resource.ConfigValidator) {
  return []resource.ConfigValidator{
  }
}
