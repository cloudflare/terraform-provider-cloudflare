package cloudflare

import (
	"fmt"
	"log"
	"strings"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
	"github.com/pkg/errors"
)

func resourceCloudFlareRateLimit() *schema.Resource {
	return &schema.Resource{
		Create: resourceCloudFlareRateLimitCreate,
		Read:   resourceCloudFlareRateLimitRead,
		Update: resourceCloudFlareRateLimitUpdate,
		Delete: resourceCloudFlareRateLimitDelete,
		Importer: &schema.ResourceImporter{
			State: resourceCloudFlareRateLimitImport,
		},

		SchemaVersion: 1,
		Schema: map[string]*schema.Schema{
			"zone": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"zone_id": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"rate_limit_id": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"threshold": {
				Type:         schema.TypeInt,
				Required:     true,
				ValidateFunc: validation.IntBetween(2, 1000000),
			},

			"period": {
				Type:         schema.TypeInt,
				Required:     true,
				ValidateFunc: validation.IntBetween(1, 86400),
			},

			"match": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"request": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MinItems: 1,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"methods": {
										Type:     schema.TypeSet,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{Type: schema.TypeString,
											ValidateFunc: validation.StringInSlice(allowedHTTPMethods, true)},
									},

									"schemes": {
										Type:     schema.TypeSet,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{Type: schema.TypeString,
											ValidateFunc: validation.StringInSlice(allowedSchemes, true)},
									},

									"url_pattern": {
										Type:         schema.TypeString,
										Optional:     true,
										Computed:     true,
										ValidateFunc: validation.StringLenBetween(0, 1024),
									},
								},
							},
						},

						"response": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MinItems: 1,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"statuses": {
										Type:     schema.TypeSet,
										Optional: true,
										Computed: true,
										Elem:     &schema.Schema{Type: schema.TypeInt},
									},

									"origin_traffic": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},

			"action": {
				Type:     schema.TypeList,
				Required: true,
				MinItems: 1,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"mode": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringInSlice([]string{"simulate", "ban"}, true),
						},

						"timeout": {
							Type:         schema.TypeInt,
							Required:     true,
							ValidateFunc: validation.IntBetween(1, 86400),
						},

						"response": {
							Type:     schema.TypeList,
							Optional: true,
							MinItems: 1,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"content_type": {
										Type:         schema.TypeString,
										Required:     true,
										ValidateFunc: validation.StringInSlice([]string{"text/plain", "text/xml", "application/json"}, true),
									},

									"body": {
										Type:         schema.TypeString,
										Required:     true,
										ValidateFunc: validation.StringLenBetween(0, 10240),
										// maybe good to hash the body before saving in state file?
									},
								},
							},
						},
					},
				},
			},

			"disabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},

			"description": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringLenBetween(0, 1024),
			},

			"bypass_url_patterns": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func resourceCloudFlareRateLimitCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)

	newRateLimit := cloudflare.RateLimit{
		Threshold: d.Get("threshold").(int),
		Period:    d.Get("period").(int),
		Action:    expandRateLimitAction(d),
	}

	newRateLimitMatch, err := expandRateLimitTrafficMatcher(d)
	if err != nil {
		return err
	}
	newRateLimit.Match = newRateLimitMatch

	if disabled, ok := d.GetOk("disabled"); ok {
		newRateLimit.Disabled = disabled.(bool)
	}

	if description, ok := d.GetOk("description"); ok {
		newRateLimit.Description = description.(string)
	}

	if bypassUrlPatterns, ok := d.GetOk("bypass_url_patterns"); ok {
		newRateLimit.Bypass = expandRateLimitBypass(bypassUrlPatterns.(*schema.Set))
	}

	zoneName := d.Get("zone").(string)
	zoneId, err := client.ZoneIDByName(zoneName)
	if err != nil {
		return fmt.Errorf("error finding zone %q: %s", zoneName, err)
	}

	log.Printf("[DEBUG] Creating CloudFlare Rate Limit from struct: %+v", newRateLimit)

	r, err := client.CreateRateLimit(zoneId, newRateLimit)
	if err != nil {
		return errors.Wrap(err, "error creating rate limit for zone")
	}

	if r.ID == "" {
		return fmt.Errorf("cailed to find id in Create response; resource was empty")
	}

	// terraform id is *not* the same as the resource id, is is the combination with the zoneId
	// this makes it easier to import and also matches the keys needed for cloudflare-go operations
	d.SetId(zoneName + "_" + r.ID)
	// assume ids are immutable, not going to look it up from the api again
	d.Set("zone_id", zoneId)
	d.Set("rate_limit_id", r.ID)

	log.Printf("[INFO] CloudFlare Rate Limit ID: %s", d.Id())

	return resourceCloudFlareRateLimitRead(d, meta)
}

func resourceCloudFlareRateLimitUpdate(d *schema.ResourceData, meta interface{}) error {
	// since api only supports replace, update looks a lot like create...
	client := meta.(*cloudflare.API)
	zoneId := d.Get("zone_id").(string)
	rateLimitId := d.Get("rate_limit_id").(string)

	updatedRateLimit := cloudflare.RateLimit{
		Threshold: d.Get("threshold").(int),
		Period:    d.Get("period").(int),
		Action:    expandRateLimitAction(d),
	}

	newRateLimitMatch, err := expandRateLimitTrafficMatcher(d)
	if err != nil {
		return err
	}
	updatedRateLimit.Match = newRateLimitMatch

	if disabled, ok := d.GetOk("disabled"); ok {
		updatedRateLimit.Disabled = disabled.(bool)
	}

	if description, ok := d.GetOk("description"); ok {
		updatedRateLimit.Description = description.(string)
	}

	if bypassUrlPatterns, ok := d.GetOk("bypass_url_patterns"); ok {
		updatedRateLimit.Bypass = expandRateLimitBypass(bypassUrlPatterns.(*schema.Set))
	}

	_, err = client.UpdateRateLimit(zoneId, rateLimitId, updatedRateLimit)
	if err != nil {
		return errors.Wrap(err, "error creating rate limit for zone")
	}
	return resourceCloudFlareRateLimitRead(d, meta)
}

func expandRateLimitTrafficMatcher(d *schema.ResourceData) (matcher cloudflare.RateLimitTrafficMatcher, err error) {
	v, ok := d.GetOk("match")
	if !ok {
		return
	}
	cfg := v.([]interface{})[0].(map[string]interface{})

	if matchReqIface, ok := cfg["request"]; ok && len(matchReqIface.([]interface{})) > 0 {
		matchReq := matchReqIface.([]interface{})[0].(map[string]interface{})

		requestMatcher := cloudflare.RateLimitRequestMatcher{
			URLPattern: matchReq["url_pattern"].(string),
		}

		if methodsSet, ok := matchReq["methods"]; ok {
			methods := make([]string, methodsSet.(*schema.Set).Len())
			for i, m := range methodsSet.(*schema.Set).List() {
				methods[i] = m.(string)
			}
			requestMatcher.Methods = methods
		}
		if schemesSet, ok := matchReq["schemes"]; ok {
			schemes := make([]string, schemesSet.(*schema.Set).Len())
			for i, s := range schemesSet.(*schema.Set).List() {
				schemes[i] = s.(string)
			}
			requestMatcher.Schemes = schemes
		}
		matcher.Request = requestMatcher
	}
	if matchRespIface, ok := cfg["response"]; ok && len(matchRespIface.([]interface{})) > 0 {
		matchResp := matchRespIface.([]interface{})[0].(map[string]interface{})

		responseMatcher := cloudflare.RateLimitResponseMatcher{}

		if statusesSet, ok := matchResp["statuses"]; ok {
			statuses := make([]int, statusesSet.(*schema.Set).Len())
			for i, s := range statusesSet.(*schema.Set).List() {
				statuses[i] = s.(int)
			}
			responseMatcher.Statuses = statuses
		}
		if originIface, ok := matchResp["origin_traffic"]; ok {
			originTraffic := originIface.(bool)
			responseMatcher.OriginTraffic = &originTraffic
		}
		matcher.Response = responseMatcher
	}
	return
}

func expandRateLimitAction(d *schema.ResourceData) cloudflare.RateLimitAction {
	// dont need to guard for array length because MinItems is set **and** action is required
	tfAction := d.Get("action").([]interface{})[0].(map[string]interface{})

	action := cloudflare.RateLimitAction{
		Mode:    tfAction["mode"].(string),
		Timeout: tfAction["timeout"].(int),
	}

	if _, ok := tfAction["response"]; ok && len(tfAction["response"].([]interface{})) > 0 {
		log.Printf("[DEBUG] CloudFlare Rate Limit specified action: %+v \n", tfAction)
		tfActionResponse := tfAction["response"].([]interface{})[0].(map[string]interface{})

		action.Response = &cloudflare.RateLimitActionResponse{
			ContentType: tfActionResponse["content_type"].(string),
			Body:        tfActionResponse["body"].(string),
		}
	}
	return action
}

func expandRateLimitBypass(bypassUrlPatterns *schema.Set) []cloudflare.RateLimitKeyValue {
	bypass := make([]cloudflare.RateLimitKeyValue, bypassUrlPatterns.Len())
	for i, urlPattern := range bypassUrlPatterns.List() {
		bypass[i] = cloudflare.RateLimitKeyValue{
			Name:  "url",
			Value: urlPattern.(string),
		}
	}
	return bypass
}

func resourceCloudFlareRateLimitRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	zoneId := d.Get("zone_id").(string)
	rateLimitId := d.Get("rate_limit_id").(string)

	rateLimit, err := client.RateLimit(zoneId, rateLimitId)
	if err != nil {
		if strings.Contains(err.Error(), "HTTP status 404") {
			log.Printf("[INFO] Resource %s in zone %s no longer exists", rateLimitId, zoneId)
			d.SetId("")
			return nil
		} else {
			return errors.Wrap(err,
				fmt.Sprintf("Error reading rate limit resource from API for resource %s in zone %s", zoneId, rateLimitId))
		}
	}
	log.Printf("[DEBUG] Read CloudFlare Rate Limit from API as struct: %+v", rateLimit)

	d.Set("threshold", rateLimit.Threshold)
	d.Set("period", rateLimit.Period)
	d.Set("match", flattenRateLimitTrafficMatcher(rateLimit.Match))
	d.Set("action", flattenRateLimitAction(rateLimit.Action))

	d.Set("description", rateLimit.Description)
	d.Set("disabled", rateLimit.Disabled)
	bypassUrlPatterns := make([]string, 0)
	for _, bypassItem := range rateLimit.Bypass {
		if bypassItem.Name == "url" {
			bypassUrlPatterns = append(bypassUrlPatterns, bypassItem.Value)
		} else {
			// maybe a new type of bypass was added to api
			log.Printf("[WARN] Unkown bypass type found in rate limit for zone: %s", bypassItem.Name)
		}
	}
	d.Set("bypass_url_patterns", bypassUrlPatterns)
	return nil
}

func flattenRateLimitTrafficMatcher(cfg cloudflare.RateLimitTrafficMatcher) []map[string]interface{} {
	data := map[string]interface{}{
		"request":  flattenRateLimitRequestMatcher(cfg.Request),
		"response": flattenRateLimitResponseMatcher(cfg.Response),
	}
	return []map[string]interface{}{data}
}

func flattenRateLimitRequestMatcher(cfg cloudflare.RateLimitRequestMatcher) []map[string]interface{} {
	data := map[string]interface{}{
		"methods":     schema.NewSet(schema.HashString, flattenStringList(cfg.Methods)),
		"schemes":     schema.NewSet(schema.HashString, flattenStringList(cfg.Schemes)),
		"url_pattern": cfg.URLPattern,
	}

	return []map[string]interface{}{data}
}

func flattenRateLimitResponseMatcher(cfg cloudflare.RateLimitResponseMatcher) []map[string]interface{} {
	data := map[string]interface{}{
		"origin_traffic": *cfg.OriginTraffic,
	}

	if len(cfg.Statuses) > 0 {
		data["statuses"] = schema.NewSet(IntIdentity, flattenIntList(cfg.Statuses))
	}

	if len(data) > 0 {
		return []map[string]interface{}{data}
	} else {
		return []map[string]interface{}{}
	}
}

func flattenRateLimitAction(cfg cloudflare.RateLimitAction) []map[string]interface{} {
	action := map[string]interface{}{
		"mode":    cfg.Mode,
		"timeout": cfg.Timeout,
	}

	if cfg.Response != nil {
		cfgResponse := *cfg.Response
		actionResponse := map[string]interface{}{
			"content_type": cfgResponse.ContentType,
			"body":         cfgResponse.Body,
		}
		action["response"] = []map[string]interface{}{actionResponse}
	}
	return []map[string]interface{}{action}
}

func resourceCloudFlareRateLimitDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	zoneId := d.Get("zone_id").(string)
	rateLimitId := d.Get("rate_limit_id").(string)

	log.Printf("[INFO] Deleting CloudFlare Rate Limit: %s for zone: %s", rateLimitId, zoneId)

	err := client.DeleteRateLimit(zoneId, rateLimitId)
	if err != nil {
		return fmt.Errorf("error deleting CloudFlare Rate Limit for zone: %s", err)
	}

	return nil
}

func resourceCloudFlareRateLimitImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	client := meta.(*cloudflare.API)

	// split the id so we can lookup
	idAttr := strings.SplitN(d.Id(), "_", 2)
	var zoneName string
	var rateLimitId string
	if len(idAttr) == 2 {
		zoneName = idAttr[0]
		rateLimitId = idAttr[1]
		d.Set("zone", zoneName)
		d.Set("rate_limit_id", rateLimitId)

	} else {
		return nil, fmt.Errorf("invalid id (\"%s\") specified, should be in format \"zoneName_rateLimitId\"", d.Id())
	}
	zoneId, err := client.ZoneIDByName(zoneName)
	d.Set("zone_id", zoneId)
	if err != nil {
		return nil, fmt.Errorf("error finding zoneName %q: %s", zoneName, err)
	}
	return []*schema.ResourceData{d}, nil
}
