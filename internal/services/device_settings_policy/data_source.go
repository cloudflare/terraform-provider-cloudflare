// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package device_settings_policy

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

type DeviceSettingsPolicyDataSource struct {
	client *cloudflare.Client
}

var _ datasource.DataSourceWithConfigure = &DeviceSettingsPolicyDataSource{}

func NewDeviceSettingsPolicyDataSource() datasource.DataSource {
	return &DeviceSettingsPolicyDataSource{}
}

func (d *DeviceSettingsPolicyDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device_settings_policy"
}

func (r *DeviceSettingsPolicyDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*cloudflare.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"unexpected resource configure type",
			fmt.Sprintf("Expected *cloudflare.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *DeviceSettingsPolicyDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *DeviceSettingsPolicyDataSource

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
