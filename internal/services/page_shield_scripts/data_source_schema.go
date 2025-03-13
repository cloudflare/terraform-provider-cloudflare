// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package page_shield_scripts

import (
  "context"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
  "github.com/hashicorp/terraform-plugin-framework/datasource"
  "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
  "github.com/hashicorp/terraform-plugin-framework/schema/validator"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*PageShieldScriptsDataSource)(nil)

func DataSourceSchema(ctx context.Context) (schema.Schema) {
  return schema.Schema{
    Attributes: map[string]schema.Attribute{
      "script_id": schema.StringAttribute{
        Description: "Identifier",
        Required: true,
      },
      "zone_id": schema.StringAttribute{
        Description: "Identifier",
        Required: true,
      },
      "added_at": schema.StringAttribute{
        Computed: true,
        CustomType: timetypes.RFC3339Type{

        },
      },
      "cryptomining_score": schema.Int64Attribute{
        Description: "The cryptomining score of the JavaScript content.",
        Computed: true,
        Validators: []validator.Int64{
        int64validator.Between(1, 99),
        },
      },
      "dataflow_score": schema.Int64Attribute{
        Description: "The dataflow score of the JavaScript content.",
        Computed: true,
        Validators: []validator.Int64{
        int64validator.Between(1, 99),
        },
      },
      "domain_reported_malicious": schema.BoolAttribute{
        Computed: true,
      },
      "fetched_at": schema.StringAttribute{
        Description: "The timestamp of when the script was last fetched.",
        Computed: true,
      },
      "first_page_url": schema.StringAttribute{
        Computed: true,
      },
      "first_seen_at": schema.StringAttribute{
        Computed: true,
        CustomType: timetypes.RFC3339Type{

        },
      },
      "hash": schema.StringAttribute{
        Description: "The computed hash of the analyzed script.",
        Computed: true,
      },
      "host": schema.StringAttribute{
        Computed: true,
      },
      "id": schema.StringAttribute{
        Description: "Identifier",
        Computed: true,
      },
      "js_integrity_score": schema.Int64Attribute{
        Description: "The integrity score of the JavaScript content.",
        Computed: true,
        Validators: []validator.Int64{
        int64validator.Between(1, 99),
        },
      },
      "last_seen_at": schema.StringAttribute{
        Computed: true,
        CustomType: timetypes.RFC3339Type{

        },
      },
      "magecart_score": schema.Int64Attribute{
        Description: "The magecart score of the JavaScript content.",
        Computed: true,
        Validators: []validator.Int64{
        int64validator.Between(1, 99),
        },
      },
      "malware_score": schema.Int64Attribute{
        Description: "The malware score of the JavaScript content.",
        Computed: true,
        Validators: []validator.Int64{
        int64validator.Between(1, 99),
        },
      },
      "obfuscation_score": schema.Int64Attribute{
        Description: "The obfuscation score of the JavaScript content.",
        Computed: true,
        Validators: []validator.Int64{
        int64validator.Between(1, 99),
        },
      },
      "url": schema.StringAttribute{
        Computed: true,
      },
      "url_contains_cdn_cgi_path": schema.BoolAttribute{
        Computed: true,
      },
      "url_reported_malicious": schema.BoolAttribute{
        Computed: true,
      },
      "malicious_domain_categories": schema.ListAttribute{
        Computed: true,
        CustomType: customfield.NewListType[types.String](ctx),
        ElementType: types.StringType,
      },
      "malicious_url_categories": schema.ListAttribute{
        Computed: true,
        CustomType: customfield.NewListType[types.String](ctx),
        ElementType: types.StringType,
      },
      "page_urls": schema.ListAttribute{
        Computed: true,
        CustomType: customfield.NewListType[types.String](ctx),
        ElementType: types.StringType,
      },
      "versions": schema.ListNestedAttribute{
        Computed: true,
        CustomType: customfield.NewNestedObjectListType[PageShieldScriptsVersionsDataSourceModel](ctx),
        NestedObject: schema.NestedAttributeObject{
          Attributes: map[string]schema.Attribute{
            "cryptomining_score": schema.Int64Attribute{
              Description: "The cryptomining score of the JavaScript content.",
              Computed: true,
              Validators: []validator.Int64{
              int64validator.Between(1, 99),
              },
            },
            "dataflow_score": schema.Int64Attribute{
              Description: "The dataflow score of the JavaScript content.",
              Computed: true,
              Validators: []validator.Int64{
              int64validator.Between(1, 99),
              },
            },
            "fetched_at": schema.StringAttribute{
              Description: "The timestamp of when the script was last fetched.",
              Computed: true,
            },
            "hash": schema.StringAttribute{
              Description: "The computed hash of the analyzed script.",
              Computed: true,
            },
            "js_integrity_score": schema.Int64Attribute{
              Description: "The integrity score of the JavaScript content.",
              Computed: true,
              Validators: []validator.Int64{
              int64validator.Between(1, 99),
              },
            },
            "magecart_score": schema.Int64Attribute{
              Description: "The magecart score of the JavaScript content.",
              Computed: true,
              Validators: []validator.Int64{
              int64validator.Between(1, 99),
              },
            },
            "malware_score": schema.Int64Attribute{
              Description: "The malware score of the JavaScript content.",
              Computed: true,
              Validators: []validator.Int64{
              int64validator.Between(1, 99),
              },
            },
            "obfuscation_score": schema.Int64Attribute{
              Description: "The obfuscation score of the JavaScript content.",
              Computed: true,
              Validators: []validator.Int64{
              int64validator.Between(1, 99),
              },
            },
          },
        },
      },
    },
  }
}

func (d *PageShieldScriptsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
  resp.Schema = DataSourceSchema(ctx)
}

func (d *PageShieldScriptsDataSource) ConfigValidators(_ context.Context) ([]datasource.ConfigValidator) {
  return []datasource.ConfigValidator{
  }
}
