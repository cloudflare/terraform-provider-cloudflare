// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_short_lived_certificate

import (
  "context"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
  "github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
  "github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
  "github.com/hashicorp/terraform-plugin-framework/datasource"
  "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
  "github.com/hashicorp/terraform-plugin-framework/path"
  "github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*ZeroTrustAccessShortLivedCertificatesDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) (schema.Schema) {
  return schema.Schema{
    Attributes: map[string]schema.Attribute{
      "account_id": schema.StringAttribute{
        Description: "The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.",
        Optional: true,
      },
      "zone_id": schema.StringAttribute{
        Description: "The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.",
        Optional: true,
      },
      "max_items": schema.Int64Attribute{
        Description: "Max items to fetch, default: 1000",
        Optional: true,
        Validators: []validator.Int64{
        int64validator.AtLeast(0),
        },
      },
      "result": schema.ListNestedAttribute{
        Description: "The items returned by the data source",
        Computed: true,
        CustomType: customfield.NewNestedObjectListType[ZeroTrustAccessShortLivedCertificatesResultDataSourceModel](ctx),
        NestedObject: schema.NestedAttributeObject{
          Attributes: map[string]schema.Attribute{
            "id": schema.StringAttribute{
              Description: "The ID of the CA.",
              Computed: true,
            },
            "aud": schema.StringAttribute{
              Description: "The Application Audience (AUD) tag. Identifies the application associated with the CA.",
              Computed: true,
            },
            "public_key": schema.StringAttribute{
              Description: "The public key to add to your SSH server configuration.",
              Computed: true,
            },
          },
        },
      },
    },
  }
}

func (d *ZeroTrustAccessShortLivedCertificatesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
  resp.Schema = ListDataSourceSchema(ctx)
}

func (d *ZeroTrustAccessShortLivedCertificatesDataSource) ConfigValidators(_ context.Context) ([]datasource.ConfigValidator) {
  return []datasource.ConfigValidator{
  datasourcevalidator.ExactlyOneOf(path.MatchRoot("account_id"), path.MatchRoot("zone_id")),
  }
}
