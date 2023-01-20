package sdkv2provider

import (
	"context"
	"fmt"
	"log"

	"time"

	"reflect"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
)

func resourceCloudflareZoneSettingsOverride() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareZoneSettingsOverrideSchema(),
		CreateContext: resourceCloudflareZoneSettingsOverrideCreate,
		ReadContext:   resourceCloudflareZoneSettingsOverrideRead,
		UpdateContext: resourceCloudflareZoneSettingsOverrideUpdate,
		DeleteContext: resourceCloudflareZoneSettingsOverrideDelete,
		Description:   "Provides a resource which customizes Cloudflare zone settings.",
	}
}

var fetchAsSingleSetting = []string{
	"binary_ast",
	"h2_prioritization",
	"image_resizing",
	"early_hints",
	"origin_max_http_version",
}

func resourceCloudflareZoneSettingsOverrideCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)
	d.SetId(zoneID)

	tflog.Info(ctx, fmt.Sprintf("Creating zone settings resource for zone ID: %s", d.Id()))

	// do extra initial read to get initial_settings before updating
	zoneSettings, err := client.ZoneSettings(ctx, d.Id())
	if err != nil {
		return diag.FromErr(errors.Wrap(err, fmt.Sprintf("Error reading initial settings for zone %q", d.Id())))
	}

	if err = updateZoneSettingsResponseWithSingleZoneSettings(ctx, zoneSettings, d.Id(), client); err != nil {
		return diag.FromErr(err)
	}

	if _, ok := d.GetOk("settings.0.universal_ssl"); ok {
		// pulling USSL status and wrapping it into a cloudflare.ZoneSetting that we can set initial_settings
		if err = updateZoneSettingsResponseWithUniversalSSLSettings(ctx, zoneSettings, d.Id(), client); err != nil {
			return diag.FromErr(err)
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("Read CloudflareZone initial settings: %#v", zoneSettings))

	if err := d.Set("initial_settings", flattenZoneSettings(ctx, d, zoneSettings.Result, true)); err != nil {
		tflog.Warn(ctx, fmt.Sprintf("Error setting initial_settings for zone %q: %s", d.Id(), err))
	}

	d.Set("initial_settings_read_at", time.Now().UTC().Format(time.RFC3339Nano))

	// set readonly setting so that update can behave correctly
	if err := d.Set("readonly_settings", flattenReadOnlyZoneSettings(ctx, zoneSettings.Result)); err != nil {
		tflog.Warn(ctx, fmt.Sprintf("Error setting readonly_settings for zone %q: %s", d.Id(), err))
	}

	tflog.Debug(ctx, fmt.Sprintf("Saved CloudflareZone initial settings: %#v", d.Get("initial_settings")))

	return resourceCloudflareZoneSettingsOverrideUpdate(ctx, d, meta)
}

func updateZoneSettingsResponseWithSingleZoneSettings(ctx context.Context, zoneSettings *cloudflare.ZoneSettingResponse, zoneId string, client *cloudflare.API) error {
	for _, settingName := range fetchAsSingleSetting {
		singleSetting, err := client.ZoneSingleSetting(ctx, zoneId, settingName)
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("Error reading setting '%q' for zone %q", settingName, zoneId))
		}
		zoneSettings.Result = append(zoneSettings.Result, singleSetting)
	}
	return nil
}

func updateZoneSettingsResponseWithUniversalSSLSettings(ctx context.Context, zoneSettings *cloudflare.ZoneSettingResponse, zoneId string, client *cloudflare.API) error {
	ussl, err := client.UniversalSSLSettingDetails(ctx, zoneId)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Error reading initial Universal SSL settings for zone %q", zoneId))
	}

	usslToZoneSetting := cloudflare.ZoneSetting{
		ID:       "universal_ssl",
		Value:    stringFromBool(ussl.Enabled),
		Editable: true,
	}

	zoneSettings.Result = append(zoneSettings.Result, usslToZoneSetting)

	return nil
}

func resourceCloudflareZoneSettingsOverrideRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	zone, err := client.ZoneDetails(ctx, d.Id())
	if err != nil {
		var notFoundError *cloudflare.NotFoundError
		if errors.As(err, &notFoundError) {
			tflog.Info(ctx, fmt.Sprintf("Zone %q not found", d.Id()))
			d.SetId("")
			return nil
		} else {
			return diag.FromErr(errors.Wrap(err, fmt.Sprintf("Error reading zone %q", d.Id())))
		}
	}

	d.Set(consts.ZoneIDSchemaKey, d.Id())

	// not all settings are visible to all users, so this might be a subset
	// assume (for now) that user can see/do everything
	zoneSettings, err := client.ZoneSettings(ctx, d.Id())
	if err != nil {
		return diag.FromErr(errors.Wrap(err, fmt.Sprintf("Error reading settings for zone %q", d.Id())))
	}

	if err = updateZoneSettingsResponseWithSingleZoneSettings(ctx, zoneSettings, d.Id(), client); err != nil {
		return diag.FromErr(err)
	}

	if _, ok := d.GetOk("settings.0.universal_ssl"); ok {
		if err = updateZoneSettingsResponseWithUniversalSSLSettings(ctx, zoneSettings, d.Id(), client); err != nil {
			return diag.FromErr(err)
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("Read CloudflareZone Settings: %#v", zoneSettings))

	d.Set("zone_status", zone.Status)
	d.Set("zone_type", zone.Type)

	newZoneSettings := flattenZoneSettings(ctx, d, zoneSettings.Result, false)
	// if polish is off (or we don't know) we need to ignore what comes back from the api for webp
	if polish, ok := newZoneSettings[0]["polish"]; !ok || polish.(string) == "" || polish.(string) == "off" {
		newZoneSettings[0]["webp"] = d.Get("settings.0.webp").(string)
	}

	if err := d.Set("settings", newZoneSettings); err != nil {
		log.Printf("[WARN] Error setting settings for zone %q: %s", d.Id(), err)
	}

	if err := d.Set("readonly_settings", flattenReadOnlyZoneSettings(ctx, zoneSettings.Result)); err != nil {
		log.Printf("[WARN] Error setting readonly_settings for zone %q: %s", d.Id(), err)
	}

	return nil
}

func flattenZoneSettings(ctx context.Context, d *schema.ResourceData, settings []cloudflare.ZoneSetting, flattenAll bool) []map[string]interface{} {
	cfg := map[string]interface{}{}
	for _, s := range settings {
		if s.ID == "0rtt" { // NOTE: 0rtt is an invalid attribute in HCLs grammar.  Remap to `zero_rtt`
			s.ID = "zero_rtt"
		}

		if !settingInSchema(s.ID) {
			log.Printf("[WARN] Value not in schema returned from API zone settings (is it new?) - %q : %#v", s.ID, s.Value)
			continue
		}
		if _, ok := d.GetOkExists(fmt.Sprintf("settings.0.%s", s.ID)); !ok && !flattenAll {
			// don't put settings that were never specified in the update request
			continue
		}

		if s.ID == "minify" || s.ID == "mobile_redirect" {
			cfg[s.ID] = []interface{}{s.Value.(map[string]interface{})}
		} else if s.ID == "security_header" {
			cfg[s.ID] = []interface{}{s.Value.(map[string]interface{})["strict_transport_security"]}
		} else if listValues, ok := s.Value.([]interface{}); ok {
			cfg[s.ID] = listValues
		} else if strValue, ok := s.Value.(string); ok {
			cfg[s.ID] = strValue
		} else if floatValue, ok := s.Value.(float64); ok {
			cfg[s.ID] = int(floatValue)
		} else {
			tflog.Warn(ctx, fmt.Sprintf("Unexpected value type found in API zone settings - %q : %#v", s.ID, s.Value))
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("Flattened Cloudflare Zone Settings: %#v", cfg))

	return []map[string]interface{}{cfg}
}

func settingInSchema(val string) bool {
	for k := range resourceCloudflareZoneSettingsSchema {
		if val == k {
			return true
		}
	}
	return false
}

func flattenReadOnlyZoneSettings(ctx context.Context, settings []cloudflare.ZoneSetting) []string {
	ids := make([]string, 0)
	for _, zs := range settings {
		if !zs.Editable {
			ids = append(ids, zs.ID)
		}
	}
	tflog.Debug(ctx, fmt.Sprintf("Flattened Cloudflare Read Only Zone Settings: %#v", ids))

	return ids
}

func updateSingleZoneSettings(ctx context.Context, zoneSettings []cloudflare.ZoneSetting, client *cloudflare.API, zoneID string) ([]cloudflare.ZoneSetting, error) {
	var indexesToCut []int
	for i, setting := range zoneSettings {
		if contains(fetchAsSingleSetting, setting.ID) {
			_, err := client.UpdateZoneSingleSetting(ctx, zoneID, setting.ID, setting)
			if err != nil {
				return zoneSettings, err
			}
			indexesToCut = append(indexesToCut, i)
		}
	}

	offset := 0
	for _, indexToCut := range indexesToCut {
		adjustedIndexToCut := indexToCut - offset
		zoneSettings = append(zoneSettings[:adjustedIndexToCut], zoneSettings[adjustedIndexToCut+1:]...)
		offset += 1
	}
	return zoneSettings, nil
}

func updateUniversalSSLSetting(ctx context.Context, zoneSettings []cloudflare.ZoneSetting, client *cloudflare.API, zoneID string) ([]cloudflare.ZoneSetting, error) {
	indexToCut := -1
	for i, setting := range zoneSettings {
		// Skipping USSL Update if value is empty, especially when we are reverting to the initial state and we did not had the information
		if setting.ID == "universal_ssl" {
			if setting.Value.(string) != "" {
				_, err := client.EditUniversalSSLSetting(ctx, zoneID, cloudflare.UniversalSSLSetting{Enabled: boolFromString(setting.Value.(string))})
				if err != nil {
					return zoneSettings, err
				}
			}
			indexToCut = i
		}
	}

	if indexToCut != -1 {
		zoneSettings = append(zoneSettings[:indexToCut], zoneSettings[indexToCut+1:]...)
	}

	return zoneSettings, nil
}

func resourceCloudflareZoneSettingsOverrideUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	if cfg, ok := d.GetOkExists("settings"); ok && cfg != nil && len(cfg.([]interface{})) > 0 {
		readOnlySettings := expandInterfaceToStringList(d.Get("readonly_settings"))
		zoneSettings, err := expandOverriddenZoneSettings(d, "settings", readOnlySettings)
		if err != nil {
			return diag.FromErr(err)
		}

		tflog.Debug(ctx, fmt.Sprintf("Cloudflare Zone Settings update configuration: %#v", zoneSettings))

		if zoneSettings, err = updateSingleZoneSettings(ctx, zoneSettings, client, d.Id()); err != nil {
			return diag.FromErr(err)
		}

		if zoneSettings, err = updateUniversalSSLSetting(ctx, zoneSettings, client, d.Id()); err != nil {
			return diag.FromErr(err)
		}

		if len(zoneSettings) > 0 {
			_, err = client.UpdateZoneSettings(ctx, d.Id(), zoneSettings)
			if err != nil {
				return diag.FromErr(err)
			}
		} else {
			tflog.Debug(ctx, fmt.Sprintf("Skipped update call because no settings were set"))
		}
	}

	return resourceCloudflareZoneSettingsOverrideRead(ctx, d, meta)
}

func expandOverriddenZoneSettings(d *schema.ResourceData, settingsKey string, readOnlySettings []string) ([]cloudflare.ZoneSetting, error) {
	zoneSettings := make([]cloudflare.ZoneSetting, 0)

	keyFormat := fmt.Sprintf("%s.0.%%s", settingsKey)

	for k := range resourceCloudflareZoneSettingsSchema {
		// we only update if the user set the value non-empty before, and its different from the read value
		// note that if user removes an attribute, we don't do anything
		if settingValue, ok := d.GetOkExists(fmt.Sprintf(keyFormat, k)); ok && d.HasChange(fmt.Sprintf(keyFormat, k)) {
			zoneSettingValue, err := expandZoneSetting(d, keyFormat, k, settingValue, readOnlySettings)
			if err != nil {
				return zoneSettings, err
			}

			// Remap zero_rtt key back to Cloudflare's setting name, 0rtt
			if k == "zero_rtt" {
				k = "0rtt"
			}

			if zoneSettingValue != nil {
				newZoneSetting := cloudflare.ZoneSetting{
					ID:    k,
					Value: zoneSettingValue,
				}
				zoneSettings = append(zoneSettings, newZoneSetting)
			}
		}
	}
	return zoneSettings, nil
}

func expandZoneSetting(d *schema.ResourceData, keyFormatString, k string, settingValue interface{}, readOnlySettings []string) (interface{}, error) {
	if contains(readOnlySettings, k) {
		return nil, fmt.Errorf("invalid zone setting %q (value: %v) found - cannot be set as it is read only", k, settingValue)
	}

	var zoneSettingValue interface{}
	switch k {
	case "webp":
		{
			// only ever set webp if polish is on
			polishKey := fmt.Sprintf(keyFormatString, "polish")
			polish := d.Get(polishKey).(string)

			if polish != "" && polish != "off" {
				zoneSettingValue = settingValue
			}
		}
	case "minify", "mobile_redirect":
		{
			listValue := settingValue.([]interface{})
			if len(listValue) > 0 && listValue != nil {
				zoneSettingValue = listValue[0].(map[string]interface{})
			}
		}
	case "security_header":
		{
			listValue := settingValue.([]interface{})
			if len(listValue) > 0 && listValue != nil {
				zoneSettingValue = map[string]interface{}{
					"strict_transport_security": listValue[0].(map[string]interface{}),
				}
			}
		}
	default:
		{
			zoneSettingValue = settingValue
		}
	}
	return zoneSettingValue, nil
}

func resourceCloudflareZoneSettingsOverrideDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	if cfg, ok := d.GetOkExists("settings"); ok && cfg != nil && len(cfg.([]interface{})) > 0 {
		readOnlySettings := expandInterfaceToStringList(d.Get("readonly_settings"))

		zoneSettings, err := expandRevertibleZoneSettings(d, readOnlySettings)
		if err != nil {
			return diag.FromErr(err)
		}

		tflog.Debug(ctx, fmt.Sprintf("Reverting Cloudflare Zone Settings to initial settings with update configuration: %#v", zoneSettings))

		if zoneSettings, err = updateSingleZoneSettings(ctx, zoneSettings, client, d.Id()); err != nil {
			return diag.FromErr(err)
		}

		if zoneSettings, err = updateUniversalSSLSetting(ctx, zoneSettings, client, d.Id()); err != nil {
			return diag.FromErr(err)
		}

		if len(zoneSettings) > 0 {
			_, err = client.UpdateZoneSettings(ctx, d.Id(), zoneSettings)
			if err != nil {
				return diag.FromErr(err)
			}
		} else {
			tflog.Debug(ctx, fmt.Sprintf("Skipped call to revert settings because no settings were changed"))
		}
	}
	return nil
}

func expandRevertibleZoneSettings(d *schema.ResourceData, readOnlySettings []string) ([]cloudflare.ZoneSetting, error) {
	zoneSettings := make([]cloudflare.ZoneSetting, 0)

	keyFormat := fmt.Sprintf("%s.0.%%s", "initial_settings")

	for k := range resourceCloudflareZoneSettingsSchema {
		initialKey := fmt.Sprintf("initial_settings.0.%s", k)
		initialVal := d.Get(initialKey)
		currentKey := fmt.Sprintf("settings.0.%s", k)

		if k == "zero_rtt" {
			k = "0rtt"
		}

		// if the value was never set we don't need to revert it
		if currentVal, ok := d.GetOk(currentKey); ok && !schemaValueEquals(initialVal, currentVal) {
			zoneSettingValue, err := expandZoneSetting(d, keyFormat, k, initialVal, readOnlySettings)
			if err != nil {
				return zoneSettings, err
			}

			if zoneSettingValue != nil {
				newZoneSetting := cloudflare.ZoneSetting{
					ID:    k,
					Value: zoneSettingValue,
				}
				zoneSettings = append(zoneSettings, newZoneSetting)
			}
		}
	}
	return zoneSettings, nil
}

func schemaValueEquals(a, b interface{}) bool {
	// this is the same equality check used in d.HasChange

	// If the type implements the Equal interface, then call that
	// instead of just doing a reflect.DeepEqual. An example where this is
	// needed is *Set
	if eq, ok := a.(schema.Equal); ok {
		return eq.Equal(b)
	}

	return reflect.DeepEqual(a, b)
}
