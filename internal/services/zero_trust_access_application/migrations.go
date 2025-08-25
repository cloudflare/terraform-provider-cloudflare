// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_application

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
)

var _ resource.ResourceWithUpgradeState = (*ZeroTrustAccessApplicationResource)(nil)

// zeroTrustAccessApplicationResourceSchemaV0 defines the v0 schema (v4 provider format)
var zeroTrustAccessApplicationResourceSchemaV0 = schema.Schema{
	Attributes: map[string]schema.Attribute{
		"account_id": schema.StringAttribute{
			Optional: true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"zone_id": schema.StringAttribute{
			Optional: true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"id": schema.StringAttribute{
			Computed: true,
		},
		"name": schema.StringAttribute{
			Optional: true,
			Computed: true,
		},
		"domain": schema.StringAttribute{
			Optional: true,
			Computed: true,
		},
		"type": schema.StringAttribute{
			Optional: true,
		},
		"session_duration": schema.StringAttribute{
			Optional: true,
		},
		"auto_redirect_to_identity": schema.BoolAttribute{
			Optional: true,
		},
		"enable_binding_cookie": schema.BoolAttribute{
			Optional: true,
		},
		"custom_deny_message": schema.StringAttribute{
			Optional: true,
		},
		"custom_deny_url": schema.StringAttribute{
			Optional: true,
		},
		"custom_non_identity_deny_url": schema.StringAttribute{
			Optional: true,
		},
		"http_only_cookie_attribute": schema.BoolAttribute{
			Optional: true,
		},
		"same_site_cookie_attribute": schema.StringAttribute{
			Optional: true,
		},
		"logo_url": schema.StringAttribute{
			Optional: true,
		},
		"skip_interstitial": schema.BoolAttribute{
			Optional: true,
		},
		"app_launcher_visible": schema.BoolAttribute{
			Optional: true,
		},
		"service_auth_401_redirect": schema.BoolAttribute{
			Optional: true,
		},
		"options_preflight_bypass": schema.BoolAttribute{
			Optional: true,
		},
		"allow_authenticate_via_warp": schema.BoolAttribute{
			Optional: true,
		},
		"app_launcher_logo_url": schema.StringAttribute{
			Optional: true,
		},
		"bg_color": schema.StringAttribute{
			Optional: true,
		},
		"header_bg_color": schema.StringAttribute{
			Optional: true,
		},
		"skip_app_launcher_login_page": schema.BoolAttribute{
			Optional: true,
		},
		"aud": schema.StringAttribute{
			Computed: true,
		},
		"domain_type": schema.StringAttribute{
			Optional: true,
		},
		// V4 Set types that become List types in V5
		"allowed_idps": schema.SetAttribute{
			ElementType: types.StringType,
			Optional:    true,
		},
		"custom_pages": schema.SetAttribute{
			ElementType: types.StringType,
			Optional:    true,
		},
		"tags": schema.SetAttribute{
			ElementType: types.StringType,
			Optional:    true,
		},
		"self_hosted_domains": schema.SetAttribute{
			ElementType: types.StringType,
			Optional:    true,
		},
		// V4 List[string] that becomes List[Object] in V5
		"policies": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
		},
	},
	// V4 Block types that become SingleNestedAttribute in V5
	Blocks: map[string]schema.Block{
		"cors_headers": schema.ListNestedBlock{
			NestedObject: schema.NestedBlockObject{
				Attributes: map[string]schema.Attribute{
					"allowed_methods": schema.SetAttribute{
						ElementType: types.StringType,
						Optional:    true,
					},
					"allowed_origins": schema.SetAttribute{
						ElementType: types.StringType,
						Optional:    true,
					},
					"allowed_headers": schema.SetAttribute{
						ElementType: types.StringType,
						Optional:    true,
					},
					"allow_credentials": schema.BoolAttribute{
						Optional: true,
					},
					"max_age": schema.Int64Attribute{
						Optional: true,
					},
				},
			},
		},
		"saas_app": schema.ListNestedBlock{
			NestedObject: schema.NestedBlockObject{
				Attributes: map[string]schema.Attribute{
					"consumer_service_url": schema.StringAttribute{
						Optional: true,
					},
					"sp_entity_id": schema.StringAttribute{
						Optional: true,
					},
					"name_id_format": schema.StringAttribute{
						Optional: true,
					},
					"client_id": schema.StringAttribute{
						Optional: true,
					},
					"client_secret": schema.StringAttribute{
						Optional: true,
					},
					"auth_type": schema.StringAttribute{
						Optional: true,
					},
					"redirect_uris": schema.SetAttribute{
						ElementType: types.StringType,
						Optional:    true,
					},
					"grant_types": schema.SetAttribute{
						ElementType: types.StringType,
						Optional:    true,
					},
					"scopes": schema.SetAttribute{
						ElementType: types.StringType,
						Optional:    true,
					},
				},
				Blocks: map[string]schema.Block{
					"custom_claim": schema.ListNestedBlock{
						NestedObject: schema.NestedBlockObject{
							Attributes: map[string]schema.Attribute{
								"source": schema.StringAttribute{
									Required: true,
								},
								"name": schema.StringAttribute{
									Required: true,
								},
								"required": schema.BoolAttribute{
									Optional: true,
								},
								"scope": schema.StringAttribute{
									Optional: true,
								},
							},
						},
					},
					"custom_attribute": schema.ListNestedBlock{
						NestedObject: schema.NestedBlockObject{
							Attributes: map[string]schema.Attribute{
								"source": schema.StringAttribute{
									Required: true,
								},
								"name": schema.StringAttribute{
									Required: true,
								},
								"name_format": schema.StringAttribute{
									Optional: true,
								},
								"friendly_name": schema.StringAttribute{
									Optional: true,
								},
							},
							Blocks: map[string]schema.Block{
								"source": schema.ListNestedBlock{
									NestedObject: schema.NestedBlockObject{
										Attributes: map[string]schema.Attribute{
											"name_by_idp": schema.MapAttribute{
												ElementType: types.StringType,
												Optional:    true,
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"landing_page_design": schema.ListNestedBlock{
			NestedObject: schema.NestedBlockObject{
				Attributes: map[string]schema.Attribute{
					"message": schema.StringAttribute{
						Optional: true,
					},
					"button_color": schema.StringAttribute{
						Optional: true,
					},
					"button_text_color": schema.StringAttribute{
						Optional: true,
					},
					"image_url": schema.StringAttribute{
						Optional: true,
					},
				},
			},
		},
		"footer_links": schema.SetNestedBlock{
			NestedObject: schema.NestedBlockObject{
				Attributes: map[string]schema.Attribute{
					"name": schema.StringAttribute{
						Required: true,
					},
					"url": schema.StringAttribute{
						Required: true,
					},
				},
			},
		},
		"scim_config": schema.ListNestedBlock{
			NestedObject: schema.NestedBlockObject{
				Attributes: map[string]schema.Attribute{
					"enabled": schema.BoolAttribute{
						Optional: true,
					},
					"remote_uri": schema.StringAttribute{
						Optional: true,
					},
					"idp_uid": schema.StringAttribute{
						Optional: true,
					},
					"deactivate_on_delete": schema.BoolAttribute{
						Optional: true,
					},
				},
				Blocks: map[string]schema.Block{
					"authentication": schema.ListNestedBlock{
						NestedObject: schema.NestedBlockObject{
							Attributes: map[string]schema.Attribute{
								"scheme": schema.StringAttribute{
									Required: true,
								},
								"token": schema.StringAttribute{
									Optional: true,
								},
								"scopes": schema.SetAttribute{
									ElementType: types.StringType,
									Optional:    true,
								},
							},
						},
					},
				},
			},
		},
		// V4 Block that becomes List[Object] in V5
		"target_criteria": schema.ListNestedBlock{
			NestedObject: schema.NestedBlockObject{
				Attributes: map[string]schema.Attribute{
					"port": schema.StringAttribute{
						Required: true,
					},
					"protocol": schema.StringAttribute{
						Required: true,
					},
				},
				Blocks: map[string]schema.Block{
					"target_attributes": schema.ListNestedBlock{
						NestedObject: schema.NestedBlockObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Required: true,
								},
								"values": schema.ListAttribute{
									ElementType: types.StringType,
									Required:    true,
								},
							},
						},
					},
				},
			},
		},
		"destinations": schema.ListNestedBlock{
			NestedObject: schema.NestedBlockObject{
				Attributes: map[string]schema.Attribute{
					"uri": schema.StringAttribute{
						Optional: true,
					},
					"type": schema.StringAttribute{
						Optional: true,
					},
				},
			},
		},
	},
}

func (r *ZeroTrustAccessApplicationResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		0: {
			PriorSchema:   &zeroTrustAccessApplicationResourceSchemaV0,
			StateUpgrader: upgradeZeroTrustAccessApplicationStateV0toV1,
		},
	}
}

// upgradeZeroTrustAccessApplicationStateV0toV1 migrates from v4 provider state format to v5
func upgradeZeroTrustAccessApplicationStateV0toV1(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	// Parse the old state using the raw state data
	var oldState map[string]interface{}
	if req.RawState == nil || len(req.RawState.JSON) == 0 {
		return
	}

	tflog.Debug(ctx, "Starting state upgrade from v0 to v1", map[string]interface{}{
		"raw_state_json_length": len(req.RawState.JSON),
	})

	// Debug: Print first part of raw JSON to see http_only_cookie_attribute
	rawJSONStr := string(req.RawState.JSON)
	if strings.Contains(rawJSONStr, "http_only_cookie_attribute") {
		startIdx := strings.Index(rawJSONStr, "http_only_cookie_attribute")
		endIdx := startIdx + 50
		if endIdx > len(rawJSONStr) {
			endIdx = len(rawJSONStr)
		}
		tflog.Debug(ctx, "Found http_only_cookie_attribute in raw JSON", map[string]interface{}{
			"http_only_snippet": rawJSONStr[startIdx:endIdx],
		})
	}

	// Parse raw JSON and clean up problematic fields
	var tempState map[string]interface{}
	err := json.Unmarshal(req.RawState.JSON, &tempState)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failed to parse raw state",
			fmt.Sprintf("Could not parse raw state during migration: %s", err),
		)
		return
	}

	// Fix v4->v5 schema incompatibilities
	// Convert empty arrays to nil for fields that changed from List[Block] to SingleNestedAttribute
	if corsHeaders, ok := tempState["cors_headers"].([]interface{}); ok && len(corsHeaders) == 0 {
		tempState["cors_headers"] = nil
	}
	if saasApp, ok := tempState["saas_app"].([]interface{}); ok && len(saasApp) == 0 {
		tempState["saas_app"] = nil
	}
	if scimConfig, ok := tempState["scim_config"].([]interface{}); ok && len(scimConfig) == 0 {
		tempState["scim_config"] = nil
	}
	if landingPage, ok := tempState["landing_page_design"].([]interface{}); ok && len(landingPage) == 0 {
		tempState["landing_page_design"] = nil
	}

	// Remove computed fields that might have different API defaults between v4 and v5
	// to prevent plan diffs during migration
	delete(tempState, "http_only_cookie_attribute")

	oldState = tempState

	// Create new state structure
	var newState ZeroTrustAccessApplicationModel

	// Migrate basic attributes
	if accountID, ok := oldState["account_id"].(string); ok && accountID != "" {
		newState.AccountID = types.StringValue(accountID)
	}
	if zoneID, ok := oldState["zone_id"].(string); ok && zoneID != "" {
		newState.ZoneID = types.StringValue(zoneID)
	}
	if id, ok := oldState["id"].(string); ok {
		newState.ID = types.StringValue(id)
	}
	if name, ok := oldState["name"].(string); ok {
		newState.Name = types.StringPointerValue(&name)
	}
	if domain, ok := oldState["domain"].(string); ok {
		newState.Domain = types.StringPointerValue(&domain)
	}
	if appType, ok := oldState["type"].(string); ok {
		newState.Type = types.StringValue(appType)
	}
	if sessionDuration, ok := oldState["session_duration"].(string); ok && sessionDuration != "" {
		newState.SessionDuration = types.StringPointerValue(&sessionDuration)
	}

	// Migrate boolean attributes
	if autoRedirect, ok := oldState["auto_redirect_to_identity"].(bool); ok {
		newState.AutoRedirectToIdentity = types.BoolValue(autoRedirect)
	}
	if enableBinding, ok := oldState["enable_binding_cookie"].(bool); ok {
		newState.EnableBindingCookie = types.BoolValue(enableBinding)
	}
	if serviceAuth401, ok := oldState["service_auth_401_redirect"].(bool); ok {
		newState.ServiceAuth401Redirect = types.BoolValue(serviceAuth401)
	}
	// Skip http_only_cookie_attribute completely - let it be computed by the API
	// The migration removes it from v4 state to avoid conflicts with v5 computed defaults
	if skipInterstitial, ok := oldState["skip_interstitial"].(bool); ok {
		newState.SkipInterstitial = types.BoolValue(skipInterstitial)
	}
	if appLauncherVisible, ok := oldState["app_launcher_visible"].(bool); ok {
		newState.AppLauncherVisible = types.BoolPointerValue(&appLauncherVisible)
	}
	if skipAppLauncherLogin, ok := oldState["skip_app_launcher_login_page"].(bool); ok {
		newState.SkipAppLauncherLoginPage = types.BoolPointerValue(&skipAppLauncherLogin)
	}
	if optionsPreflight, ok := oldState["options_preflight_bypass"].(bool); ok {
		newState.OptionsPreflightBypass = types.BoolValue(optionsPreflight)
	}
	if allowWarp, ok := oldState["allow_authenticate_via_warp"].(bool); ok {
		newState.AllowAuthenticateViaWARP = types.BoolValue(allowWarp)
	}

	// Migrate string attributes
	if customDenyMessage, ok := oldState["custom_deny_message"].(string); ok && customDenyMessage != "" {
		newState.CustomDenyMessage = types.StringPointerValue(&customDenyMessage)
	}
	if customDenyURL, ok := oldState["custom_deny_url"].(string); ok && customDenyURL != "" {
		newState.CustomDenyURL = types.StringPointerValue(&customDenyURL)
	}
	if customNonIdentityDenyURL, ok := oldState["custom_non_identity_deny_url"].(string); ok && customNonIdentityDenyURL != "" {
		newState.CustomNonIdentityDenyURL = types.StringPointerValue(&customNonIdentityDenyURL)
	}
	if sameSiteCookie, ok := oldState["same_site_cookie_attribute"].(string); ok && sameSiteCookie != "" {
		newState.SameSiteCookieAttribute = types.StringPointerValue(&sameSiteCookie)
	}
	if logoURL, ok := oldState["logo_url"].(string); ok && logoURL != "" {
		newState.LogoURL = types.StringPointerValue(&logoURL)
	}
	if appLauncherLogoURL, ok := oldState["app_launcher_logo_url"].(string); ok && appLauncherLogoURL != "" {
		newState.AppLauncherLogoURL = types.StringPointerValue(&appLauncherLogoURL)
	}
	if bgColor, ok := oldState["bg_color"].(string); ok && bgColor != "" {
		newState.BgColor = types.StringPointerValue(&bgColor)
	}
	if headerBgColor, ok := oldState["header_bg_color"].(string); ok && headerBgColor != "" {
		newState.HeaderBgColor = types.StringPointerValue(&headerBgColor)
	}
	if aud, ok := oldState["aud"].(string); ok {
		newState.AUD = types.StringValue(aud)
	}

	// Remove deprecated attributes that don't exist in v5
	// domain_type is removed completely in v5 - do not migrate

	// Migrate Set -> List/Set conversions (reverted in 5.7.0)
	if allowedIDPs, ok := oldState["allowed_idps"].([]interface{}); ok {
		migratedAllowedIDPs := migrateStringSliceToList(allowedIDPs)
		if len(migratedAllowedIDPs) > 0 {
			newState.AllowedIdPs = &migratedAllowedIDPs
		}
	}

	if customPages, ok := oldState["custom_pages"].([]interface{}); ok {
		migratedCustomPages := migrateStringSliceToList(customPages)
		if len(migratedCustomPages) > 0 {
			newState.CustomPages = &migratedCustomPages
		}
	}

	if tags, ok := oldState["tags"].([]interface{}); ok {
		migratedTags := migrateStringSliceToList(tags)
		if len(migratedTags) > 0 {
			newState.Tags, _ = customfield.NewList[types.String](ctx, migratedTags)
		}
	}

	if selfHostedDomains, ok := oldState["self_hosted_domains"].([]interface{}); ok {
		migratedDomains := migrateStringSliceToList(selfHostedDomains)
		if len(migratedDomains) > 0 {
			newState.SelfHostedDomains, _ = customfield.NewList[types.String](ctx, migratedDomains)
		}
	}

	// Migrate policies: List[string] -> List[Object] with id field
	if policies, ok := oldState["policies"].([]interface{}); ok {
		migratedPolicies := migratePoliciesStringListToObjectList(policies)
		if len(migratedPolicies) > 0 {
			newState.Policies = &migratedPolicies
		}
	}

	// Migrate Block -> SingleNestedAttribute conversions
	// CORS Headers: List[Block] (max 1) -> SingleNestedAttribute
	if corsHeadersData, ok := oldState["cors_headers"].([]interface{}); ok && len(corsHeadersData) > 0 {
		if corsHeadersMap, ok := corsHeadersData[0].(map[string]interface{}); ok {
			migratedCorsHeaders := migrateCorsHeadersBlockToObject(corsHeadersMap)
			newState.CORSHeaders = &migratedCorsHeaders
		}
	}

	// SAAS App: List[Block] (max 1) -> SingleNestedAttribute
	if saasAppData, ok := oldState["saas_app"].([]interface{}); ok && len(saasAppData) > 0 {
		if saasAppMap, ok := saasAppData[0].(map[string]interface{}); ok {
			migratedSaasApp := migrateSaasAppBlockToObject(saasAppMap)
			newState.SaaSApp, _ = customfield.NewObject(ctx, &migratedSaasApp)
		}
	}

	// Landing Page Design: List[Block] (max 1) -> SingleNestedAttribute with default title
	if landingPageData, ok := oldState["landing_page_design"].([]interface{}); ok && len(landingPageData) > 0 {
		if landingPageMap, ok := landingPageData[0].(map[string]interface{}); ok {
			migratedLandingPage := migrateLandingPageDesignBlockToObject(landingPageMap)
			newState.LandingPageDesign, _ = customfield.NewObject(ctx, &migratedLandingPage)
		}
	}

	// Footer Links: Set[Block] -> List[Object]
	if footerLinksData, ok := oldState["footer_links"].([]interface{}); ok && len(footerLinksData) > 0 {
		migratedFooterLinks := migrateFooterLinksSetBlockToListObject(footerLinksData)
		if len(migratedFooterLinks) > 0 {
			var footerLinkPtrs []*ZeroTrustAccessApplicationFooterLinksModel
			for i := range migratedFooterLinks {
				footerLinkPtrs = append(footerLinkPtrs, &migratedFooterLinks[i])
			}
			newState.FooterLinks = &footerLinkPtrs
		}
	}

	// SCIM Config: List[Block] (max 1) -> SingleNestedAttribute
	if scimConfigData, ok := oldState["scim_config"].([]interface{}); ok && len(scimConfigData) > 0 {
		if scimConfigMap, ok := scimConfigData[0].(map[string]interface{}); ok {
			migratedScimConfig := migrateScimConfigBlockToObject(scimConfigMap)
			newState.SCIMConfig = &migratedScimConfig
		}
	}

	// Target Criteria: List[Block] -> List[Object] with target_attributes restructure
	if targetCriteriaData, ok := oldState["target_criteria"].([]interface{}); ok && len(targetCriteriaData) > 0 {
		migratedTargetCriteria := migrateTargetCriteriaBlocksToObjects(targetCriteriaData)
		if len(migratedTargetCriteria) > 0 {
			var targetCriteriaPtrs []*ZeroTrustAccessApplicationTargetCriteriaModel
			for i := range migratedTargetCriteria {
				targetCriteriaPtrs = append(targetCriteriaPtrs, &migratedTargetCriteria[i])
			}
			newState.TargetCriteria = &targetCriteriaPtrs
		}
	}

	// Destinations: List[Block] -> List[Object]  
	if destinationsData, ok := oldState["destinations"].([]interface{}); ok && len(destinationsData) > 0 {
		migratedDestinations := migrateDestinationsBlocksToObjects(destinationsData)
		if len(migratedDestinations) > 0 {
			newState.Destinations = customfield.NewObjectListMust(ctx, migratedDestinations)
		}
	}

	// Add new v5 attributes with defaults where appropriate
	// allow_iframe is new in v5 - leave as null
	// read_service_tokens_from_header is new in v5 - leave as null 
	// path_cookie_attribute is new in v5 - leave as null

	// Set the upgraded state
	resp.Diagnostics.Append(resp.State.Set(ctx, &newState)...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, fmt.Sprintf("Failed to set new state: %v", resp.Diagnostics.Errors()))
		return
	}
}

// Helper functions for specific migrations

// migrateSetToList converts set data to list for attributes that reverted from List back to Set in 5.7.0
func migrateSetToList(setData []interface{}) []types.String {
	var result []types.String
	for _, item := range setData {
		if str, ok := item.(string); ok && str != "" {
			result = append(result, types.StringValue(str))
		}
	}
	return result
}

// migrateStringSliceToList converts interface slice to string slice  
func migrateStringSliceToList(sliceData []interface{}) []types.String {
	var result []types.String
	for _, item := range sliceData {
		if str, ok := item.(string); ok && str != "" {
			result = append(result, types.StringValue(str))
		}
	}
	return result
}

// migratePoliciesStringListToObjectList converts policies from List[string] to List[Object] with id field
func migratePoliciesStringListToObjectList(policies []interface{}) []ZeroTrustAccessApplicationPoliciesModel {
	var result []ZeroTrustAccessApplicationPoliciesModel
	for _, policy := range policies {
		if policyStr, ok := policy.(string); ok && policyStr != "" {
			result = append(result, ZeroTrustAccessApplicationPoliciesModel{
				ID: types.StringValue(policyStr),
			})
		}
	}
	return result
}

// migrateCorsHeadersBlockToObject converts CORS headers from List[Block] (max 1) to SingleNestedAttribute
func migrateCorsHeadersBlockToObject(corsData map[string]interface{}) ZeroTrustAccessApplicationCORSHeadersModel {
	var corsHeaders ZeroTrustAccessApplicationCORSHeadersModel

	if allowedMethods, ok := corsData["allowed_methods"].([]interface{}); ok {
		methods := migrateStringSliceToList(allowedMethods)
		if len(methods) > 0 {
			corsHeaders.AllowedMethods = &methods
		}
	}

	if allowedOrigins, ok := corsData["allowed_origins"].([]interface{}); ok {
		origins := migrateStringSliceToList(allowedOrigins)
		if len(origins) > 0 {
			corsHeaders.AllowedOrigins = &origins
		}
	}

	if allowedHeaders, ok := corsData["allowed_headers"].([]interface{}); ok {
		headers := migrateStringSliceToList(allowedHeaders)
		if len(headers) > 0 {
			corsHeaders.AllowedHeaders = &headers
		}
	}

	if allowCredentials, ok := corsData["allow_credentials"].(bool); ok {
		corsHeaders.AllowCredentials = types.BoolValue(allowCredentials)
	}

	// MaxAge: Int -> Float64 conversion
	if maxAge, ok := corsData["max_age"]; ok {
		switch v := maxAge.(type) {
		case int:
			corsHeaders.MaxAge = types.Float64Value(float64(v))
		case int64:
			corsHeaders.MaxAge = types.Float64Value(float64(v))
		case float64:
			corsHeaders.MaxAge = types.Float64Value(v)
		case string:
			if parsed, err := strconv.ParseFloat(v, 64); err == nil {
				corsHeaders.MaxAge = types.Float64Value(parsed)
			}
		}
	}

	return corsHeaders
}

// migrateSaasAppBlockToObject converts SAAS app from List[Block] (max 1) to SingleNestedAttribute
func migrateSaasAppBlockToObject(saasData map[string]interface{}) ZeroTrustAccessApplicationSaaSAppModel {
	var saasApp ZeroTrustAccessApplicationSaaSAppModel

	if consumerServiceURL, ok := saasData["consumer_service_url"].(string); ok && consumerServiceURL != "" {
		saasApp.ConsumerServiceURL = types.StringPointerValue(&consumerServiceURL)
	}
	if spEntityID, ok := saasData["sp_entity_id"].(string); ok && spEntityID != "" {
		saasApp.SPEntityID = types.StringPointerValue(&spEntityID)
	}
	if nameIDFormat, ok := saasData["name_id_format"].(string); ok && nameIDFormat != "" {
		saasApp.NameIDFormat = types.StringPointerValue(&nameIDFormat)
	}
	if clientID, ok := saasData["client_id"].(string); ok && clientID != "" {
		saasApp.ClientID = types.StringPointerValue(&clientID)
	}
	if clientSecret, ok := saasData["client_secret"].(string); ok && clientSecret != "" {
		saasApp.ClientSecret = types.StringPointerValue(&clientSecret)
	}
	if authType, ok := saasData["auth_type"].(string); ok && authType != "" {
		saasApp.AuthType = types.StringPointerValue(&authType)
	}

	// Set-to-List conversions
	if redirectURIs, ok := saasData["redirect_uris"].([]interface{}); ok {
		uris := migrateStringSliceToList(redirectURIs)
		if len(uris) > 0 {
			saasApp.RedirectURIs = &uris
		}
	}

	if grantTypes, ok := saasData["grant_types"].([]interface{}); ok {
		grants := migrateStringSliceToList(grantTypes)
		if len(grants) > 0 {
			saasApp.GrantTypes = &grants
		}
	}

	if scopes, ok := saasData["scopes"].([]interface{}); ok {
		scopesList := migrateStringSliceToList(scopes)
		if len(scopesList) > 0 {
			saasApp.Scopes = &scopesList
		}
	}

	// Attribute renames: custom_claim -> custom_claims, custom_attribute -> custom_attributes
	if customClaimData, ok := saasData["custom_claim"].([]interface{}); ok && len(customClaimData) > 0 {
		var customClaims []*ZeroTrustAccessApplicationSaaSAppCustomClaimsModel
		for _, claimItem := range customClaimData {
			if claimMap, ok := claimItem.(map[string]interface{}); ok {
				var claim ZeroTrustAccessApplicationSaaSAppCustomClaimsModel
				if source, ok := claimMap["source"].(string); ok {
					sourceModel := &ZeroTrustAccessApplicationSaaSAppCustomClaimsSourceModel{
						Name: types.StringValue(source),
					}
					claim.Source = sourceModel
				}
				if name, ok := claimMap["name"].(string); ok {
					claim.Name = types.StringValue(name)
				}
				if required, ok := claimMap["required"].(bool); ok {
					claim.Required = types.BoolPointerValue(&required)
				}
				if scope, ok := claimMap["scope"].(string); ok && scope != "" {
					claim.Scope = types.StringPointerValue(&scope)
				}
				customClaims = append(customClaims, &claim)
			}
		}
		if len(customClaims) > 0 {
			saasApp.CustomClaims = &customClaims
		}
	}

	if customAttributeData, ok := saasData["custom_attribute"].([]interface{}); ok && len(customAttributeData) > 0 {
		var customAttributes []*ZeroTrustAccessApplicationSaaSAppCustomAttributesModel
		for _, attrItem := range customAttributeData {
			if attrMap, ok := attrItem.(map[string]interface{}); ok {
				var attr ZeroTrustAccessApplicationSaaSAppCustomAttributesModel
				if source, ok := attrMap["source"].(string); ok {
					sourceModel := &ZeroTrustAccessApplicationSaaSAppCustomAttributesSourceModel{
						Name: types.StringValue(source),
					}
					attr.Source = sourceModel
				}
				if name, ok := attrMap["name"].(string); ok {
					attr.Name = types.StringValue(name)
				}
				if nameFormat, ok := attrMap["name_format"].(string); ok && nameFormat != "" {
					attr.NameFormat = types.StringPointerValue(&nameFormat)
				}
				if friendlyName, ok := attrMap["friendly_name"].(string); ok && friendlyName != "" {
					attr.FriendlyName = types.StringPointerValue(&friendlyName)
				}

				// Handle nested source block for name_by_idp - simplified for now
				// Complex name_by_idp map -> list -> map transition requires custom handling
				// Skip for now to get basic migration working

				customAttributes = append(customAttributes, &attr)
			}
		}
		if len(customAttributes) > 0 {
			saasApp.CustomAttributes = &customAttributes
		}
	}

	return saasApp
}

// migrateLandingPageDesignBlockToObject converts landing page design from List[Block] (max 1) to SingleNestedAttribute
func migrateLandingPageDesignBlockToObject(landingPageData map[string]interface{}) ZeroTrustAccessApplicationLandingPageDesignModel {
	var landingPage ZeroTrustAccessApplicationLandingPageDesignModel

	// Add default title "Welcome!" as per v5 schema
	defaultTitle := "Welcome!"
	landingPage.Title = types.StringPointerValue(&defaultTitle)

	if message, ok := landingPageData["message"].(string); ok && message != "" {
		landingPage.Message = types.StringPointerValue(&message)
	}
	if buttonColor, ok := landingPageData["button_color"].(string); ok && buttonColor != "" {
		landingPage.ButtonColor = types.StringPointerValue(&buttonColor)
	}
	if buttonTextColor, ok := landingPageData["button_text_color"].(string); ok && buttonTextColor != "" {
		landingPage.ButtonTextColor = types.StringPointerValue(&buttonTextColor)
	}
	if imageURL, ok := landingPageData["image_url"].(string); ok && imageURL != "" {
		landingPage.ImageURL = types.StringPointerValue(&imageURL)
	}

	return landingPage
}

// migrateFooterLinksSetBlockToListObject converts footer links from Set[Block] to List[Object]
func migrateFooterLinksSetBlockToListObject(footerLinksData []interface{}) []ZeroTrustAccessApplicationFooterLinksModel {
	var result []ZeroTrustAccessApplicationFooterLinksModel
	for _, linkItem := range footerLinksData {
		if linkMap, ok := linkItem.(map[string]interface{}); ok {
			var footerLink ZeroTrustAccessApplicationFooterLinksModel
			if name, ok := linkMap["name"].(string); ok {
				footerLink.Name = types.StringValue(name)
			}
			if url, ok := linkMap["url"].(string); ok {
				footerLink.URL = types.StringValue(url)
			}
			result = append(result, footerLink)
		}
	}
	return result
}

// migrateScimConfigBlockToObject converts SCIM config from List[Block] (max 1) to SingleNestedAttribute
func migrateScimConfigBlockToObject(scimData map[string]interface{}) ZeroTrustAccessApplicationSCIMConfigModel {
	var scimConfig ZeroTrustAccessApplicationSCIMConfigModel

	if enabled, ok := scimData["enabled"].(bool); ok {
		scimConfig.Enabled = types.BoolPointerValue(&enabled)
	}
	if remoteURI, ok := scimData["remote_uri"].(string); ok && remoteURI != "" {
		scimConfig.RemoteURI = types.StringPointerValue(&remoteURI)
	}
	if idpUID, ok := scimData["idp_uid"].(string); ok && idpUID != "" {
		scimConfig.IdPUID = types.StringPointerValue(&idpUID)
	}
	if deactivateOnDelete, ok := scimData["deactivate_on_delete"].(bool); ok {
		scimConfig.DeactivateOnDelete = types.BoolPointerValue(&deactivateOnDelete)
	}

	// Handle nested authentication block
	if authData, ok := scimData["authentication"].([]interface{}); ok && len(authData) > 0 {
		if authMap, ok := authData[0].(map[string]interface{}); ok {
			var auth ZeroTrustAccessApplicationSCIMConfigAuthenticationModel
			if scheme, ok := authMap["scheme"].(string); ok {
				auth.Scheme = types.StringValue(scheme)
			}
			if token, ok := authMap["token"].(string); ok && token != "" {
				auth.Password = types.StringPointerValue(&token) // Note: v5 uses Password field
			}

			scimConfig.Authentication = &auth
		}
	}

	return scimConfig
}

// migrateTargetCriteriaBlocksToObjects converts target criteria from List[Block] to List[Object]
// This includes the complex target_attributes list -> map conversion
func migrateTargetCriteriaBlocksToObjects(targetCriteriaData []interface{}) []ZeroTrustAccessApplicationTargetCriteriaModel {
	var result []ZeroTrustAccessApplicationTargetCriteriaModel
	for _, criteriaItem := range targetCriteriaData {
		if criteriaMap, ok := criteriaItem.(map[string]interface{}); ok {
			var criteria ZeroTrustAccessApplicationTargetCriteriaModel
			if port, ok := criteriaMap["port"].(string); ok {
				if portNum, err := strconv.ParseInt(port, 10, 64); err == nil {
					criteria.Port = types.Int64Value(portNum)
				}
			}
			if protocol, ok := criteriaMap["protocol"].(string); ok {
				criteria.Protocol = types.StringValue(protocol)
			}

			// Complex target_attributes transformation: List[{name, values}] -> Map[string, List[string]]
			if targetAttrsData, ok := criteriaMap["target_attributes"].([]interface{}); ok && len(targetAttrsData) > 0 {
				targetAttrsMap := make(map[string]*[]types.String)
				for _, attrItem := range targetAttrsData {
					if attrMap, ok := attrItem.(map[string]interface{}); ok {
						if name, ok := attrMap["name"].(string); ok {
							if values, ok := attrMap["values"].([]interface{}); ok {
								valuesList := migrateStringSliceToList(values)
								if len(valuesList) > 0 {
									targetAttrsMap[name] = &valuesList
								}
							}
						}
					}
				}
				if len(targetAttrsMap) > 0 {
					// Convert to the v5 format - this may need customfield handling
					// For now, create a simple structure
					criteria.TargetAttributes = &targetAttrsMap
				}
			}

			result = append(result, criteria)
		}
	}
	return result
}

// migrateDestinationsBlocksToObjects converts destinations from List[Block] to List[Object]
func migrateDestinationsBlocksToObjects(destinationsData []interface{}) []ZeroTrustAccessApplicationDestinationsModel {
	var result []ZeroTrustAccessApplicationDestinationsModel
	for _, destItem := range destinationsData {
		if destMap, ok := destItem.(map[string]interface{}); ok {
			var destination ZeroTrustAccessApplicationDestinationsModel
			if uri, ok := destMap["uri"].(string); ok && uri != "" {
				destination.URI = types.StringPointerValue(&uri)
			}
			if destType, ok := destMap["type"].(string); ok && destType != "" {
				destination.Type = types.StringPointerValue(&destType)
			}
			result = append(result, destination)
		}
	}
	return result
}
