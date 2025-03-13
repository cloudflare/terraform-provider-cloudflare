// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package web_analytics_site

import (
  "context"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
  "github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
  "github.com/hashicorp/terraform-plugin-framework/datasource"
  "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
  "github.com/hashicorp/terraform-plugin-framework/path"
  "github.com/hashicorp/terraform-plugin-framework/schema/validator"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*WebAnalyticsSiteDataSource)(nil)

func DataSourceSchema(ctx context.Context) (schema.Schema) {
  return schema.Schema{
    Attributes: map[string]schema.Attribute{
      "id": schema.StringAttribute{
        Description: "Identifier",
        Computed: true,
      },
      "site_id": schema.StringAttribute{
        Description: "Identifier",
        Optional: true,
      },
      "account_id": schema.StringAttribute{
        Description: "Identifier",
        Required: true,
      },
      "auto_install": schema.BoolAttribute{
        Description: "If enabled, the JavaScript snippet is automatically injected for orange-clouded sites.",
        Computed: true,
      },
      "created": schema.StringAttribute{
        Computed: true,
        CustomType: timetypes.RFC3339Type{

        },
      },
      "site_tag": schema.StringAttribute{
        Description: "The Web Analytics site identifier.",
        Computed: true,
      },
      "site_token": schema.StringAttribute{
        Description: "The Web Analytics site token.",
        Computed: true,
      },
      "snippet": schema.StringAttribute{
        Description: "Encoded JavaScript snippet.",
        Computed: true,
      },
      "rules": schema.ListNestedAttribute{
        Description: "A list of rules.",
        Computed: true,
        CustomType: customfield.NewNestedObjectListType[WebAnalyticsSiteRulesDataSourceModel](ctx),
        NestedObject: schema.NestedAttributeObject{
          Attributes: map[string]schema.Attribute{
            "id": schema.StringAttribute{
              Description: "The Web Analytics rule identifier.",
              Computed: true,
            },
            "created": schema.StringAttribute{
              Computed: true,
              CustomType: timetypes.RFC3339Type{

              },
            },
            "host": schema.StringAttribute{
              Description: "The hostname the rule will be applied to.",
              Computed: true,
            },
            "inclusive": schema.BoolAttribute{
              Description: "Whether the rule includes or excludes traffic from being measured.",
              Computed: true,
            },
            "is_paused": schema.BoolAttribute{
              Description: "Whether the rule is paused or not.",
              Computed: true,
            },
            "paths": schema.ListAttribute{
              Description: "The paths the rule will be applied to.",
              Computed: true,
              CustomType: customfield.NewListType[types.String](ctx),
              ElementType: types.StringType,
            },
            "priority": schema.Float64Attribute{
              Computed: true,
            },
          },
        },
      },
      "ruleset": schema.SingleNestedAttribute{
        Computed: true,
        CustomType: customfield.NewNestedObjectType[WebAnalyticsSiteRulesetDataSourceModel](ctx),
        Attributes: map[string]schema.Attribute{
          "id": schema.StringAttribute{
            Description: "The Web Analytics ruleset identifier.",
            Computed: true,
          },
          "enabled": schema.BoolAttribute{
            Description: "Whether the ruleset is enabled.",
            Computed: true,
          },
          "zone_name": schema.StringAttribute{
            Computed: true,
          },
          "zone_tag": schema.StringAttribute{
            Description: "The zone identifier.",
            Computed: true,
          },
        },
      },
      "filter": schema.SingleNestedAttribute{
        Optional: true,
        Attributes: map[string]schema.Attribute{
          "order_by": schema.StringAttribute{
            Description: "The property used to sort the list of results.\nAvailable values: \"host\", \"created\".",
            Optional: true,
            Validators: []validator.String{
            stringvalidator.OneOfCaseInsensitive("host", "created"),
            },
          },
        },
      },
    },
  }
}

func (d *WebAnalyticsSiteDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
  resp.Schema = DataSourceSchema(ctx)
}

func (d *WebAnalyticsSiteDataSource) ConfigValidators(_ context.Context) ([]datasource.ConfigValidator) {
  return []datasource.ConfigValidator{
  datasourcevalidator.ExactlyOneOf(path.MatchRoot("site_id"), path.MatchRoot("filter")),
  }
}
