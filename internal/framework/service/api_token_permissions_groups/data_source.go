package api_token_permissions_groups

import (
	"context"
	"fmt"
	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &APITokenPermissionsGroupDataSource{}

func NewDataSource() datasource.DataSource {
	return &APITokenPermissionsGroupDataSource{}
}

// APITokenPermissionsGroupDataSource defines the data source implementation.
type APITokenPermissionsGroupDataSource struct {
	client *cloudflare.API
}

func (r *APITokenPermissionsGroupDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_api_token_permissions_group"
}

func (r *APITokenPermissionsGroupDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*cloudflare.API)

	if !ok {
		resp.Diagnostics.AddError(
			"unexpected resource configure type",
			fmt.Sprintf("expected *cloudflare.API, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *APITokenPermissionsGroupDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	tflog.Debug(ctx, fmt.Sprintf("Reading API Token Permission Groups"))
	var data APITokenPermissionsGroupModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	permissions, err := r.client.ListAPITokensPermissionGroups(ctx)

	if err != nil {
		resp.Diagnostics.AddError(
			"error listing API Token Permission Groups",
			err.Error(),
		)
		return
	}

	permissionDetails := make(map[string]attr.Value)
	zoneScopes := make(map[string]attr.Value)
	accountScopes := make(map[string]attr.Value)
	userScopes := make(map[string]attr.Value)
	r2Scopes := make(map[string]attr.Value)
	var ids []string

	for _, v := range permissions {
		// This is for backwards compatibility and shouldn't be used going forward
		// due to some permissions overlapping and returning invalid IDs.
		permissionDetails[v.Name] = types.StringValue(v.ID)
		ids = append(ids, v.ID)

		switch v.Scopes[0] {
		case "com.cloudflare.api.account":
			accountScopes[v.Name] = types.StringValue(v.ID)
		case "com.cloudflare.api.account.zone":
			zoneScopes[v.Name] = types.StringValue(v.ID)
		case "com.cloudflare.api.user":
			userScopes[v.Name] = types.StringValue(v.ID)
		case "com.cloudflare.edge.r2.bucket":
			r2Scopes[v.Name] = types.StringValue(v.ID)
		default:
			tflog.Warn(ctx, fmt.Sprintf("unknown permission scope found: %s", v.Scopes[0]))
		}
	}
	permissionsMap, errDiag := types.MapValue(types.StringType, permissionDetails)
	if errDiag.HasError() {
		return
	}
	data.Permissions = permissionsMap

	zoneMap, errDiag := types.MapValue(types.StringType, zoneScopes)
	if errDiag.HasError() {
		return
	}
	data.Zone = zoneMap

	accountMap, errDiag := types.MapValue(types.StringType, accountScopes)
	if errDiag.HasError() {
		return
	}
	data.Account = accountMap

	userMap, errDiag := types.MapValue(types.StringType, userScopes)
	if errDiag.HasError() {
		return
	}
	data.User = userMap

	r2Map, errDiag := types.MapValue(types.StringType, r2Scopes)
	if errDiag.HasError() {
		return
	}
	data.R2 = r2Map

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)

}
