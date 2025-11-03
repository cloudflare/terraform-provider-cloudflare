// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sso_connector

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

type SSOConnectorDataSource struct {
	client *cloudflare.Client
}

var _ datasource.DataSourceWithConfigure = (*SSOConnectorDataSource)(nil)

func NewSSOConnectorDataSource() datasource.DataSource {
	return &SSOConnectorDataSource{}
}

func (d *SSOConnectorDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sso_connector"
}

func (d *SSOConnectorDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

	d.client = client
}

func (d *SSOConnectorDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	resp.Diagnostics.AddError(
		"SSO Connector not supported",
		"The SSO Connector data source is not currently supported as the SSO service is not available in the cloudflare-go SDK.",
	)
}
