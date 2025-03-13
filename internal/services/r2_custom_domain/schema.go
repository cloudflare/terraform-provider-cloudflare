// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2_custom_domain

import (
  "context"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
  "github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
  "github.com/hashicorp/terraform-plugin-framework/resource"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
  "github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*R2CustomDomainResource)(nil)

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
      "domain": schema.StringAttribute{
        Description: "Name of the custom domain to be added",
        Required: true,
        PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
      },
      "zone_id": schema.StringAttribute{
        Description: "Zone ID of the custom domain",
        Required: true,
        PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
      },
      "enabled": schema.BoolAttribute{
        Description: "Whether to enable public bucket access at the custom domain. If undefined, the domain will be enabled.",
        Required: true,
      },
      "min_tls": schema.StringAttribute{
        Description: "Minimum TLS Version the custom domain will accept for incoming connections. If not set, defaults to 1.0.\nAvailable values: \"1.0\", \"1.1\", \"1.2\", \"1.3\".",
        Optional: true,
        Validators: []validator.String{
        stringvalidator.OneOfCaseInsensitive(
          "1.0",
          "1.1",
          "1.2",
          "1.3",
        ),
        },
      },
      "zone_name": schema.StringAttribute{
        Description: "Zone that the custom domain resides in",
        Computed: true,
      },
      "status": schema.SingleNestedAttribute{
        Computed: true,
        CustomType: customfield.NewNestedObjectType[R2CustomDomainStatusModel](ctx),
        Attributes: map[string]schema.Attribute{
          "ownership": schema.StringAttribute{
            Description: "Ownership status of the domain\nAvailable values: \"pending\", \"active\", \"deactivated\", \"blocked\", \"error\", \"unknown\".",
            Computed: true,
            Validators: []validator.String{
            stringvalidator.OneOfCaseInsensitive(
              "pending",
              "active",
              "deactivated",
              "blocked",
              "error",
              "unknown",
            ),
            },
          },
          "ssl": schema.StringAttribute{
            Description: "SSL certificate status\nAvailable values: \"initializing\", \"pending\", \"active\", \"deactivated\", \"error\", \"unknown\".",
            Computed: true,
            Validators: []validator.String{
            stringvalidator.OneOfCaseInsensitive(
              "initializing",
              "pending",
              "active",
              "deactivated",
              "error",
              "unknown",
            ),
            },
          },
        },
      },
    },
  }
}

func (r *R2CustomDomainResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
  resp.Schema = ResourceSchema(ctx)
}

func (r *R2CustomDomainResource) ConfigValidators(_ context.Context) ([]resource.ConfigValidator) {
  return []resource.ConfigValidator{
  }
}
