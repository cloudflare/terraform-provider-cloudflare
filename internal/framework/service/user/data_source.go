package user

import (
	"context"
	"fmt"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/framework/muxclient"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSource = &CloudflareUserDataSource{}

func NewDataSource() datasource.DataSource {
	return &CloudflareUserDataSource{}
}

type CloudflareUserDataSource struct {
	client *muxclient.Client
}

func (r *CloudflareUserDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_user"
}

func (r *CloudflareUserDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *CloudflareUserDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data CloudflareUserDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	user, err := r.client.V1.UserDetails(ctx)
	if err != nil {
		resp.Diagnostics.AddError("unable to retrieve user details", err.Error())
		return
	}

	data = CloudflareUserDataSourceModel{
		ID:       types.StringValue(user.ID),
		Email:    types.StringValue(user.Email),
		Username: types.StringValue(user.Username),
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
