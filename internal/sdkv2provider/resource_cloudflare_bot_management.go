package sdkv2provider

import (
	"context"
	"fmt"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
)

func resourceCloudflareBotManagement() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareBotManagementSchema(),
		CreateContext: resourceCloudflareBotManagementCreate,
		ReadContext:   resourceCloudflareBotManagementRead,
		UpdateContext: resourceCloudflareBotManagementUpdate,
		DeleteContext: resourceCloudflareBotManagementDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceCloudflareBotManagementImport,
		},
		Description: heredoc.Doc(`Provides a resource to configure Bot Management.

		Specifically, this resource can be used to manage:

		- **Bot Fight Mode**
		- **Super Bot Fight Mode**
		- **Bot Management for Enterprise**
		`),
	}
}

func resourceCloudflareBotManagementCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	d.SetId(d.Get(consts.ZoneIDSchemaKey).(string))

	return resourceCloudflareBotManagementUpdate(ctx, d, meta)
}

func resourceCloudflareBotManagementRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	bm, err := client.GetBotManagement(ctx, cloudflare.ZoneIdentifier(d.Id()))
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to fetch bot management configuration: %w", err))
	}

	if bm.EnableJS != nil {
		d.Set("enable_js", bm.EnableJS)
	}

	if bm.FightMode != nil {
		d.Set("fight_mode", bm.FightMode)
	}

	if bm.SBFMDefinitelyAutomated != nil {
		d.Set("sbfm_definitely_automated", bm.SBFMDefinitelyAutomated)
	}

	if bm.SBFMLikelyAutomated != nil {
		d.Set("sbfm_likely_automated", bm.SBFMLikelyAutomated)
	}

	if bm.SBFMVerifiedBots != nil {
		d.Set("sbfm_verified_bots", bm.SBFMVerifiedBots)
	}

	if bm.SBFMStaticResourceProtection != nil {
		d.Set("sbfm_static_resource_protection", bm.SBFMStaticResourceProtection)
	}

	if bm.OptimizeWordpress != nil {
		d.Set("optimize_wordpress", bm.OptimizeWordpress)
	}

	if bm.SuppressSessionScore != nil {
		d.Set("suppress_session_score", bm.SuppressSessionScore)
	}

	if bm.AutoUpdateModel != nil {
		d.Set("auto_update_model", bm.AutoUpdateModel)
	}

	if bm.UsingLatestModel != nil {
		d.Set("using_latest_model", bm.UsingLatestModel)
	}

	return nil
}

func resourceCloudflareBotManagementUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	params := buildBotManagementParams(d)

	_, err := client.UpdateBotManagement(ctx, cloudflare.ZoneIdentifier(zoneID), params)
	if err != nil {
		return diag.FromErr(errors.Wrap(err, "failed to update bot management configuration"))
	}

	return resourceCloudflareBotManagementRead(ctx, d, meta)
}

// Deletion of bot management configuration is not something we support, we will
// use a dummy handler for now.
func resourceCloudflareBotManagementDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func resourceCloudflareBotManagementImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	d.SetId(d.Id())
	d.Set(consts.ZoneIDSchemaKey, d.Id())

	resourceCloudflareBotManagementRead(ctx, d, meta)

	return []*schema.ResourceData{d}, nil
}

func buildBotManagementParams(d *schema.ResourceData) cloudflare.UpdateBotManagementParams {
	bm := cloudflare.UpdateBotManagementParams{}

	if val, exists := d.GetOkExists("enable_js"); exists {
		bm.EnableJS = cloudflare.BoolPtr(val.(bool))
	}
	if val, exists := d.GetOkExists("fight_mode"); exists {
		bm.EnableJS = cloudflare.BoolPtr(val.(bool))
	}

	if val, exists := d.GetOkExists("sbfm_definitely_automated"); exists {
		bm.SBFMDefinitelyAutomated = cloudflare.StringPtr(val.(string))
	}
	if val, exists := d.GetOkExists("sbfm_likely_automated"); exists {
		bm.SBFMLikelyAutomated = cloudflare.StringPtr(val.(string))
	}
	if val, exists := d.GetOkExists("sbfm_verified_bots"); exists {
		bm.SBFMVerifiedBots = cloudflare.StringPtr(val.(string))
	}
	if val, exists := d.GetOkExists("sbfm_static_resource_protection"); exists {
		bm.SBFMStaticResourceProtection = cloudflare.BoolPtr(val.(bool))
	}
	if val, exists := d.GetOkExists("optimize_wordpress"); exists {
		bm.OptimizeWordpress = cloudflare.BoolPtr(val.(bool))
	}

	if val, exists := d.GetOkExists("suppress_session_score"); exists {
		bm.SuppressSessionScore = cloudflare.BoolPtr(val.(bool))
	}
	if val, exists := d.GetOkExists("auto_update_model"); exists {
		bm.AutoUpdateModel = cloudflare.BoolPtr(val.(bool))
	}

	return bm
}
