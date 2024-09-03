package dcv_delegation

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/framework/muxclient"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSource = &DCVDelegationDataSource{}

func NewDataSource() datasource.DataSource {
	return &DCVDelegationDataSource{}
}

type DCVDelegationDataSource struct {
	client *muxclient.Client
}

func (r *DCVDelegationDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dcv_delegation"
}

func (r *DCVDelegationDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *DCVDelegationDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data DCVDelegationModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	dcv, _, err := r.client.V1.GetDCVDelegation(ctx, cloudflare.ZoneIdentifier(data.ZoneID.ValueString()), cloudflare.GetDCVDelegationParams{})
	if err != nil {
		resp.Diagnostics.AddError("failed to fetch DCV Delegation", err.Error())
		return
	}

	data = DCVDelegationModel{
		ZoneID:   types.StringValue(data.ZoneID.ValueString()),
		ID:       types.StringValue(dcv.UUID),
		Hostname: types.StringValue(fmt.Sprintf("%s.dcv.cloudflare.com", dcv.UUID)),
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
