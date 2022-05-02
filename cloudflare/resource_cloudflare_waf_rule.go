package cloudflare

import (
	"context"
	"errors"
	"fmt"
	"strings"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareWAFRule() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareWAFRuleSchema(),
		CreateContext: resourceCloudflareWAFRuleCreate,
		ReadContext:   resourceCloudflareWAFRuleRead,
		UpdateContext: resourceCloudflareWAFRuleUpdate,
		DeleteContext: resourceCloudflareWAFRuleDelete,

		Importer: &schema.ResourceImporter{
			StateContext: resourceCloudflareWAFRuleImport,
		},
	}
}

func resourceCloudflareWAFRuleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	ruleID := d.Get("rule_id").(string)
	zoneID := d.Get("zone_id").(string)
	packageID := d.Get("package_id").(string)

	rule, err := client.WAFRule(ctx, zoneID, packageID, ruleID)
	if err != nil {
		var requestError *cloudflare.RequestError
		if errors.As(err, &requestError) && (sliceContainsInt(requestError.ErrorCodes(), 1002) || sliceContainsInt(requestError.ErrorCodes(), 1004)) {
			d.SetId("")
			return nil
		}

		return diag.FromErr(err)
	}

	// Only need to set mode as that is the only attribute that could have changed
	d.Set("mode", rule.Mode)
	d.Set("group_id", rule.Group.ID)
	d.SetId(rule.ID)

	return nil
}

func resourceCloudflareWAFRuleCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	ruleID := d.Get("rule_id").(string)
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
		var rule cloudflare.WAFRule

		rule, err = client.WAFRule(ctx, zoneID, pkg.ID, ruleID)
		if err != nil {
			continue
		}

		d.Set("rule_id", rule.ID)
		d.Set("zone_id", zoneID)
		d.Set("group_id", rule.Group.ID)
		d.Set("package_id", pkg.ID)

		if rule.Mode != mode {
			err := resourceCloudflareWAFRuleUpdate(ctx, d, meta)
			if err != nil {
				d.SetId("")
				return err
			}
		}

		return resourceCloudflareWAFRuleRead(ctx, d, meta)
	}

	return diag.FromErr(fmt.Errorf("unable to find WAF Rule %s", ruleID))
}

func resourceCloudflareWAFRuleDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	ruleID := d.Get("rule_id").(string)
	zoneID := d.Get("zone_id").(string)
	packageID := d.Get("package_id").(string)

	rule, err := client.WAFRule(ctx, zoneID, packageID, ruleID)
	if err != nil {
		return diag.FromErr(err)
	}

	// Find the default mode to be used
	defaultMode := "default"
	if !contains(rule.AllowedModes, defaultMode) {
		defaultMode = "on"
	}

	// Can't delete WAF Rule so instead reset it to default
	if rule.Mode != defaultMode {
		_, err = client.UpdateWAFRule(ctx, zoneID, packageID, ruleID, defaultMode)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	return nil
}

func resourceCloudflareWAFRuleUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	ruleID := d.Get("rule_id").(string)
	zoneID := d.Get("zone_id").(string)
	mode := d.Get("mode").(string)
	packageID := d.Get("package_id").(string)

	// We can only update the mode of a WAF Rule
	_, err := client.UpdateWAFRule(ctx, zoneID, packageID, ruleID, mode)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceCloudflareWAFRuleRead(ctx, d, meta)
}

func resourceCloudflareWAFRuleImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	client := meta.(*cloudflare.API)

	// split the id so we can lookup
	idAttr := strings.SplitN(d.Id(), "/", 2)
	var zoneID string
	var WAFID string
	if len(idAttr) == 2 {
		zoneID = idAttr[0]
		WAFID = idAttr[1]
	} else {
		return nil, fmt.Errorf("invalid id (\"%s\") specified, should be in format \"zoneID/WAFID\" for import", d.Id())
	}

	packs, err := client.ListWAFPackages(ctx, zoneID)
	if err != nil {
		return nil, fmt.Errorf("error listing WAF packages: %s", err)
	}

	for _, p := range packs {
		rule, err := client.WAFRule(ctx, zoneID, p.ID, WAFID)
		if err == nil {
			d.Set("rule_id", rule.ID)
			d.Set("zone_id", zoneID)
			d.Set("package_id", rule.PackageID)
			d.Set("group_id", rule.Group.ID)
			d.Set("mode", rule.Mode)

			// The ID is known by the user in advance
			d.SetId(WAFID)
		}
	}

	if d.Id() != WAFID {
		return nil, fmt.Errorf("Unable to find WAF Rule %s", WAFID)
	}

	resourceCloudflareWAFRuleRead(ctx, d, meta)

	return []*schema.ResourceData{d}, nil
}
