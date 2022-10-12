package provider

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"time"

	"github.com/MakeNowJust/heredoc/v2"
	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
)

func resourceCloudflareLoadBalancer() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCloudflareLoadBalancerCreate,
		ReadContext:   resourceCloudflareLoadBalancerRead,
		UpdateContext: resourceCloudflareLoadBalancerUpdate,
		DeleteContext: resourceCloudflareLoadBalancerDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceCloudflareLoadBalancerImport,
		},

		SchemaVersion: 1,

		Schema: resourceCloudflareLoadBalancerSchema(),

		StateUpgraders: []schema.StateUpgrader{
			{
				Type:    resourceCloudflareLoadBalancerV0().CoreConfigSchema().ImpliedType(),
				Upgrade: resourceCloudflareLoadBalancerStateUpgradeV1,
				Version: 0,
			},
		},

		Description: heredoc.Doc(`
			Provides a Cloudflare Load Balancer resource. This sits in front of
			a number of defined pools of origins and provides various options
			for geographically-aware load balancing. Note that the load balancing
			feature must be enabled in your Cloudflare account before you can use
			this resource.
		`),
	}
}

func resourceCloudflareLoadBalancerCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	zoneID := d.Get("zone_id").(string)

	enabled := d.Get("enabled").(bool)
	newLoadBalancer := cloudflare.LoadBalancer{
		Name:           d.Get("name").(string),
		FallbackPool:   d.Get("fallback_pool_id").(string),
		DefaultPools:   expandInterfaceToStringList(d.Get("default_pool_ids")),
		Proxied:        d.Get("proxied").(bool),
		Enabled:        &enabled,
		TTL:            d.Get("ttl").(int),
		SteeringPolicy: d.Get("steering_policy").(string),
		Persistence:    d.Get("session_affinity").(string),
	}

	if description, ok := d.GetOk("description"); ok {
		newLoadBalancer.Description = description.(string)
	}

	if ttl, ok := d.GetOk("ttl"); ok {
		newLoadBalancer.TTL = ttl.(int)
	}

	if regionPools, ok := d.GetOk("region_pools"); ok {
		expandedRegionPools, err := expandGeoPools(regionPools, "region")
		if err != nil {
			return diag.FromErr(err)
		}
		newLoadBalancer.RegionPools = expandedRegionPools
	}

	if countryPools, ok := d.GetOk("country_pools"); ok {
		expandedCountryPools, err := expandGeoPools(countryPools, "country")
		if err != nil {
			return diag.FromErr(err)
		}
		newLoadBalancer.CountryPools = expandedCountryPools
	}

	if popPools, ok := d.GetOk("pop_pools"); ok {
		expandedPopPools, err := expandGeoPools(popPools, "pop")
		if err != nil {
			return diag.FromErr(err)
		}
		newLoadBalancer.PopPools = expandedPopPools
	}

	if sessionAffinityTTL, ok := d.GetOk("session_affinity_ttl"); ok {
		newLoadBalancer.PersistenceTTL = sessionAffinityTTL.(int)
	}

	if sessionAffinityAttrs, ok := d.GetOk("session_affinity_attributes"); ok {
		sessionAffinityAttributes, err := expandSessionAffinityAttrs(sessionAffinityAttrs)
		if err != nil {
			return diag.FromErr(err)
		}
		newLoadBalancer.SessionAffinityAttributes = sessionAffinityAttributes
	}

	if adaptiveRouting, ok := d.GetOk("adaptive_routing"); ok {
		newLoadBalancer.AdaptiveRouting = expandAdaptiveRouting(adaptiveRouting)
	}

	if locationStrategy, ok := d.GetOk("location_strategy"); ok {
		newLoadBalancer.LocationStrategy = expandLocationStrategy(locationStrategy)
	}

	if randomSteering, ok := d.GetOk("random_steering"); ok {
		newLoadBalancer.RandomSteering = expandRandomSteering(randomSteering)
	}

	if rules, ok := d.GetOk("rules"); ok {
		v, err := expandRules(rules)
		if err != nil {
			return diag.FromErr(err)
		}
		newLoadBalancer.Rules = v
	}

	tflog.Info(ctx, fmt.Sprintf("Creating Cloudflare Load Balancer from struct: %+v", newLoadBalancer))

	r, err := client.CreateLoadBalancer(ctx, cloudflare.ZoneIdentifier(zoneID), cloudflare.CreateLoadBalancerParams{LoadBalancer: newLoadBalancer})
	if err != nil {
		return diag.FromErr(errors.Wrap(err, "error creating load balancer for zone"))
	}

	if r.ID == "" {
		return diag.FromErr(fmt.Errorf("failed to find id in Create response; resource was empty"))
	}

	d.SetId(r.ID)

	tflog.Info(ctx, fmt.Sprintf("Cloudflare Load Balancer ID: %s", d.Id()))

	return resourceCloudflareLoadBalancerRead(ctx, d, meta)
}

func resourceCloudflareLoadBalancerUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)

	enabled := d.Get("enabled").(bool)
	loadBalancer := cloudflare.LoadBalancer{
		ID:             d.Id(),
		Name:           d.Get("name").(string),
		FallbackPool:   d.Get("fallback_pool_id").(string),
		DefaultPools:   expandInterfaceToStringList(d.Get("default_pool_ids")),
		Proxied:        d.Get("proxied").(bool),
		Enabled:        &enabled,
		TTL:            d.Get("ttl").(int),
		SteeringPolicy: d.Get("steering_policy").(string),
		Persistence:    d.Get("session_affinity").(string),
	}

	if description, ok := d.GetOk("description"); ok {
		loadBalancer.Description = description.(string)
	}

	if regionPools, ok := d.GetOk("region_pools"); ok {
		expandedRegionPools, err := expandGeoPools(regionPools, "region")
		if err != nil {
			return diag.FromErr(err)
		}
		loadBalancer.RegionPools = expandedRegionPools
	}

	if countryPools, ok := d.GetOk("country_pools"); ok {
		expandedCountryPools, err := expandGeoPools(countryPools, "country")
		if err != nil {
			return diag.FromErr(err)
		}
		loadBalancer.CountryPools = expandedCountryPools
	}

	if popPools, ok := d.GetOk("pop_pools"); ok {
		expandedPopPools, err := expandGeoPools(popPools, "pop")
		if err != nil {
			return diag.FromErr(err)
		}
		loadBalancer.PopPools = expandedPopPools
	}

	if sessionAffinityTTL, ok := d.GetOk("session_affinity_ttl"); ok {
		loadBalancer.PersistenceTTL = sessionAffinityTTL.(int)
	}

	if sessionAffinityAttrs, ok := d.GetOk("session_affinity_attributes"); ok {
		sessionAffinityAttributes, err := expandSessionAffinityAttrs(sessionAffinityAttrs)
		if err != nil {
			return diag.FromErr(err)
		}
		loadBalancer.SessionAffinityAttributes = sessionAffinityAttributes
	}

	if adaptiveRouting, ok := d.GetOk("adaptive_routing"); ok {
		loadBalancer.AdaptiveRouting = expandAdaptiveRouting(adaptiveRouting)
	}

	if locationStrategy, ok := d.GetOk("location_strategy"); ok {
		loadBalancer.LocationStrategy = expandLocationStrategy(locationStrategy)
	}

	if randomSteering, ok := d.GetOk("random_steering"); ok {
		loadBalancer.RandomSteering = expandRandomSteering(randomSteering)
	}

	if rules, ok := d.GetOk("rules"); ok {
		v, err := expandRules(rules)
		if err != nil {
			return diag.FromErr(err)
		}
		loadBalancer.Rules = v
	}

	tflog.Info(ctx, fmt.Sprintf("Updating Cloudflare Load Balancer from struct: %+v", loadBalancer))

	_, err := client.UpdateLoadBalancer(ctx, cloudflare.ZoneIdentifier(zoneID), cloudflare.UpdateLoadBalancerParams{LoadBalancer: loadBalancer})
	if err != nil {
		return diag.FromErr(errors.Wrap(err, "error updating load balancer for zone"))
	}

	return resourceCloudflareLoadBalancerRead(ctx, d, meta)
}

func expandGeoPools(pool interface{}, geoType string) (map[string][]string, error) {
	cfg := pool.(*schema.Set).List()
	expanded := make(map[string][]string)
	for _, v := range cfg {
		locationConfig := v.(map[string]interface{})
		// lists are of type interface{} by default
		location := locationConfig[geoType].(string)
		if _, present := expanded[location]; !present {
			expanded[location] = expandInterfaceToStringList(locationConfig["pool_ids"])
		} else {
			return nil, fmt.Errorf("duplicate entry specified for %s pool in location %q. each location must only be specified once", geoType, location)
		}
	}
	return expanded, nil
}

func resourceCloudflareLoadBalancerRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)
	loadBalancerID := d.Id()

	loadBalancer, err := client.GetLoadBalancer(ctx, cloudflare.ZoneIdentifier(zoneID), loadBalancerID)
	if err != nil {
		var notFoundError *cloudflare.NotFoundError
		if errors.As(err, &notFoundError) {
			tflog.Info(ctx, fmt.Sprintf("Load balancer %s in zone %s not found", loadBalancerID, zoneID))
			d.SetId("")
			return nil
		}
		return diag.FromErr(errors.Wrap(err,
			fmt.Sprintf("Error reading load balancer resource from API for resource %s in zone %s", loadBalancerID, zoneID)))
	}

	d.Set("name", loadBalancer.Name)
	d.Set("fallback_pool_id", loadBalancer.FallbackPool)
	d.Set("proxied", loadBalancer.Proxied)
	d.Set("enabled", loadBalancer.Enabled)
	d.Set("description", loadBalancer.Description)
	d.Set("ttl", loadBalancer.TTL)
	d.Set("steering_policy", loadBalancer.SteeringPolicy)
	d.Set("session_affinity", loadBalancer.Persistence)

	d.Set("created_on", loadBalancer.CreatedOn.Format(time.RFC3339Nano))
	d.Set("modified_on", loadBalancer.ModifiedOn.Format(time.RFC3339Nano))

	if _, sessionAffinityAttrsOk := d.GetOk("session_affinity_attributes"); sessionAffinityAttrsOk {
		if err := d.Set("session_affinity_attributes", flattenSessionAffinityAttrs(loadBalancer.SessionAffinityAttributes)); err != nil {
			return diag.FromErr(fmt.Errorf("failed to set session_affinity_attributes: %w", err))
		}
	}

	if _, adaptiveRoutingOk := d.GetOk("adaptive_routing"); adaptiveRoutingOk {
		if err := d.Set("adaptive_routing", flattenAdaptiveRouting(loadBalancer.AdaptiveRouting)); err != nil {
			return diag.FromErr(fmt.Errorf("failed to set adaptive_routing: %w", err))
		}
	}

	if _, locationStrategyOk := d.GetOk("location_strategy"); locationStrategyOk {
		if err := d.Set("location_strategy", flattenLocationStrategy(loadBalancer.LocationStrategy)); err != nil {
			return diag.FromErr(fmt.Errorf("failed to set location_strategy: %w", err))
		}
	}

	if _, randomSteeringOk := d.GetOk("random_steering"); randomSteeringOk {
		if err := d.Set("random_steering", flattenRandomSteering(loadBalancer.RandomSteering)); err != nil {
			return diag.FromErr(fmt.Errorf("failed to set random_steering: %w", err))
		}
	}

	if len(loadBalancer.Rules) > 0 {
		fr, err := flattenRules(d, loadBalancer.Rules)
		if err != nil {
			return diag.FromErr(fmt.Errorf("failed to flatten rules: %w", err))
		}
		if err := d.Set("rules", fr); err != nil {
			return diag.FromErr(fmt.Errorf("failed to set rules: %w\n %v", err, fr))
		}
	}

	if err := d.Set("default_pool_ids", loadBalancer.DefaultPools); err != nil {
		tflog.Warn(ctx, fmt.Sprintf("Error setting default_pool_ids on load balancer %q: %s", d.Id(), err))
	}

	if err := d.Set("pop_pools", flattenGeoPools(loadBalancer.PopPools, "pop", loadBalancerLocalPoolElems)); err != nil {
		tflog.Warn(ctx, fmt.Sprintf("Error setting pop_pools on load balancer %q: %s", d.Id(), err))
	}

	if err := d.Set("country_pools", flattenGeoPools(loadBalancer.CountryPools, "country", loadBalancerLocalPoolElems)); err != nil {
		tflog.Warn(ctx, fmt.Sprintf("Error setting country_pools on load balancer %q: %s", d.Id(), err))
	}

	if err := d.Set("region_pools", flattenGeoPools(loadBalancer.RegionPools, "region", loadBalancerLocalPoolElems)); err != nil {
		tflog.Warn(ctx, fmt.Sprintf("Error setting region_pools on load balancer %q: %s", d.Id(), err))
	}

	if loadBalancer.PersistenceTTL != 0 {
		d.Set("session_affinity_ttl", loadBalancer.PersistenceTTL)
	}

	return nil
}

func flattenGeoPools(pools map[string][]string, geoType string, hashResourceMap map[string]*schema.Resource) *schema.Set {
	flattened := make([]interface{}, 0)
	for k, v := range pools {
		geoConf := map[string]interface{}{
			geoType:    k,
			"pool_ids": flattenStringList(v),
		}
		flattened = append(flattened, geoConf)
	}
	return schema.NewSet(schema.HashResource(hashResourceMap[geoType]), flattened)
}

func flattenSessionAffinityAttrs(attrs *cloudflare.SessionAffinityAttributes) map[string]interface{} {
	return map[string]interface{}{
		"drain_duration":         strconv.Itoa(attrs.DrainDuration),
		"samesite":               attrs.SameSite,
		"secure":                 attrs.Secure,
		"zero_downtime_failover": attrs.ZeroDowntimeFailover,
	}
}

func flattenAdaptiveRouting(properties *cloudflare.AdaptiveRouting) *schema.Set {
	flattened := []interface{}{
		map[string]interface{}{
			"failover_across_pools": bool(properties.FailoverAcrossPools != nil && *properties.FailoverAcrossPools),
		},
	}
	return schema.NewSet(schema.HashResource(loadBalancerAdaptiveRoutingElem), flattened)
}

func flattenLocationStrategy(properties *cloudflare.LocationStrategy) *schema.Set {
	flattened := []interface{}{
		map[string]interface{}{
			"prefer_ecs": properties.PreferECS,
			"mode":       properties.Mode,
		},
	}
	return schema.NewSet(schema.HashResource(loadBalancerLocationStrategyElem), flattened)
}

func flattenRandomSteering(properties *cloudflare.RandomSteering) *schema.Set {
	poolWeights := make(map[string]interface{})
	for poolID, poolWeight := range properties.PoolWeights {
		poolWeights[poolID] = poolWeight
	}
	flattened := []interface{}{
		map[string]interface{}{
			"pool_weights":   poolWeights,
			"default_weight": properties.DefaultWeight,
		},
	}
	return schema.NewSet(schema.HashResource(loadBalancerRandomSteeringElem), flattened)
}

func resourceCloudflareLoadBalancerDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)
	loadBalancerID := d.Id()

	tflog.Info(ctx, fmt.Sprintf("Deleting Cloudflare Load Balancer: %s in zone: %s", loadBalancerID, zoneID))

	err := client.DeleteLoadBalancer(ctx, cloudflare.ZoneIdentifier(zoneID), loadBalancerID)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error deleting Cloudflare Load Balancer: %w", err))
	}

	return nil
}

func resourceCloudflareLoadBalancerImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	// split the id so we can lookup
	idAttr := strings.SplitN(d.Id(), "/", 2)
	var zoneID string
	var loadBalancerID string
	if len(idAttr) == 2 {
		zoneID = idAttr[0]
		loadBalancerID = idAttr[1]
	} else {
		return nil, fmt.Errorf("invalid id (\"%s\") specified, should be in format \"zoneID/loadBalancerID\"", d.Id())
	}

	d.Set("zone_id", zoneID)
	d.SetId(loadBalancerID)

	resourceCloudflareLoadBalancerRead(ctx, d, meta)

	return []*schema.ResourceData{d}, nil
}

func flattenRules(d *schema.ResourceData, rules []*cloudflare.LoadBalancerRule) (interface{}, error) {
	if len(rules) == 0 {
		return nil, nil
	}

	cfResources := []map[string]interface{}{}
	for idx, r := range rules {
		m := map[string]interface{}{
			"name":      r.Name,
			"condition": r.Condition,
			"disabled":  r.Disabled,
		}

		if _, ok := d.GetOkExists(fmt.Sprintf("rules.%d.priority", idx)); ok {
			m["priority"] = r.Priority
		}
		if _, ok := d.GetOkExists(fmt.Sprintf("rules.%d.terminates", idx)); ok {
			m["terminates"] = r.Terminates
		}

		if fr := r.FixedResponse; fr != nil {
			frm := map[string]interface{}{}
			if _, ok := d.GetOkExists(fmt.Sprintf("rules.%d.fixed_response.0.message_body", idx)); ok {
				frm["message_body"] = fr.MessageBody
				m["fixed_response"] = []interface{}{frm} // only set if one of these has is true
			}
			if _, ok := d.GetOkExists(fmt.Sprintf("rules.%d.fixed_response.0.status_code", idx)); ok {
				frm["status_code"] = fr.StatusCode
				m["fixed_response"] = []interface{}{frm} // only set if one of these has is true
			}
			if _, ok := d.GetOkExists(fmt.Sprintf("rules.%d.fixed_response.0.content_type", idx)); ok {
				frm["content_type"] = fr.ContentType
				m["fixed_response"] = []interface{}{frm} // only set if one of these has is true
			}
			if _, ok := d.GetOkExists(fmt.Sprintf("rules.%d.fixed_response.0.location", idx)); ok {
				frm["location"] = fr.Location
				m["fixed_response"] = []interface{}{frm} // only set if one of these has is true
			}
		}

		if _, ok := d.GetOkExists(fmt.Sprintf("rules.%d.overrides", idx)); ok {
			o := r.Overrides
			om := map[string]interface{}{}

			if _, ok := d.GetOkExists(fmt.Sprintf("rules.%d.overrides.0.session_affinity", idx)); ok {
				om["session_affinity"] = o.Persistence
				m["overrides"] = []interface{}{om}
			}
			if _, ok := d.GetOkExists(fmt.Sprintf("rules.%d.overrides.0.session_affinity_ttl", idx)); ok {
				om["session_affinity"] = o.PersistenceTTL
				m["overrides"] = []interface{}{om}
			}
			if _, ok := d.GetOkExists(fmt.Sprintf("rules.%d.overrides.0.ttl", idx)); ok {
				om["ttl"] = o.TTL
				m["overrides"] = []interface{}{om}
			}
			if _, ok := d.GetOkExists(fmt.Sprintf("rules.%d.overrides.0.steering_policy", idx)); ok {
				om["steering_policy"] = o.SteeringPolicy
				m["overrides"] = []interface{}{om}
			}
			if _, ok := d.GetOkExists(fmt.Sprintf("rules.%d.overrides.0.fallback_pool", idx)); ok {
				om["fallback_pool"] = o.FallbackPool
				m["overrides"] = []interface{}{om}
			}
			if _, ok := d.GetOkExists(fmt.Sprintf("rules.%d.overrides.0.default_pools", idx)); ok {
				om["default_pools"] = o.DefaultPools
				m["overrides"] = []interface{}{om}
			}
			if _, ok := d.GetOkExists(fmt.Sprintf("rules.%d.overrides.0.pop_pools", idx)); ok {
				om["pop_pools"] = flattenGeoPools(o.PoPPools, "pop", loadBalancerOverridesLocalPoolElems)
				m["overrides"] = []interface{}{om}
			}
			if _, ok := d.GetOkExists(fmt.Sprintf("rules.%d.overrides.0.country_pools", idx)); ok {
				om["country_pools"] = flattenGeoPools(o.CountryPools, "country", loadBalancerOverridesLocalPoolElems)
				m["overrides"] = []interface{}{om}
			}
			if _, ok := d.GetOkExists(fmt.Sprintf("rules.%d.overrides.0.region_pools", idx)); ok {
				om["region_pools"] = flattenGeoPools(o.RegionPools, "region", loadBalancerOverridesLocalPoolElems)
				m["overrides"] = []interface{}{om}
			}
			if _, ok := d.GetOkExists(fmt.Sprintf("rules.%d.overrides.0.session_affinity_attributes", idx)); o.SessionAffinityAttrs != nil && ok {
				saa := map[string]interface{}{}
				om["session_affinity_attributes"] = saa
				m["overrides"] = []interface{}{om}
				if _, ok := d.GetOkExists(fmt.Sprintf("rules.%d.overrides.0.session_affinity_attributes.samesite", idx)); ok {
					saa["samesite"] = o.SessionAffinityAttrs.SameSite
				}
				if _, ok := d.GetOkExists(fmt.Sprintf("rules.%d.overrides.0.session_affinity_attributes.secure", idx)); ok {
					saa["secure"] = o.SessionAffinityAttrs.Secure
				}
				if _, ok := d.GetOkExists(fmt.Sprintf("rules.%d.overrides.0.session_affinity_attributes.zero_downtime_failover", idx)); ok {
					saa["zero_downtime_failover"] = o.SessionAffinityAttrs.ZeroDowntimeFailover
				}
			}
			if arOk, ok := d.GetOk(fmt.Sprintf("rules.%d.overrides.0.adaptive_routing", idx)); o.AdaptiveRouting != nil && ok {
				ar := map[string]interface{}{}
				if l := arOk.(*schema.Set).List(); len(l) > 0 {
					for k := range l[0].(map[string]interface{}) {
						switch k {
						case "failover_across_pools":
							ar[k] = bool(o.AdaptiveRouting.FailoverAcrossPools != nil && *o.AdaptiveRouting.FailoverAcrossPools)
						}
					}
				}
				flattened := []interface{}{ar}
				om["adaptive_routing"] = schema.NewSet(schema.HashResource(loadBalancerOverridesAdaptiveRoutingElem), flattened)
				m["overrides"] = []interface{}{om}
			}
			if lsOk, ok := d.GetOk(fmt.Sprintf("rules.%d.overrides.0.location_strategy", idx)); o.LocationStrategy != nil && ok {
				ls := map[string]interface{}{}
				if l := lsOk.(*schema.Set).List(); len(l) > 0 {
					for k := range l[0].(map[string]interface{}) {
						switch k {
						case "prefer_ecs":
							ls[k] = o.LocationStrategy.PreferECS
						case "mode":
							ls[k] = o.LocationStrategy.Mode
						}
					}
				}
				flattened := []interface{}{ls}
				om["location_strategy"] = schema.NewSet(schema.HashResource(loadBalancerOverridesLocationStrategyElem), flattened)
				m["overrides"] = []interface{}{om}
			}
			if rsOk, ok := d.GetOk(fmt.Sprintf("rules.%d.overrides.0.random_steering", idx)); o.RandomSteering != nil && ok {
				rs := map[string]interface{}{}
				if l := rsOk.(*schema.Set).List(); len(l) > 0 {
					for k := range l[0].(map[string]interface{}) {
						switch k {
						case "pool_weights":
							poolWeights := make(map[string]interface{})
							for poolID, poolWeight := range o.RandomSteering.PoolWeights {
								poolWeights[poolID] = poolWeight
							}
							rs[k] = poolWeights
						case "default_weight":
							rs[k] = o.RandomSteering.DefaultWeight
						}
					}
				}
				flattened := []interface{}{rs}
				om["random_steering"] = schema.NewSet(schema.HashResource(loadBalancerOverridesRandomSteeringElem), flattened)
				m["overrides"] = []interface{}{om}
			}
		}

		cfResources = append(cfResources, m)
	}
	return cfResources, nil
}

func expandRules(rdata interface{}) ([]*cloudflare.LoadBalancerRule, error) {
	var rules []*cloudflare.LoadBalancerRule
	for _, ele := range rdata.([]interface{}) {
		r := ele.(map[string]interface{})
		lbr := &cloudflare.LoadBalancerRule{
			Name: r["name"].(string),
		}
		if v, ok := r["priority"]; ok {
			lbr.Priority = v.(int)
		}
		if d, ok := r["disabled"]; ok {
			lbr.Disabled = d.(bool)
		}
		if c, ok := r["condition"]; ok {
			lbr.Condition = c.(string)
		}
		if t, ok := r["terminates"]; ok {
			lbr.Terminates = t.(bool)
		}

		if overridesData, ok := r["overrides"]; ok && len(overridesData.([]interface{})) > 0 {
			ov := overridesData.([]interface{})[0].(map[string]interface{})

			if sa, ok := ov["session_affinity"]; ok {
				lbr.Overrides.Persistence = sa.(string)
			}

			if sattl, ok := ov["session_affinity_ttl"]; ok {
				v := uint(sattl.(int))
				// a default value of seem to be set into this field bypassing
				// the IntBetween(1800, 604800) validation check ignore
				// this zero values here
				if v != 0 {
					lbr.Overrides.PersistenceTTL = &v
				}
			}

			if saattr, ok := ov["session_affinity_attributes"]; ok {
				attr := saattr.(map[string]interface{})
				v := &cloudflare.LoadBalancerRuleOverridesSessionAffinityAttrs{}
				if ss, ok := attr["samesite"]; ok {
					v.SameSite = ss.(string)
					lbr.Overrides.SessionAffinityAttrs = v
				}
				if sec, ok := attr["secure"]; ok {
					v.Secure = sec.(string)
					lbr.Overrides.SessionAffinityAttrs = v
				}
				if zdf, ok := attr["zero_downtime_failover"]; ok {
					v.ZeroDowntimeFailover = zdf.(string)
					lbr.Overrides.SessionAffinityAttrs = v
				}
			}

			if ar, ok := ov["adaptive_routing"]; ok {
				if l := ar.(*schema.Set).List(); len(l) > 0 {
					arOverride := &cloudflare.AdaptiveRouting{}
					for k, v := range l[0].(map[string]interface{}) {
						switch k {
						case "failover_across_pools":
							arOverride.FailoverAcrossPools = cloudflare.BoolPtr(v.(bool))
							lbr.Overrides.AdaptiveRouting = arOverride
						}
					}
				}
			}

			if ls, ok := ov["location_strategy"]; ok {
				if l := ls.(*schema.Set).List(); len(l) > 0 {
					lsOverride := &cloudflare.LocationStrategy{}
					for k, v := range l[0].(map[string]interface{}) {
						switch k {
						case "prefer_ecs":
							lsOverride.PreferECS = v.(string)
							lbr.Overrides.LocationStrategy = lsOverride
						case "mode":
							lsOverride.Mode = v.(string)
							lbr.Overrides.LocationStrategy = lsOverride
						}
					}
				}
			}

			if rs, ok := ov["random_steering"]; ok {
				if l := rs.(*schema.Set).List(); len(l) > 0 {
					rsOverride := &cloudflare.RandomSteering{}
					for k, v := range l[0].(map[string]interface{}) {
						switch k {
						case "pool_weights":
							poolWeights := make(map[string]float64)
							for poolID, poolWeight := range v.(map[string]interface{}) {
								poolWeights[poolID] = poolWeight.(float64)
							}
							rsOverride.PoolWeights = poolWeights
							lbr.Overrides.RandomSteering = rsOverride
						case "default_weight":
							rsOverride.DefaultWeight = v.(float64)
							lbr.Overrides.RandomSteering = rsOverride
						}
					}
				}
			}

			if ttl, ok := ov["ttl"]; ok {
				lbr.Overrides.TTL = uint(ttl.(int))
			}

			if sp, ok := ov["steering_policy"]; ok {
				lbr.Overrides.SteeringPolicy = sp.(string)
			}

			if fb, ok := ov["fallback_pool"]; ok {
				lbr.Overrides.FallbackPool = fb.(string)
			}

			if dp, ok := ov["default_pools"]; ok {
				lbr.Overrides.DefaultPools = expandInterfaceToStringList(dp)
			}

			if pp, ok := ov["pop_pools"]; ok {
				expandedPopPools, err := expandGeoPools(pp, "pop")
				if err != nil {
					return nil, err
				}
				lbr.Overrides.PoPPools = expandedPopPools
			}

			if cp, ok := ov["country_pools"]; ok {
				expandedCountryPools, err := expandGeoPools(cp, "country")
				if err != nil {
					return nil, err
				}
				lbr.Overrides.CountryPools = expandedCountryPools
			}

			if rp, ok := ov["region_pools"]; ok {
				expandedRegionPools, err := expandGeoPools(rp, "region")
				if err != nil {
					return nil, err
				}
				lbr.Overrides.RegionPools = expandedRegionPools
			}
		}

		for _, fixedResponseData := range r["fixed_response"].([]interface{}) {
			frd := fixedResponseData.(map[string]interface{})
			// we don't add this into our LB unless one of the cases below is true
			fr := &cloudflare.LoadBalancerFixedResponseData{}

			if mb, ok := frd["message_body"]; ok {
				fr.MessageBody = mb.(string)
				lbr.FixedResponse = fr
			}
			if sc, ok := frd["status_code"]; ok {
				fr.StatusCode = sc.(int)
				lbr.FixedResponse = fr
			}
			if ct, ok := frd["content_type"]; ok {
				fr.ContentType = ct.(string)
				lbr.FixedResponse = fr
			}
			if l, ok := frd["location"]; ok {
				fr.Location = l.(string)
				lbr.FixedResponse = fr
			}
		}

		rules = append(rules, lbr)
	}

	return rules, nil
}

func expandSessionAffinityAttrs(attrs interface{}) (*cloudflare.SessionAffinityAttributes, error) {
	var cfSessionAffinityAttrs cloudflare.SessionAffinityAttributes

	for k, v := range attrs.(map[string]interface{}) {
		switch k {
		case "secure":
			cfSessionAffinityAttrs.Secure = v.(string)
		case "samesite":
			cfSessionAffinityAttrs.SameSite = v.(string)
		case "drain_duration":
			var err error
			if cfSessionAffinityAttrs.DrainDuration, err = strconv.Atoi(v.(string)); err != nil {
				return nil, err
			}
		case "zero_downtime_failover":
			cfSessionAffinityAttrs.ZeroDowntimeFailover = v.(string)
		}
	}

	return &cfSessionAffinityAttrs, nil
}

func expandAdaptiveRouting(set interface{}) *cloudflare.AdaptiveRouting {
	var cfAdaptiveRouting cloudflare.AdaptiveRouting

	if l := set.(*schema.Set).List(); len(l) > 0 {
		for k, v := range l[0].(map[string]interface{}) {
			switch k {
			case "failover_across_pools":
				cfAdaptiveRouting.FailoverAcrossPools = cloudflare.BoolPtr(v.(bool))
			}
		}
	}

	return &cfAdaptiveRouting
}

func expandLocationStrategy(set interface{}) *cloudflare.LocationStrategy {
	var cfLocationStrategy cloudflare.LocationStrategy

	if l := set.(*schema.Set).List(); len(l) > 0 {
		for k, v := range l[0].(map[string]interface{}) {
			switch k {
			case "prefer_ecs":
				cfLocationStrategy.PreferECS = v.(string)
			case "mode":
				cfLocationStrategy.Mode = v.(string)
			}
		}
	}

	return &cfLocationStrategy
}

func expandRandomSteering(set interface{}) *cloudflare.RandomSteering {
	var cfRandomSteering cloudflare.RandomSteering

	if l := set.(*schema.Set).List(); len(l) > 0 {
		for k, v := range l[0].(map[string]interface{}) {
			switch k {
			case "pool_weights":
				poolWeights := make(map[string]float64)
				for poolID, poolWeight := range v.(map[string]interface{}) {
					poolWeights[poolID] = poolWeight.(float64)
				}
				cfRandomSteering.PoolWeights = poolWeights
			case "default_weight":
				cfRandomSteering.DefaultWeight = v.(float64)
			}
		}
	}

	return &cfRandomSteering
}
