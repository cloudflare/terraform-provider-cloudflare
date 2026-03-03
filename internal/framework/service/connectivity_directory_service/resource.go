package connectivity_directory_service

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	cfv6 "github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/connectivity"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/framework/flatteners"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/framework/muxclient"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &ConnectivityDirectoryServiceResource{}
var _ resource.ResourceWithImportState = &ConnectivityDirectoryServiceResource{}

func NewResource() resource.Resource {
	return &ConnectivityDirectoryServiceResource{}
}

type ConnectivityDirectoryServiceResource struct {
	client *muxclient.Client
}

func (r *ConnectivityDirectoryServiceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_connectivity_directory_service"
}

func (r *ConnectivityDirectoryServiceResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*muxclient.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *muxclient.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	r.client = client
}

func (r *ConnectivityDirectoryServiceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ConnectivityDirectoryServiceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	host, err := expandNewHost(ctx, data.Host)
	if err != nil {
		resp.Diagnostics.AddError("Error building host params", err.Error())
		return
	}

	params := connectivity.DirectoryServiceNewParams{
		AccountID: cfv6.F(data.AccountID.ValueString()),
		Name:      cfv6.F(data.Name.ValueString()),
		Type:      cfv6.F(connectivity.DirectoryServiceNewParamsType(data.Type.ValueString())),
		Host:      cfv6.F(host),
	}

	if !data.HTTPPort.IsNull() && !data.HTTPPort.IsUnknown() {
		params.HTTPPort = cfv6.F(data.HTTPPort.ValueInt64())
	}
	if !data.HTTPSPort.IsNull() && !data.HTTPSPort.IsUnknown() {
		params.HTTPSPort = cfv6.F(data.HTTPSPort.ValueInt64())
	}

	tflog.Debug(ctx, fmt.Sprintf("Creating connectivity directory service: name=%s", data.Name.ValueString()))

	result, err := r.client.V6.Connectivity.Directory.Services.New(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError("Error creating connectivity directory service", err.Error())
		return
	}

	data = flattenNewResponse(data, result)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ConnectivityDirectoryServiceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ConnectivityDirectoryServiceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	serviceID := resolveServiceID(data)

	result, err := r.client.V6.Connectivity.Directory.Services.Get(
		ctx,
		serviceID,
		connectivity.DirectoryServiceGetParams{
			AccountID: cfv6.F(data.AccountID.ValueString()),
		},
	)
	if err != nil {
		if isNotFound(err) {
			tflog.Warn(ctx, fmt.Sprintf("Connectivity directory service %s not found, removing from state", serviceID))
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Error reading connectivity directory service", err.Error())
		return
	}

	data = flattenGetResponse(data, result)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ConnectivityDirectoryServiceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data ConnectivityDirectoryServiceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	host, err := expandUpdateHost(ctx, data.Host)
	if err != nil {
		resp.Diagnostics.AddError("Error building host params", err.Error())
		return
	}

	serviceID := resolveServiceID(data)

	params := connectivity.DirectoryServiceUpdateParams{
		AccountID: cfv6.F(data.AccountID.ValueString()),
		Name:      cfv6.F(data.Name.ValueString()),
		Type:      cfv6.F(connectivity.DirectoryServiceUpdateParamsType(data.Type.ValueString())),
		Host:      cfv6.F(host),
	}

	if !data.HTTPPort.IsNull() && !data.HTTPPort.IsUnknown() {
		params.HTTPPort = cfv6.F(data.HTTPPort.ValueInt64())
	}
	if !data.HTTPSPort.IsNull() && !data.HTTPSPort.IsUnknown() {
		params.HTTPSPort = cfv6.F(data.HTTPSPort.ValueInt64())
	}

	tflog.Debug(ctx, fmt.Sprintf("Updating connectivity directory service: id=%s", serviceID))

	result, err := r.client.V6.Connectivity.Directory.Services.Update(ctx, serviceID, params)
	if err != nil {
		resp.Diagnostics.AddError("Error updating connectivity directory service", err.Error())
		return
	}

	data = flattenUpdateResponse(data, result)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ConnectivityDirectoryServiceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ConnectivityDirectoryServiceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	serviceID := resolveServiceID(data)

	tflog.Debug(ctx, fmt.Sprintf("Deleting connectivity directory service: id=%s", serviceID))

	err := r.client.V6.Connectivity.Directory.Services.Delete(
		ctx,
		serviceID,
		connectivity.DirectoryServiceDeleteParams{
			AccountID: cfv6.F(data.AccountID.ValueString()),
		},
	)
	if err != nil {
		// If the resource is already gone, the delete is a success.
		if isNotFound(err) {
			return
		}
		resp.Diagnostics.AddError("Error deleting connectivity directory service", err.Error())
	}
}

func (r *ConnectivityDirectoryServiceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, "/")
	if len(idParts) != 2 {
		resp.Diagnostics.AddError(
			"Error importing connectivity directory service",
			"Invalid ID format. Please specify the ID as \"account_id/service_id\".",
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("account_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), idParts[1])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("service_id"), idParts[1])...)
}

// resolveServiceID returns the service ID from the model, preferring ServiceID over ID.
func resolveServiceID(data ConnectivityDirectoryServiceModel) string {
	if v := data.ServiceID.ValueString(); v != "" {
		return v
	}
	return data.ID.ValueString()
}

// isNotFound checks whether an error from the v6 SDK represents a 404 response.
func isNotFound(err error) bool {
	var apiErr *connectivity.Error
	if errors.As(err, &apiErr) {
		return apiErr.StatusCode == http.StatusNotFound
	}
	return false
}

// flattenTimestamp formats a time.Time as RFC3339 or returns null for zero values.
func flattenTimestamp(t time.Time) types.String {
	if t.IsZero() {
		return types.StringNull()
	}
	return types.StringValue(t.Format(time.RFC3339))
}

// expandNewHost converts the Terraform model host into the v6 SDK create params host union.
func expandNewHost(_ context.Context, host *ConnectivityDirectoryServiceHostModel) (connectivity.DirectoryServiceNewParamsHostUnion, error) {
	if host == nil {
		return nil, fmt.Errorf("host block is required")
	}

	hasIPv4 := !host.IPV4.IsNull() && !host.IPV4.IsUnknown() && host.IPV4.ValueString() != ""
	hasIPv6 := !host.IPV6.IsNull() && !host.IPV6.IsUnknown() && host.IPV6.ValueString() != ""
	hasHostname := !host.Hostname.IsNull() && !host.Hostname.IsUnknown() && host.Hostname.ValueString() != ""

	switch {
	case hasHostname:
		h := connectivity.DirectoryServiceNewParamsHostInfraHostnameHost{
			Hostname: cfv6.F(host.Hostname.ValueString()),
		}
		if host.ResolverNetwork != nil {
			rn := connectivity.DirectoryServiceNewParamsHostInfraHostnameHostResolverNetwork{
				TunnelID: cfv6.F(host.ResolverNetwork.TunnelID.ValueString()),
			}
			if !host.ResolverNetwork.ResolverIPs.IsNull() && !host.ResolverNetwork.ResolverIPs.IsUnknown() {
				var ips []string
				for _, elem := range host.ResolverNetwork.ResolverIPs.Elements() {
					if sv, ok := elem.(types.String); ok {
						ips = append(ips, sv.ValueString())
					}
				}
				if len(ips) > 0 {
					rn.ResolverIPs = cfv6.F(ips)
				}
			}
			h.ResolverNetwork = cfv6.F(rn)
		}
		return h, nil

	case hasIPv4 && hasIPv6:
		h := connectivity.DirectoryServiceNewParamsHostInfraDualStackHost{
			IPV4: cfv6.F(host.IPV4.ValueString()),
			IPV6: cfv6.F(host.IPV6.ValueString()),
		}
		if host.Network != nil {
			h.Network = cfv6.F(connectivity.DirectoryServiceNewParamsHostInfraDualStackHostNetwork{
				TunnelID: cfv6.F(host.Network.TunnelID.ValueString()),
			})
		}
		return h, nil

	case hasIPv4:
		h := connectivity.DirectoryServiceNewParamsHostInfraIPv4Host{
			IPV4: cfv6.F(host.IPV4.ValueString()),
		}
		if host.Network != nil {
			h.Network = cfv6.F(connectivity.DirectoryServiceNewParamsHostInfraIPv4HostNetwork{
				TunnelID: cfv6.F(host.Network.TunnelID.ValueString()),
			})
		}
		return h, nil

	case hasIPv6:
		h := connectivity.DirectoryServiceNewParamsHostInfraIPv6Host{
			IPV6: cfv6.F(host.IPV6.ValueString()),
		}
		if host.Network != nil {
			h.Network = cfv6.F(connectivity.DirectoryServiceNewParamsHostInfraIPv6HostNetwork{
				TunnelID: cfv6.F(host.Network.TunnelID.ValueString()),
			})
		}
		return h, nil

	default:
		return nil, fmt.Errorf("host must specify at least one of: hostname, ipv4, ipv6")
	}
}

// expandUpdateHost converts the Terraform model host into the v6 SDK update params host union.
func expandUpdateHost(_ context.Context, host *ConnectivityDirectoryServiceHostModel) (connectivity.DirectoryServiceUpdateParamsHostUnion, error) {
	if host == nil {
		return nil, fmt.Errorf("host block is required")
	}

	hasIPv4 := !host.IPV4.IsNull() && !host.IPV4.IsUnknown() && host.IPV4.ValueString() != ""
	hasIPv6 := !host.IPV6.IsNull() && !host.IPV6.IsUnknown() && host.IPV6.ValueString() != ""
	hasHostname := !host.Hostname.IsNull() && !host.Hostname.IsUnknown() && host.Hostname.ValueString() != ""

	switch {
	case hasHostname:
		h := connectivity.DirectoryServiceUpdateParamsHostInfraHostnameHost{
			Hostname: cfv6.F(host.Hostname.ValueString()),
		}
		if host.ResolverNetwork != nil {
			rn := connectivity.DirectoryServiceUpdateParamsHostInfraHostnameHostResolverNetwork{
				TunnelID: cfv6.F(host.ResolverNetwork.TunnelID.ValueString()),
			}
			if !host.ResolverNetwork.ResolverIPs.IsNull() && !host.ResolverNetwork.ResolverIPs.IsUnknown() {
				var ips []string
				for _, elem := range host.ResolverNetwork.ResolverIPs.Elements() {
					if sv, ok := elem.(types.String); ok {
						ips = append(ips, sv.ValueString())
					}
				}
				if len(ips) > 0 {
					rn.ResolverIPs = cfv6.F(ips)
				}
			}
			h.ResolverNetwork = cfv6.F(rn)
		}
		return h, nil

	case hasIPv4 && hasIPv6:
		h := connectivity.DirectoryServiceUpdateParamsHostInfraDualStackHost{
			IPV4: cfv6.F(host.IPV4.ValueString()),
			IPV6: cfv6.F(host.IPV6.ValueString()),
		}
		if host.Network != nil {
			h.Network = cfv6.F(connectivity.DirectoryServiceUpdateParamsHostInfraDualStackHostNetwork{
				TunnelID: cfv6.F(host.Network.TunnelID.ValueString()),
			})
		}
		return h, nil

	case hasIPv4:
		h := connectivity.DirectoryServiceUpdateParamsHostInfraIPv4Host{
			IPV4: cfv6.F(host.IPV4.ValueString()),
		}
		if host.Network != nil {
			h.Network = cfv6.F(connectivity.DirectoryServiceUpdateParamsHostInfraIPv4HostNetwork{
				TunnelID: cfv6.F(host.Network.TunnelID.ValueString()),
			})
		}
		return h, nil

	case hasIPv6:
		h := connectivity.DirectoryServiceUpdateParamsHostInfraIPv6Host{
			IPV6: cfv6.F(host.IPV6.ValueString()),
		}
		if host.Network != nil {
			h.Network = cfv6.F(connectivity.DirectoryServiceUpdateParamsHostInfraIPv6HostNetwork{
				TunnelID: cfv6.F(host.Network.TunnelID.ValueString()),
			})
		}
		return h, nil

	default:
		return nil, fmt.Errorf("host must specify at least one of: hostname, ipv4, ipv6")
	}
}

// flattenNewResponse maps a DirectoryServiceNewResponse to the Terraform model.
func flattenNewResponse(data ConnectivityDirectoryServiceModel, resp *connectivity.DirectoryServiceNewResponse) ConnectivityDirectoryServiceModel {
	data.ID = flatteners.String(resp.ServiceID)
	data.ServiceID = flatteners.String(resp.ServiceID)
	data.Name = flatteners.String(resp.Name)
	data.Type = flatteners.String(string(resp.Type))
	data.HTTPPort = flatteners.Int64(resp.HTTPPort)
	data.HTTPSPort = flatteners.Int64(resp.HTTPSPort)
	data.CreatedAt = flattenTimestamp(resp.CreatedAt)
	data.UpdatedAt = flattenTimestamp(resp.UpdatedAt)
	data.Host = flattenNewResponseHost(resp.Host)
	return data
}

// flattenUpdateResponse maps a DirectoryServiceUpdateResponse to the Terraform model.
func flattenUpdateResponse(data ConnectivityDirectoryServiceModel, resp *connectivity.DirectoryServiceUpdateResponse) ConnectivityDirectoryServiceModel {
	data.ID = flatteners.String(resp.ServiceID)
	data.ServiceID = flatteners.String(resp.ServiceID)
	data.Name = flatteners.String(resp.Name)
	data.Type = flatteners.String(string(resp.Type))
	data.HTTPPort = flatteners.Int64(resp.HTTPPort)
	data.HTTPSPort = flatteners.Int64(resp.HTTPSPort)
	data.CreatedAt = flattenTimestamp(resp.CreatedAt)
	data.UpdatedAt = flattenTimestamp(resp.UpdatedAt)
	data.Host = flattenUpdateResponseHost(resp.Host)
	return data
}

// flattenGetResponse maps a DirectoryServiceGetResponse to the Terraform model.
func flattenGetResponse(data ConnectivityDirectoryServiceModel, resp *connectivity.DirectoryServiceGetResponse) ConnectivityDirectoryServiceModel {
	data.ID = flatteners.String(resp.ServiceID)
	data.ServiceID = flatteners.String(resp.ServiceID)
	data.Name = flatteners.String(resp.Name)
	data.Type = flatteners.String(string(resp.Type))
	data.HTTPPort = flatteners.Int64(resp.HTTPPort)
	data.HTTPSPort = flatteners.Int64(resp.HTTPSPort)
	data.CreatedAt = flattenTimestamp(resp.CreatedAt)
	data.UpdatedAt = flattenTimestamp(resp.UpdatedAt)
	data.Host = flattenGetResponseHost(resp.Host)
	return data
}

// flattenNetworkIfPresent returns a network model only if the tunnel ID is non-empty.
func flattenNetworkIfPresent(tunnelID string) *ConnectivityDirectoryServiceHostNetworkModel {
	if tunnelID == "" {
		return nil
	}
	return &ConnectivityDirectoryServiceHostNetworkModel{
		TunnelID: flatteners.String(tunnelID),
	}
}

// flattenNewResponseHost uses the union type switch to flatten the response host.
func flattenNewResponseHost(host connectivity.DirectoryServiceNewResponseHost) *ConnectivityDirectoryServiceHostModel {
	m := &ConnectivityDirectoryServiceHostModel{}
	switch h := host.AsUnion().(type) {
	case connectivity.DirectoryServiceNewResponseHostInfraIPv4Host:
		m.IPV4 = flatteners.String(h.IPV4)
		m.Network = flattenNetworkIfPresent(h.Network.TunnelID)
	case connectivity.DirectoryServiceNewResponseHostInfraIPv6Host:
		m.IPV6 = flatteners.String(h.IPV6)
		m.Network = flattenNetworkIfPresent(h.Network.TunnelID)
	case connectivity.DirectoryServiceNewResponseHostInfraDualStackHost:
		m.IPV4 = flatteners.String(h.IPV4)
		m.IPV6 = flatteners.String(h.IPV6)
		m.Network = flattenNetworkIfPresent(h.Network.TunnelID)
	case connectivity.DirectoryServiceNewResponseHostInfraHostnameHost:
		m.Hostname = flatteners.String(h.Hostname)
		m.ResolverNetwork = flattenResolverNetwork(h.ResolverNetwork.TunnelID, h.ResolverNetwork.ResolverIPs)
	default:
		m.IPV4 = flatteners.String(host.IPV4)
		m.IPV6 = flatteners.String(host.IPV6)
		m.Hostname = flatteners.String(host.Hostname)
	}
	return m
}

// flattenUpdateResponseHost uses the union type switch to flatten the response host.
func flattenUpdateResponseHost(host connectivity.DirectoryServiceUpdateResponseHost) *ConnectivityDirectoryServiceHostModel {
	m := &ConnectivityDirectoryServiceHostModel{}
	switch h := host.AsUnion().(type) {
	case connectivity.DirectoryServiceUpdateResponseHostInfraIPv4Host:
		m.IPV4 = flatteners.String(h.IPV4)
		m.Network = flattenNetworkIfPresent(h.Network.TunnelID)
	case connectivity.DirectoryServiceUpdateResponseHostInfraIPv6Host:
		m.IPV6 = flatteners.String(h.IPV6)
		m.Network = flattenNetworkIfPresent(h.Network.TunnelID)
	case connectivity.DirectoryServiceUpdateResponseHostInfraDualStackHost:
		m.IPV4 = flatteners.String(h.IPV4)
		m.IPV6 = flatteners.String(h.IPV6)
		m.Network = flattenNetworkIfPresent(h.Network.TunnelID)
	case connectivity.DirectoryServiceUpdateResponseHostInfraHostnameHost:
		m.Hostname = flatteners.String(h.Hostname)
		m.ResolverNetwork = flattenResolverNetwork(h.ResolverNetwork.TunnelID, h.ResolverNetwork.ResolverIPs)
	default:
		m.IPV4 = flatteners.String(host.IPV4)
		m.IPV6 = flatteners.String(host.IPV6)
		m.Hostname = flatteners.String(host.Hostname)
	}
	return m
}

// flattenGetResponseHost uses the union type switch to flatten the response host.
func flattenGetResponseHost(host connectivity.DirectoryServiceGetResponseHost) *ConnectivityDirectoryServiceHostModel {
	m := &ConnectivityDirectoryServiceHostModel{}
	switch h := host.AsUnion().(type) {
	case connectivity.DirectoryServiceGetResponseHostInfraIPv4Host:
		m.IPV4 = flatteners.String(h.IPV4)
		m.Network = flattenNetworkIfPresent(h.Network.TunnelID)
	case connectivity.DirectoryServiceGetResponseHostInfraIPv6Host:
		m.IPV6 = flatteners.String(h.IPV6)
		m.Network = flattenNetworkIfPresent(h.Network.TunnelID)
	case connectivity.DirectoryServiceGetResponseHostInfraDualStackHost:
		m.IPV4 = flatteners.String(h.IPV4)
		m.IPV6 = flatteners.String(h.IPV6)
		m.Network = flattenNetworkIfPresent(h.Network.TunnelID)
	case connectivity.DirectoryServiceGetResponseHostInfraHostnameHost:
		m.Hostname = flatteners.String(h.Hostname)
		m.ResolverNetwork = flattenResolverNetwork(h.ResolverNetwork.TunnelID, h.ResolverNetwork.ResolverIPs)
	default:
		m.IPV4 = flatteners.String(host.IPV4)
		m.IPV6 = flatteners.String(host.IPV6)
		m.Hostname = flatteners.String(host.Hostname)
	}
	return m
}

// flattenResolverNetwork builds a resolver network model, or nil if the tunnel ID is empty.
func flattenResolverNetwork(tunnelID string, resolverIPs []string) *ConnectivityDirectoryServiceHostResolverNetworkModel {
	if tunnelID == "" {
		return nil
	}
	return &ConnectivityDirectoryServiceHostResolverNetworkModel{
		TunnelID:    flatteners.String(tunnelID),
		ResolverIPs: flattenStringSliceToList(context.Background(), resolverIPs),
	}
}

// flattenStringSliceToList converts a []string to a types.List of StringType.
func flattenStringSliceToList(ctx context.Context, in []string) types.List {
	if len(in) == 0 {
		return types.ListNull(types.StringType)
	}
	list, diags := types.ListValueFrom(ctx, types.StringType, in)
	if diags.HasError() {
		tflog.Error(ctx, fmt.Sprintf("Error converting string slice to list: %s", diags.Errors()))
		return types.ListNull(types.StringType)
	}
	return list
}
