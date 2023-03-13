package sdkv2provider

import (
	"context"
	"fmt"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCloudflareList() *schema.Resource {
	return &schema.Resource{
		Description: heredoc.Doc("Data source for looking up Cloudflare Lists."),
		ReadContext: dataSourceCloudflareListRead,

		Schema: map[string]*schema.Schema{
			consts.AccountIDSchemaKey: {
				Description: "The account identifier to target for the resource.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"name": {
				Description: "The list name to target for the resource.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"description": {
				Description: "List description.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"kind": {
				Description: "List kind.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"numitems": {
				Description: "Number of items in list.",
				Type:        schema.TypeInt,
				Computed:    true,
			},
		},
	}
}

func dataSourceCloudflareListRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)
	name := d.Get("name").(string)

	tflog.Debug(ctx, "reading list")

	lists, err := client.ListLists(ctx, cloudflare.AccountIdentifier(accountID), cloudflare.ListListsParams{})
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to fetch Cloudflare lists: %w", err))
	}

	for _, l := range lists {
		if l.Name == name {
			d.SetId(l.ID)
			d.Set("name", l.Name)
			d.Set("description", l.Description)
			d.Set("kind", l.Kind)
			d.Set("numitems", l.NumItems)

			return nil
		}
	}

	return diag.Errorf("unable to find list named %s", name)
}
