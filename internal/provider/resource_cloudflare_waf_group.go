package provider

import (
	"context"
	"errors"
	"fmt"
	"strings"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareWAFGroup() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareWAFGroupSchema(),
		CreateContext: resourceCloudflareWAFGroupCreate,
		ReadContext:   resourceCloudflareWAFGroupRead,
		UpdateContext: resourceCloudflareWAFGroupUpdate,
		DeleteContext: resourceCloudflareWAFGroupDelete,

		Importer: &schema.ResourceImporter{
			StateContext: resourceCloudflareWAFGroupImport,
		},
	}
}

func resourceCloudflareWAFGroupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	groupID := d.Get("group_id").(string)
	zoneID := d.Get("zone_id").(string)
	packageID := d.Get("package_id").(string)

	group, err := client.WAFGroup(ctx, zoneID, packageID, groupID)
	if err != nil {
		var requestError *cloudflare.RequestError
		if errors.As(err, &requestError) && (sliceContainsInt(requestError.ErrorCodes(), 1002) || sliceContainsInt(requestError.ErrorCodes(), 1003)) {
			d.SetId("")
			return nil
		}

		return diag.FromErr(err)
	}

	// Only need to set mode as that is the only attribute that could have changed
	d.Set("mode", group.Mode)
	d.SetId(group.ID)

	return nil
}

func resourceCloudflareWAFGroupCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	groupID := d.Get("group_id").(string)
	zoneID := d.Get("zone_id").(string)
	packageID := d.Get("package_id").(string)
	mode := d.Get("mode").(string)

	// If no package ID is given try to resolve it
	var pkgList []cloudflare.WAFPackage
	if packageID == "" {
		var err error
		pkgList, err = client.ListWAFPackages(ctx, zoneID)
		if err != nil {
			return diag.FromErr(err)
		}
	} else {
		pkgList = append(pkgList, cloudflare.WAFPackage{ID: packageID})
	}

	for _, pkg := range pkgList {
		var err error
		var group cloudflare.WAFGroup

		group, err = client.WAFGroup(ctx, zoneID, pkg.ID, groupID)
		if err != nil {
			continue
		}

		d.Set("group_id", group.ID)
		d.Set("zone_id", zoneID)
		d.Set("package_id", pkg.ID)

		if group.Mode != mode {
			err := resourceCloudflareWAFGroupUpdate(ctx, d, meta)
			if err != nil {
				d.SetId("")
				return err
			}
		}

		return resourceCloudflareWAFGroupRead(ctx, d, meta)
	}

	return diag.FromErr(fmt.Errorf("unable to find WAF Group %s", groupID))
}

func resourceCloudflareWAFGroupDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	groupID := d.Get("group_id").(string)
	zoneID := d.Get("zone_id").(string)
	packageID := d.Get("package_id").(string)

	group, err := client.WAFGroup(ctx, zoneID, packageID, groupID)
	if err != nil {
		return diag.FromErr(err)
	}

	// Can't delete WAF Group so instead reset it to default
	schema := resourceCloudflareWAFGroup().Schema
	defaultMode := schema["mode"].Default.(string)

	if group.Mode != defaultMode {
		_, err = client.UpdateWAFGroup(ctx, zoneID, packageID, groupID, defaultMode)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	return nil
}

func resourceCloudflareWAFGroupUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	groupID := d.Get("group_id").(string)
	zoneID := d.Get("zone_id").(string)
	mode := d.Get("mode").(string)
	packageID := d.Get("package_id").(string)

	// We can only update the mode of a WAF Group
	_, err := client.UpdateWAFGroup(ctx, zoneID, packageID, groupID, mode)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceCloudflareWAFGroupRead(ctx, d, meta)
}

func resourceCloudflareWAFGroupImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
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

	pkgList, err := client.ListWAFPackages(ctx, zoneID)
	if err != nil {
		return nil, fmt.Errorf("error listing WAF packages: %w", err)
	}

	for _, pkg := range pkgList {
		group, err := client.WAFGroup(ctx, zoneID, pkg.ID, groupID)
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
