package cloudflare

import (
	"context"
	"fmt"
	"strings"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

const CLOUDFLARE_INVALID_OR_REMOVED_WAF_RULE_SET_ID_ERROR = 1003

func resourceCloudflareWAFGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceCloudflareWAFGroupCreate,
		Read:   resourceCloudflareWAFGroupRead,
		Update: resourceCloudflareWAFGroupUpdate,
		Delete: resourceCloudflareWAFGroupDelete,

		Importer: &schema.ResourceImporter{
			State: resourceCloudflareWAFGroupImport,
		},

		SchemaVersion: 0,
		Schema: map[string]*schema.Schema{
			"group_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"zone_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"package_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"mode": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "on",
				ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
			},
		},
	}
}

func errorIsWAFGroupNotFound(err error) bool {
	return cloudflareErrorIsOneOfCodes(err, []int{
		CLOUDFLARE_INVALID_OR_REMOVED_WAF_PACKAGE_ID_ERROR,
		CLOUDFLARE_INVALID_OR_REMOVED_WAF_RULE_SET_ID_ERROR,
	})
}

func resourceCloudflareWAFGroupRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)

	groupID := d.Get("group_id").(string)
	zoneID := d.Get("zone_id").(string)
	packageID := d.Get("package_id").(string)

	group, err := client.WAFGroup(context.Background(), zoneID, packageID, groupID)
	if err != nil {
		if errorIsWAFGroupNotFound(err) {
			d.SetId("")
			return nil
		}

		return err
	}

	// Only need to set mode as that is the only attribute that could have changed
	d.Set("mode", group.Mode)
	d.SetId(group.ID)

	return nil
}

func resourceCloudflareWAFGroupCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	groupID := d.Get("group_id").(string)
	zoneID := d.Get("zone_id").(string)
	packageID := d.Get("package_id").(string)
	mode := d.Get("mode").(string)

	// If no package ID is given try to resolve it
	var pkgList []cloudflare.WAFPackage
	if packageID == "" {
		var err error
		pkgList, err = client.ListWAFPackages(context.Background(), zoneID)
		if err != nil {
			return err
		}
	} else {
		pkgList = append(pkgList, cloudflare.WAFPackage{ID: packageID})
	}

	for _, pkg := range pkgList {
		var err error
		var group cloudflare.WAFGroup

		group, err = client.WAFGroup(context.Background(), zoneID, pkg.ID, groupID)
		if err != nil {
			continue
		}

		d.Set("group_id", group.ID)
		d.Set("zone_id", zoneID)
		d.Set("package_id", pkg.ID)

		if group.Mode != mode {
			err = resourceCloudflareWAFGroupUpdate(d, meta)
			if err != nil {
				d.SetId("")
				return err
			}
		}

		return resourceCloudflareWAFGroupRead(d, meta)
	}

	return fmt.Errorf("Unable to find WAF Group %s", groupID)
}

func resourceCloudflareWAFGroupDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)

	groupID := d.Get("group_id").(string)
	zoneID := d.Get("zone_id").(string)
	packageID := d.Get("package_id").(string)

	group, err := client.WAFGroup(context.Background(), zoneID, packageID, groupID)
	if err != nil {
		return err
	}

	// Can't delete WAF Group so instead reset it to default
	schema := resourceCloudflareWAFGroup().Schema
	defaultMode := schema["mode"].Default.(string)

	if group.Mode != defaultMode {
		_, err = client.UpdateWAFGroup(context.Background(), zoneID, packageID, groupID, defaultMode)
		if err != nil {
			return err
		}
	}

	return nil
}

func resourceCloudflareWAFGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)

	groupID := d.Get("group_id").(string)
	zoneID := d.Get("zone_id").(string)
	mode := d.Get("mode").(string)
	packageID := d.Get("package_id").(string)

	// We can only update the mode of a WAF Group
	_, err := client.UpdateWAFGroup(context.Background(), zoneID, packageID, groupID, mode)
	if err != nil {
		return err
	}

	return resourceCloudflareWAFGroupRead(d, meta)
}

func resourceCloudflareWAFGroupImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	client := meta.(*cloudflare.API)

	// split the id so we can lookup
	idAttr := strings.SplitN(d.Id(), "/", 2)
	var zoneID string
	var groupID string
	if len(idAttr) == 2 {
		zoneID = idAttr[0]
		groupID = idAttr[1]
	} else {
		return nil, fmt.Errorf("invalid id (\"%s\") specified, should be in format \"zoneID/GroupID\" for import", d.Id())
	}

	pkgList, err := client.ListWAFPackages(context.Background(), zoneID)
	if err != nil {
		return nil, fmt.Errorf("error listing WAF packages: %s", err)
	}

	for _, pkg := range pkgList {
		group, err := client.WAFGroup(context.Background(), zoneID, pkg.ID, groupID)
		if err != nil {
			continue
		}

		d.Set("group_id", group.ID)
		d.Set("zone_id", zoneID)
		d.Set("package_id", pkg.ID)
		d.Set("mode", group.Mode)

		d.SetId(group.ID)

		return []*schema.ResourceData{d}, nil
	}

	return nil, fmt.Errorf("Unable to find WAF Group %s", groupID)
}
