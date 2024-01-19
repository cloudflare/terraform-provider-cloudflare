package sdkv2provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-log/tflog"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCloudflareAccessGroup() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCloudflareAccessGroupRead,
		Schema: map[string]*schema.Schema{
			consts.AccountIDSchemaKey: {
				Description:   consts.AccountIDSchemaDescription,
				Type:          schema.TypeString,
				Optional:      true,
				ForceNew:      true,
				ConflictsWith: []string{consts.ZoneIDSchemaKey},
			},
			consts.ZoneIDSchemaKey: {
				Description:   consts.ZoneIDSchemaDescription,
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{consts.AccountIDSchemaKey},
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceCloudflareAccessGroupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)
	d.SetId(accountID)
	searchAccessGroupName := d.Get("name").(string)

	tflog.Debug(ctx, "reading accounts")
	identifier, err := initIdentifier(d)
	if err != nil {
		return diag.FromErr(err)
	}

	// Call the Cloudflare API to retrieve the access group by name
	accessGroups, _, err := client.ListAccessGroups(ctx, identifier, cloudflare.ListAccessGroupsParams{})
	if err != nil {
		return diag.FromErr(fmt.Errorf("error listing Access Groups: %w", err))
	}
	if len(accessGroups) == 0 {
		return diag.Errorf("no Access groups found")
	}
	var accessGroup cloudflare.AccessGroup
	for _, group := range accessGroups {
		if group.Name == searchAccessGroupName {
			accessGroup = group
			break
		}
	}
	if accessGroup.ID == "" {
		return diag.Errorf("No Access groups matching name %q", searchAccessGroupName)
	}
	d.SetId(accessGroup.ID)
	d.Set("name", accessGroup.Name)

	return nil
}
