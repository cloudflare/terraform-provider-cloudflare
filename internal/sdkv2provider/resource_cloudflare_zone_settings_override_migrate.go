package sdkv2provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareZoneSettingsOverrideV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"settings": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: resourceCloudflareZoneSettingsSchemaV0,
				},
			},

			"initial_settings": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: resourceCloudflareZoneSettingsSchemaV0,
				},
			},
		},
	}
}

func resourceCloudflareZoneSettingsOverrideStateUpgradeV1(
	_ context.Context,
	rawState map[string]interface{},
	_ interface{},
) (map[string]interface{}, error) {
	errMsg := "could not upgrade cloudflare_zone_settings_override from v0 to v1"

	if rawState == nil {
		return nil, fmt.Errorf("%s: state is nil", errMsg)
	}

	upgrade := func(state map[string]interface{}, name string) (map[string]interface{}, error) {
		settings, ok := state[name].([]map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("%s: field %q is not an array of objects", errMsg, name)
		}

		if len(settings) == 0 {
			return nil, fmt.Errorf("%s: field %q is an empty array", errMsg, name)
		}

		delete(settings[0], "mobile_redirect")

		if len(settings[0]) == 0 {
			state[name] = []map[string]interface{}{}
		} else {
			state[name] = settings
		}

		return state, nil
	}

	state, err := upgrade(rawState, "settings")
	if err != nil {
		return nil, err
	}

	state, err = upgrade(state, "initial_settings")
	if err != nil {
		return nil, err
	}

	return state, nil
}
