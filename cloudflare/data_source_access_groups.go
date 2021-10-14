package cloudflare

import (
	"context"
	"errors"
	"fmt"
	"log"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCloudflareAccessGroups() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceCloudflareAccessGroupsRead,

		Schema: map[string]*schema.Schema{
			"account_id": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"zone_id"},
			},
			"zone_id": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"account_id"},
			},
			"groups": {
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
						"require": {
							Type:     schema.TypeList,
							Optional: true,
							Elem:     AccessGroupOptionSchemaElement,
						},
						"exclude": {
							Type:     schema.TypeList,
							Optional: true,
							Elem:     AccessGroupOptionSchemaElement,
						},
						"include": {
							Type:     schema.TypeList,
							Required: true,
							Elem:     AccessGroupOptionSchemaElement,
						},
					},
				},
			},
		},
	}
}

func paginateAllGroups(readAccessGroupHandler func (ctx context.Context, accountID string, pageOpts cloudflare.PaginationOptions) ([]cloudflare.AccessGroup, cloudflare.ResultInfo, error), filter string) ([]cloudflare.AccessGroup, error) {
	var groups []cloudflare.AccessGroup
	paginationOptions := cloudflare.PaginationOptions{}

	for {
		groupPage, resultInfo, err := readAccessGroupHandler(context.Background(), filter, paginationOptions)
		if err != nil {
			return nil, err
		}

		for _, g := range groupPage {
			groups = append(groups, g)
		}
		if resultInfo.Page < resultInfo.TotalPages {
			paginationOptions.Page = resultInfo.Page + 1
		} else {
			return groups, nil
		}
	}
}

func dataSourceCloudflareAccessGroupsRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] Reading Access Groups")
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)
	accountID := d.Get("account_id").(string)

	var filter string
	var handler func (ctx context.Context, accountID string, pageOpts cloudflare.PaginationOptions) ([]cloudflare.AccessGroup, cloudflare.ResultInfo, error)

	if accountID != "" {
		filter = accountID
		handler = client.AccessGroups
	} else if zoneID != "" {
		filter = zoneID
		handler = client.ZoneLevelAccessGroups
	} else {
		return errors.New("One of Zone ID or Account ID are required")
	}

	groups, err := paginateAllGroups(handler, filter)

	if err != nil {
		return nil
	}

	if err != nil {
		return fmt.Errorf("error setting access groups: %s", err)
	}

	groupIds := make([]string, 0)
	groupDetails := make([]interface{}, 0)
	for _, g := range groups {
		groupDetails = append(groupDetails, map[string]interface{}{
			"id": g.ID,
			"name": g.Name,
			"include": TransformAccessGroupForSchema(g.Include),
			"exclude": TransformAccessGroupForSchema(g.Exclude),
			"require": TransformAccessGroupForSchema(g.Require),
		})
		groupIds = append(groupIds, g.ID)
	}
	err = d.Set("groups", groupDetails)
	if err != nil {
		return fmt.Errorf("error setting access groups: %s", err)
	}
	d.SetId(stringListChecksum(groupIds))
	return nil
}
