// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package firewall_rule

import (
  "context"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
  "github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
  "github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
  "github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
  "github.com/hashicorp/terraform-plugin-framework/resource"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
  "github.com/hashicorp/terraform-plugin-framework/schema/validator"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*FirewallRuleResource)(nil)

func ResourceSchema(ctx context.Context) (schema.Schema) {
  return schema.Schema{
    Attributes: map[string]schema.Attribute{
      "id": schema.StringAttribute{
        Description: "The unique identifier of the firewall rule.",
        Computed: true,
        PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
      },
      "zone_id": schema.StringAttribute{
        Description: "Identifier",
        Required: true,
        PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
      },
      "action": schema.SingleNestedAttribute{
        Description: "The action to perform when the threshold of matched traffic within the configured period is exceeded.",
        Required: true,
        Attributes: map[string]schema.Attribute{
          "mode": schema.StringAttribute{
            Description: "The action to perform.\nAvailable values: \"simulate\", \"ban\", \"challenge\", \"js_challenge\", \"managed_challenge\".",
            Optional: true,
            Validators: []validator.String{
            stringvalidator.OneOfCaseInsensitive(
              "simulate",
              "ban",
              "challenge",
              "js_challenge",
              "managed_challenge",
            ),
            },
          },
          "response": schema.SingleNestedAttribute{
            Description: "A custom content type and reponse to return when the threshold is exceeded. The custom response configured in this object will override the custom error for the zone. This object is optional.\nNotes: If you omit this object, Cloudflare will use the default HTML error page. If \"mode\" is \"challenge\", \"managed_challenge\", or \"js_challenge\", Cloudflare will use the zone challenge pages and you should not provide the \"response\" object.",
            Optional: true,
            Attributes: map[string]schema.Attribute{
              "body": schema.StringAttribute{
                Description: "The response body to return. The value must conform to the configured content type.",
                Optional: true,
              },
              "content_type": schema.StringAttribute{
                Description: "The content type of the body. Must be one of the following: `text/plain`, `text/xml`, or `application/json`.",
                Optional: true,
              },
            },
          },
          "timeout": schema.Float64Attribute{
            Description: "The time in seconds during which Cloudflare will perform the mitigation action. Must be an integer value greater than or equal to the period.\nNotes: If \"mode\" is \"challenge\", \"managed_challenge\", or \"js_challenge\", Cloudflare will use the zone's Challenge Passage time and you should not provide this value.",
            Optional: true,
            Validators: []validator.Float64{
            float64validator.Between(1, 86400),
            },
          },
        },
      },
      "filter": schema.SingleNestedAttribute{
        Required: true,
        Attributes: map[string]schema.Attribute{
          "id": schema.StringAttribute{
            Description: "The unique identifier of the filter.",
            Computed: true,
          },
          "description": schema.StringAttribute{
            Description: "An informative summary of the filter.",
            Optional: true,
          },
          "expression": schema.StringAttribute{
            Description: "The filter expression. For more information, refer to [Expressions](https://developers.cloudflare.com/ruleset-engine/rules-language/expressions/).",
            Optional: true,
          },
          "paused": schema.BoolAttribute{
            Description: "When true, indicates that the filter is currently paused.",
            Optional: true,
          },
          "ref": schema.StringAttribute{
            Description: "A short reference tag. Allows you to select related filters.",
            Optional: true,
          },
        },
      },
      "description": schema.StringAttribute{
        Description: "An informative summary of the firewall rule.",
        Computed: true,
      },
      "paused": schema.BoolAttribute{
        Description: "When true, indicates that the firewall rule is currently paused.",
        Computed: true,
      },
      "priority": schema.Float64Attribute{
        Description: "The priority of the rule. Optional value used to define the processing order. A lower number indicates a higher priority. If not provided, rules with a defined priority will be processed before rules without a priority.",
        Computed: true,
        Validators: []validator.Float64{
        float64validator.Between(0, 2147483647),
        },
      },
      "ref": schema.StringAttribute{
        Description: "A short reference tag. Allows you to select related firewall rules.",
        Computed: true,
      },
      "products": schema.ListAttribute{
        Computed: true,
        Validators: []validator.List{
        listvalidator.ValueStringsAre(
          stringvalidator.OneOfCaseInsensitive(
            "zoneLockdown",
            "uaBlock",
            "bic",
            "hot",
            "securityLevel",
            "rateLimit",
            "waf",
          ),
        ),
        },
        CustomType: customfield.NewListType[types.String](ctx),
        ElementType: types.StringType,
      },
    },
  }
}

func (r *FirewallRuleResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
  resp.Schema = ResourceSchema(ctx)
}

func (r *FirewallRuleResource) ConfigValidators(_ context.Context) ([]resource.ConfigValidator) {
  return []resource.ConfigValidator{
  }
}
