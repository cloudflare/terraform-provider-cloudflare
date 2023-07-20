package sdkv2provider

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCloudflareAccountRoles() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCloudflareAccountRolesRead,

		Schema: map[string]*schema.Schema{
			consts.AccountIDSchemaKey: {
				Description: consts.AccountIDSchemaDescription,
				Type:        schema.TypeString,
				Required:    true,
			},

			"roles": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "A list of roles object.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Role identifier tag.",
						},
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Role Name.",
						},
						"description": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Description of role's permissions.",
						},
					},
				},
			},
		},
		Description: "Use this data source to lookup [Account Roles](https://api.cloudflare.com/#account-roles-properties).",
	}
}
