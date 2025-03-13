// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2_managed_domain

import (
  "context"

  "github.com/hashicorp/terraform-plugin-framework/resource"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

var _ resource.ResourceWithConfigValidators = (*R2ManagedDomainResource)(nil)

func ResourceSchema(ctx context.Context) (schema.Schema) {
  return schema.Schema{
    Attributes: map[string]schema.Attribute{
      "account_id": schema.StringAttribute{
        Description: "Account ID",
        Required: true,
        PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
      },
      "bucket_name": schema.StringAttribute{
        Description: "Name of the bucket",
        Required: true,
        PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
      },
      "enabled": schema.BoolAttribute{
        Description: "Whether to enable public bucket access at the r2.dev domain",
        Required: true,
        PlanModifiers: []planmodifier.Bool{boolplanmodifier.RequiresReplace()},
      },
      "bucket_id": schema.StringAttribute{
        Description: "Bucket ID",
        Computed: true,
      },
      "domain": schema.StringAttribute{
        Description: "Domain name of the bucket's r2.dev domain",
        Computed: true,
      },
    },
  }
}

func (r *R2ManagedDomainResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
  resp.Schema = ResourceSchema(ctx)
}

func (r *R2ManagedDomainResource) ConfigValidators(_ context.Context) ([]resource.ConfigValidator) {
  return []resource.ConfigValidator{
  }
}
