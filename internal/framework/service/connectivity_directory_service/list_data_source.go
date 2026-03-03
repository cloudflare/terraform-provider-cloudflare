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

var _ datasource.DataSource = &ConnectivityDirectoryServicesDataSource{}

func NewListDataSource() datasource.DataSource {
	return &ConnectivityDirectoryServicesDataSource{}
}

type ConnectivityDirectoryServicesDataSource struct {
	client *muxclient.Client
}

func (d *ConnectivityDirectoryServicesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_connectivity_directory_services"
}

func (d *ConnectivityDirectoryServicesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *ConnectivityDirectoryServicesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data ConnectivityDirectoryServicesDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	params := connectivity.DirectoryServiceListParams{
		AccountID: cfv6.F(data.AccountID.ValueString()),
	}

	if !data.Type.IsNull() && !data.Type.IsUnknown() {
		params.Type = cfv6.F(connectivity.DirectoryServiceListParamsType(data.Type.ValueString()))
	}

	iter := d.client.V6.Connectivity.Directory.Services.ListAutoPaging(ctx, params)

	services := make([]ConnectivityDirectoryServiceDataSourceResultModel, 0)
	for iter.Next() {
		svc := iter.Current()
		item := ConnectivityDirectoryServiceDataSourceResultModel{
			ID:        flatteners.String(svc.ServiceID),
			ServiceID: flatteners.String(svc.ServiceID),
			Name:      flatteners.String(svc.Name),
			Type:      flatteners.String(string(svc.Type)),
			HTTPPort:  flatteners.Int64(svc.HTTPPort),
			HTTPSPort: flatteners.Int64(svc.HTTPSPort),
			CreatedAt: flattenTimestamp(svc.CreatedAt),
			UpdatedAt: flattenTimestamp(svc.UpdatedAt),
			Host:      flattenListResponseHost(ctx, svc.Host),
		}
		services = append(services, item)
	}
	if err := iter.Err(); err != nil {
		resp.Diagnostics.AddError("Error listing connectivity directory services", err.Error())
		return
	}

	data.ID = data.AccountID
	data.Services = services
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// flattenListResponseHost uses the union type switch to flatten a list response host.
func flattenListResponseHost(_ context.Context, host connectivity.DirectoryServiceListResponseHost) *ConnectivityDirectoryServiceHostModel {
	m := &ConnectivityDirectoryServiceHostModel{}
	switch h := host.AsUnion().(type) {
	case connectivity.DirectoryServiceListResponseHostInfraIPv4Host:
		m.IPV4 = flatteners.String(h.IPV4)
		m.Network = flattenNetworkIfPresent(h.Network.TunnelID)
	case connectivity.DirectoryServiceListResponseHostInfraIPv6Host:
		m.IPV6 = flatteners.String(h.IPV6)
		m.Network = flattenNetworkIfPresent(h.Network.TunnelID)
	case connectivity.DirectoryServiceListResponseHostInfraDualStackHost:
		m.IPV4 = flatteners.String(h.IPV4)
		m.IPV6 = flatteners.String(h.IPV6)
		m.Network = flattenNetworkIfPresent(h.Network.TunnelID)
	case connectivity.DirectoryServiceListResponseHostInfraHostnameHost:
		m.Hostname = flatteners.String(h.Hostname)
		m.ResolverNetwork = flattenResolverNetwork(h.ResolverNetwork.TunnelID, h.ResolverNetwork.ResolverIPs)
	default:
		m.IPV4 = flatteners.String(host.IPV4)
		m.IPV6 = flatteners.String(host.IPV6)
		m.Hostname = flatteners.String(host.Hostname)
	}
	return m
}
