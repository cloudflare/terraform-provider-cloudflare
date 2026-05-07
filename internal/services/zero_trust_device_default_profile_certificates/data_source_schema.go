// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_device_default_profile_certificates

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/schemata"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = (*ZeroTrustDeviceDefaultProfileCertificatesDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: schemata.Description{
			Scopes: []string{
				"SSL and Certificates Read",
				"SSL and Certificates Write",
			},
		}.String(),
		Attributes: map[string]schema.Attribute{
			"zone_id": schema.StringAttribute{
				Optional: true,
			},
			"enabled": schema.BoolAttribute{
				Description: "The current status of the device policy certificate provisioning feature for WARP clients.",
				Computed:    true,
			},
		},
	}
}

func (d *ZeroTrustDeviceDefaultProfileCertificatesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *ZeroTrustDeviceDefaultProfileCertificatesDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
