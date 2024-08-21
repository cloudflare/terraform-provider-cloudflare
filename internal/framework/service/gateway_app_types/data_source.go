package gateway_app_types

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/framework/muxclient"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSource = &CloudflareGatewayAppTypesDataSource{}

func NewDataSource() datasource.DataSource {
	return &CloudflareGatewayAppTypesDataSource{}
}

type CloudflareGatewayAppTypesDataSource struct {
	client *muxclient.Client
}

func (d *CloudflareGatewayAppTypesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gateway_app_types"
}

func (d *CloudflareGatewayAppTypesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*muxclient.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected resource configure type",
			fmt.Sprintf("Expected *muxclient.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	d.client = client
}

func (d *CloudflareGatewayAppTypesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data GatewayAppTypesModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	params := zero_trust.GatewayAppTypeListParams{
		AccountID: cloudflare.F(data.AccountID.ValueString()),
	}

	iter := d.client.V2.ZeroTrust.Gateway.AppTypes.ListAutoPaging(ctx, params)
	var appTypes []GatewayAppTypeModel

	for iter.Next() {
		appType := iter.Current()

		appTypes = append(appTypes, GatewayAppTypeModel{
			ID:                types.Int64Value(appType.ID),
			ApplicationTypeID: types.Int64Value(appType.ApplicationTypeID),
			Name:              types.StringValue(appType.Name),
			Description:       types.StringValue(appType.Description),
		})
	}
	if err := iter.Err(); err != nil {
		resp.Diagnostics.AddError("Failed to fetch Gateway App Types", err.Error())
		return
	}

	data.AppTypes = appTypes
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
