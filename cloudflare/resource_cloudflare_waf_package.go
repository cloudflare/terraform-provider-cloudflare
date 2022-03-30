package cloudflare

import (
	"context"
	"errors"
	"fmt"
	"strings"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareWAFPackage() *schema.Resource {
	return &schema.Resource{
		Schema: resourceCloudflareWAFPackageSchema(),
		Create: resourceCloudflareWAFPackageCreate,
		Read:   resourceCloudflareWAFPackageRead,
		Update: resourceCloudflareWAFPackageUpdate,
		Delete: resourceCloudflareWAFPackageDelete,

		Importer: &schema.ResourceImporter{
			State: resourceCloudflareWAFPackageImport,
		},
	}
}

func resourceCloudflareWAFPackageRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)

	packageID := d.Get("package_id").(string)
	zoneID := d.Get("zone_id").(string)

	pkg, err := client.WAFPackage(context.Background(), zoneID, packageID)
	if err != nil {
		var requestError *cloudflare.RequestError
		if errors.As(err, &requestError) && sliceContainsInt(requestError.ErrorCodes(), 1002) {
			d.SetId("")
			return nil
		}

		return err
	}

	d.Set("sensitivity", pkg.Sensitivity)
	d.Set("action_mode", pkg.ActionMode)
	d.SetId(pkg.ID)

	return nil
}

func resourceCloudflareWAFPackageCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)

	packageID := d.Get("package_id").(string)
	zoneID := d.Get("zone_id").(string)
	sensitivity := d.Get("sensitivity").(string)
	actionMode := d.Get("action_mode").(string)

	pkg, err := client.WAFPackage(context.Background(), zoneID, packageID)
	if err != nil {
		return fmt.Errorf("unable to find WAF Package %s", packageID)
	}

	d.Set("zone_id", zoneID)
	d.Set("package_id", packageID)
	d.Set("sensitivity", sensitivity)
	d.Set("action_mode", actionMode)

	// Set the ID to the package_id parameter passed in from the user.
	// All WAF packages already exist so we already know the package_id.
	//
	// This is a work around as we are not really "creating" a WAF Package,
	// only associating it with our terraform config for future updates.
	d.SetId(packageID)

	if pkg.Sensitivity != sensitivity || pkg.ActionMode != actionMode {
		err = resourceCloudflareWAFPackageUpdate(d, meta)
		if err != nil {
			d.SetId("")
			return err
		}
	}

	return nil
}

func resourceCloudflareWAFPackageDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)

	packageID := d.Get("package_id").(string)
	zoneID := d.Get("zone_id").(string)

	pkg, err := client.WAFPackage(context.Background(), zoneID, packageID)
	if err != nil {
		return err
	}

	// Can't delete WAF Package so instead reset it to default
	schema := resourceCloudflareWAFPackage().Schema
	defaultSensitivity := schema["sensitivity"].Default.(string)
	defaultActionMode := schema["action_mode"].Default.(string)

	if pkg.Sensitivity != defaultSensitivity || pkg.ActionMode != defaultActionMode {
		options := cloudflare.WAFPackageOptions{
			Sensitivity: defaultSensitivity,
			ActionMode:  defaultActionMode,
		}

		_, err = client.UpdateWAFPackage(context.Background(), zoneID, packageID, options)
		if err != nil {
			return err
		}
	}

	return nil
}

func resourceCloudflareWAFPackageUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)

	packageID := d.Get("package_id").(string)
	zoneID := d.Get("zone_id").(string)
	sensitivity := d.Get("sensitivity").(string)
	actionMode := d.Get("action_mode").(string)

	options := cloudflare.WAFPackageOptions{
		Sensitivity: sensitivity,
		ActionMode:  actionMode,
	}

	_, err := client.UpdateWAFPackage(context.Background(), zoneID, packageID, options)
	if err != nil {
		return err
	}

	return nil
}

func resourceCloudflareWAFPackageImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	client := meta.(*cloudflare.API)

	// split the id so we can lookup
	idAttr := strings.SplitN(d.Id(), "/", 2)
	var zoneID string
	var packageID string
	if len(idAttr) == 2 {
		zoneID = idAttr[0]
		packageID = idAttr[1]
	} else {
		return nil, fmt.Errorf("invalid id (\"%s\") specified, should be in format \"zoneID/PackageID\" for import", d.Id())
	}

	pkg, err := client.WAFPackage(context.Background(), zoneID, packageID)
	if err != nil {
		return nil, err
	}

	d.Set("package_id", pkg.ID)
	d.Set("zone_id", zoneID)
	d.Set("sensitivity", pkg.Sensitivity)
	d.Set("action_mode", pkg.ActionMode)

	d.SetId(pkg.ID)

	return []*schema.ResourceData{d}, nil
}
