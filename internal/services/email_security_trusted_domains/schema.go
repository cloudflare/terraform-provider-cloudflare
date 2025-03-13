// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_security_trusted_domains

import (
  "context"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework/resource"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

var _ resource.ResourceWithConfigValidators = (*EmailSecurityTrustedDomainsResource)(nil)

func ResourceSchema(ctx context.Context) (schema.Schema) {
  return schema.Schema{
    Attributes: map[string]schema.Attribute{
      "id": schema.Int64Attribute{
        Description: "The unique identifier for the trusted domain.",
        Computed: true,
        PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
      },
      "account_id": schema.StringAttribute{
        Description: "Account Identifier",
        Required: true,
        PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
      },
      "body": schema.ListNestedAttribute{
        Computed: true,
        Optional: true,
        CustomType: customfield.NewNestedObjectListType[EmailSecurityTrustedDomainsBodyModel](ctx),
        NestedObject: schema.NestedAttributeObject{
          Attributes: map[string]schema.Attribute{
            "is_recent": schema.BoolAttribute{
              Description: "Select to prevent recently registered domains from triggering a\nSuspicious or Malicious disposition.",
              Required: true,
            },
            "is_regex": schema.BoolAttribute{
              Required: true,
            },
            "is_similarity": schema.BoolAttribute{
              Description: "Select for partner or other approved domains that have similar\nspelling to your connected domains. Prevents listed domains from\ntriggering a Spoof disposition.",
              Required: true,
            },
            "pattern": schema.StringAttribute{
              Required: true,
            },
            "comments": schema.StringAttribute{
              Optional: true,
            },
          },
        },
        PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
      },
      "comments": schema.StringAttribute{
        Optional: true,
      },
      "is_recent": schema.BoolAttribute{
        Description: "Select to prevent recently registered domains from triggering a\nSuspicious or Malicious disposition.",
        Optional: true,
      },
      "is_regex": schema.BoolAttribute{
        Optional: true,
      },
      "is_similarity": schema.BoolAttribute{
        Description: "Select for partner or other approved domains that have similar\nspelling to your connected domains. Prevents listed domains from\ntriggering a Spoof disposition.",
        Optional: true,
      },
      "pattern": schema.StringAttribute{
        Optional: true,
      },
      "created_at": schema.StringAttribute{
        Computed: true,
        CustomType: timetypes.RFC3339Type{

        },
      },
      "last_modified": schema.StringAttribute{
        Computed: true,
        CustomType: timetypes.RFC3339Type{

        },
      },
    },
  }
}

func (r *EmailSecurityTrustedDomainsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
  resp.Schema = ResourceSchema(ctx)
}

func (r *EmailSecurityTrustedDomainsResource) ConfigValidators(_ context.Context) ([]resource.ConfigValidator) {
  return []resource.ConfigValidator{
  }
}
