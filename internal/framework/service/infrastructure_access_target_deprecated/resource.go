package infrastructure_access_target_deprecated

import (
	"context"
	"errors"
	"fmt"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/framework/flatteners"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/framework/muxclient"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	tftypes "github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &InfrastructureAccessTargetDeprecatedResource{}

func NewResource() resource.Resource {
	return &InfrastructureAccessTargetDeprecatedResource{}
}

// InfrastructureAccessTargetDeprecatedResource defines the resource implementation.
type InfrastructureAccessTargetDeprecatedResource struct {
	client *muxclient.Client
}

func (r *InfrastructureAccessTargetDeprecatedResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_infrastructure_access_target"
}

func (r *InfrastructureAccessTargetDeprecatedResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*muxclient.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"unexpected resource configure type",
			fmt.Sprintf("Expected *muxclient.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	r.client = client
}

func (r *InfrastructureAccessTargetDeprecatedResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *InfrastructureAccessTargetDeprecatedModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	accountId := data.AccountID.ValueString()
	if accountId == "" {
		resp.Diagnostics.AddError("failed to create infrastructure access target", "account id cannot be an empty string")
		return
	}
	ipInfo, err := buildCreateIPInfoFromDetails(ctx, data.IP, resp)
	if err != nil {
		resp.Diagnostics.AddError("failed to create infrastructure access target", "account id cannot be an empty string")
		return
	}
	createTargetParams := cloudflare.CreateInfrastructureAccessTargetParams{
		InfrastructureAccessTargetParams: cloudflare.InfrastructureAccessTargetParams{
			Hostname: data.Hostname.ValueString(),
			IP:       ipInfo,
		},
	}

	tflog.Debug(ctx, fmt.Sprintf("Creating Cloudflare Infrastructure Access Target from struct %+v", createTargetParams))
	target, err := r.client.V1.CreateInfrastructureAccessTarget(ctx, cloudflare.AccountIdentifier(accountId), createTargetParams)
	if err != nil {
		resp.Diagnostics.AddError(fmt.Sprintf("error creating Infrastructure Access Target for account %q", accountId), err.Error())
		return
	}

	data = buildTargetModelFromResponse(data.AccountID, target)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *InfrastructureAccessTargetDeprecatedResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *InfrastructureAccessTargetDeprecatedModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Retrieving Cloudflare Infrastructure Access Target with ID %s", data.ID))
	target, err := r.client.V1.GetInfrastructureAccessTarget(ctx, cloudflare.AccountIdentifier(data.AccountID.ValueString()), data.ID.ValueString())
	if err != nil {
		var notFoundError *cloudflare.NotFoundError
		if errors.As(err, &notFoundError) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError(fmt.Sprintf("error finding Infrastructure Access Target with ID %s", data.ID), err.Error())
		return
	}

	data = buildTargetModelFromResponse(data.AccountID, target)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *InfrastructureAccessTargetDeprecatedResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *InfrastructureAccessTargetDeprecatedModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	accountId := data.AccountID.ValueString()
	if accountId == "" {
		resp.Diagnostics.AddError("failed to update infrastructure access target", "account id cannot be an empty string")
		return
	}
	ipInfo, err := buildUpdateIPInfoFromDetails(ctx, data.IP, resp)
	if err != nil {
		resp.Diagnostics.AddError("failed to create infrastructure access target", "account id cannot be an empty string")
		return
	}
	updatedTargetParams := cloudflare.UpdateInfrastructureAccessTargetParams{
		ID: data.ID.ValueString(),
		ModifyParams: cloudflare.InfrastructureAccessTargetParams{
			Hostname: data.Hostname.ValueString(),
			IP:       ipInfo,
		},
	}

	tflog.Debug(ctx, fmt.Sprintf("Updating Cloudflare Infrastructure Access Target from struct: %+v", updatedTargetParams))
	updatedTarget, err := r.client.V1.UpdateInfrastructureAccessTarget(ctx, cloudflare.AccountIdentifier(accountId), updatedTargetParams)
	if err != nil {
		resp.Diagnostics.AddError(fmt.Sprintf("error updating Infrastructure Access Target with ID %s for account %q", data.ID, accountId), err.Error())
		return
	}

	data = buildTargetModelFromResponse(data.AccountID, updatedTarget)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *InfrastructureAccessTargetDeprecatedResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *InfrastructureAccessTargetDeprecatedModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Deleting Cloudflare Infrastructure Access Target with ID: %s", data.ID))
	err := r.client.V1.DeleteInfrastructureAccessTarget(ctx, cloudflare.AccountIdentifier(data.AccountID.ValueString()), data.ID.ValueString())
	var notFoundError *cloudflare.NotFoundError
	if errors.As(err, &notFoundError) {
		// Return early without error if target is already deleted
		return
	}
	if err != nil {
		resp.Diagnostics.AddError(fmt.Sprintf("error deleting Infrastructure Access Target with ID %s for account %q", data.ID, data.AccountID.ValueString()), err.Error())
		return
	}
}

func buildCreateIPInfoFromDetails(ctx context.Context, ipInfoModel basetypes.ObjectValue, resp *resource.CreateResponse) (cloudflare.InfrastructureAccessTargetIPInfo, error) {
	if ipInfoModel.IsNull() || ipInfoModel.IsUnknown() {
		return cloudflare.InfrastructureAccessTargetIPInfo{}, fmt.Errorf("failed: ip info model is empty")
	}
	var ipInfo *InfrastructureAccessTargetIPInfoDeprecatedModel
	resp.Diagnostics.Append(ipInfoModel.As(ctx, &ipInfo, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)

	if (ipInfo.IPV4.IsNull() || ipInfo.IPV4.IsUnknown()) && (ipInfo.IPV6.IsNull() || ipInfo.IPV6.IsUnknown()) {
		return cloudflare.InfrastructureAccessTargetIPInfo{}, fmt.Errorf("error creating target resource: one of ipv4 or ipv6 must be configured")
	}

	if !(ipInfo.IPV4.IsNull() || ipInfo.IPV4.IsUnknown()) && !(ipInfo.IPV6.IsNull() || ipInfo.IPV6.IsUnknown()) {
		var ipv4Details *InfrastructureAccessTargetIPDetailsDeprecatedModel
		resp.Diagnostics.Append(ipInfo.IPV4.As(ctx, &ipv4Details, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		var ipv6Details *InfrastructureAccessTargetIPDetailsDeprecatedModel
		resp.Diagnostics.Append(ipInfo.IPV6.As(ctx, &ipv6Details, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		return buildIPInfoFromAttributes(ipv4Details.IPAddr.ValueString(), ipv6Details.IPAddr.ValueString(), ipv4Details.VirtualNetworkId.ValueString(), ipv6Details.VirtualNetworkId.ValueString()), nil
	} else if !(ipInfo.IPV4.IsNull() || ipInfo.IPV4.IsUnknown()) {
		var ipv4Details *InfrastructureAccessTargetIPDetailsDeprecatedModel
		resp.Diagnostics.Append(ipInfo.IPV4.As(ctx, &ipv4Details, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		return buildIPV4InfoFromAttributes(ipv4Details.IPAddr.ValueString(), ipv4Details.VirtualNetworkId.ValueString()), nil
	} else {
		var ipv6Details *InfrastructureAccessTargetIPDetailsDeprecatedModel
		resp.Diagnostics.Append(ipInfo.IPV6.As(ctx, &ipv6Details, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		return buildIPV6InfoFromAttributes(ipv6Details.IPAddr.ValueString(), ipv6Details.VirtualNetworkId.ValueString()), nil
	}
}

func buildUpdateIPInfoFromDetails(ctx context.Context, ipInfoModel basetypes.ObjectValue, resp *resource.UpdateResponse) (cloudflare.InfrastructureAccessTargetIPInfo, error) {
	if ipInfoModel.IsNull() || ipInfoModel.IsUnknown() {
		return cloudflare.InfrastructureAccessTargetIPInfo{}, fmt.Errorf("failed: ip info model is empty")
	}
	var ipInfo *InfrastructureAccessTargetIPInfoDeprecatedModel
	resp.Diagnostics.Append(ipInfoModel.As(ctx, &ipInfo, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)

	if (ipInfo.IPV4.IsNull() || ipInfo.IPV4.IsUnknown()) && (ipInfo.IPV6.IsNull() || ipInfo.IPV6.IsUnknown()) {
		return cloudflare.InfrastructureAccessTargetIPInfo{}, fmt.Errorf("error creating target resource: one of ipv4 or ipv6 must be configured")
	}

	if !(ipInfo.IPV4.IsNull() || ipInfo.IPV4.IsUnknown()) && !(ipInfo.IPV6.IsNull() || ipInfo.IPV6.IsUnknown()) {
		var ipv4Details *InfrastructureAccessTargetIPDetailsDeprecatedModel
		resp.Diagnostics.Append(ipInfo.IPV4.As(ctx, &ipv4Details, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		var ipv6Details *InfrastructureAccessTargetIPDetailsDeprecatedModel
		resp.Diagnostics.Append(ipInfo.IPV6.As(ctx, &ipv6Details, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		return buildIPInfoFromAttributes(ipv4Details.IPAddr.ValueString(), ipv6Details.IPAddr.ValueString(), ipv4Details.VirtualNetworkId.ValueString(), ipv6Details.VirtualNetworkId.ValueString()), nil
	} else if !(ipInfo.IPV4.IsNull() || ipInfo.IPV4.IsUnknown()) {
		var ipv4Details *InfrastructureAccessTargetIPDetailsDeprecatedModel
		resp.Diagnostics.Append(ipInfo.IPV4.As(ctx, &ipv4Details, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		return buildIPV4InfoFromAttributes(ipv4Details.IPAddr.ValueString(), ipv4Details.VirtualNetworkId.ValueString()), nil
	} else {
		var ipv6Details *InfrastructureAccessTargetIPDetailsDeprecatedModel
		resp.Diagnostics.Append(ipInfo.IPV6.As(ctx, &ipv6Details, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		return buildIPV6InfoFromAttributes(ipv6Details.IPAddr.ValueString(), ipv6Details.VirtualNetworkId.ValueString()), nil
	}
}

func buildIPInfoFromAttributes(ipv4Addr string, ipv6Addr string, ipv4VirtualNetworkId string, ipv6VirtualNetworkId string) cloudflare.InfrastructureAccessTargetIPInfo {
	return cloudflare.InfrastructureAccessTargetIPInfo{
		IPV4: &cloudflare.InfrastructureAccessTargetIPDetails{
			IPAddr:           ipv4Addr,
			VirtualNetworkId: ipv4VirtualNetworkId,
		},
		IPV6: &cloudflare.InfrastructureAccessTargetIPDetails{
			IPAddr:           ipv6Addr,
			VirtualNetworkId: ipv6VirtualNetworkId,
		},
	}
}

func buildIPV4InfoFromAttributes(ipAddr string, virtualNetworkId string) cloudflare.InfrastructureAccessTargetIPInfo {
	return cloudflare.InfrastructureAccessTargetIPInfo{
		IPV4: &cloudflare.InfrastructureAccessTargetIPDetails{
			IPAddr:           ipAddr,
			VirtualNetworkId: virtualNetworkId,
		},
	}
}

func buildIPV6InfoFromAttributes(ipAddr string, virtualNetworkId string) cloudflare.InfrastructureAccessTargetIPInfo {
	return cloudflare.InfrastructureAccessTargetIPInfo{
		IPV6: &cloudflare.InfrastructureAccessTargetIPDetails{
			IPAddr:           ipAddr,
			VirtualNetworkId: virtualNetworkId,
		},
	}
}

func buildTargetModelFromResponse(accountID tftypes.String, target cloudflare.InfrastructureAccessTarget) *InfrastructureAccessTargetDeprecatedModel {
	built := InfrastructureAccessTargetDeprecatedModel{
		AccountID:  accountID,
		Hostname:   flatteners.String(target.Hostname),
		ID:         flatteners.String(target.ID),
		IP:         convertIPInfoToBaseTypeObject(target.IP),
		CreatedAt:  flatteners.String(target.CreatedAt),
		ModifiedAt: flatteners.String(target.ModifiedAt),
	}
	return &built
}

func convertIPInfoToBaseTypeObject(ipInfo cloudflare.InfrastructureAccessTargetIPInfo) basetypes.ObjectValue {
	if ipInfo.IPV4 != nil && ipInfo.IPV6 != nil {
		ipv4Object := buildObjectFromIpDetails(ipInfo.IPV4.IPAddr, ipInfo.IPV4.VirtualNetworkId)
		ipv6Object := buildObjectFromIpDetails(ipInfo.IPV6.IPAddr, ipInfo.IPV6.VirtualNetworkId)
		parentObjectMap := map[string]attr.Value{
			"ipv4": ipv4Object,
			"ipv6": ipv6Object,
		}
		parentObjectValue, _ := tftypes.ObjectValue(map[string]attr.Type{
			"ipv4": tftypes.ObjectType{
				AttrTypes: map[string]attr.Type{
					"ip_addr":            tftypes.StringType,
					"virtual_network_id": tftypes.StringType,
				},
			},
			"ipv6": tftypes.ObjectType{
				AttrTypes: map[string]attr.Type{
					"ip_addr":            tftypes.StringType,
					"virtual_network_id": tftypes.StringType,
				},
			},
		}, parentObjectMap)
		return parentObjectValue
	} else if ipInfo.IPV4 != nil {
		ipv4Object := buildObjectFromIpDetails(ipInfo.IPV4.IPAddr, ipInfo.IPV4.VirtualNetworkId)
		return buildObjectFromIpInfoV4(ipv4Object)
	} else {
		ipv6Object := buildObjectFromIpDetails(ipInfo.IPV6.IPAddr, ipInfo.IPV6.VirtualNetworkId)
		return buildObjectFromIpInfoV6(ipv6Object)
	}
}

func buildObjectFromIpDetails(ipAddr string, virtualNetworkId string) basetypes.ObjectValue {
	ipDetailsAttributes := map[string]attr.Value{
		"ip_addr":            flatteners.String(ipAddr),
		"virtual_network_id": flatteners.String(virtualNetworkId),
	}
	ipDetailsObjectType, _ := tftypes.ObjectValue(map[string]attr.Type{
		"ip_addr":            tftypes.StringType,
		"virtual_network_id": tftypes.StringType,
	}, ipDetailsAttributes)

	return ipDetailsObjectType
}

func buildObjectFromIpInfoV4(baseObjectMap basetypes.ObjectValue) basetypes.ObjectValue {
	parentObjectMap := map[string]attr.Value{
		"ipv4": baseObjectMap,
		"ipv6": basetypes.NewObjectNull(map[string]attr.Type{
			"ip_addr":            tftypes.StringType,
			"virtual_network_id": tftypes.StringType,
		}),
	}
	return buildIPInfoObjectValue(parentObjectMap)
}

func buildObjectFromIpInfoV6(baseObjectMap basetypes.ObjectValue) basetypes.ObjectValue {
	parentObjectMap := map[string]attr.Value{
		"ipv4": basetypes.NewObjectNull(map[string]attr.Type{
			"ip_addr":            tftypes.StringType,
			"virtual_network_id": tftypes.StringType,
		}),
		"ipv6": baseObjectMap,
	}
	return buildIPInfoObjectValue(parentObjectMap)
}

func buildIPInfoObjectValue(objectMap map[string]attr.Value) basetypes.ObjectValue {
	ipInfoObjectValue, _ := tftypes.ObjectValue(map[string]attr.Type{
		"ipv4": tftypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"ip_addr":            tftypes.StringType,
				"virtual_network_id": tftypes.StringType,
			},
		},
		"ipv6": tftypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"ip_addr":            tftypes.StringType,
				"virtual_network_id": tftypes.StringType,
			},
		},
	}, objectMap)
	return ipInfoObjectValue
}
