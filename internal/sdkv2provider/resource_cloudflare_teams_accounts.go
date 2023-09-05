package sdkv2provider

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareTeamsAccount() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareTeamsAccountSchema(),
		ReadContext:   resourceCloudflareTeamsAccountRead,
		UpdateContext: resourceCloudflareTeamsAccountUpdate,
		CreateContext: resourceCloudflareTeamsAccountUpdate,
		// This resource is a top-level account configuration and cant be "deleted"
		Delete: func(_ *schema.ResourceData, _ interface{}) error { return nil },
		Importer: &schema.ResourceImporter{
			StateContext: resourceCloudflareTeamsAccountImport,
		},
		Description: heredoc.Doc(`
			Provides a Cloudflare Teams Account resource. The Teams Account
			resource defines configuration for secure web gateway.
		`),
	}
}

func resourceCloudflareTeamsAccountRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)

	configuration, err := client.TeamsAccountConfiguration(ctx, accountID)
	if err != nil {
		if strings.Contains(err.Error(), "HTTP status 400") {
			tflog.Info(ctx, fmt.Sprintf("Teams Account config %s does not exists", d.Id()))
			d.SetId("")
			return nil
		}
		return diag.FromErr(fmt.Errorf("error finding Teams Account config %q: %w", d.Id(), err))
	}

	if configuration.Settings.BlockPage != nil {
		if err := d.Set("block_page", flattenBlockPageConfig(configuration.Settings.BlockPage)); err != nil {
			return diag.FromErr(fmt.Errorf("error parsing account block page config: %w", err))
		}
	}

	if configuration.Settings.Antivirus != nil {
		if err := d.Set("antivirus", flattenAntivirusConfig(configuration.Settings.Antivirus)); err != nil {
			return diag.FromErr(fmt.Errorf("error parsing account antivirus config: %w", err))
		}
	}

	if configuration.Settings.TLSDecrypt != nil {
		if err := d.Set("tls_decrypt_enabled", configuration.Settings.TLSDecrypt.Enabled); err != nil {
			return diag.FromErr(fmt.Errorf("error parsing account tls decrypt enablement: %w", err))
		}
	}

	if configuration.Settings.ProtocolDetection != nil {
		if err := d.Set("protocol_detection_enabled", configuration.Settings.ProtocolDetection.Enabled); err != nil {
			return diag.FromErr(fmt.Errorf("error parsing account protocol detection enablement: %w", err))
		}
	}

	if err := d.Set("activity_log_enabled", configuration.Settings.ActivityLog.Enabled); err != nil {
		return diag.FromErr(fmt.Errorf("error parsing account activity log enablement: %w", err))
	}

	if configuration.Settings.FIPS != nil {
		if err := d.Set("fips", flattenFIPSConfig(configuration.Settings.FIPS)); err != nil {
			return diag.FromErr(fmt.Errorf("error parsing account FIPS config: %w", err))
		}
	}

	if configuration.Settings.BrowserIsolation != nil {
		if err := d.Set("url_browser_isolation_enabled", configuration.Settings.BrowserIsolation.UrlBrowserIsolationEnabled); err != nil {
			return diag.FromErr(fmt.Errorf("error parsing account url browser isolation enablement: %w", err))
		}
	}

	logSettings, err := client.TeamsAccountLoggingConfiguration(ctx, accountID)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error finding Teams Account log settings %q: %w", d.Id(), err))
	}

	if logSettings.LoggingSettingsByRuleType != nil {
		if err := d.Set("logging", flattenTeamsLoggingSettings(&logSettings)); err != nil {
			return diag.FromErr(fmt.Errorf("error parsing teams account log settings: %w", err))
		}
	}

	deviceSettings, err := client.TeamsAccountDeviceConfiguration(ctx, accountID)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error finding Teams Account device settings %q: %w", d.Id(), err))
	}

	if err := d.Set("proxy", flattenTeamsDeviceSettings(&deviceSettings)); err != nil {
		return diag.FromErr(fmt.Errorf("error parsing teams account device settings: %w", err))
	}

	payloadLogSettings, err := client.GetDLPPayloadLogSettings(ctx, cloudflare.AccountIdentifier(accountID), cloudflare.GetDLPPayloadLogSettingsParams{})
	if err == nil {
		if err := d.Set("payload_log", flattenPayloadLogSettings(&payloadLogSettings)); err != nil {
			return diag.FromErr(fmt.Errorf("error parsing payload log settings: %w", err))
		}
	} else {
		var notFoundError *cloudflare.NotFoundError
		if !errors.As(err, &notFoundError) {
			return diag.FromErr(fmt.Errorf("error finding DLP Account config %q: %w", d.Id(), err))
		}
	}

	return nil
}

func resourceCloudflareTeamsAccountUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)
	blockPageConfig := inflateBlockPageConfig(d.Get("block_page"))
	fipsConfig := inflateFIPSConfig(d.Get("fips"))
	antivirusConfig := inflateAntivirusConfig(d.Get("antivirus"))
	loggingConfig := inflateLoggingSettings(d.Get("logging"))
	deviceConfig := inflateDeviceSettings(d.Get("proxy"))
	payloadLogSettings := inflatePayloadLogSettings(d.Get("payload_log"))
	updatedTeamsAccount := cloudflare.TeamsConfiguration{
		Settings: cloudflare.TeamsAccountSettings{
			Antivirus: antivirusConfig,
			BlockPage: blockPageConfig,
			FIPS:      fipsConfig,
		},
	}

	//nolint:staticcheck
	tlsDecrypt, ok := d.GetOkExists("tls_decrypt_enabled")
	if ok {
		updatedTeamsAccount.Settings.TLSDecrypt = &cloudflare.TeamsTLSDecrypt{Enabled: tlsDecrypt.(bool)}
	}

	//nolint:staticcheck
	protocolDetection, ok := d.GetOkExists("protocol_detection_enabled")
	if ok {
		updatedTeamsAccount.Settings.ProtocolDetection = &cloudflare.TeamsProtocolDetection{Enabled: protocolDetection.(bool)}
	}

	//nolint:staticcheck
	activtyLog, ok := d.GetOkExists("activity_log_enabled")
	if ok {
		updatedTeamsAccount.Settings.ActivityLog = &cloudflare.TeamsActivityLog{Enabled: activtyLog.(bool)}
	}

	//nolint:staticcheck
	browserIsolation, ok := d.GetOkExists("url_browser_isolation_enabled")
	if ok {
		updatedTeamsAccount.Settings.BrowserIsolation = &cloudflare.BrowserIsolation{UrlBrowserIsolationEnabled: browserIsolation.(bool)}
	}

	tflog.Debug(ctx, fmt.Sprintf("Updating Cloudflare Teams Account configuration from struct: %+v", updatedTeamsAccount))

	if _, err := client.TeamsAccountUpdateConfiguration(ctx, accountID, updatedTeamsAccount); err != nil {
		return diag.FromErr(fmt.Errorf("error updating Teams Account configuration for account %q: %w", accountID, err))
	}

	if loggingConfig != nil {
		if _, err := client.TeamsAccountUpdateLoggingConfiguration(ctx, accountID, *loggingConfig); err != nil {
			return diag.FromErr(fmt.Errorf("error updating Teams Account logging settings for account %q: %w", accountID, err))
		}
	}

	if deviceConfig != nil {
		if _, err := client.TeamsAccountDeviceUpdateConfiguration(ctx, accountID, *deviceConfig); err != nil {
			return diag.FromErr(fmt.Errorf("error updating Teams Account proxy settings for account %q: %w", accountID, err))
		}
	}

	if payloadLogSettings != nil {
		if _, err := client.UpdateDLPPayloadLogSettings(ctx, cloudflare.AccountIdentifier(accountID), *payloadLogSettings); err != nil {
			return diag.FromErr(fmt.Errorf("error updating DLP Account configuration for account %q: %w", accountID, err))
		}
	}

	d.SetId(accountID)
	return resourceCloudflareTeamsAccountRead(ctx, d, meta)
}

func resourceCloudflareTeamsAccountImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	d.SetId(d.Id())
	d.Set(consts.AccountIDSchemaKey, d.Id())

	resourceCloudflareTeamsAccountRead(ctx, d, meta)

	return []*schema.ResourceData{d}, nil
}

func flattenBlockPageConfig(blockPage *cloudflare.TeamsBlockPage) []interface{} {
	return []interface{}{map[string]interface{}{
		"enabled":          *blockPage.Enabled,
		"footer_text":      blockPage.FooterText,
		"header_text":      blockPage.HeaderText,
		"logo_path":        blockPage.LogoPath,
		"background_color": blockPage.BackgroundColor,
		"name":             blockPage.Name,
		"mailto_address":   blockPage.MailtoAddress,
		"mailto_subject":   blockPage.MailtoSubject,
	}}
}

func flattenTeamsLoggingSettings(logSettings *cloudflare.TeamsLoggingSettings) []interface{} {
	return []interface{}{map[string]interface{}{
		"redact_pii": logSettings.RedactPii,
		"settings_by_rule_type": []interface{}{map[string]interface{}{
			"dns": []interface{}{map[string]bool{
				"log_all":    logSettings.LoggingSettingsByRuleType[cloudflare.TeamsDnsRuleType].LogAll,
				"log_blocks": logSettings.LoggingSettingsByRuleType[cloudflare.TeamsDnsRuleType].LogBlocks,
			}},
			"http": []interface{}{map[string]bool{
				"log_all":    logSettings.LoggingSettingsByRuleType[cloudflare.TeamsHttpRuleType].LogAll,
				"log_blocks": logSettings.LoggingSettingsByRuleType[cloudflare.TeamsHttpRuleType].LogBlocks,
			}},
			"l4": []interface{}{map[string]bool{
				"log_all":    logSettings.LoggingSettingsByRuleType[cloudflare.TeamsL4RuleType].LogAll,
				"log_blocks": logSettings.LoggingSettingsByRuleType[cloudflare.TeamsL4RuleType].LogBlocks,
			}},
		},
		},
	},
	}
}

func inflateBlockPageConfig(blockPage interface{}) *cloudflare.TeamsBlockPage {
	blockPageList := blockPage.([]interface{})
	if len(blockPageList) != 1 {
		return nil
	}

	blockPageMap := blockPageList[0].(map[string]interface{})
	enabled := blockPageMap["enabled"].(bool)
	return &cloudflare.TeamsBlockPage{
		Enabled:         &enabled,
		FooterText:      blockPageMap["footer_text"].(string),
		HeaderText:      blockPageMap["header_text"].(string),
		LogoPath:        blockPageMap["logo_path"].(string),
		BackgroundColor: blockPageMap["background_color"].(string),
		Name:            blockPageMap["name"].(string),
		MailtoSubject:   blockPageMap["mailto_subject"].(string),
		MailtoAddress:   blockPageMap["mailto_address"].(string),
	}
}

func flattenAntivirusConfig(antivirusConfig *cloudflare.TeamsAntivirus) []interface{} {
	return []interface{}{map[string]interface{}{
		"enabled_download_phase": antivirusConfig.EnabledDownloadPhase,
		"enabled_upload_phase":   antivirusConfig.EnabledUploadPhase,
		"fail_closed":            antivirusConfig.FailClosed,
	}}
}

func flattenTeamsDeviceSettings(deviceSettings *cloudflare.TeamsDeviceSettings) []interface{} {
	return []interface{}{map[string]interface{}{
		"tcp":     deviceSettings.GatewayProxyEnabled,
		"udp":     deviceSettings.GatewayProxyUDPEnabled,
		"root_ca": deviceSettings.RootCertificateInstallationEnabled,
	}}
}

func inflateAntivirusConfig(antivirus interface{}) *cloudflare.TeamsAntivirus {
	avList := antivirus.([]interface{})

	if len(avList) != 1 {
		return nil
	}

	avMap := avList[0].(map[string]interface{})
	return &cloudflare.TeamsAntivirus{
		EnabledDownloadPhase: avMap["enabled_download_phase"].(bool),
		EnabledUploadPhase:   avMap["enabled_upload_phase"].(bool),
		FailClosed:           avMap["fail_closed"].(bool),
	}
}

func flattenFIPSConfig(fips *cloudflare.TeamsFIPS) []interface{} {
	return []interface{}{map[string]interface{}{
		"tls": fips.TLS,
	}}
}

func inflateFIPSConfig(fipsList interface{}) *cloudflare.TeamsFIPS {
	list := fipsList.([]interface{})
	if len(list) != 1 {
		return nil
	}

	m := list[0].(map[string]interface{})
	return &cloudflare.TeamsFIPS{
		TLS: m["tls"].(bool),
	}
}

func inflateLoggingSettings(log interface{}) *cloudflare.TeamsLoggingSettings {
	logList := log.([]interface{})

	if len(logList) != 1 {
		return nil
	}

	logSettings, ok := logList[0].(map[string]interface{})
	if !ok {
		return nil
	}
	logRuleSettingsList, ok := logSettings["settings_by_rule_type"].([]interface{})
	if !ok {
		return nil
	}

	logRuleSettings, ok := logRuleSettingsList[0].(map[string]interface{})
	if !ok {
		return nil
	}

	dnsRuleSettingsList, ok := logRuleSettings["dns"].([]interface{})
	if !ok {
		return nil
	}
	dnsRuleSettings := dnsRuleSettingsList[0].(map[string]interface{})

	httpRuleSettingsList, ok := logRuleSettings["http"].([]interface{})
	if !ok {
		return nil
	}
	httpRuleSettings := httpRuleSettingsList[0].(map[string]interface{})

	l4RuleSettingsList, ok := logRuleSettings["l4"].([]interface{})
	if !ok {
		return nil
	}
	l4RuleSettings := l4RuleSettingsList[0].(map[string]interface{})

	return &cloudflare.TeamsLoggingSettings{
		LoggingSettingsByRuleType: map[cloudflare.TeamsRuleType]cloudflare.TeamsAccountLoggingConfiguration{
			cloudflare.TeamsDnsRuleType: {
				LogAll:    dnsRuleSettings["log_all"].(bool),
				LogBlocks: dnsRuleSettings["log_blocks"].(bool),
			},
			cloudflare.TeamsHttpRuleType: {
				LogAll:    httpRuleSettings["log_all"].(bool),
				LogBlocks: httpRuleSettings["log_blocks"].(bool),
			},
			cloudflare.TeamsL4RuleType: {
				LogAll:    l4RuleSettings["log_all"].(bool),
				LogBlocks: l4RuleSettings["log_blocks"].(bool),
			},
		},
		RedactPii: logSettings["redact_pii"].(bool),
	}
}

func inflateDeviceSettings(device interface{}) *cloudflare.TeamsDeviceSettings {
	deviceList := device.([]interface{})

	if len(deviceList) != 1 {
		return nil
	}

	deviceSettings, ok := deviceList[0].(map[string]interface{})
	if !ok {
		return nil
	}

	return &cloudflare.TeamsDeviceSettings{
		GatewayProxyEnabled:                deviceSettings["tcp"].(bool),
		GatewayProxyUDPEnabled:             deviceSettings["udp"].(bool),
		RootCertificateInstallationEnabled: deviceSettings["root_ca"].(bool),
	}
}

func flattenPayloadLogSettings(payloadLogSettings *cloudflare.DLPPayloadLogSettings) []interface{} {
	return []interface{}{map[string]interface{}{
		"public_key": payloadLogSettings.PublicKey,
	}}
}

func inflatePayloadLogSettings(payloadLog interface{}) *cloudflare.DLPPayloadLogSettings {
	payloadLogList := payloadLog.([]interface{})
	if len(payloadLogList) != 1 {
		return nil
	}

	payloadLogMap := payloadLogList[0].(map[string]interface{})
	publicKey := payloadLogMap["public_key"].(string)
	return &cloudflare.DLPPayloadLogSettings{
		PublicKey: publicKey,
	}
}
