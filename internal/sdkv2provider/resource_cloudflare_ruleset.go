package sdkv2provider

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"regexp"
	"sort"
	"strings"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
)

const (
	accountLevelRulesetDeleteURL = "https://api.cloudflare.com/#account-rulesets-delete-account-ruleset"
	zoneLevelRulesetDeleteURL    = "https://api.cloudflare.com/#zone-rulesets-delete-zone-ruleset"
	duplicateRulesetError        = "failed to create ruleset %q as a similar configuration with rules already exists and overwriting will have unintended consequences. If you are migrating from the Dashboard, you will need to first remove the existing rules otherwise you can remove the existing phase yourself using the API (%s)."
)

func resourceCloudflareRuleset() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareRulesetSchema(),
		CreateContext: resourceCloudflareRulesetCreate,
		ReadContext:   resourceCloudflareRulesetRead,
		UpdateContext: resourceCloudflareRulesetUpdate,
		DeleteContext: resourceCloudflareRulesetDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceCloudflareRulesetImport,
		},
		SchemaVersion: 1,
		StateUpgraders: []schema.StateUpgrader{
			{
				Type:    resourceCloudflareRulesetSchemaV0().CoreConfigSchema().ImpliedType(),
				Upgrade: resourceCloudflareRulesetStateUpgradeV0ToV1,
				Version: 0,
			},
		},
		Description: heredoc.Doc(`
			The [Cloudflare Ruleset Engine](https://developers.cloudflare.com/firewall/cf-rulesets)
			allows you to create and deploy rules and rulesets.

			The engine syntax, inspired by the Wireshark Display Filter language, is the
			same syntax used in custom Firewall Rules. Cloudflare uses the Ruleset Engine
			in different products, allowing you to configure several products using the same
			basic syntax.
		`),
	}
}

func resourceCloudflareRulesetCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)
	rulesetPhase := d.Get("phase").(string)

	var ruleset cloudflare.Ruleset
	var sempahoreErr error
	if accountID != "" {
		ruleset, sempahoreErr = client.GetAccountRulesetPhase(ctx, accountID, rulesetPhase)
	} else {
		ruleset, sempahoreErr = client.GetZoneRulesetPhase(ctx, zoneID, rulesetPhase)
	}

	if len(ruleset.Rules) > 0 {
		deleteRulesetURL := accountLevelRulesetDeleteURL
		if accountID == "" {
			deleteRulesetURL = zoneLevelRulesetDeleteURL
		}
		return diag.FromErr(fmt.Errorf(duplicateRulesetError, rulesetPhase, deleteRulesetURL))
	}

	rulesetName := d.Get("name").(string)
	rulesetDescription := d.Get("description").(string)
	rulesetKind := d.Get("kind").(string)
	rs := cloudflare.Ruleset{
		Name:        rulesetName,
		Description: rulesetDescription,
		Kind:        rulesetKind,
		Phase:       rulesetPhase,
	}

	rules, err := buildRulesetRulesFromResource(d)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error building ruleset rules from resource: %w", err))
	}

	if len(rules) > 0 {
		rs.Rules = rules
	}

	if sempahoreErr == nil && len(ruleset.Rules) == 0 && ruleset.Description == "" {
		log.Print("[DEBUG] default ruleset created by the UI with empty rules found, recreating from scratch")
		var deleteRulesetErr error
		if accountID != "" {
			deleteRulesetErr = client.DeleteAccountRuleset(ctx, accountID, ruleset.ID)
		} else {
			deleteRulesetErr = client.DeleteZoneRuleset(ctx, zoneID, ruleset.ID)
		}

		if deleteRulesetErr != nil {
			return diag.FromErr(fmt.Errorf("failed to delete ruleset: %w", deleteRulesetErr))
		}
	}

	var rulesetCreateErr error
	if accountID != "" {
		ruleset, rulesetCreateErr = client.CreateAccountRuleset(ctx, accountID, rs)
	} else {
		ruleset, rulesetCreateErr = client.CreateZoneRuleset(ctx, zoneID, rs)
	}

	if rulesetCreateErr != nil {
		return diag.FromErr(fmt.Errorf("error creating ruleset %s: %w", rulesetName, rulesetCreateErr))
	}

	rulesetEntryPoint := cloudflare.Ruleset{
		Description: rulesetDescription,
		Rules:       rules,
	}

	// For "custom" rulesets, we don't send a follow up PUT it to the entrypoint
	// endpoint.
	if rulesetKind != string(cloudflare.RulesetKindCustom) {
		if accountID != "" {
			_, err = client.UpdateAccountRulesetPhase(ctx, accountID, rulesetPhase, rulesetEntryPoint)
		} else {
			_, err = client.UpdateZoneRulesetPhase(ctx, zoneID, rulesetPhase, rulesetEntryPoint)
		}

		if err != nil {
			return diag.FromErr(fmt.Errorf("error updating ruleset phase entrypoint %s: %w", rulesetName, err))
		}
	}

	d.SetId(ruleset.ID)

	return resourceCloudflareRulesetRead(ctx, d, meta)
}

func resourceCloudflareRulesetImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	attributes := strings.SplitN(d.Id(), "/", 3)

	if len(attributes) != 3 {
		return nil, fmt.Errorf(`invalid id (%q) specified, should be in format "resourceType/resourceTypeID/rulesetID`, d.Id())
	}

	resourceType, resourceTypeID, rulesetID := attributes[0], attributes[1], attributes[2]

	if resourceType == "account" {
		d.Set(consts.AccountIDSchemaKey, resourceTypeID)
	} else {
		d.Set(consts.ZoneIDSchemaKey, resourceTypeID)
	}
	d.SetId(rulesetID)

	resourceCloudflareRulesetRead(ctx, d, meta)

	return []*schema.ResourceData{d}, nil
}

func resourceCloudflareRulesetRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	var ruleset cloudflare.Ruleset
	var err error

	if accountID != "" {
		ruleset, err = client.GetAccountRuleset(ctx, accountID, d.Id())
	} else {
		ruleset, err = client.GetZoneRuleset(ctx, zoneID, d.Id())
	}

	if err != nil {
		if strings.Contains(err.Error(), "could not find ruleset") {
			log.Printf("[INFO] Ruleset %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return diag.FromErr(fmt.Errorf("error reading ruleset ID %q: %w", d.Id(), err))
	}

	d.Set("name", ruleset.Name)
	d.Set("description", ruleset.Description)
	d.Set("kind", ruleset.Kind)
	d.Set("phase", ruleset.Phase)

	if err := d.Set("rules", buildStateFromRulesetRules(ruleset.Rules)); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceCloudflareRulesetUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	rules, err := buildRulesetRulesFromResource(d)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error building ruleset from resource: %w", err))
	}

	description := d.Get("description").(string)
	if accountID != "" {
		_, err = client.UpdateAccountRuleset(ctx, accountID, d.Id(), description, rules)
	} else {
		_, err = client.UpdateZoneRuleset(ctx, zoneID, d.Id(), description, rules)
	}

	if err != nil {
		return diag.FromErr(fmt.Errorf("error updating ruleset with ID %q: %w", d.Id(), err))
	}

	return resourceCloudflareRulesetRead(ctx, d, meta)
}

func resourceCloudflareRulesetDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)
	var err error

	if accountID != "" {
		err = client.DeleteAccountRuleset(ctx, accountID, d.Id())
	} else {
		err = client.DeleteZoneRuleset(ctx, zoneID, d.Id())
	}

	if err != nil {
		return diag.FromErr(fmt.Errorf("error deleting ruleset with ID %q: %w", d.Id(), err))
	}

	return nil
}

// buildStateFromRulesetRules receives the current ruleset rules and returns an
// interface for the state file.
func buildStateFromRulesetRules(rules []cloudflare.RulesetRule) interface{} {
	var rulesData []map[string]interface{}
	for _, r := range rules {
		rule := map[string]interface{}{
			"id":         r.ID,
			"expression": r.Expression,
			"action":     r.Action,
			"enabled":    r.Enabled,
		}

		if r.Description != "" {
			rule["description"] = r.Description
		}

		if !reflect.ValueOf(r.ActionParameters).IsNil() {
			var (
				actionParameters       []map[string]interface{}
				overrides              []map[string]interface{}
				idBasedOverrides       []map[string]interface{}
				categoryBasedOverrides []map[string]interface{}
				headers                []map[string]interface{}
				uri                    []map[string]interface{}
				matchedData            []map[string]interface{}
				response               []map[string]interface{}
				origin                 []map[string]interface{}
				sni                    []map[string]interface{}
				requestFields          []string
				responseFields         []string
				cookieFields           []string
				edgeTTLFields          []map[string]interface{}
				browserTTLFields       []map[string]interface{}
				serveStaleFields       []map[string]interface{}
				cacheKeyFields         []map[string]interface{}
				fromListFields         []map[string]interface{}
				fromValueFields        []map[string]interface{}
				autoMinifyFields       []map[string]interface{}
				polishSetting          string
				sslSetting             string
				securityLevel          string
			)
			actionParameterRules := make(map[string]string)

			if !reflect.ValueOf(r.ActionParameters.Overrides).IsNil() {
				for _, overrideRule := range r.ActionParameters.Overrides.Rules {
					idBasedOverrides = append(idBasedOverrides, map[string]interface{}{
						"id":                overrideRule.ID,
						"action":            overrideRule.Action,
						"status":            apiEnabledToStatusFieldConversion(overrideRule.Enabled),
						"score_threshold":   overrideRule.ScoreThreshold,
						"sensitivity_level": overrideRule.SensitivityLevel,
					})
				}

				for _, overrideRule := range r.ActionParameters.Overrides.Categories {
					categoryBasedOverrides = append(categoryBasedOverrides, map[string]interface{}{
						"category": overrideRule.Category,
						"action":   overrideRule.Action,
						"status":   apiEnabledToStatusFieldConversion(overrideRule.Enabled),
					})
				}

				overrides = append(overrides, map[string]interface{}{
					"categories":        categoryBasedOverrides,
					"rules":             idBasedOverrides,
					"status":            apiEnabledToStatusFieldConversion(r.ActionParameters.Overrides.Enabled),
					"action":            r.ActionParameters.Overrides.Action,
					"sensitivity_level": r.ActionParameters.Overrides.SensitivityLevel,
				})
			}

			if !reflect.ValueOf(r.ActionParameters.Headers).IsNil() {
				sortedHeaders := make([]string, 0, len(r.ActionParameters.Headers))
				for headerName := range r.ActionParameters.Headers {
					sortedHeaders = append(sortedHeaders, headerName)
				}
				sort.Strings(sortedHeaders)

				for _, headerName := range sortedHeaders {
					headers = append(headers, map[string]interface{}{
						"name":       headerName,
						"value":      r.ActionParameters.Headers[headerName].Value,
						"expression": r.ActionParameters.Headers[headerName].Expression,
						"operation":  r.ActionParameters.Headers[headerName].Operation,
					})
				}
			}

			if !reflect.ValueOf(r.ActionParameters.URI).IsNil() {
				var query []map[string]interface{}
				var path []map[string]interface{}
				originValue := false
				if r.ActionParameters.URI.Origin {
					originValue = true
				}

				if !reflect.ValueOf(r.ActionParameters.URI.Query).IsNil() {
					query = append(query, map[string]interface{}{
						"value":      r.ActionParameters.URI.Query.Value,
						"expression": r.ActionParameters.URI.Query.Expression,
					})
				}

				if !reflect.ValueOf(r.ActionParameters.URI.Path).IsNil() {
					path = append(path, map[string]interface{}{
						"value":      r.ActionParameters.URI.Path.Value,
						"expression": r.ActionParameters.URI.Path.Expression,
					})
				}

				uri = append(uri, map[string]interface{}{
					"origin": originValue,
					"query":  query,
					"path":   path,
				})
			}

			if !reflect.ValueOf(r.ActionParameters.MatchedData).IsNil() {
				matchedData = append(matchedData, map[string]interface{}{
					"public_key": r.ActionParameters.MatchedData.PublicKey,
				})
			}

			if !reflect.ValueOf(r.ActionParameters.Rules).IsNil() {
				for k, v := range r.ActionParameters.Rules {
					actionParameterRules[k] = strings.Join(v, ",")
				}
			}

			if !reflect.ValueOf(r.ActionParameters.Response).IsNil() {
				response = append(response, map[string]interface{}{
					"status_code":  r.ActionParameters.Response.StatusCode,
					"content_type": r.ActionParameters.Response.ContentType,
					"content":      r.ActionParameters.Response.Content,
				})
			}

			if !reflect.ValueOf(r.ActionParameters.Origin).IsNil() {
				origin = append(origin, map[string]interface{}{
					"host": r.ActionParameters.Origin.Host,
					"port": r.ActionParameters.Origin.Port,
				})
			}

			if !reflect.ValueOf(r.ActionParameters.SNI).IsNil() {
				sni = append(sni, map[string]interface{}{
					"value": r.ActionParameters.SNI.Value,
				})
			}

			if !reflect.ValueOf(r.ActionParameters.RequestFields).IsNil() {
				requestFields = make([]string, 0)
				for _, v := range r.ActionParameters.RequestFields {
					requestFields = append(requestFields, v.Name)
				}
			}

			if !reflect.ValueOf(r.ActionParameters.ResponseFields).IsNil() {
				responseFields = make([]string, 0)
				for _, v := range r.ActionParameters.ResponseFields {
					responseFields = append(responseFields, v.Name)
				}
			}

			if !reflect.ValueOf(r.ActionParameters.CookieFields).IsNil() {
				cookieFields = make([]string, 0)
				for _, v := range r.ActionParameters.CookieFields {
					cookieFields = append(cookieFields, v.Name)
				}
			}

			if !reflect.ValueOf(r.ActionParameters.EdgeTTL).IsNil() {
				edgeTTL := map[string]interface{}{
					"mode":    r.ActionParameters.EdgeTTL.Mode,
					"default": r.ActionParameters.EdgeTTL.Default,
				}
				if !reflect.ValueOf(r.ActionParameters.EdgeTTL.StatusCodeTTL).IsNil() {
					edgeTTL["status_code_ttl"] = []interface{}{}
					for _, sc := range r.ActionParameters.EdgeTTL.StatusCodeTTL {
						scTTL := map[string]interface{}{
							"status_code": sc.StatusCodeValue,
							"value":       sc.Value,
						}
						if !reflect.ValueOf(sc.StatusCodeRange).IsNil() {
							scTTL["status_code_range"] = []interface{}{map[string]interface{}{
								"from": sc.StatusCodeRange.From,
								"to":   sc.StatusCodeRange.To,
							}}
						}
						edgeTTL["status_code_ttl"] = append(edgeTTL["status_code_ttl"].([]interface{}), scTTL)
					}
				}
				edgeTTLFields = append(edgeTTLFields, edgeTTL)
			}

			if !reflect.ValueOf(r.ActionParameters.BrowserTTL).IsNil() {
				browserTTLFields = append(browserTTLFields, map[string]interface{}{
					"mode":    r.ActionParameters.BrowserTTL.Mode,
					"default": r.ActionParameters.BrowserTTL.Default,
				})
			}

			if !reflect.ValueOf(r.ActionParameters.ServeStale).IsNil() {
				serveStaleFields = append(serveStaleFields, map[string]interface{}{
					"disable_stale_while_updating": r.ActionParameters.ServeStale.DisableStaleWhileUpdating,
				})
			}

			if !reflect.ValueOf(r.ActionParameters.CacheKey).IsNil() {
				cacheKey := map[string]interface{}{
					"cache_by_device_type":       r.ActionParameters.CacheKey.CacheByDeviceType,
					"ignore_query_strings_order": r.ActionParameters.CacheKey.IgnoreQueryStringsOrder,
					"cache_deception_armor":      r.ActionParameters.CacheKey.CacheDeceptionArmor,
				}
				if !reflect.ValueOf(r.ActionParameters.CacheKey.CustomKey).IsNil() {
					customKey := map[string]interface{}{}
					if !reflect.ValueOf(r.ActionParameters.CacheKey.CustomKey.Query).IsNil() {
						query := map[string]interface{}{}
						if !reflect.ValueOf(r.ActionParameters.CacheKey.CustomKey.Query.Include).IsNil() {
							query["include"] = []interface{}{}
							for _, v := range r.ActionParameters.CacheKey.CustomKey.Query.Include.List {
								query["include"] = append(query["include"].([]interface{}), v)
							}
							if r.ActionParameters.CacheKey.CustomKey.Query.Include.All {
								query["include"] = append(query["include"].([]interface{}), "*")
							}
						}
						if !reflect.ValueOf(r.ActionParameters.CacheKey.CustomKey.Query.Exclude).IsNil() {
							query["exclude"] = []interface{}{}
							for _, v := range r.ActionParameters.CacheKey.CustomKey.Query.Exclude.List {
								query["exclude"] = append(query["exclude"].([]interface{}), v)
							}
							if r.ActionParameters.CacheKey.CustomKey.Query.Exclude.All {
								query["exclude"] = append(query["exclude"].([]interface{}), "*")
							}
						}
						customKey["query_string"] = []interface{}{query}
					}
					if !reflect.ValueOf(r.ActionParameters.CacheKey.CustomKey.Header).IsNil() {
						header := map[string]interface{}{
							"exclude_origin": r.ActionParameters.CacheKey.CustomKey.Header.ExcludeOrigin,
							"include":        []interface{}{},
							"check_presence": []interface{}{},
						}
						for _, h := range r.ActionParameters.CacheKey.CustomKey.Header.Include {
							header["include"] = append(header["include"].([]interface{}), h)
						}
						for _, h := range r.ActionParameters.CacheKey.CustomKey.Header.CheckPresence {
							header["check_presence"] = append(header["check_presence"].([]interface{}), h)
						}
						customKey["header"] = []interface{}{header}
					}
					if !reflect.ValueOf(r.ActionParameters.CacheKey.CustomKey.Cookie).IsNil() {
						cookie := map[string]interface{}{
							"include":        []interface{}{},
							"check_presence": []interface{}{},
						}
						for _, h := range r.ActionParameters.CacheKey.CustomKey.Cookie.Include {
							cookie["include"] = append(cookie["include"].([]interface{}), h)
						}
						for _, h := range r.ActionParameters.CacheKey.CustomKey.Cookie.CheckPresence {
							cookie["check_presence"] = append(cookie["check_presence"].([]interface{}), h)
						}
						customKey["cookie"] = []interface{}{cookie}
					}
					if !reflect.ValueOf(r.ActionParameters.CacheKey.CustomKey.User).IsNil() {
						customKey["user"] = []interface{}{map[string]interface{}{
							"device_type": r.ActionParameters.CacheKey.CustomKey.User.DeviceType,
							"geo":         r.ActionParameters.CacheKey.CustomKey.User.Geo,
							"lang":        r.ActionParameters.CacheKey.CustomKey.User.Lang,
						}}
					}
					if !reflect.ValueOf(r.ActionParameters.CacheKey.CustomKey.Host).IsNil() {
						customKey["host"] = []interface{}{map[string]interface{}{
							"resolved": r.ActionParameters.CacheKey.CustomKey.Host.Resolved,
						}}
					}
					cacheKey["custom_key"] = []interface{}{customKey}
				}
				cacheKeyFields = append(cacheKeyFields, cacheKey)
			}

			if !reflect.ValueOf(r.ActionParameters.FromList).IsNil() {
				fromListFields = append(fromListFields, map[string]interface{}{
					"name": r.ActionParameters.FromList.Name,
					"key":  r.ActionParameters.FromList.Key,
				})
			}

			if !reflect.ValueOf(r.ActionParameters.FromValue).IsNil() {
				fromValueFields = append(fromValueFields, map[string]interface{}{
					"status_code": r.ActionParameters.FromValue.StatusCode,
					"target_url": []interface{}{map[string]interface{}{
						"value":      r.ActionParameters.FromValue.TargetURL.Value,
						"expression": r.ActionParameters.FromValue.TargetURL.Expression,
					}},
					"preserve_query_string": r.ActionParameters.FromValue.PreserveQueryString,
				})
			}

			if !reflect.ValueOf(r.ActionParameters.AutoMinify).IsNil() {
				autoMinifyFields = append(autoMinifyFields, map[string]interface{}{
					"html": r.ActionParameters.AutoMinify.HTML,
					"css":  r.ActionParameters.AutoMinify.CSS,
					"js":   r.ActionParameters.AutoMinify.JS,
				})
			}

			if !reflect.ValueOf(r.ActionParameters.Polish).IsNil() {
				polishSetting = r.ActionParameters.Polish.String()
			}

			if !reflect.ValueOf(r.ActionParameters.SecurityLevel).IsNil() {
				securityLevel = r.ActionParameters.SecurityLevel.String()
			}

			if !reflect.ValueOf(r.ActionParameters.SSL).IsNil() {
				sslSetting = r.ActionParameters.SSL.String()
			}

			actionParameters = append(actionParameters, map[string]interface{}{
				"id":                         r.ActionParameters.ID,
				"increment":                  r.ActionParameters.Increment,
				"headers":                    headers,
				"overrides":                  overrides,
				"products":                   r.ActionParameters.Products,
				"phases":                     r.ActionParameters.Phases,
				"ruleset":                    r.ActionParameters.Ruleset,
				"rulesets":                   r.ActionParameters.Rulesets,
				"rules":                      actionParameterRules,
				"uri":                        uri,
				"matched_data":               matchedData,
				"response":                   response,
				"version":                    r.ActionParameters.Version,
				"host_header":                r.ActionParameters.HostHeader,
				"sni":                        sni,
				"origin":                     origin,
				"request_fields":             requestFields,
				"response_fields":            responseFields,
				"cookie_fields":              cookieFields,
				"cache":                      r.ActionParameters.Cache,
				"edge_ttl":                   edgeTTLFields,
				"browser_ttl":                browserTTLFields,
				"serve_stale":                serveStaleFields,
				"respect_strong_etags":       r.ActionParameters.RespectStrongETags,
				"cache_key":                  cacheKeyFields,
				"origin_error_page_passthru": r.ActionParameters.OriginErrorPagePassthru,
				"from_list":                  fromListFields,
				"from_value":                 fromValueFields,
				"content":                    r.ActionParameters.Content,
				"content_type":               r.ActionParameters.ContentType,
				"status_code":                r.ActionParameters.StatusCode,
				"automatic_https_rewrites":   r.ActionParameters.AutomaticHTTPSRewrites,
				"autominify":                 autoMinifyFields,
				"bic":                        r.ActionParameters.BrowserIntegrityCheck,
				"disable_apps":               r.ActionParameters.DisableApps,
				"disable_zaraz":              r.ActionParameters.DisableZaraz,
				"disable_railgun":            r.ActionParameters.DisableRailgun,
				"email_obfuscation":          r.ActionParameters.EmailObfuscation,
				"mirage":                     r.ActionParameters.Mirage,
				"opportunistic_encryption":   r.ActionParameters.OpportunisticEncryption,
				"polish":                     polishSetting,
				"rocket_loader":              r.ActionParameters.RocketLoader,
				"security_level":             securityLevel,
				"server_side_excludes":       r.ActionParameters.ServerSideExcludes,
				"ssl":                        sslSetting,
				"sxg":                        r.ActionParameters.SXG,
				"hotlink_protection":         r.ActionParameters.HotLinkProtection,
			})

			rule["action_parameters"] = actionParameters
		}

		if !reflect.ValueOf(r.RateLimit).IsNil() {
			var rateLimit []map[string]interface{}

			rateLimit = append(rateLimit, map[string]interface{}{
				"characteristics":            r.RateLimit.Characteristics,
				"period":                     r.RateLimit.Period,
				"requests_per_period":        r.RateLimit.RequestsPerPeriod,
				"score_per_period":           r.RateLimit.ScorePerPeriod,
				"score_response_header_name": r.RateLimit.ScoreResponseHeaderName,
				"mitigation_timeout":         r.RateLimit.MitigationTimeout,
				"counting_expression":        r.RateLimit.CountingExpression,
				"requests_to_origin":         r.RateLimit.RequestsToOrigin,
			})

			rule["ratelimit"] = rateLimit
		}

		if !reflect.ValueOf(r.ExposedCredentialCheck).IsNil() {
			var exposedCredentialCheck []map[string]interface{}

			exposedCredentialCheck = append(exposedCredentialCheck, map[string]interface{}{
				"username_expression": r.ExposedCredentialCheck.UsernameExpression,
				"password_expression": r.ExposedCredentialCheck.PasswordExpression,
			})

			rule["exposed_credential_check"] = exposedCredentialCheck
		}

		if !reflect.ValueOf(r.Logging).IsNil() {
			var logging []map[string]interface{}

			logging = append(logging, map[string]interface{}{
				"status": apiEnabledToStatusFieldConversion(r.Logging.Enabled),
			})

			rule["logging"] = logging
		}

		rulesData = append(rulesData, rule)
	}

	return rulesData
}

func buildRule(d *schema.ResourceData, resourceRule map[string]interface{}, rulesCounter int) cloudflare.RulesetRule {
	var rule cloudflare.RulesetRule

	if len(resourceRule["action_parameters"].([]interface{})) > 0 {
		rule.ActionParameters = &cloudflare.RulesetRuleActionParameters{}
		for _, parameter := range resourceRule["action_parameters"].([]interface{}) {
			for pKey, pValue := range parameter.(map[string]interface{}) {
				switch pKey {
				case "id":
					rule.ActionParameters.ID = pValue.(string)
				case "version":
					if rule.ActionParameters.Version != "" {
						rule.ActionParameters.Version = pValue.(string)
					}
				case "products":
					var products []string
					for _, product := range pValue.(*schema.Set).List() {
						products = append(products, product.(string))
					}

					if len(products) > 0 {
						rule.ActionParameters.Products = products
					}
				case "phases":
					var phases []string
					for _, phase := range pValue.(*schema.Set).List() {
						phases = append(phases, phase.(string))
					}

					if len(phases) > 0 {
						rule.ActionParameters.Phases = phases
					}
				case "ruleset":
					rule.ActionParameters.Ruleset = pValue.(string)
				case "rulesets":
					var rulesetsValues []string
					for _, v := range pValue.(*schema.Set).List() {
						rulesetsValues = append(rulesetsValues, v.(string))
					}
					rule.ActionParameters.Rulesets = rulesetsValues
				case "rules":
					apRules := make(map[string][]string)
					for name, data := range pValue.(map[string]interface{}) {
						// regex (not string.Split) needs to be used here to account for
						// whitespace from the end user in the map value.
						split := regexp.MustCompile("\\s*,\\s*").Split(data.(string), -1)
						ruleValues := []string{}

						for i := range split {
							ruleValues = append(ruleValues, split[i])
						}

						apRules[name] = ruleValues
					}

					rule.ActionParameters.Rules = apRules
				case "increment":
					rule.ActionParameters.Increment = pValue.(int)
				case "overrides":
					var overrideConfiguration cloudflare.RulesetRuleActionParametersOverrides
					var categories []cloudflare.RulesetRuleActionParametersCategories
					var rules []cloudflare.RulesetRuleActionParametersRules

					for overrideCounter, overrideParamValue := range pValue.([]interface{}) {
						if value, ok := d.GetOk(fmt.Sprintf("rules.%d.action_parameters.0.overrides.%d.status", rulesCounter, overrideCounter)); ok {
							if value.(string) != "" {
								overrideConfiguration.Enabled = statusToAPIEnabledFieldConversion(value.(string))
							}
						}

						if val, ok := overrideParamValue.(map[string]interface{})["action"]; ok {
							overrideConfiguration.Action = val.(string)
						}

						if val, ok := overrideParamValue.(map[string]interface{})["sensitivity_level"]; ok {
							overrideConfiguration.SensitivityLevel = val.(string)
						}

						// Category based overrides
						if val, ok := overrideParamValue.(map[string]interface{})["categories"]; ok {
							for categoryCounter, category := range val.([]interface{}) {
								cData := category.(map[string]interface{})
								categoryOverride := cloudflare.RulesetRuleActionParametersCategories{
									Category: cData["category"].(string),
									Action:   cData["action"].(string),
								}

								if value, ok := d.GetOk(fmt.Sprintf("rules.%d.action_parameters.0.overrides.%d.categories.%d.status", rulesCounter, overrideCounter, categoryCounter)); ok {
									if value != "" {
										categoryOverride.Enabled = statusToAPIEnabledFieldConversion(value.(string))
									}
								}

								categories = append(categories, categoryOverride)
							}
						}

						// Rule ID based overrides
						if val, ok := overrideParamValue.(map[string]interface{})["rules"]; ok {
							for ruleOverrideCounter, rule := range val.([]interface{}) {
								rData := rule.(map[string]interface{})
								ruleOverride := cloudflare.RulesetRuleActionParametersRules{
									ID:               rData["id"].(string),
									Action:           rData["action"].(string),
									ScoreThreshold:   rData["score_threshold"].(int),
									SensitivityLevel: rData["sensitivity_level"].(string),
								}

								if value, ok := d.GetOk(fmt.Sprintf("rules.%d.action_parameters.0.overrides.%d.rules.%d.status", rulesCounter, overrideCounter, ruleOverrideCounter)); ok {
									if value != "" {
										ruleOverride.Enabled = statusToAPIEnabledFieldConversion(value.(string))
									}
								}

								rules = append(rules, ruleOverride)
							}
						}
					}

					if len(categories) > 0 || len(rules) > 0 {
						overrideConfiguration.Categories = categories
						overrideConfiguration.Rules = rules
					}

					if !reflect.DeepEqual(overrideConfiguration, cloudflare.RulesetRuleActionParametersOverrides{}) {
						rule.ActionParameters.Overrides = &overrideConfiguration
					}

				case "matched_data":
					for i := range pValue.([]interface{}) {
						rule.ActionParameters.MatchedData = &cloudflare.RulesetRuleActionParametersMatchedData{
							PublicKey: pValue.([]interface{})[i].(map[string]interface{})["public_key"].(string),
						}
					}

				case "response":
					for i := range pValue.([]interface{}) {
						rule.ActionParameters.Response = &cloudflare.RulesetRuleActionParametersBlockResponse{
							StatusCode:  uint16(pValue.([]interface{})[i].(map[string]interface{})["status_code"].(int)),
							ContentType: pValue.([]interface{})[i].(map[string]interface{})["content_type"].(string),
							Content:     pValue.([]interface{})[i].(map[string]interface{})["content"].(string),
						}
					}

				case "uri":
					var uriParameterConfig cloudflare.RulesetRuleActionParametersURI
					for _, uriValue := range pValue.([]interface{}) {
						if val, ok := uriValue.(map[string]interface{})["path"]; ok && len(val.([]interface{})) > 0 {
							uriPathConfig := val.([]interface{})[0].(map[string]interface{})
							uriParameterConfig.Path = &cloudflare.RulesetRuleActionParametersURIPath{
								Value:      uriPathConfig["value"].(string),
								Expression: uriPathConfig["expression"].(string),
							}
						}

						if val, ok := uriValue.(map[string]interface{})["query"]; ok && len(val.([]interface{})) > 0 {
							uriQueryConfig := val.([]interface{})[0].(map[string]interface{})
							uriParameterConfig.Query = &cloudflare.RulesetRuleActionParametersURIQuery{
								Value:      uriQueryConfig["value"].(string),
								Expression: uriQueryConfig["expression"].(string),
							}
						}

						rule.ActionParameters.URI = &uriParameterConfig
					}

				case "headers":
					headers := make(map[string]cloudflare.RulesetRuleActionParametersHTTPHeader)
					for _, headerList := range pValue.([]interface{}) {
						name := headerList.(map[string]interface{})["name"].(string)

						headers[name] = cloudflare.RulesetRuleActionParametersHTTPHeader{
							Value:      headerList.(map[string]interface{})["value"].(string),
							Expression: headerList.(map[string]interface{})["expression"].(string),
							Operation:  headerList.(map[string]interface{})["operation"].(string),
						}
					}

					rule.ActionParameters.Headers = headers

				case "content":
					rule.ActionParameters.Content = pValue.(string)

				case "content_type":
					rule.ActionParameters.ContentType = pValue.(string)

				case "status_code":
					rule.ActionParameters.StatusCode = uint16(pValue.(int))

				case "host_header":
					rule.ActionParameters.HostHeader = pValue.(string)

				case "origin":
					for i := range pValue.([]interface{}) {
						rule.ActionParameters.Origin = &cloudflare.RulesetRuleActionParametersOrigin{
							Host: pValue.([]interface{})[i].(map[string]interface{})["host"].(string),
							Port: uint16(pValue.([]interface{})[i].(map[string]interface{})["port"].(int)),
						}
					}
				case "automatic_https_rewrites":
					if value, ok := d.GetOk(fmt.Sprintf("rules.%d.action_parameters.0.automatic_https_rewrites", rulesCounter)); ok {
						rule.ActionParameters.AutomaticHTTPSRewrites = cloudflare.BoolPtr(value.(bool))
					}
				case "autominify":
					for i := range pValue.([]interface{}) {
						rule.ActionParameters.AutoMinify = &cloudflare.RulesetRuleActionParametersAutoMinify{
							HTML: pValue.([]interface{})[i].(map[string]interface{})["html"].(bool),
							CSS:  pValue.([]interface{})[i].(map[string]interface{})["css"].(bool),
							JS:   pValue.([]interface{})[i].(map[string]interface{})["js"].(bool),
						}
					}

				case "bic":
					if value, ok := d.GetOk(fmt.Sprintf("rules.%d.action_parameters.0.bic", rulesCounter)); ok {
						rule.ActionParameters.BrowserIntegrityCheck = cloudflare.BoolPtr(value.(bool))
					}
				case "disable_apps":
					if value, ok := d.GetOk(fmt.Sprintf("rules.%d.action_parameters.0.disable_apps", rulesCounter)); ok {
						rule.ActionParameters.DisableApps = cloudflare.BoolPtr(value.(bool))
					}
				case "disable_zaraz":
					if value, ok := d.GetOk(fmt.Sprintf("rules.%d.action_parameters.0.disable_zaraz", rulesCounter)); ok {
						rule.ActionParameters.DisableZaraz = cloudflare.BoolPtr(value.(bool))
					}
				case "disable_railgun":
					if value, ok := d.GetOk(fmt.Sprintf("rules.%d.action_parameters.0.disable_zaraz", rulesCounter)); ok {
						rule.ActionParameters.DisableRailgun = cloudflare.BoolPtr(value.(bool))
					}
				case "email_obfuscation":
					if value, ok := d.GetOk(fmt.Sprintf("rules.%d.action_parameters.0.email_obfuscation", rulesCounter)); ok {
						rule.ActionParameters.EmailObfuscation = cloudflare.BoolPtr(value.(bool))
					}
				case "mirage":
					if value, ok := d.GetOk(fmt.Sprintf("rules.%d.action_parameters.0.mirage", rulesCounter)); ok {
						rule.ActionParameters.Mirage = cloudflare.BoolPtr(value.(bool))
					}
				case "opportunistic_encryption":
					if value, ok := d.GetOk(fmt.Sprintf("rules.%d.action_parameters.0.opportunistic_encryption", rulesCounter)); ok {
						rule.ActionParameters.OpportunisticEncryption = cloudflare.BoolPtr(value.(bool))
					}
				case "polish":
					if value, ok := d.GetOk(fmt.Sprintf("rules.%d.action_parameters.0.polish", rulesCounter)); ok {
						p, _ := cloudflare.PolishFromString(value.(string))
						rule.ActionParameters.Polish = p
					}
				case "rocket_loader":
					if value, ok := d.GetOk(fmt.Sprintf("rules.%d.action_parameters.0.rocket_loader", rulesCounter)); ok {
						rule.ActionParameters.RocketLoader = cloudflare.BoolPtr(value.(bool))
					}
				case "security_level":
					if value, ok := d.GetOk(fmt.Sprintf("rules.%d.action_parameters.0.security_level", rulesCounter)); ok {
						sl, _ := cloudflare.SecurityLevelFromString(value.(string))
						rule.ActionParameters.SecurityLevel = sl
					}
				case "server_side_excludes":
					if value, ok := d.GetOk(fmt.Sprintf("rules.%d.action_parameters.0.server_side_excludes", rulesCounter)); ok {
						rule.ActionParameters.ServerSideExcludes = cloudflare.BoolPtr(value.(bool))
					}
				case "ssl":
					if value, ok := d.GetOk(fmt.Sprintf("rules.%d.action_parameters.0.ssl", rulesCounter)); ok {
						ssl, _ := cloudflare.SSLFromString(value.(string))
						rule.ActionParameters.SSL = ssl
					}
				case "sxg":
					if value, ok := d.GetOk(fmt.Sprintf("rules.%d.action_parameters.0.sxg", rulesCounter)); ok {
						rule.ActionParameters.SXG = cloudflare.BoolPtr(value.(bool))
					}
				case "hotlink_protection":
					if value, ok := d.GetOk(fmt.Sprintf("rules.%d.action_parameters.0.hotlink_protection", rulesCounter)); ok {
						rule.ActionParameters.HotLinkProtection = cloudflare.BoolPtr(value.(bool))
					}
				case "sni":
					for i := range pValue.([]interface{}) {
						rule.ActionParameters.SNI = &cloudflare.RulesetRuleActionParametersSni{
							Value: pValue.([]interface{})[i].(map[string]interface{})["value"].(string),
						}
					}
				case "cache":
					if value, ok := d.GetOk(fmt.Sprintf("rules.%d.action_parameters.0.cache", rulesCounter)); ok {
						rule.ActionParameters.Cache = cloudflare.BoolPtr(value.(bool))
					}

				case "edge_ttl":
					for i := range pValue.([]interface{}) {
						rule.ActionParameters.EdgeTTL = &cloudflare.RulesetRuleActionParametersEdgeTTL{}
						for pKey, pValue := range pValue.([]interface{})[i].(map[string]interface{}) {
							switch pKey {
							case "default":
								if value, ok := d.GetOk(fmt.Sprintf("rules.%d.action_parameters.0.edge_ttl.0.default", rulesCounter)); ok {
									rule.ActionParameters.EdgeTTL.Default = cloudflare.UintPtr(uint(value.(int)))
								}
							case "mode":
								rule.ActionParameters.EdgeTTL.Mode = pValue.(string)
							case "status_code_ttl":
								for statusCodesCounter := range pValue.([]interface{}) {
									sc := cloudflare.RulesetRuleActionParametersStatusCodeTTL{}
									if pValue.([]interface{})[statusCodesCounter] == nil {
										continue
									}
									for pKey, pValue := range pValue.([]interface{})[statusCodesCounter].(map[string]interface{}) {
										switch pKey {
										case "status_code":
											if value, ok := d.GetOk(fmt.Sprintf("rules.%d.action_parameters.0.edge_ttl.0.status_code_ttl.%d.status_code", rulesCounter, statusCodesCounter)); ok {
												sc.StatusCodeValue = cloudflare.UintPtr(uint(value.(int)))
											}
										case "value":
											sc.Value = cloudflare.IntPtr(pValue.(int))
										case "status_code_range":
											for i := range pValue.([]interface{}) {
												sc.StatusCodeRange = &cloudflare.RulesetRuleActionParametersStatusCodeRange{}
												if pValue.([]interface{})[i] == nil {
													continue
												}
												for pKey := range pValue.([]interface{})[i].(map[string]interface{}) {
													switch pKey {
													case "from":
														if value, ok := d.GetOk(fmt.Sprintf("rules.%d.action_parameters.0.edge_ttl.0.status_code_ttl.%d.status_code_range.0.from", rulesCounter, statusCodesCounter)); ok {
															sc.StatusCodeRange.From = cloudflare.UintPtr(uint(value.(int)))
														}
													case "to":
														if value, ok := d.GetOk(fmt.Sprintf("rules.%d.action_parameters.0.edge_ttl.0.status_code_ttl.%d.status_code_range.0.to", rulesCounter, statusCodesCounter)); ok {
															sc.StatusCodeRange.To = cloudflare.UintPtr(uint(value.(int)))
														}
													}
												}
											}
										}
									}
									rule.ActionParameters.EdgeTTL.StatusCodeTTL = append(rule.ActionParameters.EdgeTTL.StatusCodeTTL, sc)
								}
							}
						}
					}

				case "browser_ttl":
					for i := range pValue.([]interface{}) {
						rule.ActionParameters.BrowserTTL = &cloudflare.RulesetRuleActionParametersBrowserTTL{}
						for pKey, pValue := range pValue.([]interface{})[i].(map[string]interface{}) {
							switch pKey {
							case "default":
								if value, ok := d.GetOk(fmt.Sprintf("rules.%d.action_parameters.0.browser_ttl.0.default", rulesCounter)); ok {
									rule.ActionParameters.BrowserTTL.Default = cloudflare.UintPtr(uint(value.(int)))
								}
							case "mode":
								rule.ActionParameters.BrowserTTL.Mode = pValue.(string)
							}
						}
					}

				case "serve_stale":
					for i := range pValue.([]interface{}) {
						rule.ActionParameters.ServeStale = &cloudflare.RulesetRuleActionParametersServeStale{}
						if pValue.([]interface{})[i] == nil {
							continue
						}
						for pKey := range pValue.([]interface{})[i].(map[string]interface{}) {
							switch pKey {
							case "disable_stale_while_updating":
								if value, ok := d.GetOk(fmt.Sprintf("rules.%d.action_parameters.0.serve_stale.0.disable_stale_while_updating", rulesCounter)); ok {
									rule.ActionParameters.ServeStale.DisableStaleWhileUpdating = cloudflare.BoolPtr(value.(bool))
								}
							}
						}
					}

				case "respect_strong_etags":
					if value, ok := d.GetOk(fmt.Sprintf("rules.%d.action_parameters.0.respect_strong_etags", rulesCounter)); ok {
						rule.ActionParameters.RespectStrongETags = cloudflare.BoolPtr(value.(bool))
					}

				case "cache_key":
					for i := range pValue.([]interface{}) {
						rule.ActionParameters.CacheKey = &cloudflare.RulesetRuleActionParametersCacheKey{}
						if pValue.([]interface{})[i] == nil {
							continue
						}
						for pKey, pValue := range pValue.([]interface{})[i].(map[string]interface{}) {
							switch pKey {
							case "cache_by_device_type":
								if value, ok := d.GetOk(fmt.Sprintf("rules.%d.action_parameters.0.cache_key.0.cache_by_device_type", rulesCounter)); ok {
									rule.ActionParameters.CacheKey.CacheByDeviceType = cloudflare.BoolPtr(value.(bool))
								}
							case "ignore_query_strings_order":
								if value, ok := d.GetOk(fmt.Sprintf("rules.%d.action_parameters.0.cache_key.0.ignore_query_strings_order", rulesCounter)); ok {
									rule.ActionParameters.CacheKey.IgnoreQueryStringsOrder = cloudflare.BoolPtr(value.(bool))
								}
							case "cache_deception_armor":
								if value, ok := d.GetOk(fmt.Sprintf("rules.%d.action_parameters.0.cache_key.0.cache_deception_armor", rulesCounter)); ok {
									rule.ActionParameters.CacheKey.CacheDeceptionArmor = cloudflare.BoolPtr(value.(bool))
								}
							case "custom_key":
								for i := range pValue.([]interface{}) {
									rule.ActionParameters.CacheKey.CustomKey = &cloudflare.RulesetRuleActionParametersCustomKey{}
									if pValue.([]interface{})[i] == nil {
										continue
									}
									for pKey, pValue := range pValue.([]interface{})[i].(map[string]interface{}) {
										switch pKey {
										case "query_string":
											for i := range pValue.([]interface{}) {
												rule.ActionParameters.CacheKey.CustomKey.Query = &cloudflare.RulesetRuleActionParametersCustomKeyQuery{}
												if pValue.([]interface{})[i] == nil {
													continue
												}
												for pKey, pValue := range pValue.([]interface{})[i].(map[string]interface{}) {
													switch pKey {
													case "include":
														if len(pValue.([]interface{})) == 0 {
															continue
														}
														rule.ActionParameters.CacheKey.CustomKey.Query.Include = &cloudflare.RulesetRuleActionParametersCustomKeyList{}
														for i := range pValue.([]interface{}) {
															entry := pValue.([]interface{})[i].(string)
															if entry == "*" {
																rule.ActionParameters.CacheKey.CustomKey.Query.Include.All = true
																break
															}
															rule.ActionParameters.CacheKey.CustomKey.Query.Include.List = append(rule.ActionParameters.CacheKey.CustomKey.Query.Include.List, entry)
														}
													case "exclude":
														if len(pValue.([]interface{})) == 0 {
															continue
														}
														rule.ActionParameters.CacheKey.CustomKey.Query.Exclude = &cloudflare.RulesetRuleActionParametersCustomKeyList{}
														for i := range pValue.([]interface{}) {
															entry := pValue.([]interface{})[i].(string)
															if entry == "*" {
																rule.ActionParameters.CacheKey.CustomKey.Query.Exclude.All = true
																break
															}
															rule.ActionParameters.CacheKey.CustomKey.Query.Exclude.List = append(rule.ActionParameters.CacheKey.CustomKey.Query.Exclude.List, entry)
														}
													}
												}
											}
										case "header":
											for i := range pValue.([]interface{}) {
												rule.ActionParameters.CacheKey.CustomKey.Header = &cloudflare.RulesetRuleActionParametersCustomKeyHeader{}
												if pValue.([]interface{})[i] == nil {
													continue
												}
												for pKey, pValue := range pValue.([]interface{})[i].(map[string]interface{}) {
													switch pKey {
													case "include":
														for i := range pValue.([]interface{}) {
															rule.ActionParameters.CacheKey.CustomKey.Header.Include = append(rule.ActionParameters.CacheKey.CustomKey.Header.Include, pValue.([]interface{})[i].(string))
														}
													case "check_presence":
														for i := range pValue.([]interface{}) {
															rule.ActionParameters.CacheKey.CustomKey.Header.CheckPresence = append(rule.ActionParameters.CacheKey.CustomKey.Header.CheckPresence, pValue.([]interface{})[i].(string))
														}
													case "exclude_origin":
														if value, ok := d.GetOk(fmt.Sprintf("rules.%d.action_parameters.0.cache_key.0.custom_key.0.header.0.exclude_origin", rulesCounter)); ok {
															rule.ActionParameters.CacheKey.CustomKey.Header.ExcludeOrigin = cloudflare.BoolPtr(value.(bool))
														}
													}
												}
											}
										case "cookie":
											for i := range pValue.([]interface{}) {
												rule.ActionParameters.CacheKey.CustomKey.Cookie = &cloudflare.RulesetRuleActionParametersCustomKeyCookie{}
												if pValue.([]interface{})[i] == nil {
													continue
												}
												for pKey, pValue := range pValue.([]interface{})[i].(map[string]interface{}) {
													switch pKey {
													case "include":
														for i := range pValue.([]interface{}) {
															rule.ActionParameters.CacheKey.CustomKey.Cookie.Include = append(rule.ActionParameters.CacheKey.CustomKey.Cookie.Include, pValue.([]interface{})[i].(string))
														}
													case "check_presence":
														for i := range pValue.([]interface{}) {
															rule.ActionParameters.CacheKey.CustomKey.Cookie.CheckPresence = append(rule.ActionParameters.CacheKey.CustomKey.Cookie.CheckPresence, pValue.([]interface{})[i].(string))
														}
													}
												}
											}
										case "user":
											for i := range pValue.([]interface{}) {
												rule.ActionParameters.CacheKey.CustomKey.User = &cloudflare.RulesetRuleActionParametersCustomKeyUser{}
												if pValue.([]interface{})[i] == nil {
													continue
												}
												for pKey := range pValue.([]interface{})[i].(map[string]interface{}) {
													switch pKey {
													case "device_type":
														if value, ok := d.GetOk(fmt.Sprintf("rules.%d.action_parameters.0.cache_key.0.custom_key.0.user.0.device_type", rulesCounter)); ok {
															rule.ActionParameters.CacheKey.CustomKey.User.DeviceType = cloudflare.BoolPtr(value.(bool))
														}
													case "geo":
														if value, ok := d.GetOk(fmt.Sprintf("rules.%d.action_parameters.0.cache_key.0.custom_key.0.user.0.geo", rulesCounter)); ok {
															rule.ActionParameters.CacheKey.CustomKey.User.Geo = cloudflare.BoolPtr(value.(bool))
														}
													case "lang":
														if value, ok := d.GetOk(fmt.Sprintf("rules.%d.action_parameters.0.cache_key.0.custom_key.0.user.0.lang", rulesCounter)); ok {
															rule.ActionParameters.CacheKey.CustomKey.User.Lang = cloudflare.BoolPtr(value.(bool))
														}
													}
												}
											}
										case "host":
											for i := range pValue.([]interface{}) {
												rule.ActionParameters.CacheKey.CustomKey.Host = &cloudflare.RulesetRuleActionParametersCustomKeyHost{}
												if pValue.([]interface{})[i] == nil {
													continue
												}
												for pKey := range pValue.([]interface{})[i].(map[string]interface{}) {
													switch pKey {
													case "resolved":
														if value, ok := d.GetOk(fmt.Sprintf("rules.%d.action_parameters.0.cache_key.0.custom_key.0.host.0.resolved", rulesCounter)); ok {
															rule.ActionParameters.CacheKey.CustomKey.Host.Resolved = cloudflare.BoolPtr(value.(bool))
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}

				case "origin_error_page_passthru":
					if value, ok := d.GetOk(fmt.Sprintf("rules.%d.action_parameters.0.origin_error_page_passthru", rulesCounter)); ok {
						rule.ActionParameters.OriginErrorPagePassthru = cloudflare.BoolPtr(value.(bool))
					}

				case "request_fields":
					fields := make([]cloudflare.RulesetActionParametersLogCustomField, 0)
					for _, v := range pValue.(*schema.Set).List() {
						fields = append(fields, cloudflare.RulesetActionParametersLogCustomField{
							Name: v.(string),
						})
					}
					rule.ActionParameters.RequestFields = fields

				case "response_fields":
					fields := make([]cloudflare.RulesetActionParametersLogCustomField, 0)
					for _, v := range pValue.(*schema.Set).List() {
						fields = append(fields, cloudflare.RulesetActionParametersLogCustomField{
							Name: v.(string),
						})
					}
					rule.ActionParameters.ResponseFields = fields

				case "cookie_fields":
					fields := make([]cloudflare.RulesetActionParametersLogCustomField, 0)
					for _, v := range pValue.(*schema.Set).List() {
						fields = append(fields, cloudflare.RulesetActionParametersLogCustomField{
							Name: v.(string),
						})
					}
					rule.ActionParameters.CookieFields = fields

				case "from_list":
					for i := range pValue.([]interface{}) {
						rule.ActionParameters.FromList = &cloudflare.RulesetRuleActionParametersFromList{
							Name: pValue.([]interface{})[i].(map[string]interface{})["name"].(string),
							Key:  pValue.([]interface{})[i].(map[string]interface{})["key"].(string),
						}
					}

				case "from_value":
					for i := range pValue.([]interface{}) {
						var targetURL cloudflare.RulesetRuleActionParametersTargetURL
						for _, pValue := range pValue.([]interface{})[i].(map[string]interface{})["target_url"].([]interface{}) {
							for pKey, pValue := range pValue.(map[string]interface{}) {
								switch pKey {
								case "value":
									targetURL.Value = pValue.(string)
								case "expression":
									targetURL.Expression = pValue.(string)
								}
							}
						}

						rule.ActionParameters.FromValue = &cloudflare.RulesetRuleActionParametersFromValue{
							StatusCode:          uint16(pValue.([]interface{})[i].(map[string]interface{})["status_code"].(int)),
							TargetURL:           targetURL,
							PreserveQueryString: pValue.([]interface{})[i].(map[string]interface{})["preserve_query_string"].(bool),
						}
					}

				default:
					log.Printf("[DEBUG] unknown key encountered in buildRulesetRulesFromResource for action parameters: %s", pKey)
				}
			}
		}
	}

	if len(resourceRule["ratelimit"].([]interface{})) > 0 {
		rule.RateLimit = &cloudflare.RulesetRuleRateLimit{}
		for _, parameter := range resourceRule["ratelimit"].([]interface{}) {
			for pKey, pValue := range parameter.(map[string]interface{}) {
				switch pKey {
				case "characteristics":
					characteristicKeys := make([]string, 0)
					for _, v := range pValue.(*schema.Set).List() {
						characteristicKeys = append(characteristicKeys, v.(string))
					}
					rule.RateLimit.Characteristics = characteristicKeys
				case "period":
					rule.RateLimit.Period = pValue.(int)
				case "requests_per_period":
					rule.RateLimit.RequestsPerPeriod = pValue.(int)
				case "score_per_period":
					rule.RateLimit.ScorePerPeriod = pValue.(int)
				case "score_response_header_name":
					rule.RateLimit.ScoreResponseHeaderName = pValue.(string)
				case "mitigation_timeout":
					rule.RateLimit.MitigationTimeout = pValue.(int)
				case "counting_expression":
					rule.RateLimit.CountingExpression = pValue.(string)
				case "requests_to_origin":
					rule.RateLimit.RequestsToOrigin = pValue.(bool)

				default:
					log.Printf("[DEBUG] unknown key encountered in buildRulesetRulesFromResource for ratelimit: %s", pKey)
				}
			}
		}
	}

	if len(resourceRule["exposed_credential_check"].([]interface{})) > 0 {
		rule.ExposedCredentialCheck = &cloudflare.RulesetRuleExposedCredentialCheck{}
		for _, parameter := range resourceRule["exposed_credential_check"].([]interface{}) {
			for pKey, pValue := range parameter.(map[string]interface{}) {
				switch pKey {
				case "username_expression":
					rule.ExposedCredentialCheck.UsernameExpression = pValue.(string)
				case "password_expression":
					rule.ExposedCredentialCheck.PasswordExpression = pValue.(string)

				default:
					log.Printf("[DEBUG] unknown key encountered in buildRulesetRulesFromResource for exposed_credential_check: %s", pKey)
				}
			}
		}
	}

	if len(resourceRule["logging"].([]interface{})) > 0 {
		rule.Logging = &cloudflare.RulesetRuleLogging{}
		for _, parameter := range resourceRule["logging"].([]interface{}) {
			for pKey, pValue := range parameter.(map[string]interface{}) {
				switch pKey {
				case "status":
					rule.Logging.Enabled = statusToAPIEnabledFieldConversion(pValue.(string))
				default:
					log.Printf("[DEBUG] unknown key encountered in buildRulesetRulesFromResource for logging: %s", pKey)
				}
			}
		}
	}

	rule.Action = resourceRule["action"].(string)
	rule.Enabled = resourceRule["enabled"].(bool)

	if resourceRule["expression"] != nil {
		rule.Expression = resourceRule["expression"].(string)
	}

	if resourceRule["description"] != nil {
		rule.Description = resourceRule["description"].(string)
	}

	// Rule IDs are only reliable for the old values of rules (those stored
	// in Terraform state after a GET). For new values (those actually sent
	// with a POST or PUT request), IDs may be empty or, worse, assigned
	// based on position. Assigning IDs based on position is wrong when the
	// user reorders rules and when the user inserts or removes rules "in
	// the middle". We fix this in two steps: First, we use the IDs of old
	// rules to build a lookup table (see ruleIDs). Second, we rewrite the
	// IDs for new rules based on that lookup table. We don't attempt to
	// guess the ID of modified rules--those get reset to the empty string.
	if resourceRule["id"] != nil {
		rule.ID = resourceRule["id"].(string)
	}

	return rule
}

func buildRules(d *schema.ResourceData, value interface{}) ([]cloudflare.RulesetRule, error) {
	var rulesetRules []cloudflare.RulesetRule

	rules, ok := value.([]interface{})
	if !ok {
		return nil, errors.New("unable to create interface array type assertion")
	}

	for rulesCounter, v := range rules {
		resourceRule, ok := v.(map[string]interface{})
		if !ok {
			return nil, errors.New("unable to create interface map type assertion for rule")
		}

		rule := buildRule(d, resourceRule, rulesCounter)

		rulesetRules = append(rulesetRules, rule)
	}

	return rulesetRules, nil
}

// ruleIDs is a lookup table for rule IDs with two operations, add and pop. We
// use add to populate the table from the old value of rules. We use pop to look
// up the ID for the new value of a rule (and remove it from the table).
// Internally, both operations serialize the rule to JSON and use the resulting
// string as the lookup key; the ID itself is excluded from the JSON. If a
// ruleset has multiple copies of the same rule, the copies have a single lookup
// key associated with multiple IDs; we preserve order when adding and popping
// the IDs.
type ruleIDs struct {
	ids map[string][]string
}

// add stores an ID for the given rule.
func (r *ruleIDs) add(rule cloudflare.RulesetRule) error {
	if rule.ID == "" {
		// This is unexpected. We only invoke this function for the old
		// values of rules, which have their IDs populated.
		return errors.New("unable to determine ID of existing rule")
	}

	id := rule.ID
	rule.ID = ""

	data, err := json.Marshal(rule)
	if err != nil {
		return err
	}

	key := string(data[:])

	r.ids[key] = append(r.ids[key], id)
	return nil
}

// pop removes an ID for the given rule and returns it. Multiple IDs are
// returned in the order they were added. If no ID was found for the rule, pop
// returns an empty string.
func (r *ruleIDs) pop(rule cloudflare.RulesetRule) (string, error) {
	rule.ID = ""

	data, err := json.Marshal(rule)
	if err != nil {
		return "", err
	}

	key := string(data[:])
	ids := r.ids[key]
	if len(ids) == 0 {
		return "", nil
	}

	id, ids := ids[0], ids[1:]
	r.ids[key] = ids

	return id, nil
}

// empty returns true if the store does not contain any rule IDs.
func (r *ruleIDs) empty() bool {
	return len(r.ids) == 0
}

func newRuleIDs(rulesetRules []cloudflare.RulesetRule) (ruleIDs, error) {
	r := ruleIDs{make(map[string][]string)}

	for _, rule := range rulesetRules {
		err := r.add(rule)
		if err != nil {
			return ruleIDs{}, err
		}
	}

	return r, nil
}

// receives the resource config and builds a ruleset rule array.
func buildRulesetRulesFromResource(d *schema.ResourceData) ([]cloudflare.RulesetRule, error) {
	oldValue, newValue := d.GetChange("rules")

	oldRules, err := buildRules(d, oldValue)
	if err != nil {
		return nil, err
	}

	newRules, err := buildRules(d, newValue)
	if err != nil {
		return nil, err
	}

	ids, err := newRuleIDs(oldRules)
	if err != nil {
		return nil, err
	}

	if ids.empty() {
		// There are no rule IDs when the ruleset is first created.
		return newRules, nil
	}

	for i := range newRules {
		rule := &newRules[i]
		rule.ID, err = ids.pop(*rule)

		if err != nil {
			return nil, err
		}
	}

	return newRules, nil
}

// statusToAPIEnabledFieldConversion takes the "status" field from the Terraform
// schema/state and converts it to the API equivalent for the "enabled" field.
func statusToAPIEnabledFieldConversion(s string) *bool {
	if s == "enabled" {
		return cloudflare.BoolPtr(true)
	} else if s == "disabled" {
		return cloudflare.BoolPtr(false)
	} else {
		return nil
	}
}

// apiEnabledToStatusFieldConversion takes the "enabled" field from the API and
// converts it to the Terraform schema/state key "status".
func apiEnabledToStatusFieldConversion(s *bool) string {
	if s == nil {
		return ""
	}

	if *s == true {
		return "enabled"
	} else if *s == false {
		return "disabled"
	} else {
		return ""
	}
}
