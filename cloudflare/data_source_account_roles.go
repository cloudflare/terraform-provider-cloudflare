package cloudflare

import (
	"context"
	"fmt"
	"log"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCloudflareAcountRoles() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceCloudflareAcountRolesRead,

		Schema: map[string]*schema.Schema{
			"account_id": {
				Type:     schema.TypeString,
				Required: true,
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

func dataSourceCloudflareAcountRolesRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	accountID := d.Get("account_id").(string)

	log.Printf("[DEBUG] Reading Account Roles")
	roles, err := client.AccountRoles(context.Background(), accountID)
	if err != nil {
		return fmt.Errorf("error listing Account Roles: %s", err)
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
		return fmt.Errorf("error setting roles: %s", err)
	}

	d.SetId(stringListChecksum(roleIds))
	return nil
}
