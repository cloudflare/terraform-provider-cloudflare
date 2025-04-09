// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2_bucket

import (
  "context"

  "github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
  "github.com/hashicorp/terraform-plugin-framework/resource"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
  "github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*R2BucketResource)(nil)

func ResourceSchema(ctx context.Context) (schema.Schema) {
  return schema.Schema{
    Attributes: map[string]schema.Attribute{
      "id": schema.StringAttribute{
        Description: "Name of the bucket",
        Computed: true,
        PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
      },
      "name": schema.StringAttribute{
        Description: "Name of the bucket",
        Required: true,
        PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
      },
      "account_id": schema.StringAttribute{
        Description: "Account ID",
        Required: true,
        PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
      },
      "location": schema.StringAttribute{
        Description: "Location of the bucket\nAvailable values: \"apac\", \"eeur\", \"enam\", \"weur\", \"wnam\", \"oc\".",
        Optional: true,
        Validators: []validator.String{
        stringvalidator.OneOfCaseInsensitive(
          "apac",
          "eeur",
          "enam",
          "weur",
          "wnam",
          "oc",
        ),
        },
      },
      "storage_class": schema.StringAttribute{
        Description: "Storage class for newly uploaded objects, unless specified otherwise.\nAvailable values: \"Standard\", \"InfrequentAccess\".",
        Computed: true,
        Optional: true,
        Validators: []validator.String{
        stringvalidator.OneOfCaseInsensitive("Standard", "InfrequentAccess"),
        },
        Default: stringdefault.  StaticString("Standard"),
      },
      "creation_date": schema.StringAttribute{
        Description: "Creation timestamp",
        Computed: true,
      },
    },
  }
}

func (r *R2BucketResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
  resp.Schema = ResourceSchema(ctx)
}

func (r *R2BucketResource) ConfigValidators(_ context.Context) ([]resource.ConfigValidator) {
  return []resource.ConfigValidator{
  }
}
