package cloudflare

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
)

func resourceCloudflareTeamsAccount() *schema.Resource {
	return &schema.Resource{
		Read:   resourceCloudflareTeamsAccountRead,
		Update: resourceCloudflareTeamsAccountUpdate,
		Create: resourceCloudflareTeamsAccountUpdate,
		// This resource is a top-level account configuration and cant be "deleted"
		Delete: func(_ *schema.ResourceData, _ interface{}) error { return nil },
		Importer: &schema.ResourceImporter{
			State: resourceCloudflareTeamsAccountImport,
		},

		Schema: map[string]*schema.Schema{
			"account_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"block_page": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: blockPageSchema,
				},
			},
			"antivirus": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: antivirusSchema,
				},
			},
			"tls_decrypt_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"activity_log_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
	}
}

var blockPageSchema = map[string]*schema.Schema{
	"enabled": {
		Type:     schema.TypeBool,
		Optional: true,
	},
	"footer_text": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"header_text": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"logo_path": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"background_color": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"name": {
		Type:     schema.TypeString,
		Optional: true,
	},
}

var antivirusSchema = map[string]*schema.Schema{
	"enabled_download_phase": {
		Type:     schema.TypeBool,
		Required: true,
	},
	"enabled_upload_phase": {
		Type:     schema.TypeBool,
		Required: true,
	},
	"fail_closed": {
		Type:     schema.TypeBool,
		Required: true,
	},
}

func resourceCloudflareTeamsAccountRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	accountID := d.Get("account_id").(string)

	configuration, err := client.TeamsAccountConfiguration(context.Background(), accountID)
	if err != nil {
		if strings.Contains(err.Error(), "HTTP status 400") {
			log.Printf("[INFO] Teams Account config %s does not exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("error finding Teams Account config %q: %s", d.Id(), err)
	}

	if configuration.Settings.BlockPage != nil {
		if err := d.Set("block_page", flattenBlockPageConfig(configuration.Settings.BlockPage)); err != nil {
			return errors.Wrap(err, "error parsing account block page config")
		}
	}

	if configuration.Settings.Antivirus != nil {
		if err := d.Set("antivirus", flattenAntivirusConfig(configuration.Settings.Antivirus)); err != nil {
			return errors.Wrap(err, "error parsing account antivirus config")
		}
	}

	if configuration.Settings.TLSDecrypt != nil {
		if err := d.Set("tls_decrypt_enabled", configuration.Settings.TLSDecrypt.Enabled); err != nil {
			return errors.Wrap(err, "error parsing account tls decrypt enablement")
		}
	}

	if configuration.Settings.ActivityLog != nil {
		if err := d.Set("activity_log_enabled", configuration.Settings.ActivityLog.Enabled); err != nil {
			return errors.Wrap(err, "error parsing account activity log enablement")
		}
	}
	return nil
}

func resourceCloudflareTeamsAccountUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	accountID := d.Get("account_id").(string)
	blockPageConfig := inflateBlockPageConfig(d.Get("block_page"))
	antivirusConfig := inflateAntivirusConfig(d.Get("antivirus"))
	updatedTeamsAccount := cloudflare.TeamsConfiguration{
		Settings: cloudflare.TeamsAccountSettings{
			Antivirus: antivirusConfig,
			BlockPage: blockPageConfig,
		},
	}

	tlsDecrypt, ok := d.GetOkExists("tls_decrypt_enabled")
	if ok {
		updatedTeamsAccount.Settings.TLSDecrypt = &cloudflare.TeamsTLSDecrypt{Enabled: tlsDecrypt.(bool)}
	}

	activtyLog, ok := d.GetOkExists("activity_log_enabled")
	if ok {
		updatedTeamsAccount.Settings.ActivityLog = &cloudflare.TeamsActivityLog{Enabled: activtyLog.(bool)}
	}
	log.Printf("[DEBUG] Updating Cloudflare Teams Account configuration from struct: %+v", updatedTeamsAccount)

	if _, err := client.TeamsAccountUpdateConfiguration(context.Background(), accountID, updatedTeamsAccount); err != nil {
		return fmt.Errorf("error updating Teams Account configuration for account %q: %s", accountID, err)
	}

	d.SetId(accountID)
	return resourceCloudflareTeamsAccountRead(d, meta)
}

func resourceCloudflareTeamsAccountImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	d.SetId(d.Id())
	d.Set("account_id", d.Id())

	err := resourceCloudflareTeamsAccountRead(d, meta)
	return []*schema.ResourceData{d}, err
}

func flattenBlockPageConfig(blockPage *cloudflare.TeamsBlockPage) []interface{} {
	return []interface{}{map[string]interface{}{
		"enabled":          *blockPage.Enabled,
		"footer_text":      blockPage.FooterText,
		"header_text":      blockPage.HeaderText,
		"logo_path":        blockPage.LogoPath,
		"background_color": blockPage.BackgroundColor,
		"name":             blockPage.Name,
	}}
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
	}
}

func flattenAntivirusConfig(antivirusConfig *cloudflare.TeamsAntivirus) []interface{} {
	return []interface{}{map[string]interface{}{
		"enabled_download_phase": antivirusConfig.EnabledDownloadPhase,
		"enabled_upload_phase":   antivirusConfig.EnabledUploadPhase,
		"fail_closed":            antivirusConfig.FailClosed,
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
