package sdkv2provider

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/MakeNowJust/heredoc/v2"
	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareFilter() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareFilterSchema(),
		CreateContext: resourceCloudflareFilterCreate,
		ReadContext:   resourceCloudflareFilterRead,
		UpdateContext: resourceCloudflareFilterUpdate,
		DeleteContext: resourceCloudflareFilterDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceCloudflareFilterImport,
		},
		Description: heredoc.Doc(`
			Filter expressions that can be referenced across multiple features,
			e.g. Firewall Rules. See [what is a filter](https://developers.cloudflare.com/firewall/api/cf-filters/what-is-a-filter/)
			for more details and available fields and operators.
		`),
		DeprecationMessage: heredoc.Doc(fmt.Sprintf(`
			%s resource is in a deprecation phase that will
			last for one year (May 1st, 2024). During this time period, this
			resource is still fully supported but you are strongly advised
			to move to the %s resource. For more information, see
			https://developers.cloudflare.com/waf/reference/migration-guides/firewall-rules-to-custom-rules/#relevant-changes-for-terraform-users.
		`, "`cloudflare_filter`", "`cloudflare_ruleset`")),
	}
}

func resourceCloudflareFilterCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	var err error

	var newFilter cloudflare.FilterCreateParams

	if paused, ok := d.GetOk("paused"); ok {
		newFilter.Paused = paused.(bool)
	}

	if description, ok := d.GetOk("description"); ok {
		newFilter.Description = description.(string)
	}

	if expression, ok := d.GetOk("expression"); ok {
		newFilter.Expression = expression.(string)
	}

	if ref, ok := d.GetOk("ref"); ok {
		newFilter.Ref = ref.(string)
	}

	tflog.Debug(ctx, fmt.Sprintf("Creating Cloudflare Filter from struct: %+v", newFilter))

	r, err := client.CreateFilters(ctx, cloudflare.ZoneIdentifier(zoneID), []cloudflare.FilterCreateParams{newFilter})

	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating Filter for zone %q: %w", zoneID, err))
	}

	if len(r) == 0 {
		return diag.FromErr(fmt.Errorf("failed to find id in Create response; resource was empty"))
	}

	d.SetId(r[0].ID)

	tflog.Info(ctx, fmt.Sprintf("Cloudflare Filter ID: %s", d.Id()))

	return resourceCloudflareFilterRead(ctx, d, meta)
}

func resourceCloudflareFilterRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	tflog.Debug(ctx, fmt.Sprintf("Getting a Filter record for zone %q, id %s", zoneID, d.Id()))
	filter, err := client.Filter(ctx, cloudflare.ZoneIdentifier(zoneID), d.Id())

	tflog.Debug(ctx, fmt.Sprintf("filter: %#v", filter))
	tflog.Debug(ctx, fmt.Sprintf("filter error: %#v", err))

	if err != nil {
		var notFoundError *cloudflare.NotFoundError
		if errors.As(err, &notFoundError) {
			tflog.Info(ctx, fmt.Sprintf("Filter %s no longer exists", d.Id()))
			d.SetId("")
			return nil
		}
		return diag.FromErr(fmt.Errorf("error finding Filter %q: %w", d.Id(), err))
	}

	tflog.Debug(ctx, fmt.Sprintf("Cloudflare Filter read configuration: %#v", filter))

	d.Set("paused", filter.Paused)
	d.Set("description", filter.Description)
	d.Set("expression", filter.Expression)
	d.Set("ref", filter.Ref)

	return nil
}

func resourceCloudflareFilterUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	var newFilter cloudflare.FilterUpdateParams
	newFilter.ID = d.Id()

	if paused, ok := d.GetOk("paused"); ok {
		newFilter.Paused = paused.(bool)
	}

	if description, ok := d.GetOk("description"); ok {
		newFilter.Description = description.(string)
	}

	if expression, ok := d.GetOk("expression"); ok {
		newFilter.Expression = expression.(string)
	}

	if ref, ok := d.GetOk("ref"); ok {
		newFilter.Ref = ref.(string)
	}

	tflog.Debug(ctx, fmt.Sprintf("Updating Cloudflare Filter from struct: %+v", newFilter))

	r, err := client.UpdateFilter(ctx, cloudflare.ZoneIdentifier(zoneID), newFilter)

	if err != nil {
		return diag.FromErr(fmt.Errorf("error updating Filter for zone %q: %w", zoneID, err))
	}

	if r.ID == "" {
		return diag.FromErr(fmt.Errorf("failed to find id in Update response; resource was empty"))
	}

	return resourceCloudflareFilterRead(ctx, d, meta)
}

func resourceCloudflareFilterDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	tflog.Info(ctx, fmt.Sprintf("Deleting Cloudflare Filter: id %s for zone %s", d.Id(), zoneID))

	err := client.DeleteFilter(ctx, cloudflare.ZoneIdentifier(zoneID), d.Id())

	if err != nil {
		return diag.FromErr(fmt.Errorf("error deleting Cloudflare Filter: %w", err))
	}

	return nil
}

func resourceCloudflareFilterImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	// split the id so we can lookup
	idAttr := strings.SplitN(d.Id(), "/", 2)

	if len(idAttr) != 2 {
		return nil, fmt.Errorf("invalid id (\"%s\") specified, should be in format \"zoneID/filterID\"", d.Id())
	}

	zoneID, filterID := idAttr[0], idAttr[1]

	tflog.Debug(ctx, fmt.Sprintf("Importing Cloudflare Filter: id %s for zone %s", filterID, zoneID))

	d.Set(consts.ZoneIDSchemaKey, zoneID)
	d.SetId(filterID)

	resourceCloudflareFilterRead(ctx, d, meta)

	return []*schema.ResourceData{d}, nil
}
