package sdkv2provider

import (
	"context"
	"fmt"
	"regexp"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceCloudflareZones() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCloudflareZonesRead,

		Schema: map[string]*schema.Schema{
			"filter": {
				Type:        schema.TypeList,
				Required:    true,
				MaxItems:    1,
				Description: "One or more values used to look up zone records. If more than one value is given all values must match in order to be included.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						consts.AccountIDSchemaKey: {
							Description: consts.AccountIDSchemaDescription,
							Type:        schema.TypeString,
							Optional:    true,
						},
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "A string value to search for.",
						},
						"match": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "A RE2 compatible regular expression to filter the	results. This is performed client side whereas the `name` and `lookup_type`	are performed on the Cloudflare server side.",
						},
						"lookup_type": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.StringInSlice([]string{"contains", "exact"}, false),
							Description:  fmt.Sprintf("The type of search to perform for the `name` value when querying the zone API. %s", renderAvailableDocumentationValuesStringSlice([]string{"contains", "exact"})),
							Default:      "exact",
						},
						"status": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Status of the zone to lookup.",
						},
						"paused": {
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     false,
							Description: "Paused status of the zone to lookup.",
						},
					},
				},
			},
			"zones": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "A list of zone objects.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The zone ID.",
						},
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Zone name.",
						},
					},
				},
			},
		},
		Description: "Use this data source to look up Zone results for use in other resources.",
	}
}

func dataSourceCloudflareZonesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	tflog.Debug(ctx, fmt.Sprintf("Reading Zones"))
	client := meta.(*cloudflare.API)
	filter, err := expandFilter(d.Get("filter"))
	if err != nil {
		return diag.FromErr(err)
	}

	zoneLookupValue := filter.name
	if filter.lookupType == "contains" {
		zoneLookupValue = "contains:" + zoneLookupValue
	}

	zoneFilter := cloudflare.WithZoneFilters(
		zoneLookupValue,
		filter.accountID,
		filter.status,
	)

	zones, err := client.ListZonesContext(ctx, zoneFilter)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error listing Zone: %w", err))
	}

	zoneIds := make([]string, 0)
	zoneDetails := make([]interface{}, 0)
	for _, v := range zones.Result {
		if filter.regexValue != nil {
			if !filter.regexValue.Match([]byte(v.Name)) {
				continue
			}
		}

		if filter.paused != v.Paused {
			continue
		}

		zoneDetails = append(zoneDetails, map[string]interface{}{
			"id":   v.ID,
			"name": v.Name,
		})
		zoneIds = append(zoneIds, v.ID)
	}

	err = d.Set("zones", zoneDetails)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error setting zones: %w", err))
	}

	d.SetId(stringListChecksum(zoneIds))
	return nil
}

func expandFilter(d interface{}) (*searchFilter, error) {
	cfg := d.([]interface{})
	filter := &searchFilter{}

	m := cfg[0].(map[string]interface{})
	name, ok := m["name"]
	if ok {
		filter.name = name.(string)
	}

	match, ok := m["match"]
	if ok {
		match, err := regexp.Compile(match.(string))
		if err != nil {
			return nil, err
		}

		filter.regexValue = match
	}

	lookupType, ok := m["lookup_type"]
	if ok {
		filter.lookupType = lookupType.(string)
	}

	paused, ok := m["paused"]
	if ok {
		filter.paused = paused.(bool)
	}

	status, ok := m["status"]
	if ok {
		filter.status = status.(string)
	}

	accountID, ok := m[consts.AccountIDSchemaKey]
	if ok {
		filter.accountID = accountID.(string)
	}

	return filter, nil
}

type searchFilter struct {
	accountID  string
	name       string
	regexValue *regexp.Regexp
	lookupType string
	status     string
	paused     bool
}
