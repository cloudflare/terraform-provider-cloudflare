package provider

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCloudflareAccountRoles() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCloudflareAccountRolesRead,

		Schema: map[string]*schema.Schema{
			"account_id": {
				Description: "The account identifier to target for the resource.",
				Type:        schema.TypeString,
				Required:    true,
			},

			"roles": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"description": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceCloudflareAccountRolesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get("account_id").(string)

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
