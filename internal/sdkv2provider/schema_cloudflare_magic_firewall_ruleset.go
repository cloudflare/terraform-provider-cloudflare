package sdkv2provider

import (
	"fmt"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareMagicFirewallRulesetSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.AccountIDSchemaKey: {
			Description: "The account identifier to target for the resource.",
			Type:        schema.TypeString,
			Required:    true,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		"description": {
			Type:     schema.TypeString,
			Optional: true,
		},
		// the order of firewall rules has to be maintained, so we are using a list of maps here and validate the
		// map using a custom validator
		"rules": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ruleElem,
		},
	}
}

var ruleElem = &schema.Schema{
	Type: schema.TypeMap,
	Elem: &schema.Schema{
		Type: schema.TypeString,
	},
	ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
		ruleElemValidators := ruleElemValidators()
		ruleFields, ok := val.(map[string]interface{})

		if !ok {
			errs = append(errs, fmt.Errorf("got invalid map for rule element"))
			return
		}

		for k, v := range ruleFields {
			if _, ok := ruleElemValidators[k]; !ok {
				errs = append(errs, fmt.Errorf("%s is not supported in a rule", k))
			}

			validationFunc := ruleElemValidators[k]
			delete(ruleElemValidators, k)
			if validationFunc == nil {
				continue
			}

			w, e := validationFunc(v, k)
			warns = append(warns, w...)
			errs = append(errs, e...)
		}

		// attributes with non-nil validators must be set
		for k, v := range ruleElemValidators {
			if v == nil {
				continue
			}
			errs = append(errs, fmt.Errorf("%s must be set in a rule", k))
		}

		return
	},
}
