package sdkv2provider

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCloudflareAccountRolesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)

	tflog.Debug(ctx, fmt.Sprintf("Reading Account Roles"))
	roles, err := client.AccountRoles(ctx, accountID)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error listing Account Roles: %w", err))
	}

	roleIds := make([]string, 0)
	roleDetails := make([]interface{}, 0)

	for _, v := range roles {
		roleDetails = append(roleDetails, map[string]interface{}{
			"id":          v.ID,
			"name":        v.Name,
			"description": v.Description,
		})
		roleIds = append(roleIds, v.ID)
	}

	err = d.Set("roles", roleDetails)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error setting roles: %w", err))
	}

	d.SetId(stringListChecksum(roleIds))
	return nil
}
