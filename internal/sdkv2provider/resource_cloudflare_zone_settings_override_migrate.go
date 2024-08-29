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

	upgrade := func(state map[string]interface{}, name string) map[string]interface{} {
		if val, ok := state[name]; ok && val != nil {
			delete(state[name].([]interface{})[0].(map[string]interface{}), "mobile_redirect")
		}
		return state
	}

	state := upgrade(rawState, "settings")
	state = upgrade(state, "initial_settings")

	return state, nil
}

func resourceCloudflareZoneSettingsOverrideV1() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"settings": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: resourceCloudflareZoneSettingsSchemaV1,
				},
			},

			"initial_settings": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: resourceCloudflareZoneSettingsSchemaV1,
				},
			},
		},
	}
}

func resourceCloudflareZoneSettingsOverrideStateUpgradeV2(
	_ context.Context,
	rawState map[string]interface{},
	_ interface{},
) (map[string]interface{}, error) {
	errMsg := "could not upgrade cloudflare_zone_settings_override from v1 to v2"

	if rawState == nil {
		return nil, fmt.Errorf("%s: state is nil", errMsg)
	}

	upgrade := func(state map[string]interface{}, name string) map[string]interface{} {
		if val, ok := state[name]; ok && val != nil {
			delete(state[name].([]interface{})[0].(map[string]interface{}), "minify")
		}
		return state
	}

	state := upgrade(rawState, "settings")
	state = upgrade(state, "initial_settings")

	return state, nil
}
