package provider

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareRateLimitSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"zone_id": {
			Description: "The zone identifier to target for the resource.",
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},

		"threshold": {
			Type:         schema.TypeInt,
			Required:     true,
			ValidateFunc: validation.IntBetween(1, 1000000),
		},

		"period": {
			Type:         schema.TypeInt,
			Required:     true,
			ValidateFunc: validation.IntBetween(1, 86400),
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
						ValidateFunc: validation.StringInSlice([]string{"simulate", "ban", "challenge", "js_challenge", "managed_challenge"}, true),
					},

					"timeout": {
						Type:         schema.TypeInt,
						Optional:     true,
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

								"headers": {
									Type:     schema.TypeList,
									Optional: true,
									Elem:     headersElem,
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

		"correlate": {
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"by": {
						Type:         schema.TypeString,
						Optional:     true,
						ValidateFunc: validation.StringInSlice([]string{"nat"}, true),
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
