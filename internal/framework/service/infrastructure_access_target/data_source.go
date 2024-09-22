package infrastructure_access_target

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/framework/muxclient"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSource = &InfrastructureAccessTargetDataSource{}

func NewDataSource() datasource.DataSource {
	return &InfrastructureAccessTargetDataSource{}
}

type InfrastructureAccessTargetDataSource struct {
	client *muxclient.Client
}

func (d *InfrastructureAccessTargetDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_infrastructure_access_targets"
}

func (d *InfrastructureAccessTargetDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *InfrastructureAccessTargetDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *InfrastructureAccessTargetsModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	accountId := data.AccountID.ValueString()
	if accountId == "" {
		resp.Diagnostics.AddError("failed to update infrastructure access target", "account id cannot be an empty string")
		return
	}
	checkSetNil := func(s string) *string {
		if s == "" {
			return nil
		} else {
			return &s
		}
	}
	params := cloudflare.TargetListParams{
		Hostname:         *checkSetNil(data.Hostname.String()),
		HostnameContains: *checkSetNil(data.HostnameContains.String()),
		IPV4:             *checkSetNil(data.IPV4.String()),
		IPV6:             *checkSetNil(data.IPV6.String()),
		CreatedAfter:     *checkSetNil(data.CreatedAfter.String()),
		ModifedAfter:     *checkSetNil(data.ModifiedAfter.String()),
		VirtualNetworkId: *checkSetNil(data.VirtualNetworkId.String()),
	}

	allTargets, _, err := d.client.V1.ListInfrastructureAccessTargets(ctx, cloudflare.AccountIdentifier(accountId), params)
	if err != nil {
		resp.Diagnostics.AddError("failed to fetch Infrastructure Targets: %w", err.Error())
		return
	}
	if len(allTargets) == 0 {
		resp.Diagnostics.AddError("failed to fetch Infrastructure Targets", "no Infrastructure Targets matching given query parameters")
	}

	var targets []InfrastructureAccessTargetModel
	for _, target := range allTargets {
		targets = append(targets, InfrastructureAccessTargetModel{
			AccountID:  types.StringValue(accountId),
			Hostname:   types.StringValue(target.Hostname),
			ID:         types.StringValue(target.ID),
			IP:         target.IP,
			CreatedAt:  types.StringValue(target.CreatedAt),
			ModifiedAt: types.StringValue(target.ModifiedAt),
		})
	}

	data.Targets = targets
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
