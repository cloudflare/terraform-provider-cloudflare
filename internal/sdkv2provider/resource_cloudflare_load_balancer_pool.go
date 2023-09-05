package sdkv2provider

import (
	"context"
	"fmt"
	"math"
	"strings"

	"time"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
)

func resourceCloudflareLoadBalancerPool() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareLoadBalancerPoolSchema(),
		CreateContext: resourceCloudflareLoadBalancerPoolCreate,
		UpdateContext: resourceCloudflareLoadBalancerPoolUpdate,
		ReadContext:   resourceCloudflareLoadBalancerPoolRead,
		DeleteContext: resourceCloudflareLoadBalancerPoolDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceCloudflareLoadBalancerPoolImport,
		},
		Description: heredoc.Doc(`
			Provides a Cloudflare Load Balancer pool resource. This provides a
			pool of origins that can be used by a Cloudflare Load Balancer.
		`),
	}
}

func resourceCloudflareLoadBalancerPoolCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	loadBalancerPool := cloudflare.LoadBalancerPool{
		Name:           d.Get("name").(string),
		Origins:        expandLoadBalancerOrigins(d.Get("origins").(*schema.Set)),
		Enabled:        d.Get("enabled").(bool),
		MinimumOrigins: cloudflare.IntPtr(d.Get("minimum_origins").(int)),
	}

	if lat, ok := d.GetOk("latitude"); ok {
		f := float32(lat.(float64))
		loadBalancerPool.Latitude = &f
	}
	if long, ok := d.GetOk("longitude"); ok {
		f := float32(long.(float64))
		loadBalancerPool.Longitude = &f
	}

	if checkRegions, ok := d.GetOk("check_regions"); ok {
		loadBalancerPool.CheckRegions = expandInterfaceToStringList(checkRegions.(*schema.Set).List())
	}

	if description, ok := d.GetOk("description"); ok {
		loadBalancerPool.Description = description.(string)
	}

	if monitor, ok := d.GetOk("monitor"); ok {
		loadBalancerPool.Monitor = monitor.(string)
	}

	if shed, ok := d.GetOk("load_shedding"); ok {
		loadBalancerPool.LoadShedding = expandLoadBalancerLoadShedding(shed.(*schema.Set))
	}

	if steering, ok := d.GetOk("origin_steering"); ok {
		loadBalancerPool.OriginSteering = expandLoadBalancerOriginSteering(steering.(*schema.Set))
	}

	if notificationEmail, ok := d.GetOk("notification_email"); ok {
		loadBalancerPool.NotificationEmail = notificationEmail.(string)
	}

	tflog.Debug(ctx, fmt.Sprintf("Creating Cloudflare Load Balancer Pool from struct: %+v", loadBalancerPool))

	accountID := d.Get(consts.AccountIDSchemaKey).(string)

	r, err := client.CreateLoadBalancerPool(ctx, cloudflare.AccountIdentifier(accountID), cloudflare.CreateLoadBalancerPoolParams{LoadBalancerPool: loadBalancerPool})
	if err != nil {
		return diag.FromErr(errors.Wrap(err, "error creating load balancer pool"))
	}

	if r.ID == "" {
		return diag.FromErr(fmt.Errorf("cailed to find id in create response; resource was empty"))
	}

	d.SetId(r.ID)

	tflog.Info(ctx, fmt.Sprintf("New Cloudflare Load Balancer Pool created with  ID: %s", d.Id()))

	return resourceCloudflareLoadBalancerPoolRead(ctx, d, meta)
}

func resourceCloudflareLoadBalancerPoolUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	loadBalancerPool := cloudflare.LoadBalancerPool{
		ID:             d.Id(),
		Name:           d.Get("name").(string),
		Origins:        expandLoadBalancerOrigins(d.Get("origins").(*schema.Set)),
		Enabled:        d.Get("enabled").(bool),
		MinimumOrigins: cloudflare.IntPtr(d.Get("minimum_origins").(int)),
	}

	if lat, ok := d.GetOk("latitude"); ok {
		f := float32(lat.(float64))
		loadBalancerPool.Latitude = &f
	}
	if long, ok := d.GetOk("longitude"); ok {
		f := float32(long.(float64))
		loadBalancerPool.Longitude = &f
	}

	if checkRegions, ok := d.GetOk("check_regions"); ok {
		loadBalancerPool.CheckRegions = expandInterfaceToStringList(checkRegions.(*schema.Set).List())
	}

	if description, ok := d.GetOk("description"); ok {
		loadBalancerPool.Description = description.(string)
	}

	if monitor, ok := d.GetOk("monitor"); ok {
		loadBalancerPool.Monitor = monitor.(string)
	}

	if shed, ok := d.GetOk("load_shedding"); ok {
		loadBalancerPool.LoadShedding = expandLoadBalancerLoadShedding(shed.(*schema.Set))
	}

	if steering, ok := d.GetOk("origin_steering"); ok {
		loadBalancerPool.OriginSteering = expandLoadBalancerOriginSteering(steering.(*schema.Set))
	}

	if notificationEmail, ok := d.GetOk("notification_email"); ok {
		loadBalancerPool.NotificationEmail = notificationEmail.(string)
	}

	tflog.Debug(ctx, fmt.Sprintf("Updating Cloudflare Load Balancer Pool from struct: %+v", loadBalancerPool))

	accountID := d.Get(consts.AccountIDSchemaKey).(string)
	_, err := client.UpdateLoadBalancerPool(ctx, cloudflare.AccountIdentifier(accountID), cloudflare.UpdateLoadBalancerPoolParams{LoadBalancer: loadBalancerPool})
	if err != nil {
		return diag.FromErr(errors.Wrap(err, "error updating load balancer pool"))
	}

	return resourceCloudflareLoadBalancerPoolRead(ctx, d, meta)
}

func expandLoadBalancerPoolHeader(cfgSet interface{}) map[string][]string {
	header := make(map[string][]string)
	cfgList := cfgSet.(*schema.Set).List()
	for _, item := range cfgList {
		cfg := item.(map[string]interface{})
		header[cfg["header"].(string)] = expandInterfaceToStringList(cfg["values"].(*schema.Set).List())
	}
	return header
}

func flattenLoadBalancerPoolHeader(header map[string][]string) *schema.Set {
	flattened := make([]interface{}, 0)
	for k, v := range header {
		cfg := map[string]interface{}{
			"header": k,
			"values": schema.NewSet(schema.HashString, flattenStringList(v)),
		}
		flattened = append(flattened, cfg)
	}
	return schema.NewSet(HashByMapKey("header"), flattened)
}

func expandLoadBalancerLoadShedding(s *schema.Set) *cloudflare.LoadBalancerLoadShedding {
	if s == nil {
		return nil
	}
	for _, iface := range s.List() {
		o := iface.(map[string]interface{})
		return &cloudflare.LoadBalancerLoadShedding{
			DefaultPercent: float32(o["default_percent"].(float64)),
			DefaultPolicy:  o["default_policy"].(string),
			SessionPercent: float32(o["session_percent"].(float64)),
			SessionPolicy:  o["session_policy"].(string),
		}
	}
	return nil
}

func expandLoadBalancerOriginSteering(s *schema.Set) *cloudflare.LoadBalancerOriginSteering {
	if s == nil {
		return nil
	}
	for _, iface := range s.List() {
		o := iface.(map[string]interface{})
		return &cloudflare.LoadBalancerOriginSteering{
			Policy: o["policy"].(string),
		}
	}
	return nil
}

func expandLoadBalancerOrigins(originSet *schema.Set) (origins []cloudflare.LoadBalancerOrigin) {
	for _, iface := range originSet.List() {
		o := iface.(map[string]interface{})
		origin := cloudflare.LoadBalancerOrigin{
			Name:    o["name"].(string),
			Address: o["address"].(string),
			Enabled: o["enabled"].(bool),
			Weight:  o["weight"].(float64),
		}

		if header, ok := o["header"]; ok {
			origin.Header = expandLoadBalancerPoolHeader(header)
		}

		origins = append(origins, origin)
	}
	return
}

func resourceCloudflareLoadBalancerPoolRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	accountID := d.Get(consts.AccountIDSchemaKey).(string)

	loadBalancerPool, err := client.GetLoadBalancerPool(ctx, cloudflare.AccountIdentifier(accountID), d.Id())
	if err != nil {
		var notFoundError *cloudflare.NotFoundError
		if errors.As(err, &notFoundError) {
			tflog.Info(ctx, fmt.Sprintf("Load balancer pool %s no longer exists", d.Id()))
			d.SetId("")
			return nil
		} else {
			return diag.FromErr(errors.Wrap(err,
				fmt.Sprintf("Error reading load balancer pool from API for resource %s ", d.Id())))
		}
	}
	tflog.Debug(ctx, fmt.Sprintf("Read Cloudflare Load Balancer Pool from API as struct: %+v", loadBalancerPool))

	d.Set("name", loadBalancerPool.Name)
	d.Set("enabled", loadBalancerPool.Enabled)
	d.Set("minimum_origins", loadBalancerPool.MinimumOrigins)
	d.Set("description", loadBalancerPool.Description)
	d.Set("monitor", loadBalancerPool.Monitor)
	d.Set("notification_email", loadBalancerPool.NotificationEmail)
	d.Set("created_on", loadBalancerPool.CreatedOn.Format(time.RFC3339Nano))
	d.Set("modified_on", loadBalancerPool.ModifiedOn.Format(time.RFC3339Nano))

	if lat := loadBalancerPool.Latitude; lat != nil {
		f := math.Round(float64(*lat)*10000) / 10000 // set precision
		d.Set("latitude", &f)
	}
	if long := loadBalancerPool.Longitude; long != nil {
		f := math.Round(float64(*long)*10000) / 10000 // set precision
		d.Set("longitude", &f)
	}

	if err := d.Set("origins", flattenLoadBalancerOrigins(d, loadBalancerPool.Origins)); err != nil {
		tflog.Warn(ctx, fmt.Sprintf("Error setting origins on load balancer pool %q: %s", d.Id(), err))
	}

	if err := d.Set("load_shedding", flattenLoadBalancerLoadShedding(loadBalancerPool.LoadShedding)); err != nil {
		tflog.Warn(ctx, fmt.Sprintf("Error setting load_shedding on load balancer pool %q: %s", d.Id(), err))
	}

	if err := d.Set("origin_steering", flattenLoadBalancerOriginSteering(loadBalancerPool.OriginSteering)); err != nil {
		tflog.Warn(ctx, fmt.Sprintf("Error setting origin_steering on load balancer pool %q: %s", d.Id(), err))
	}

	if err := d.Set("check_regions", schema.NewSet(schema.HashString, flattenStringList(loadBalancerPool.CheckRegions))); err != nil {
		tflog.Warn(ctx, fmt.Sprintf("Error setting check_regions on load balancer pool %q: %s", d.Id(), err))
	}

	return nil
}

func flattenLoadBalancerLoadShedding(ls *cloudflare.LoadBalancerLoadShedding) *schema.Set {
	if ls == nil {
		return nil
	}
	return schema.NewSet(schema.HashResource(loadShedElem), []interface{}{map[string]interface{}{
		"default_percent": math.Round(float64(ls.DefaultPercent)*1000) / 1000,
		"default_policy":  ls.DefaultPolicy,
		"session_percent": math.Round(float64(ls.SessionPercent)*1000) / 1000,
		"session_policy":  ls.SessionPolicy,
	}})
}

func flattenLoadBalancerOriginSteering(os *cloudflare.LoadBalancerOriginSteering) *schema.Set {
	if os == nil {
		return nil
	}
	return schema.NewSet(schema.HashResource(originSteeringElem), []interface{}{map[string]interface{}{
		"policy": os.Policy,
	}})
}

func flattenLoadBalancerOrigins(d *schema.ResourceData, origins []cloudflare.LoadBalancerOrigin) *schema.Set {
	flattened := make([]interface{}, 0)
	for _, o := range origins {
		cfg := map[string]interface{}{
			"name":    o.Name,
			"address": o.Address,
			"enabled": o.Enabled,
			"weight":  o.Weight,
			"header":  flattenLoadBalancerPoolHeader(o.Header),
		}

		flattened = append(flattened, cfg)
	}
	return schema.NewSet(schema.HashResource(originsElem), flattened)
}

func resourceCloudflareLoadBalancerPoolDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	tflog.Info(ctx, fmt.Sprintf("Deleting Cloudflare Load Balancer Pool: %s ", d.Id()))

	accountID := d.Get(consts.AccountIDSchemaKey).(string)
	err := client.DeleteLoadBalancerPool(ctx, cloudflare.AccountIdentifier(accountID), d.Id())
	if err != nil {
		return diag.FromErr(errors.Wrap(err, "error deleting Cloudflare Load Balancer Pool"))
	}

	return nil
}

func resourceCloudflareLoadBalancerPoolImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	// split the id so we can lookup
	idAttr := strings.SplitN(d.Id(), "/", 2)
	var accountID string
	var lbPoolID string
	if len(idAttr) == 2 {
		accountID = idAttr[0]
		lbPoolID = idAttr[1]
	} else {
		return nil, fmt.Errorf("invalid id (\"%s\") specified, should be in format \"accountID/loadBalancerPoolID\"", d.Id())
	}

	d.Set(consts.AccountIDSchemaKey, accountID)
	d.SetId(lbPoolID)

	resourceCloudflareLoadBalancerPoolRead(ctx, d, meta)

	return []*schema.ResourceData{d}, nil
}
