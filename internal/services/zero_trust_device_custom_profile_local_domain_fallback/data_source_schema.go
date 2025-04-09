// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_device_custom_profile_local_domain_fallback

import (
  "context"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
  "github.com/hashicorp/terraform-plugin-framework/datasource"
  "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*ZeroTrustDeviceCustomProfileLocalDomainFallbackDataSource)(nil)

func DataSourceSchema(ctx context.Context) (schema.Schema) {
  return schema.Schema{
    Attributes: map[string]schema.Attribute{
      "account_id": schema.StringAttribute{
        Required: true,
      },
      "policy_id": schema.StringAttribute{
        Required: true,
      },
      "description": schema.StringAttribute{
        Description: "A description of the fallback domain, displayed in the client UI.",
        Computed: true,
      },
      "suffix": schema.StringAttribute{
        Description: "The domain suffix to match when resolving locally.",
        Computed: true,
      },
      "dns_server": schema.ListAttribute{
        Description: "A list of IP addresses to handle domain resolution.",
        Computed: true,
        CustomType: customfield.NewListType[types.String](ctx),
        ElementType: types.StringType,
      },
    },
  }
}

func (d *ZeroTrustDeviceCustomProfileLocalDomainFallbackDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
  resp.Schema = DataSourceSchema(ctx)
}

func (d *ZeroTrustDeviceCustomProfileLocalDomainFallbackDataSource) ConfigValidators(_ context.Context) ([]datasource.ConfigValidator) {
  return []datasource.ConfigValidator{
  }
}
