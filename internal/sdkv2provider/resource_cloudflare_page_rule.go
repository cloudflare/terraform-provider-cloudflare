package sdkv2provider

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflarePageRule() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflarePageRuleSchema(),
		CreateContext: resourceCloudflarePageRuleCreate,
		ReadContext:   resourceCloudflarePageRuleRead,
		UpdateContext: resourceCloudflarePageRuleUpdate,
		DeleteContext: resourceCloudflarePageRuleDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceCloudflarePageRuleImport,
		},
	}
}

func suppressEquivalentURLs(k, old, new string, d *schema.ResourceData) bool {
	// this is probably due to RFC3986 normalization but its unspecified
	if strings.Trim(new, "/") == strings.Trim(old, "/") {
		return true
	}
	return false
}

func resourceCloudflarePageRuleCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	newPageRuleTargets := []cloudflare.PageRuleTarget{
		{
			Target: "url",
			Constraint: struct {
				Operator string `json:"operator"`
				Value    string `json:"value"`
			}{
				Operator: "matches",
				Value:    d.Get("target").(string),
			},
		},
	}

	actions := d.Get("actions").([]interface{})
	newPageRuleActions := make([]cloudflare.PageRuleAction, 0, len(actions))

	tflog.Debug(ctx, fmt.Sprintf("Actions found in config: %#v", actions))
	for _, action := range actions {
		for id, value := range action.(map[string]interface{}) {
			newPageRuleAction, err := transformToCloudflarePageRuleAction(ctx, id, value, d)
			if err != nil {
				return diag.FromErr(err)
			} else if newPageRuleAction.Value == nil || newPageRuleAction.Value == "" {
				continue
			}
			newPageRuleActions = append(newPageRuleActions, newPageRuleAction)
		}
	}
	pageRulesActionMap := pageRuleActionsToMap(newPageRuleActions)
	if _, ok := pageRulesActionMap["forwarding_url"]; ok && len(pageRulesActionMap) > 1 {
		return diag.FromErr(fmt.Errorf("\"forwarding_url\" cannot be set with any other actions"))
	}

	newPageRule := cloudflare.PageRule{
		Targets:  newPageRuleTargets,
		Actions:  newPageRuleActions,
		Priority: d.Get("priority").(int),
		Status:   d.Get("status").(string),
	}

	tflog.Debug(ctx, fmt.Sprintf("Cloudflare Page Rule create configuration: %#v", newPageRule))

	r, err := client.CreatePageRule(ctx, zoneID, newPageRule)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to create page rule: %w", err))
	}

	if r.ID == "" {
		return diag.FromErr(fmt.Errorf("Failed to find page rule in Create response; ID was empty"))
	}

	d.SetId(r.ID)

	return resourceCloudflarePageRuleRead(ctx, d, meta)
}

func pageRuleActionsToMap(vs []cloudflare.PageRuleAction) map[string]interface{} {
	vsm := make(map[string]interface{})
	for _, v := range vs {
		vsm[v.ID] = v.Value
	}
	return vsm
}

func resourceCloudflarePageRuleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	pageRule, err := client.PageRule(ctx, zoneID, d.Id())
	if err != nil {
		var notFoundError *cloudflare.NotFoundError
		if strings.Contains(err.Error(), "Invalid Page Rule identifier") || // api bug - this indicates non-existing resource
			errors.As(err, &notFoundError) {
			tflog.Info(ctx, fmt.Sprintf("Page Rule %s no longer exists", d.Id()))
			d.SetId("")
			return nil
		} else {
			return diag.FromErr(fmt.Errorf("error finding page rule %q: %w", d.Id(), err))
		}
	}
	tflog.Debug(ctx, fmt.Sprintf("Cloudflare Page Rule read configuration: %#v", pageRule))

	// Cloudflare presently only has one target type, and its Operator is always
	// "matches"; so we can just read the first element's Value.
	d.Set("target", pageRule.Targets[0].Constraint.Value)

	d.Set("priority", pageRule.Priority)
	d.Set("status", pageRule.Status)

	actions := map[string]interface{}{}
	for _, pageRuleAction := range pageRule.Actions {
		key, value, err := transformFromCloudflarePageRuleAction(&pageRuleAction)
		if err != nil {
			return diag.FromErr(fmt.Errorf("failed to parse page rule action: %w", err))
		}
		actions[key] = value
	}
	log.Printf("[DEBUG] Cloudflare Page Rule actions configuration: %#v", actions)

	if err := d.Set("actions", []map[string]interface{}{actions}); err != nil {
		log.Printf("[WARN] Error setting actions in page rule %q: %s", d.Id(), err)
	}

	return nil
}

func resourceCloudflarePageRuleUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	updatePageRule := cloudflare.PageRule{}

	if target, ok := d.GetOk("target"); ok {
		updatePageRule.Targets = []cloudflare.PageRuleTarget{
			{
				Target: "url",
				Constraint: struct {
					Operator string `json:"operator"`
					Value    string `json:"value"`
				}{
					Operator: "matches",
					Value:    target.(string),
				},
			},
		}
	}

	if v, ok := d.GetOk("actions"); ok {
		actions := v.([]interface{})
		newPageRuleActions := make([]cloudflare.PageRuleAction, 0, len(actions))

		for _, action := range actions {
			for id, value := range action.(map[string]interface{}) {
				newPageRuleAction, err := transformToCloudflarePageRuleAction(ctx, id, value, d)
				if err != nil {
					return diag.FromErr(err)
				} else if newPageRuleAction.Value == nil {
					continue
				}
				newPageRuleActions = append(newPageRuleActions, newPageRuleAction)
			}
		}

		updatePageRule.Actions = newPageRuleActions
	}

	if priority, ok := d.GetOk("priority"); ok {
		updatePageRule.Priority = priority.(int)
	}

	if status, ok := d.GetOk("status"); ok {
		updatePageRule.Status = status.(string)
	}

	log.Printf("[DEBUG] Cloudflare Page Rule update configuration: %#v", updatePageRule)

	if err := client.UpdatePageRule(ctx, zoneID, d.Id(), updatePageRule); err != nil {
		return diag.FromErr(fmt.Errorf("failed to update Cloudflare Page Rule: %w", err))
	}

	return resourceCloudflarePageRuleRead(ctx, d, meta)
}

func resourceCloudflarePageRuleDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	log.Printf("[INFO] Deleting Cloudflare Page Rule: %s, %s", zoneID, d.Id())

	if err := client.DeletePageRule(ctx, zoneID, d.Id()); err != nil {
		return diag.FromErr(fmt.Errorf("error deleting Cloudflare Page Rule: %w", err))
	}

	return nil
}

var pageRuleAPIOnOffFields = []string{
	"automatic_https_rewrites",
	"browser_check",
	"cache_by_device_type",
	"cache_deception_armor",
	"email_obfuscation",
	"explicit_cache_control",
	"ip_geolocation",
	"mirage",
	"opportunistic_encryption",
	"origin_error_page_pass_thru",
	"respect_strong_etag",
	"response_buffering",
	"rocket_loader",
	"server_side_exclude",
	"sort_query_string_for_cache",
	"true_client_ip_header",
	"waf",
}
var pageRuleAPINilFields = []string{
	"always_use_https",
	"disable_apps",
	"disable_performance",
	"disable_railgun",
	"disable_security",
	"disable_zaraz",
}
var pageRuleAPIStringFields = []string{
	"bypass_cache_on_cookie",
	"cache_key",
	"cache_level",
	"cache_on_cookie",
	"host_header_override",
	"polish",
	"resolve_override",
	"security_level",
	"ssl",
}

func transformFromCloudflarePageRuleAction(pageRuleAction *cloudflare.PageRuleAction) (key string, value interface{}, err error) {
	key = pageRuleAction.ID

	switch {
	case contains(pageRuleAPIOnOffFields, pageRuleAction.ID):
		value = pageRuleAction.Value.(string)
		break

	case contains(pageRuleAPINilFields, pageRuleAction.ID):
		// api returns a nil value so set the value ourselves
		value = true
		break

	case contains(pageRuleAPIStringFields, pageRuleAction.ID):
		value = pageRuleAction.Value.(string)
		break

	case pageRuleAction.ID == "edge_cache_ttl":
		value = pageRuleAction.Value.(float64) // we use TypeInt but terraform seems to do the right thing converting from float
		break

	case pageRuleAction.ID == "browser_cache_ttl":
		value = fmt.Sprintf("%.0f", pageRuleAction.Value.(float64))
		break

	case pageRuleAction.ID == "forwarding_url" || pageRuleAction.ID == "minify":
		value = []interface{}{pageRuleAction.Value.(map[string]interface{})}
		break

	case pageRuleAction.ID == "cache_key_fields":
		output := make(map[string]interface{})

		for sectionID, sectionValue := range pageRuleAction.Value.(map[string]interface{}) {
			switch sectionID {
			case "host", "user":
				output[sectionID] = []interface{}{sectionValue}

			case "cookie", "header":
				fieldOutput := map[string]interface{}{}
				for fieldID, fieldValue := range sectionValue.(map[string]interface{}) {
					switch fieldValue.(type) {
					case []interface{}:
						if len(fieldValue.([]interface{})) > 0 {
							fieldOutput[fieldID] = fieldValue
						}
					default:
						fieldOutput[fieldID] = fieldValue
					}
				}
				if len(fieldOutput) > 0 {
					output[sectionID] = []interface{}{fieldOutput}
				}

			case "query_string":
				fieldOutput := map[string]interface{}{}

				for fieldID, fieldValue := range sectionValue.(map[string]interface{}) {
					fieldOutput[fieldID] = fieldValue
				}

				if reflect.TypeOf(fieldOutput["exclude"]).Kind() == reflect.String && fieldOutput["exclude"] == "*" {
					fieldOutput["ignore"] = true
					fieldOutput["exclude"] = []interface{}{}
				}

				if reflect.TypeOf(fieldOutput["include"]).Kind() == reflect.String && fieldOutput["include"] == "*" {
					fieldOutput["ignore"] = false
					fieldOutput["include"] = []interface{}{}
				}

				output[sectionID] = []interface{}{fieldOutput}
			}
		}

		value = []interface{}{output}
		break

	case pageRuleAction.ID == "cache_ttl_by_status":
		output := make([]map[string]interface{}, 0)

		for key, value := range pageRuleAction.Value.(map[string]interface{}) {
			entry := map[string]interface{}{"codes": key}

			switch value := value.(type) {
			case float64:
				entry["ttl"] = int32(value)
			case string:
				switch value {
				case "no-cache":
					entry["ttl"] = 0
				case "no-store":
					entry["ttl"] = -1
				}
			}

			output = append(output, entry)
		}

		value = output
		break

	default:
		// User supplied ID is already validated, so this is always an internal error
		err = fmt.Errorf("Unimplemented action ID %q - this is always an internal error", pageRuleAction.ID)
	}
	return
}

func transformToCloudflarePageRuleAction(ctx context.Context, id string, value interface{}, d *schema.ResourceData) (pageRuleAction cloudflare.PageRuleAction, err error) {
	pageRuleAction.ID = id

	if strValue, ok := value.(string); ok {
		if id == "browser_cache_ttl" {
			intValue, err := strconv.Atoi(strValue)
			if err == nil {
				pageRuleAction.Value = intValue
			}
		} else if strValue == "" {
			pageRuleAction.Value = nil
		} else {
			pageRuleAction.Value = strValue
		}
	} else if unitValue, ok := value.(bool); ok {
		if !unitValue {
			if contains(pageRuleAPIOnOffFields, id) {
				pageRuleAction.Value = "off"
			} else {
				pageRuleAction.Value = nil
			}
		} else {
			if contains(pageRuleAPIOnOffFields, id) {
				pageRuleAction.Value = "on"
			} else {
				pageRuleAction.Value = true
			}
		}
	} else if intValue, ok := value.(int); ok {
		if id == "edge_cache_ttl" && intValue > 0 {
			pageRuleAction.Value = intValue
		} else {
			pageRuleAction.Value = nil
		}
	} else if id == "forwarding_url" {
		forwardActionSchema := value.([]interface{})

		log.Printf("[DEBUG] forwarding_url action to be applied: %#v", forwardActionSchema)

		if len(forwardActionSchema) != 0 {
			fwd := forwardActionSchema[0].(map[string]interface{})

			pageRuleAction.Value = map[string]interface{}{
				"url":         fwd["url"].(string),
				"status_code": fwd["status_code"].(int),
			}
		}
	} else if id == "minify" {
		minifyActionSchema := value.([]interface{})

		log.Printf("[DEBUG] minify action to be applied: %#v", minifyActionSchema)

		if len(minifyActionSchema) != 0 {
			minify := minifyActionSchema[0].(map[string]interface{})

			pageRuleAction.Value = map[string]interface{}{
				"css":  minify["css"].(string),
				"js":   minify["js"].(string),
				"html": minify["html"].(string),
			}
		}
	} else if id == "cache_key_fields" {
		cacheKeyActionSchema := value.([]interface{})

		log.Printf("[DEBUG] cache_key_fields action to be applied: %#v", cacheKeyActionSchema)

		if len(cacheKeyActionSchema) != 0 {
			output := make(map[string]interface{})

			for sectionID, sectionValue := range cacheKeyActionSchema[0].(map[string]interface{}) {
				sectionOutput := map[string]interface{}{}

				switch sectionID {
				case "cookie", "header":
					if len(sectionValue.([]interface{})) > 0 && sectionValue.([]interface{})[0] != nil {
						for fieldID, fieldValue := range sectionValue.([]interface{})[0].(map[string]interface{}) {
							sectionOutput[fieldID] = fieldValue.(*schema.Set).List()
						}
						output[sectionID] = sectionOutput
					}
				case "query_string":
					if len(sectionValue.([]interface{})) > 0 && sectionValue.([]interface{})[0] != nil {
						for fieldID, fieldValue := range sectionValue.([]interface{})[0].(map[string]interface{}) {
							switch fieldID {
							case "exclude", "include":
								if fieldValue.(*schema.Set).Len() > 0 {
									sectionOutput[fieldID] = fieldValue.(*schema.Set).List()
								}
							case "ignore":
								sectionOutput[fieldID] = fieldValue
							default:
								sectionOutput[fieldID] = fieldValue.(*schema.Set).List()
							}
						}

						if sectionOutput["ignore"].(bool) {
							sectionOutput["exclude"] = "*"
						}
					}

					exclude, ok1 := sectionOutput["exclude"]
					include, ok2 := sectionOutput["include"]
					ignore := sectionOutput["ignore"]

					// Ensure that if no `include`, `exclude` or `ignore` attributes are
					// set, we default to including all query string parameters in the
					// cache key.
					if ignore == nil || !ignore.(bool) {
						if (!ok1 || len(exclude.([]interface{})) == 0) && (!ok2 || len(include.([]interface{})) == 0) {
							sectionOutput["include"] = "*"
						}
					}

					// Clean up the payload and ensure we don't send `ignore` property
					// despite using it in the schema.
					delete(sectionOutput, "ignore")

					output[sectionID] = sectionOutput
				default:
					if len(sectionValue.([]interface{})) > 0 {
						for fieldID, fieldValue := range sectionValue.([]interface{})[0].(map[string]interface{}) {
							sectionOutput[fieldID] = fieldValue
						}
						output[sectionID] = sectionOutput
					}
				}
			}

			pageRuleAction.Value = output
		}
	} else if id == "cache_ttl_by_status" {
		cacheTTLActionSchema := value.(*schema.Set)

		tflog.Debug(ctx, fmt.Sprintf("cache_ttl_by_status action to be applied: %#v", cacheTTLActionSchema))

		if cacheTTLActionSchema.Len() != 0 {
			output := make(map[string]int)

			for _, code := range cacheTTLActionSchema.List() {
				code := code.(map[string]interface{})
				output[code["codes"].(string)] = code["ttl"].(int)
			}

			pageRuleAction.Value = output
		}
	} else {
		err = fmt.Errorf("Bad value for %s: %s", id, value)
	}

	tflog.Debug(ctx, fmt.Sprintf("Page Rule Action to be applied: %#v", pageRuleAction))

	return
}

func resourceCloudflarePageRuleImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	// split the id so we can lookup
	idAttr := strings.SplitN(d.Id(), "/", 2)
	var zoneID string
	var pageRuleID string
	if len(idAttr) == 2 {
		zoneID = idAttr[0]
		pageRuleID = idAttr[1]
		d.Set(consts.ZoneIDSchemaKey, zoneID)
		d.SetId(pageRuleID)
	} else {
		return nil, fmt.Errorf("invalid id (%q) specified, should be in format \"zoneID/pageRuleID\"", d.Id())
	}

	resourceCloudflarePageRuleRead(ctx, d, meta)

	return []*schema.ResourceData{d}, nil
}
