// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package image_variant

import (
  "context"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
  "github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
  "github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
  "github.com/hashicorp/terraform-plugin-framework/resource"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
  "github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*ImageVariantResource)(nil)

func ResourceSchema(ctx context.Context) (schema.Schema) {
  return schema.Schema{
    Attributes: map[string]schema.Attribute{
      "id": schema.StringAttribute{
        Required: true,
        PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
      },
      "account_id": schema.StringAttribute{
        Description: "Account identifier tag.",
        Required: true,
        PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
      },
      "options": schema.SingleNestedAttribute{
        Description: "Allows you to define image resizing sizes for different use cases.",
        Required: true,
        Attributes: map[string]schema.Attribute{
          "fit": schema.StringAttribute{
            Description: "The fit property describes how the width and height dimensions should be interpreted.\nAvailable values: \"scale-down\", \"contain\", \"cover\", \"crop\", \"pad\".",
            Required: true,
            Validators: []validator.String{
            stringvalidator.OneOfCaseInsensitive(
              "scale-down",
              "contain",
              "cover",
              "crop",
              "pad",
            ),
            },
          },
          "height": schema.Float64Attribute{
            Description: "Maximum height in image pixels.",
            Required: true,
            Validators: []validator.Float64{
            float64validator.AtLeast(1),
            },
          },
          "metadata": schema.StringAttribute{
            Description: "What EXIF data should be preserved in the output image.\nAvailable values: \"keep\", \"copyright\", \"none\".",
            Required: true,
            Validators: []validator.String{
            stringvalidator.OneOfCaseInsensitive(
              "keep",
              "copyright",
              "none",
            ),
            },
          },
          "width": schema.Float64Attribute{
            Description: "Maximum width in image pixels.",
            Required: true,
            Validators: []validator.Float64{
            float64validator.AtLeast(1),
            },
          },
        },
      },
      "never_require_signed_urls": schema.BoolAttribute{
        Description: "Indicates whether the variant can access an image without a signature, regardless of image access control.",
        Computed: true,
        Optional: true,
        Default: booldefault.  StaticBool(false),
      },
      "variant": schema.SingleNestedAttribute{
        Computed: true,
        CustomType: customfield.NewNestedObjectType[ImageVariantVariantModel](ctx),
        Attributes: map[string]schema.Attribute{
          "id": schema.StringAttribute{
            Computed: true,
          },
          "options": schema.SingleNestedAttribute{
            Description: "Allows you to define image resizing sizes for different use cases.",
            Computed: true,
            CustomType: customfield.NewNestedObjectType[ImageVariantVariantOptionsModel](ctx),
            Attributes: map[string]schema.Attribute{
              "fit": schema.StringAttribute{
                Description: "The fit property describes how the width and height dimensions should be interpreted.\nAvailable values: \"scale-down\", \"contain\", \"cover\", \"crop\", \"pad\".",
                Computed: true,
                Validators: []validator.String{
                stringvalidator.OneOfCaseInsensitive(
                  "scale-down",
                  "contain",
                  "cover",
                  "crop",
                  "pad",
                ),
                },
              },
              "height": schema.Float64Attribute{
                Description: "Maximum height in image pixels.",
                Computed: true,
                Validators: []validator.Float64{
                float64validator.AtLeast(1),
                },
              },
              "metadata": schema.StringAttribute{
                Description: "What EXIF data should be preserved in the output image.\nAvailable values: \"keep\", \"copyright\", \"none\".",
                Computed: true,
                Validators: []validator.String{
                stringvalidator.OneOfCaseInsensitive(
                  "keep",
                  "copyright",
                  "none",
                ),
                },
              },
              "width": schema.Float64Attribute{
                Description: "Maximum width in image pixels.",
                Computed: true,
                Validators: []validator.Float64{
                float64validator.AtLeast(1),
                },
              },
            },
          },
          "never_require_signed_urls": schema.BoolAttribute{
            Description: "Indicates whether the variant can access an image without a signature, regardless of image access control.",
            Computed: true,
            Default: booldefault.  StaticBool(false),
          },
        },
      },
    },
  }
}

func (r *ImageVariantResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
  resp.Schema = ResourceSchema(ctx)
}

func (r *ImageVariantResource) ConfigValidators(_ context.Context) ([]resource.ConfigValidator) {
  return []resource.ConfigValidator{
  }
}
