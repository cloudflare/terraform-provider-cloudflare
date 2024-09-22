package zero_trust_access_groups

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/framework/muxclient"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &ZeroTrustAccessGroupsDataSource{}

func NewDataSource() datasource.DataSource {
	return &ZeroTrustAccessGroupsDataSource{}
}

// ZeroTrustAccessGroupsDataSource defines the data source implementation.
type ZeroTrustAccessGroupsDataSource struct {
	client *muxclient.Client
}

func (d *ZeroTrustAccessGroupsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_zero_trust_access_groups"
}

func (d *ZeroTrustAccessGroupsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

	d.client = client
}

func (d *ZeroTrustAccessGroupsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	tflog.Debug(ctx, fmt.Sprintf("Reading Zero Trust Access Group"))
	var data ZeroTrustAccessGroupsModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	params := zero_trust.AccessGroupListParams{
		AccountID: cloudflare.F(data.AccountID.ValueString()),
	}

	iter := d.client.V2.ZeroTrust.Access.Groups.ListAutoPaging(ctx, params)
	var groups []ZeroTrustAccessGroupModel

	for iter.Next() {
		group := iter.Current()

		groups = append(groups, ZeroTrustAccessGroupModel{
			ID:   types.StringValue(group.ID),
			Name: types.StringValue(group.Name),
		})
	}
	if err := iter.Err(); err != nil {
		resp.Diagnostics.AddError("Failed to fetch Zero Trust Access Groups", err.Error())
		return
	}

	data.Groups = groups

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
