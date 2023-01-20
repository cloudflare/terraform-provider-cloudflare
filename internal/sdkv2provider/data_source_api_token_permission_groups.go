package sdkv2provider

import (
	"context"
	"fmt"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCloudflareApiTokenPermissionGroups() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCloudflareApiTokenPermissionGroupsRead,
		Description: heredoc.Docf(`
			Use this data source to look up [API Token Permission Groups](https://developers.cloudflare.com/api/tokens/create/permissions).
			Commonly used as references within [%s](/docs/providers/cloudflare/r/api_token.html) resources.
		`, "`cloudflare_token`"),
		Schema: map[string]*schema.Schema{
			"permissions": {
				Computed:    true,
				Type:        schema.TypeMap,
				Deprecated:  "Use specific account, zone or user attributes instead.",
				Description: "Map of all permissions available. Should not be used as some permissions will overlap resource scope. Instead, use resource level specific attributes.",
			},
			"zone": {
				Computed:    true,
				Type:        schema.TypeMap,
				Description: "Map of permissions for zone level resources.",
			},
			"account": {
				Computed:    true,
				Type:        schema.TypeMap,
				Description: "Map of permissions for account level resources.",
			},
			"user": {
				Computed:    true,
				Type:        schema.TypeMap,
				Description: "Map of permissions for user level resources.",
			},
		},
	}
}

func dataSourceCloudflareApiTokenPermissionGroupsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	tflog.Debug(ctx, fmt.Sprintf("Reading API Token Permission Groups"))
	client := meta.(*cloudflare.API)

	permissions, err := client.ListAPITokensPermissionGroups(ctx)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error listing API Token Permission Groups: %w", err))
	}

	permissionDetails := make(map[string]interface{}, 0)
	zoneScopes := make(map[string]interface{}, 0)
	accountScopes := make(map[string]interface{}, 0)
	userScopes := make(map[string]interface{}, 0)
	ids := []string{}

	for _, v := range permissions {
		// This is for backwards compatibility and shouldn't be used going forward
		// due to some permissions overlapping and returning invalid IDs.
		permissionDetails[v.Name] = v.ID
		ids = append(ids, v.ID)

		switch v.Scopes[0] {
		case "com.cloudflare.api.account":
			accountScopes[v.Name] = v.ID
		case "com.cloudflare.api.account.zone":
			zoneScopes[v.Name] = v.ID
		case "com.cloudflare.api.user":
			userScopes[v.Name] = v.ID
		default:
			tflog.Warn(ctx, fmt.Sprintf("unknown permission scope found: %s", v.Scopes[0]))
		}
	}

	if err = d.Set("account", accountScopes); err != nil {
		return diag.FromErr(fmt.Errorf("error setting API Token Permission Groups for accounts: %w", err))
	}

	if err = d.Set("zone", zoneScopes); err != nil {
		return diag.FromErr(fmt.Errorf("error setting API Token Permission Groups for zones: %w", err))
	}

	if err = d.Set("user", userScopes); err != nil {
		return diag.FromErr(fmt.Errorf("error setting API Token Permission Groups for user: %w", err))
	}

	if err = d.Set("permissions", permissionDetails); err != nil {
		return diag.FromErr(fmt.Errorf("error setting API Token Permission Groups: %w", err))
	}

	d.SetId(stringListChecksum(ids))

	return nil
}
