package provider

import (
	"context"
	"fmt"
	"log"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCloudflareApiTokenPermissionGroups() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCloudflareApiTokenPermissionGroupsRead,

		Schema: map[string]*schema.Schema{
			"permissions": {
				Computed: true,
				Type:     schema.TypeMap,
			},
		},
	}
}

func dataSourceCloudflareApiTokenPermissionGroupsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Reading API Token Permission Groups")
	client := meta.(*cloudflare.API)

	permissions, err := client.ListAPITokensPermissionGroups(ctx)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error listing API Token Permission Groups: %w", err))
	}

	permissionDetails := make(map[string]interface{}, 0)
	ids := []string{}
	for _, v := range permissions {
		permissionDetails[v.Name] = v.ID
		ids = append(ids, v.ID)
	}

	err = d.Set("permissions", permissionDetails)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error setting API Token Permission Groups: %w", err))
	}

	d.SetId(stringListChecksum(ids))

	return nil
}
