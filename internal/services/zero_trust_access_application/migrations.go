// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_application

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithUpgradeState = (*ZeroTrustAccessApplicationResource)(nil)

func (r *ZeroTrustAccessApplicationResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		0: {
			// PriorSchema includes all unchanged attributes from V4 to V5
			// This will work with both v4 and early v5 states
			PriorSchema: &schema.Schema{
				Attributes: map[string]schema.Attribute{
					// Core attributes - unchanged
					"id": schema.StringAttribute{
						Computed: true,
					},
					"account_id": schema.StringAttribute{
						Optional: true,
					},
					"zone_id": schema.StringAttribute{
						Optional: true,
					},
					"name": schema.StringAttribute{
						Optional: true,
					},
					"domain": schema.StringAttribute{
						Optional: true,
					},
					"type": schema.StringAttribute{
						Optional: true,
					},

					// Session and cookie attributes
					"session_duration": schema.StringAttribute{
						Optional: true,
					},
					"http_only_cookie_attribute": schema.BoolAttribute{
						Optional: true,
					},
					"same_site_cookie_attribute": schema.StringAttribute{
						Optional: true,
					},

					// Visual/UI attributes - unchanged
					"logo_url": schema.StringAttribute{
						Optional: true,
					},
					"app_launcher_logo_url": schema.StringAttribute{
						Optional: true,
					},
					"header_bg_color": schema.StringAttribute{
						Optional: true,
					},
					"bg_color": schema.StringAttribute{
						Optional: true,
					},

					// Custom messaging attributes - unchanged
					"custom_deny_message": schema.StringAttribute{
						Optional: true,
					},
					"custom_deny_url": schema.StringAttribute{
						Optional: true,
					},
					"custom_non_identity_deny_url": schema.StringAttribute{
						Optional: true,
					},

					// Boolean flags - unchanged
					"allow_authenticate_via_warp": schema.BoolAttribute{
						Optional: true,
					},
					"app_launcher_visible": schema.BoolAttribute{
						Optional: true,
					},
					"auto_redirect_to_identity": schema.BoolAttribute{
						Optional: true,
					},
					"enable_binding_cookie": schema.BoolAttribute{
						Optional: true,
					},
					"skip_interstitial": schema.BoolAttribute{
						Optional: true,
					},
					"service_auth_401_redirect": schema.BoolAttribute{
						Optional: true,
					},
					"skip_app_launcher_login_page": schema.BoolAttribute{
						Optional: true,
					},
					"options_preflight_bypass": schema.BoolAttribute{
						Optional: true,
					},

					// Sets/Lists - unchanged but changed from Set to List in v5
					"allowed_idps": schema.ListAttribute{
						Optional:    true,
						ElementType: types.StringType,
					},
					"custom_pages": schema.ListAttribute{
						Optional:    true,
						ElementType: types.StringType,
					},
					"tags": schema.ListAttribute{
						Optional:    true,
						ElementType: types.StringType,
						CustomType:  customfield.NewListType[types.String](ctx),
					},
					"self_hosted_domains": schema.ListAttribute{
						Optional:    true,
						ElementType: types.StringType,
						CustomType:  customfield.NewListType[types.String](ctx),
					},
				},
			},

			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				fmt.Println("========================================")
				fmt.Println("STATE UPGRADER CALLED FOR zero_trust_access_application")
				fmt.Println("========================================")
				fmt.Printf("DEBUG: RawState JSON length: %d\n", len(req.RawState.JSON))

				// Log first 500 chars of raw state
				if len(req.RawState.JSON) > 0 {
					preview := string(req.RawState.JSON)
					if len(preview) > 500 {
						preview = preview[:500] + "..."
					}
					fmt.Printf("DEBUG: RawState preview: %s\n", preview)
				}

				// Define struct matching all attributes in PriorSchema
				var priorStateData struct {
					// Core attributes
					ID        types.String `tfsdk:"id"`
					AccountID types.String `tfsdk:"account_id"`
					ZoneID    types.String `tfsdk:"zone_id"`
					Name      types.String `tfsdk:"name"`
					Domain    types.String `tfsdk:"domain"`
					Type      types.String `tfsdk:"type"`

					// Session and cookie attributes
					SessionDuration         types.String `tfsdk:"session_duration"`
					HTTPOnlyCookieAttribute types.Bool   `tfsdk:"http_only_cookie_attribute"`
					SameSiteCookieAttribute types.String `tfsdk:"same_site_cookie_attribute"`

					// Visual/UI attributes
					LogoURL            types.String `tfsdk:"logo_url"`
					AppLauncherLogoURL types.String `tfsdk:"app_launcher_logo_url"`
					HeaderBgColor      types.String `tfsdk:"header_bg_color"`
					BgColor            types.String `tfsdk:"bg_color"`

					// Custom messaging attributes
					CustomDenyMessage        types.String `tfsdk:"custom_deny_message"`
					CustomDenyURL            types.String `tfsdk:"custom_deny_url"`
					CustomNonIdentityDenyURL types.String `tfsdk:"custom_non_identity_deny_url"`

					// Boolean flags
					AllowAuthenticateViaWARP types.Bool `tfsdk:"allow_authenticate_via_warp"`
					AppLauncherVisible       types.Bool `tfsdk:"app_launcher_visible"`
					AutoRedirectToIdentity   types.Bool `tfsdk:"auto_redirect_to_identity"`
					EnableBindingCookie      types.Bool `tfsdk:"enable_binding_cookie"`
					SkipInterstitial         types.Bool `tfsdk:"skip_interstitial"`
					ServiceAuth401Redirect   types.Bool `tfsdk:"service_auth_401_redirect"`
					SkipAppLauncherLoginPage types.Bool `tfsdk:"skip_app_launcher_login_page"`
					OptionsPreflightBypass   types.Bool `tfsdk:"options_preflight_bypass"`

					// Sets/Lists
					AllowedIdPs       *[]types.String                `tfsdk:"allowed_idps"`
					CustomPages       *[]types.String                `tfsdk:"custom_pages"`
					Tags              customfield.List[types.String] `tfsdk:"tags"`
					SelfHostedDomains customfield.List[types.String] `tfsdk:"self_hosted_domains"`
				}

				resp.Diagnostics.Append(req.State.Get(ctx, &priorStateData)...)
				if resp.Diagnostics.HasError() {
					return
				}

				// Copy all unchanged attributes to new state
				newState := ZeroTrustAccessApplicationModel{
					// Core attributes
					ID:        priorStateData.ID,
					AccountID: priorStateData.AccountID,
					ZoneID:    priorStateData.ZoneID,
					Name:      priorStateData.Name,
					Domain:    priorStateData.Domain,
					Type:      priorStateData.Type,

					// Session and cookie attributes
					SessionDuration: priorStateData.SessionDuration,
					//HTTPOnlyCookieAttribute: priorStateData.HTTPOnlyCookieAttribute,
					SameSiteCookieAttribute: priorStateData.SameSiteCookieAttribute,

					// Visual/UI attributes
					LogoURL:            priorStateData.LogoURL,
					AppLauncherLogoURL: priorStateData.AppLauncherLogoURL,
					HeaderBgColor:      priorStateData.HeaderBgColor,
					BgColor:            priorStateData.BgColor,

					// Custom messaging attributes
					CustomDenyMessage:        priorStateData.CustomDenyMessage,
					CustomDenyURL:            priorStateData.CustomDenyURL,
					CustomNonIdentityDenyURL: priorStateData.CustomNonIdentityDenyURL,

					// Boolean flags
					AllowAuthenticateViaWARP: priorStateData.AllowAuthenticateViaWARP,
					AppLauncherVisible:       priorStateData.AppLauncherVisible,
					// AutoRedirectToIdentity:   priorStateData.AutoRedirectToIdentity,
					// EnableBindingCookie:      priorStateData.EnableBindingCookie,
					SkipInterstitial:         priorStateData.SkipInterstitial,
					ServiceAuth401Redirect:   priorStateData.ServiceAuth401Redirect,
					SkipAppLauncherLoginPage: priorStateData.SkipAppLauncherLoginPage,
					OptionsPreflightBypass:   priorStateData.OptionsPreflightBypass,

					// Lists
					CustomPages:       priorStateData.CustomPages,
					Tags:              priorStateData.Tags,
					SelfHostedDomains: priorStateData.SelfHostedDomains,
					AllowedIdPs:       priorStateData.AllowedIdPs,
				}

				fmt.Printf("DEBUG: Account ID value: %v, isNull: %v\n", newState.AccountID.ValueString(), newState.AccountID.IsNull())
				// Handle cors_headers transformation from RawState
				// Get the raw state as JSON to access cors_headers directly
				rawStateJSON := req.RawState.JSON

				// Log the raw state for debugging
				fmt.Printf("DEBUG: Raw state JSON length: %d\n", len(rawStateJSON))

				var rawState map[string]interface{}
				if err := json.Unmarshal(rawStateJSON, &rawState); err == nil && rawState != nil {
					fmt.Printf("DEBUG: Raw state keys: %v\n", len(rawState))
					for k, v := range rawState {
						fmt.Printf("DEBUG: State key %s: type=%T value=%v\n", k, v, v)
					}

					if corsHeadersRaw, exists := rawState["cors_headers"]; exists && corsHeadersRaw != nil {
						fmt.Printf("DEBUG: cors_headers found: type=%T value=%v\n", corsHeadersRaw, corsHeadersRaw)
						var rawCorsHeaders interface{} = corsHeadersRaw
						switch v := rawCorsHeaders.(type) {
						case []interface{}:
							// It's an array (v4 format)
							fmt.Printf("DEBUG: cors_headers is array with %d elements\n", len(v))
							if len(v) > 0 {
								// Take the first element
								if corsObj, ok := v[0].(map[string]interface{}); ok {
									// Convert to the proper model structure
									corsHeaders := &ZeroTrustAccessApplicationCORSHeadersModel{}

									// Map the fields from the object
									if val, ok := corsObj["allow_all_headers"].(bool); ok {
										corsHeaders.AllowAllHeaders = types.BoolValue(val)
									}
									if val, ok := corsObj["allow_all_methods"].(bool); ok {
										corsHeaders.AllowAllMethods = types.BoolValue(val)
									}
									if val, ok := corsObj["allow_all_origins"].(bool); ok {
										corsHeaders.AllowAllOrigins = types.BoolValue(val)
									}
									if val, ok := corsObj["allow_credentials"].(bool); ok {
										corsHeaders.AllowCredentials = types.BoolValue(val)
									}
									if val, ok := corsObj["max_age"].(float64); ok {
										corsHeaders.MaxAge = types.Float64Value(val)
									}

									// Handle array fields
									if arr, ok := corsObj["allowed_headers"].([]interface{}); ok {
										headers := make([]types.String, len(arr))
										for i, h := range arr {
											if str, ok := h.(string); ok {
												headers[i] = types.StringValue(str)
											}
										}
										corsHeaders.AllowedHeaders = &headers
									}
									if arr, ok := corsObj["allowed_methods"].([]interface{}); ok {
										methods := make([]types.String, len(arr))
										for i, m := range arr {
											if str, ok := m.(string); ok {
												methods[i] = types.StringValue(str)
											}
										}
										corsHeaders.AllowedMethods = &methods
									}
									if arr, ok := corsObj["allowed_origins"].([]interface{}); ok {
										origins := make([]types.String, len(arr))
										for i, o := range arr {
											if str, ok := o.(string); ok {
												origins[i] = types.StringValue(str)
											}
										}
										corsHeaders.AllowedOrigins = &origins
									}

									newState.CORSHeaders = corsHeaders
								}
							} else {
								// Empty array, set to null
								newState.CORSHeaders = nil
							}
						case map[string]interface{}:
							// It's already an object (v5 format), need to convert to proper model
							corsHeaders := &ZeroTrustAccessApplicationCORSHeadersModel{}

							// Map the fields from the object
							if val, ok := v["allow_all_headers"].(bool); ok {
								corsHeaders.AllowAllHeaders = types.BoolValue(val)
							}
							if val, ok := v["allow_all_methods"].(bool); ok {
								corsHeaders.AllowAllMethods = types.BoolValue(val)
							}
							if val, ok := v["allow_all_origins"].(bool); ok {
								corsHeaders.AllowAllOrigins = types.BoolValue(val)
							}
							if val, ok := v["allow_credentials"].(bool); ok {
								corsHeaders.AllowCredentials = types.BoolValue(val)
							}
							if val, ok := v["max_age"].(float64); ok {
								corsHeaders.MaxAge = types.Float64Value(val)
							}

							// Handle array fields
							if arr, ok := v["allowed_headers"].([]interface{}); ok {
								headers := make([]types.String, len(arr))
								for i, h := range arr {
									if str, ok := h.(string); ok {
										headers[i] = types.StringValue(str)
									}
								}
								corsHeaders.AllowedHeaders = &headers
							}
							if arr, ok := v["allowed_methods"].([]interface{}); ok {
								methods := make([]types.String, len(arr))
								for i, m := range arr {
									if str, ok := m.(string); ok {
										methods[i] = types.StringValue(str)
									}
								}
								corsHeaders.AllowedMethods = &methods
							}
							if arr, ok := v["allowed_origins"].([]interface{}); ok {
								origins := make([]types.String, len(arr))
								for i, o := range arr {
									if str, ok := o.(string); ok {
										origins[i] = types.StringValue(str)
									}
								}
								corsHeaders.AllowedOrigins = &origins
							}

							newState.CORSHeaders = corsHeaders
						}
					}
				}

				// Handle other SingleNestedAttribute fields that need array-to-null conversion
				// Handle scim_config
				if scimConfigRaw, exists := rawState["scim_config"]; exists && scimConfigRaw != nil {
					switch v := scimConfigRaw.(type) {
					case []interface{}:
						if len(v) == 0 {
							// Empty array, SCIMConfig is already nil by default
						}
						// Non-empty arrays would need full conversion logic here
					}
				}

				// Handle landing_page_design
				if landingPageRaw, exists := rawState["landing_page_design"]; exists && landingPageRaw != nil {
					switch v := landingPageRaw.(type) {
					case []interface{}:
						if len(v) == 0 {
							// Empty array, LandingPageDesign uses customfield.NestedObject
							// Leave as default (zero value)
						}
						// Non-empty arrays would need full conversion logic here
					}
				}

				// Handle saas_app
				if saasAppRaw, exists := rawState["saas_app"]; exists && saasAppRaw != nil {
					switch v := saasAppRaw.(type) {
					case []interface{}:
						if len(v) == 0 {
							// Empty array, SaaSApp uses customfield.NestedObject
							// Leave as default (zero value)
						}
						// Non-empty arrays would need full conversion logic here
					}
				}

				// Log the new state for debugging
				fmt.Printf("DEBUG: Setting new state with CORSHeaders=%v\n", newState.CORSHeaders)

				// Handle policies transformation from list of strings to list of objects
				if policiesRaw, exists := rawState["policies"]; exists && policiesRaw != nil {
					fmt.Printf("DEBUG: policies found: type=%T value=%v\n", policiesRaw, policiesRaw)
					switch v := policiesRaw.(type) {
					case []interface{}:
						// It's an array of policy IDs (v4 format)
						fmt.Printf("DEBUG: policies is array with %d elements\n", len(v))
						if len(v) > 0 {
							policies := make([]ZeroTrustAccessApplicationPoliciesModel, len(v))
							for i, policyID := range v {
								if idStr, ok := policyID.(string); ok {
									policies[i] = ZeroTrustAccessApplicationPoliciesModel{
										ID: types.StringValue(idStr),
									}
								}
							}
							newState.Policies = &policies
						}
					default:
						fmt.Printf("DEBUG: policies has unexpected type: %T\n", v)
					}
				}

				// Dump the entire new state as JSON for debugging
				newStateJSON, err := json.MarshalIndent(newState, "", "  ")
				if err != nil {
					fmt.Printf("DEBUG: Failed to marshal new state to JSON: %v\n", err)
				} else {
					fmt.Printf("DEBUG: New state JSON:\n%s\n", string(newStateJSON))
				}

				// Marshal the upgraded state
				resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)

				fmt.Printf("DIAGNOSTICS")
				spew.Dump(resp.Diagnostics)

			},
		},
	}
}
