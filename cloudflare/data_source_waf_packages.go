package cloudflare

import (
	"fmt"
	"log"
	"regexp"
	"time"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func dataSourceCloudflareWAFPackages() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceCloudflareWAFPackagesRead,

		Schema: map[string]*schema.Schema{
			"zone_id": {
				Type:     schema.TypeString,
				Required: true,
			},

			"filter": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"detection_mode": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"sensitivity": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.StringInSlice([]string{"high", "medium", "low", "off"}, false),
						},
						"action_mode": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.StringInSlice([]string{"simulate", "block", "challenge"}, false),
						},
					},
				},
			},

			"packages": {
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
						"detection_mode": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"sensitivity": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"action_mode": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceCloudflareWAFPackagesRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)

	// Prepare the filters to be applied to the search
	filter, err := expandFilterWAFPackages(d.Get("filter"))
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Reading WAF Packages")
	packageDetails := make([]interface{}, 0)
	pkgList, err := client.ListWAFPackages(zoneID)
	if err != nil {
		return err
	}

	for _, pkg := range pkgList {
		if filter.Name != nil && !filter.Name.Match([]byte(pkg.Name)) {
			continue
		}

		if filter.DetectionMode != "" && filter.DetectionMode != pkg.DetectionMode {
			continue
		}

		if filter.Sensitivity != "" && filter.Sensitivity != pkg.Sensitivity {
			continue
		}

		if filter.ActionMode != "" && filter.ActionMode != pkg.ActionMode {
			continue
		}

		packageDetails = append(packageDetails, map[string]interface{}{
			"id":             pkg.ID,
			"name":           pkg.Name,
			"description":    pkg.Description,
			"detection_mode": pkg.DetectionMode,
			"sensitivity":    pkg.Sensitivity,
			"action_mode":    pkg.ActionMode,
		})
	}

	err = d.Set("packages", packageDetails)
	if err != nil {
		return fmt.Errorf("Error setting WAF packages: %s", err)
	}

	d.SetId("WAFPackages " + time.Now().UTC().String())
	return nil
}

func expandFilterWAFPackages(d interface{}) (*searchFilterWAFPackages, error) {
	cfg := d.([]interface{})
	filter := &searchFilterWAFPackages{}
	if len(cfg) == 0 || cfg[0] == nil {
		return filter, nil
	}

	m := cfg[0].(map[string]interface{})
	name, ok := m["name"]
	if ok {
		match, err := regexp.Compile(name.(string))
		if err != nil {
			return nil, err
		}

		filter.Name = match
	}

	detectionMode, ok := m["detection_mode"]
	if ok {
		filter.DetectionMode = detectionMode.(string)
	}

	sensitivity, ok := m["sensitivity"]
	if ok {
		filter.Sensitivity = sensitivity.(string)
	}

	actionMode, ok := m["action_mode"]
	if ok {
		filter.ActionMode = actionMode.(string)
	}

	return filter, nil
}

type searchFilterWAFPackages struct {
	Name          *regexp.Regexp
	DetectionMode string
	Sensitivity   string
	ActionMode    string
}
