package api_token_permissions_groups

import (
	"context"
	"crypto/md5"
	"fmt"
	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"sort"
	"strings"
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
	resp.TypeName = req.ProviderTypeName + "_api_token_permission_groups"
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

	permissionDetails := make(map[string]types.String)
	zoneScopes := make(map[string]types.String)
	accountScopes := make(map[string]types.String)
	userScopes := make(map[string]types.String)
	r2Scopes := make(map[string]types.String)
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
	data.ID = types.StringValue(stringListChecksum(ids))
	data.Permissions = permissionDetails
	data.Zone = zoneScopes
	data.Account = accountScopes
	data.User = userScopes
	data.R2 = r2Scopes

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func stringListChecksum(s []string) string {
	sort.Strings(s)
	return stringChecksum(strings.Join(s, ""))
}

// stringChecksum takes a string and returns the checksum of the string.
func stringChecksum(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)

	return fmt.Sprintf("%x", bs)
}
