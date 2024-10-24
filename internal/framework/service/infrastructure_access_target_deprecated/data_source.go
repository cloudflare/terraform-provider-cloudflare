package infrastructure_access_target_deprecated

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/framework/muxclient"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSource = &InfrastructureAccessTargetDeprecatedDataSource{}

func NewDataSource() datasource.DataSource {
	return &InfrastructureAccessTargetDeprecatedDataSource{}
}

type InfrastructureAccessTargetDeprecatedDataSource struct {
	client *muxclient.Client
}

func (d *InfrastructureAccessTargetDeprecatedDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_infrastructure_access_targets"
}

func (d *InfrastructureAccessTargetDeprecatedDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *InfrastructureAccessTargetDeprecatedDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *InfrastructureAccessTargetsDeprecatedModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	accountId := data.AccountID.ValueString()
	if accountId == "" {
		resp.Diagnostics.AddError("failed to update infrastructure access target", "account id cannot be an empty string")
		return
	}
	params := cloudflare.InfrastructureAccessTargetListParams{
		Hostname:         data.Hostname.ValueString(),
		HostnameContains: data.HostnameContains.ValueString(),
		IPV4:             data.IPV4.ValueString(),
		IPV6:             data.IPV6.ValueString(),
		CreatedAfter:     data.CreatedAfter.ValueString(),
		ModifedAfter:     data.ModifiedAfter.ValueString(),
		VirtualNetworkId: data.VirtualNetworkId.ValueString(),
	}

	allTargets, _, err := d.client.V1.ListInfrastructureAccessTargets(ctx, cloudflare.AccountIdentifier(accountId), params)
	if err != nil {
		resp.Diagnostics.AddError("failed to fetch Infrastructure Access Targets: %w", err.Error())
		return
	}
	if len(allTargets) == 0 {
		resp.Diagnostics.AddError("failed to fetch Infrastructure Access Targets", "no Infrastructure Access Targets matching given query parameters")
	}

	var targets []InfrastructureAccessTargetDeprecatedModel
	for _, target := range allTargets {
		targets = append(targets, InfrastructureAccessTargetDeprecatedModel{
			AccountID:  types.StringValue(accountId),
			Hostname:   types.StringValue(target.Hostname),
			ID:         types.StringValue(target.ID),
			IP:         convertIPInfoToBaseTypeObject(target.IP),
			CreatedAt:  types.StringValue(target.CreatedAt),
			ModifiedAt: types.StringValue(target.ModifiedAt),
		})
	}

	data.Targets = targets
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
