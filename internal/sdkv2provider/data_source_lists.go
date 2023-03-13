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

func dataSourceCloudflareLists() *schema.Resource {
	return &schema.Resource{
		Description: heredoc.Doc("Data source for looking up Cloudflare Lists."),
		ReadContext: dataSourceCloudflareListsRead,

		Schema: map[string]*schema.Schema{
			consts.AccountIDSchemaKey: {
				Description: "The account identifier to target for the resource.",
				Type:        schema.TypeString,
				Required:    true,
			},

			"lists": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Description: "The list name to target for the resource.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"id": {
							Description: "List ID.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"description": {
							Description: "List description.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"kind": {
							Description: "List kind.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"numitems": {
							Description: "Number of items in list.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceCloudflareListsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)

	tflog.Debug(ctx, "reading list")

	lists, err := client.ListLists(ctx, cloudflare.AccountIdentifier(accountID), cloudflare.ListListsParams{})
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to fetch Cloudflare lists: %w", err))
	}

	listIds := make([]string, 0)
	listDetails := make([]interface{}, 0)

	for _, l := range lists {
		listDetails = append(listDetails, map[string]interface{}{
			"id":          l.ID,
			"name":        l.Name,
			"description": l.Description,
			"kind":        l.Kind,
			"numitems":    l.NumItems,
		})

		listIds = append(listIds, l.ID)
	}

	err = d.Set("lists", listDetails)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error setting lists: %w", err))
	}

	d.SetId(stringListChecksum(listIds))
	return nil
}
