package cloudflare

import (
	"fmt"
	"log"
	"time"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceCloudflareApiTokenPermissionGroups() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceCloudflareApiTokenPermissionGroupsRead,

		Schema: map[string]*schema.Schema{
			"permissions": {
				Computed: true,
				Type:     schema.TypeMap,
			},
		},
	}
}

func dataSourceCloudflareApiTokenPermissionGroupsRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] Reading User Token Permission Groups")
	client := meta.(*cloudflare.API)

	permissions, err := client.ListAPITokensPermissionGroups()
	if err != nil {
		return fmt.Errorf("error listing User Token Permission Groups: %s", err)
	}

	permissionDetails := make(map[string]interface{}, 0)
	for _, v := range permissions {
		permissionDetails[v.Name] = v.ID
	}

	err = d.Set("permissions", permissionDetails)
	if err != nil {
		return fmt.Errorf("Error setting User Token Permission Groups: %s", err)
	}

	d.SetId(time.Now().UTC().String())
	return nil
}
