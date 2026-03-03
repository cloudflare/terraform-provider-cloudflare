package connectivity_directory_service

import (
	"context"
	"fmt"

	cfv6 "github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/connectivity"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/framework/flatteners"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/framework/muxclient"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = &ConnectivityDirectoryServiceDataSource{}

func NewDataSource() datasource.DataSource {
	return &ConnectivityDirectoryServiceDataSource{}
}

type ConnectivityDirectoryServiceDataSource struct {
	client *muxclient.Client
}

func (d *ConnectivityDirectoryServiceDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_connectivity_directory_service"
}

func (d *ConnectivityDirectoryServiceDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*muxclient.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *muxclient.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	d.client = client
}

func (d *ConnectivityDirectoryServiceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data ConnectivityDirectoryServiceDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	result, err := d.client.V6.Connectivity.Directory.Services.Get(
		ctx,
		data.ServiceID.ValueString(),
		connectivity.DirectoryServiceGetParams{
			AccountID: cfv6.F(data.AccountID.ValueString()),
		},
	)
	if err != nil {
		resp.Diagnostics.AddError("Error reading connectivity directory service", err.Error())
		return
	}

	data.ID = flatteners.String(result.ServiceID)
	data.ServiceID = flatteners.String(result.ServiceID)
	data.Name = flatteners.String(result.Name)
	data.Type = flatteners.String(string(result.Type))
	data.HTTPPort = flatteners.Int64(result.HTTPPort)
	data.HTTPSPort = flatteners.Int64(result.HTTPSPort)
	data.CreatedAt = flattenTimestamp(result.CreatedAt)
	data.UpdatedAt = flattenTimestamp(result.UpdatedAt)
	data.Host = flattenGetResponseHost(result.Host)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
