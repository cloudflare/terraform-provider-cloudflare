package infrastructure_access_target

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
var _ resource.Resource = &InfrastructureAccessTargetResource{}

func NewResource() resource.Resource {
	return &InfrastructureAccessTargetResource{}
}

// InfrastructureAccessTargetResource defines the resource implementation.
type InfrastructureAccessTargetResource struct {
	client *muxclient.Client
}

func (r *InfrastructureAccessTargetResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_infrastructure_access_target"
}

func (r *InfrastructureAccessTargetResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *InfrastructureAccessTargetResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *InfrastructureAccessTargetModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	accountId := data.AccountID.ValueString()
	if accountId == "" {
		resp.Diagnostics.AddError("failed to create infrastructure access target", "account id cannot be an empty string")
		return
	}
	ipInfo, err := validateParseIPInfoCreate(ctx, data, resp)
	if err != nil {
		resp.Diagnostics.AddError("failed to create infrastructure access target", "account id cannot be an empty string")
		return
	}
	createTargetParams := cloudflare.CreateInfrastructureAccessTargetParams{
		InfrastructureAccessTargetParams: cloudflare.InfrastructureAccessTargetParams{
			Hostname: data.Hostname.String(),
			IP:       ipInfo,
		},
	}

	tflog.Debug(ctx, fmt.Sprintf("Creating Cloudflare Infrastructure Target from struct %+v", createTargetParams))
	target, err := r.client.V1.CreateInfrastructureAccessTarget(ctx, cloudflare.AccountIdentifier(accountId), createTargetParams)
	if err != nil {
		resp.Diagnostics.AddError(fmt.Sprintf("error creating Infrastructure Target for account %q", accountId), err.Error())
		return
	}

	data = buildTargetModelFromResponse(data.AccountID, target)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *InfrastructureAccessTargetResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *InfrastructureAccessTargetModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Retrieving Cloudflare Infrastructure Target with ID %s", data.ID))
	target, err := r.client.V1.GetInfrastructureAccessTarget(ctx, cloudflare.AccountIdentifier(data.AccountID.String()), data.ID.String())
	if err != nil {
		var notFoundError *cloudflare.NotFoundError
		if errors.As(err, &notFoundError) {
			resp.Diagnostics.AddError(fmt.Sprintf("Infrastructure Target with ID %s does not exist", data.ID), err.Error())
			return
		}
		resp.Diagnostics.AddError(fmt.Sprintf("error finding Infrastructure Target with ID %s", data.ID), err.Error())
		return
	}

	data = buildTargetModelFromResponse(data.AccountID, target)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *InfrastructureAccessTargetResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *InfrastructureAccessTargetModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	accountId := data.AccountID.ValueString()
	if accountId == "" {
		resp.Diagnostics.AddError("failed to update infrastructure access target", "account id cannot be an empty string")
		return
	}
	ipInfo, err := validateParseIPInfoUpdate(ctx, data, resp)
	if err != nil {
		resp.Diagnostics.AddError("failed to create infrastructure access target", "account id cannot be an empty string")
		return
	}
	updatedTargetParams := cloudflare.UpdateInfrastructureAccessTargetParams{
		ID: data.ID.String(),
		ModifyParams: cloudflare.InfrastructureAccessTargetParams{
			Hostname: data.Hostname.String(),
			IP:       ipInfo,
		},
	}

	tflog.Debug(ctx, fmt.Sprintf("Updating Cloudflare Infrastructure Target from struct: %+v", updatedTargetParams))
	updatedTarget, err := r.client.V1.UpdateInfrastructureAccessTarget(ctx, cloudflare.AccountIdentifier(accountId), updatedTargetParams)
	if err != nil {
		resp.Diagnostics.AddError(fmt.Sprintf("error updating Infrastructure Target with ID %s for account %q", data.ID, accountId), err.Error())
		return
	}

	data = buildTargetModelFromResponse(data.AccountID, updatedTarget)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *InfrastructureAccessTargetResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *InfrastructureAccessTargetModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Deleting Cloudflare Infrastructure Target with ID: %s", data.ID))
	err := r.client.V1.DeleteInfrastructureAccessTarget(ctx, cloudflare.AccountIdentifier(data.AccountID.String()), data.ID.String())
	if err != nil {
		resp.Diagnostics.AddError(fmt.Sprintf("error deleting Infrastructure Target with ID %s for account %q", data.ID, data.AccountID.String()), err.Error())
		return
	}
}

func validateParseIPInfoCreate(ctx context.Context, data *InfrastructureAccessTargetModel, resp *resource.CreateResponse) (cloudflare.InfrastructureAccessTargetIPInfo, error) {
	var ipInfo *InfrastructureAccessTargetIPInfoModel
	resp.Diagnostics.Append(data.IP.As(ctx, &ipInfo, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)

	var ipv4Details *InfrastructureAccessTargetIPDetailsModel
	resp.Diagnostics.Append(ipInfo.IPV4.As(ctx, &ipv4Details, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
	var ipv6Details *InfrastructureAccessTargetIPDetailsModel
	resp.Diagnostics.Append(ipInfo.IPV6.As(ctx, &ipv6Details, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)

	return validateParseIPInfo(ipv4Details, ipv6Details)
}

func validateParseIPInfoUpdate(ctx context.Context, data *InfrastructureAccessTargetModel, resp *resource.UpdateResponse) (cloudflare.InfrastructureAccessTargetIPInfo, error) {
	var ipInfo *InfrastructureAccessTargetIPInfoModel
	resp.Diagnostics.Append(data.IP.As(ctx, &ipInfo, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)

	var ipv4Details *InfrastructureAccessTargetIPDetailsModel
	resp.Diagnostics.Append(ipInfo.IPV4.As(ctx, &ipv4Details, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
	var ipv6Details *InfrastructureAccessTargetIPDetailsModel
	resp.Diagnostics.Append(ipInfo.IPV6.As(ctx, &ipv6Details, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)

	return validateParseIPInfo(ipv4Details, ipv6Details)
}

func validateParseIPInfo(ipv4Details *InfrastructureAccessTargetIPDetailsModel, ipv6Details *InfrastructureAccessTargetIPDetailsModel) (cloudflare.InfrastructureAccessTargetIPInfo, error) {
	if ipv4Details == nil && ipv6Details == nil {
		return cloudflare.InfrastructureAccessTargetIPInfo{}, fmt.Errorf("error creating target resource: one of ipv4 or ipv6 must be configured")
	}

	if ipv4Details != nil && ipv6Details != nil {
		return cloudflare.InfrastructureAccessTargetIPInfo{
			IPV4: &cloudflare.InfrastructureAccessTargetIPDetails{
				IPAddr:           ipv4Details.IPAddr,
				VirtualNetworkId: ipv4Details.VirtualNetworkId,
			},
			IPV6: &cloudflare.InfrastructureAccessTargetIPDetails{
				IPAddr:           ipv6Details.IPAddr,
				VirtualNetworkId: ipv6Details.VirtualNetworkId,
			},
		}, nil
	} else if ipv4Details != nil {
		return cloudflare.InfrastructureAccessTargetIPInfo{
			IPV4: &cloudflare.InfrastructureAccessTargetIPDetails{
				IPAddr:           ipv4Details.IPAddr,
				VirtualNetworkId: ipv4Details.VirtualNetworkId,
			},
		}, nil
	} else {
		return cloudflare.InfrastructureAccessTargetIPInfo{
			IPV6: &cloudflare.InfrastructureAccessTargetIPDetails{
				IPAddr:           ipv6Details.IPAddr,
				VirtualNetworkId: ipv6Details.VirtualNetworkId,
			},
		}, nil
	}
}

func buildTargetModelFromResponse(accountID tftypes.String, target cloudflare.InfrastructureAccessTarget) *InfrastructureAccessTargetModel {
	built := InfrastructureAccessTargetModel{
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
		return buildObjectFromIpInfo("ipv4", ipv4Object)
	} else {
		ipv6Object := buildObjectFromIpDetails(ipInfo.IPV6.IPAddr, ipInfo.IPV6.VirtualNetworkId)
		return buildObjectFromIpInfo("ipv6", ipv6Object)
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

func buildObjectFromIpInfo(ipv string, baseObjectMap basetypes.ObjectValue) basetypes.ObjectValue {
	parentObjectMap := map[string]attr.Value{
		ipv: baseObjectMap,
	}
	parentObjectValue, _ := tftypes.ObjectValue(map[string]attr.Type{
		ipv: tftypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"ip_addr":            tftypes.StringType,
				"virtual_network_id": tftypes.StringType,
			},
		},
	}, parentObjectMap)
	return parentObjectValue
}
