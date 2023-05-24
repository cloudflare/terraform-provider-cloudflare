package sdkv2provider

import (
	"fmt"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareRateLimitSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.ZoneIDSchemaKey: {
			Description: consts.ZoneIDSchemaDescription,
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},

		"threshold": {
			Type:         schema.TypeInt,
			Required:     true,
			ValidateFunc: validation.IntBetween(1, 1000000),
			Description:  "The threshold that triggers the rate limit mitigations, combine with period.",
		},

		"period": {
			Type:         schema.TypeInt,
			Required:     true,
			ValidateFunc: validation.IntBetween(1, 86400),
			Description:  "The time in seconds to count matching traffic. If the count exceeds threshold within this period the action will be performed.",
		},

		"action": {
			Type:        schema.TypeList,
			Required:    true,
			MinItems:    1,
			MaxItems:    1,
			Description: "The action to be performed when the threshold of matched traffic within the period defined is exceeded.",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"mode": {
						Type:         schema.TypeString,
						Required:     true,
						ValidateFunc: validation.StringInSlice([]string{"simulate", "ban", "challenge", "js_challenge", "managed_challenge"}, true),
						Description:  fmt.Sprintf("The type of action to perform. %s", renderAvailableDocumentationValuesStringSlice([]string{"simulate", "ban", "challenge", "js_challenge", "managed_challenge"})),
					},

					"timeout": {
						Type:         schema.TypeInt,
						Optional:     true,
						ValidateFunc: validation.IntBetween(1, 86400),
						Description:  "The time in seconds as an integer to perform the mitigation action. This field is required if the `mode` is either `simulate` or `ban`. Must be the same or greater than the period.",
					},

					"response": {
						Type:        schema.TypeList,
						Optional:    true,
						MinItems:    1,
						MaxItems:    1,
						Description: "Custom content-type and body to return, this overrides the custom error for the zone. This field is not required. Omission will result in default HTML error page.",
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"content_type": {
									Type:         schema.TypeString,
									Required:     true,
									ValidateFunc: validation.StringInSlice([]string{"text/plain", "text/xml", "application/json"}, true),
									Description:  fmt.Sprintf("The content-type of the body. %s", renderAvailableDocumentationValuesStringSlice([]string{"text/plain", "text/xml", "application/json"})),
								},

								"body": {
									Type:         schema.TypeString,
									Required:     true,
									ValidateFunc: validation.StringLenBetween(0, 10240),
									// maybe good to hash the body before saving in state file?
									Description: "The body to return, the content here should conform to the `content_type`.",
								},
							},
						},
					},
				},
			},
		},

		"match": {
			Type:        schema.TypeList,
			Optional:    true,
			Computed:    true,
			MaxItems:    1,
			Description: "Determines which traffic the rate limit counts towards the threshold. By default matches all traffic in the zone.",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"request": {
						Type:        schema.TypeList,
						Optional:    true,
						Computed:    true,
						MinItems:    1,
						MaxItems:    1,
						Description: "Matches HTTP requests (from the client to Cloudflare).",
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"methods": {
									Type:        schema.TypeSet,
									Optional:    true,
									Computed:    true,
									Description: fmt.Sprintf("HTTP Methods to match traffic on. %s", renderAvailableDocumentationValuesStringSlice(allowedHTTPMethods)),
									Elem: &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice(allowedHTTPMethods, true),
									},
								},

								"schemes": {
									Type:        schema.TypeSet,
									Optional:    true,
									Computed:    true,
									Description: fmt.Sprintf("HTTP schemes to match traffic on. %s", renderAvailableDocumentationValuesStringSlice(allowedSchemes)),
									Elem: &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice(allowedSchemes, true),
									},
								},

								"url_pattern": {
									Type:         schema.TypeString,
									Optional:     true,
									Computed:     true,
									ValidateFunc: validation.StringLenBetween(0, 1024),
									Description:  "The URL pattern to match comprised of the host and path, i.e. example.org/path. Wildcard are expanded to match applicable traffic, query strings are not matched. Use _ for all traffic to your zone.",
								},
							},
						},
					},

					"response": {
						Type:        schema.TypeList,
						Optional:    true,
						Computed:    true,
						MinItems:    1,
						MaxItems:    1,
						Description: "Matches HTTP responses before they are returned to the client from Cloudflare. If this is defined, then the entire counting of traffic occurs at this stage.",
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"statuses": {
									Type:        schema.TypeSet,
									Optional:    true,
									Computed:    true,
									Elem:        &schema.Schema{Type: schema.TypeInt},
									Description: "HTTP Status codes, can be one, many or indicate all by not providing this value.",
								},

								"origin_traffic": {
									Type:        schema.TypeBool,
									Optional:    true,
									Computed:    true,
									Description: "Only count traffic that has come from your origin servers. If true, cached items that Cloudflare serve will not count towards rate limiting.",
								},

								"headers": {
									Type:        schema.TypeList,
									Optional:    true,
									Elem:        headersElem,
									Description: "List of HTTP headers maps to match the origin response on.",
								},
							},
						},
					},
				},
			},
		},

		"disabled": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			Description: "Whether this ratelimit is currently disabled.",
		},

		"description": {
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringLenBetween(0, 1024),
			Description:  "A note that you can use to describe the reason for a rate limit. This value is sanitized and all tags are removed.",
		},

		"bypass_url_patterns": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},

		"correlate": {
			Type:        schema.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "Determines how rate limiting is applied. By default if not specified, rate limiting applies to the clients IP address.",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"by": {
						Type:         schema.TypeString,
						Optional:     true,
						ValidateFunc: validation.StringInSlice([]string{"nat"}, true),
						Description:  fmt.Sprintf("If set to 'nat', NAT support will be enabled for rate limiting. %s", renderAvailableDocumentationValuesStringSlice([]string{"nat"})),
					},
				},
			},
		},
	}
}

var headersElem = &schema.Schema{
	Type: schema.TypeMap,
	Elem: &schema.Schema{
		Type: schema.TypeString,
	},
	ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
		headerElemValidators := headersElemValidators()
		headerFields, ok := val.(map[string]interface{})

		if !ok {
			errs = append(errs, fmt.Errorf("got invalid map for rule element"))
			return
		}

		for k, v := range headerFields {
			if _, ok := headerElemValidators[k]; !ok {
				errs = append(errs, fmt.Errorf("%s is not supported in a response header", k))
			}

			validationFunc := headerElemValidators[k]
			delete(headerElemValidators, k)
			if validationFunc == nil {
				continue
			}

			w, e := validationFunc(v, k)
			warns = append(warns, w...)
			errs = append(errs, e...)
		}

		// attributes with non-nil validators must be set
		for k, v := range headerElemValidators {
			if v == nil {
				continue
			}
			errs = append(errs, fmt.Errorf("%s must be set in a response header", k))
		}

		return
	},
}

func headersElemValidators() map[string]schema.SchemaValidateFunc {
	v := make(map[string]schema.SchemaValidateFunc)

	v["name"] = validation.StringIsNotEmpty
	v["op"] = validation.StringInSlice([]string{"eq", "ne"}, false)
	v["value"] = validation.StringIsNotEmpty
	return v
}
