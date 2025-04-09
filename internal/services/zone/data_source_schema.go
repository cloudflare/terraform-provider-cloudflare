// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone

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

var _ datasource.DataSourceWithConfigValidators = (*ZoneDataSource)(nil)

func DataSourceSchema(ctx context.Context) (schema.Schema) {
  return schema.Schema{
    Attributes: map[string]schema.Attribute{
      "id": schema.StringAttribute{
        Description: "Identifier",
        Computed: true,
      },
      "zone_id": schema.StringAttribute{
        Description: "Identifier",
        Optional: true,
      },
      "activated_on": schema.StringAttribute{
        Description: "The last time proof of ownership was detected and the zone was made\nactive",
        Computed: true,
        CustomType: timetypes.RFC3339Type{

        },
      },
      "created_on": schema.StringAttribute{
        Description: "When the zone was created",
        Computed: true,
        CustomType: timetypes.RFC3339Type{

        },
      },
      "development_mode": schema.Float64Attribute{
        Description: "The interval (in seconds) from when development mode expires\n(positive integer) or last expired (negative integer) for the\ndomain. If development mode has never been enabled, this value is 0.",
        Computed: true,
      },
      "modified_on": schema.StringAttribute{
        Description: "When the zone was last modified",
        Computed: true,
        CustomType: timetypes.RFC3339Type{

        },
      },
      "name": schema.StringAttribute{
        Description: "The domain name",
        Computed: true,
      },
      "original_dnshost": schema.StringAttribute{
        Description: "DNS host at the time of switching to Cloudflare",
        Computed: true,
      },
      "original_registrar": schema.StringAttribute{
        Description: "Registrar for the domain at the time of switching to Cloudflare",
        Computed: true,
      },
      "paused": schema.BoolAttribute{
        Description: "Indicates whether the zone is only using Cloudflare DNS services. A\ntrue value means the zone will not receive security or performance\nbenefits.",
        Computed: true,
      },
      "status": schema.StringAttribute{
        Description: "The zone status on Cloudflare.\nAvailable values: \"initializing\", \"pending\", \"active\", \"moved\".",
        Computed: true,
        Validators: []validator.String{
        stringvalidator.OneOfCaseInsensitive(
          "initializing",
          "pending",
          "active",
          "moved",
        ),
        },
      },
      "type": schema.StringAttribute{
        Description: "A full zone implies that DNS is hosted with Cloudflare. A partial zone is\ntypically a partner-hosted zone or a CNAME setup.\nAvailable values: \"full\", \"partial\", \"secondary\", \"internal\".",
        Computed: true,
        Validators: []validator.String{
        stringvalidator.OneOfCaseInsensitive(
          "full",
          "partial",
          "secondary",
          "internal",
        ),
        },
      },
      "verification_key": schema.StringAttribute{
        Description: "Verification key for partial zone setup.",
        Computed: true,
      },
      "name_servers": schema.ListAttribute{
        Description: "The name servers Cloudflare assigns to a zone",
        Computed: true,
        CustomType: customfield.NewListType[types.String](ctx),
        ElementType: types.StringType,
      },
      "original_name_servers": schema.ListAttribute{
        Description: "Original name servers before moving to Cloudflare",
        Computed: true,
        CustomType: customfield.NewListType[types.String](ctx),
        ElementType: types.StringType,
      },
      "vanity_name_servers": schema.ListAttribute{
        Description: "An array of domains used for custom name servers. This is only available for Business and Enterprise plans.",
        Computed: true,
        CustomType: customfield.NewListType[types.String](ctx),
        ElementType: types.StringType,
      },
      "account": schema.SingleNestedAttribute{
        Description: "The account the zone belongs to",
        Computed: true,
        CustomType: customfield.NewNestedObjectType[ZoneAccountDataSourceModel](ctx),
        Attributes: map[string]schema.Attribute{
          "id": schema.StringAttribute{
            Description: "Identifier",
            Computed: true,
          },
          "name": schema.StringAttribute{
            Description: "The name of the account",
            Computed: true,
          },
        },
      },
      "meta": schema.SingleNestedAttribute{
        Description: "Metadata about the zone",
        Computed: true,
        CustomType: customfield.NewNestedObjectType[ZoneMetaDataSourceModel](ctx),
        Attributes: map[string]schema.Attribute{
          "cdn_only": schema.BoolAttribute{
            Description: "The zone is only configured for CDN",
            Computed: true,
          },
          "custom_certificate_quota": schema.Int64Attribute{
            Description: "Number of Custom Certificates the zone can have",
            Computed: true,
          },
          "dns_only": schema.BoolAttribute{
            Description: "The zone is only configured for DNS",
            Computed: true,
          },
          "foundation_dns": schema.BoolAttribute{
            Description: "The zone is setup with Foundation DNS",
            Computed: true,
          },
          "page_rule_quota": schema.Int64Attribute{
            Description: "Number of Page Rules a zone can have",
            Computed: true,
          },
          "phishing_detected": schema.BoolAttribute{
            Description: "The zone has been flagged for phishing",
            Computed: true,
          },
          "step": schema.Int64Attribute{
            Computed: true,
          },
        },
      },
      "owner": schema.SingleNestedAttribute{
        Description: "The owner of the zone",
        Computed: true,
        CustomType: customfield.NewNestedObjectType[ZoneOwnerDataSourceModel](ctx),
        Attributes: map[string]schema.Attribute{
          "id": schema.StringAttribute{
            Description: "Identifier",
            Computed: true,
          },
          "name": schema.StringAttribute{
            Description: "Name of the owner",
            Computed: true,
          },
          "type": schema.StringAttribute{
            Description: "The type of owner",
            Computed: true,
          },
        },
      },
      "filter": schema.SingleNestedAttribute{
        Optional: true,
        Attributes: map[string]schema.Attribute{
          "account": schema.SingleNestedAttribute{
            Optional: true,
            Attributes: map[string]schema.Attribute{
              "id": schema.StringAttribute{
                Description: "An account ID",
                Optional: true,
              },
              "name": schema.StringAttribute{
                Description: "An account Name. Optional filter operators can be provided to extend refine the search:\n  * `equal` (default)\n  * `not_equal`\n  * `starts_with`\n  * `ends_with`\n  * `contains`\n  * `starts_with_case_sensitive`\n  * `ends_with_case_sensitive`\n  * `contains_case_sensitive`",
                Optional: true,
              },
            },
          },
          "direction": schema.StringAttribute{
            Description: "Direction to order zones.\nAvailable values: \"asc\", \"desc\".",
            Optional: true,
            Validators: []validator.String{
            stringvalidator.OneOfCaseInsensitive("asc", "desc"),
            },
          },
          "match": schema.StringAttribute{
            Description: "Whether to match all search requirements or at least one (any).\nAvailable values: \"any\", \"all\".",
            Computed: true,
            Optional: true,
            Validators: []validator.String{
            stringvalidator.OneOfCaseInsensitive("any", "all"),
            },
          },
          "name": schema.StringAttribute{
            Description: "A domain name. Optional filter operators can be provided to extend refine the search:\n  * `equal` (default)\n  * `not_equal`\n  * `starts_with`\n  * `ends_with`\n  * `contains`\n  * `starts_with_case_sensitive`\n  * `ends_with_case_sensitive`\n  * `contains_case_sensitive`",
            Optional: true,
          },
          "order": schema.StringAttribute{
            Description: "Field to order zones by.\nAvailable values: \"name\", \"status\", \"account.id\", \"account.name\".",
            Optional: true,
            Validators: []validator.String{
            stringvalidator.OneOfCaseInsensitive(
              "name",
              "status",
              "account.id",
              "account.name",
            ),
            },
          },
          "status": schema.StringAttribute{
            Description: "A zone status\nAvailable values: \"initializing\", \"pending\", \"active\", \"moved\".",
            Optional: true,
            Validators: []validator.String{
            stringvalidator.OneOfCaseInsensitive(
              "initializing",
              "pending",
              "active",
              "moved",
            ),
            },
          },
        },
      },
    },
  }
}

func (d *ZoneDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
  resp.Schema = DataSourceSchema(ctx)
}

func (d *ZoneDataSource) ConfigValidators(_ context.Context) ([]datasource.ConfigValidator) {
  return []datasource.ConfigValidator{
  datasourcevalidator.ExactlyOneOf(path.MatchRoot("zone_id"), path.MatchRoot("filter")),
  }
}
