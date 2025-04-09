// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_short_lived_certificate

import (
  "context"

  "github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
  "github.com/hashicorp/terraform-plugin-framework/datasource"
  "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
  "github.com/hashicorp/terraform-plugin-framework/path"
)

var _ datasource.DataSourceWithConfigValidators = (*ZeroTrustAccessShortLivedCertificateDataSource)(nil)

func DataSourceSchema(ctx context.Context) (schema.Schema) {
  return schema.Schema{
    Attributes: map[string]schema.Attribute{
      "app_id": schema.StringAttribute{
        Description: "UUID.",
        Required: true,
      },
      "account_id": schema.StringAttribute{
        Description: "The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.",
        Optional: true,
      },
      "zone_id": schema.StringAttribute{
        Description: "The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.",
        Optional: true,
      },
      "aud": schema.StringAttribute{
        Description: "The Application Audience (AUD) tag. Identifies the application associated with the CA.",
        Computed: true,
      },
      "id": schema.StringAttribute{
        Description: "The ID of the CA.",
        Computed: true,
      },
      "public_key": schema.StringAttribute{
        Description: "The public key to add to your SSH server configuration.",
        Computed: true,
      },
    },
  }
}

func (d *ZeroTrustAccessShortLivedCertificateDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
  resp.Schema = DataSourceSchema(ctx)
}

func (d *ZeroTrustAccessShortLivedCertificateDataSource) ConfigValidators(_ context.Context) ([]datasource.ConfigValidator) {
  return []datasource.ConfigValidator{
  datasourcevalidator.Conflicting(path.MatchRoot("account_id"), path.MatchRoot("zone_id")),
  }
}
