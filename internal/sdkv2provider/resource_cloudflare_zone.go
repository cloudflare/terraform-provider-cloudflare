package sdkv2provider

import (
	"context"
	"errors"
	"fmt"
	"log"

	"golang.org/x/net/idna"

	"github.com/MakeNowJust/heredoc/v2"
	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// we enforce the use of the Cloudflare API 'legacy_id' field until the mapping of plan is fixed in cloudflare-go.
const (
	planIDFree       = "free"
	planIDLite       = "lite"
	planIDPro        = "pro"
	planIDProPlus    = "pro_plus"
	planIDBusiness   = "business"
	planIDEnterprise = "enterprise"

	planIDPartnerFree       = "partners_free"
	planIDPartnerPro        = "partners_pro"
	planIDPartnerBusiness   = "partners_business"
	planIDPartnerEnterprise = "partners_enterprise"
)

type subscriptionData struct {
	ID, Name, Description string
}

var ratePlans = map[string]subscriptionData{
	planIDFree: {
		Name:        "CF_FREE",
		ID:          planIDFree,
		Description: "Free Website",
	},
	planIDLite: {
		Name:        "CF_LITE",
		ID:          planIDLite,
		Description: "Lite Website",
	},
	planIDPro: {
		Name:        "CF_PRO_20_20",
		ID:          planIDPro,
		Description: "Pro Website",
	},
	planIDProPlus: {
		Name:        "CF_PRO_PLUS",
		ID:          planIDProPlus,
		Description: "Pro Plus Website",
	},
	planIDBusiness: {
		Name:        "CF_BIZ",
		ID:          planIDBusiness,
		Description: "Business Website",
	},
	planIDEnterprise: {
		Name:        "CF_ENT",
		ID:          planIDEnterprise,
		Description: "Enterprise Website",
	},
	planIDPartnerFree: {
		Name:        "PARTNERS_FREE",
		ID:          planIDPartnerFree,
		Description: "Free Website",
	},
	planIDPartnerPro: {
		Name:        "PARTNERS_PRO",
		ID:          planIDPartnerPro,
		Description: "Pro Website",
	},
	planIDPartnerBusiness: {
		Name:        "PARTNERS_BIZ",
		ID:          planIDPartnerBusiness,
		Description: "Business Website",
	},
	planIDPartnerEnterprise: {
		Name:        "PARTNERS_ENT",
		ID:          planIDPartnerEnterprise,
		Description: "Enterprise Website",
	},
}

func resourceCloudflareZone() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareZoneSchema(),
		CreateContext: resourceCloudflareZoneCreate,
		ReadContext:   resourceCloudflareZoneRead,
		UpdateContext: resourceCloudflareZoneUpdate,
		DeleteContext: resourceCloudflareZoneDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Description: heredoc.Doc(`
			Provides a Cloudflare Zone resource. Zone is the basic resource for
			working with Cloudflare and is roughly equivalent to a domain name
			that the user purchases.
		`),
	}
}

func resourceCloudflareZoneCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)
	zoneName := d.Get("zone").(string)
	jumpstart := d.Get("jump_start").(bool)
	zoneType := d.Get("type").(string)
	account := cloudflare.Account{
		ID: accountID,
	}

	tflog.Info(ctx, fmt.Sprintf("Creating Cloudflare Zone: name %s", zoneName))

	zone, err := client.CreateZone(ctx, zoneName, jumpstart, account, zoneType)

	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating zone %q: %w", zoneName, err))
	}

	d.SetId(zone.ID)

	if paused, ok := d.GetOk("paused"); ok {
		if paused.(bool) == true {
			_, err := client.ZoneSetPaused(ctx, zone.ID, paused.(bool))

			if err != nil {
				return diag.FromErr(fmt.Errorf("error updating zone_id %q: %w", zone.ID, err))
			}
		}
	}

	if plan, ok := d.GetOk("plan"); ok {
		if err := setRatePlan(ctx, client, zone.ID, plan.(string), true, d); err != nil {
			return diag.FromErr(err)
		}
	}

	if ztype, ok := d.GetOk("type"); ok {
		_, err := client.ZoneSetType(ctx, zone.ID, ztype.(string))
		if err != nil {
			return diag.FromErr(fmt.Errorf("error setting type on zone ID %q: %w", zone.ID, err))
		}
	}

	return resourceCloudflareZoneRead(ctx, d, meta)
}

func resourceCloudflareZoneRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Id()

	zone, err := client.ZoneDetails(ctx, zoneID)

	tflog.Debug(ctx, fmt.Sprintf("ZoneDetails: %#v", zone))
	tflog.Debug(ctx, fmt.Sprintf("ZoneDetails error: %#v", err))

	if err != nil {
		var notFoundError *cloudflare.NotFoundError
		if errors.As(err, &notFoundError) {
			tflog.Info(ctx, fmt.Sprintf("Zone %s no longer exists", d.Id()))
			d.SetId("")
			return nil
		}
		return diag.FromErr(fmt.Errorf("error finding Zone %q: %w", d.Id(), err))
	}

	// In the cases where the zone isn't completely setup yet, we need to
	// check the `status` field and should it be pending, use the `LegacyID`
	// from `zone.PlanPending` instead to account for paid plans.
	var plan string
	if zone.Status == "pending" && zone.PlanPending.LegacyID != "" {
		plan = zone.PlanPending.LegacyID
	} else {
		plan = zone.Plan.LegacyID
	}

	d.Set(consts.AccountIDSchemaKey, zone.Account.ID)
	d.Set("paused", zone.Paused)
	d.Set("vanity_name_servers", zone.VanityNS)
	d.Set("status", zone.Status)
	d.Set("type", zone.Type)
	d.Set("name_servers", zone.NameServers)
	d.Set("meta", flattenMeta(d, zone.Meta))
	d.Set("zone", zone.Name)
	d.Set("plan", plan)
	d.Set("verification_key", zone.VerificationKey)

	return nil
}

func resourceCloudflareZoneUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Id()
	zone, _ := client.ZoneDetails(ctx, zoneID)

	log.Printf("[INFO] Updating Cloudflare Zone: id %s", zoneID)

	if paused, ok := d.GetOkExists("paused"); ok && d.HasChange("paused") {
		log.Printf("[DEBUG] _ paused")

		_, err := client.ZoneSetPaused(ctx, zoneID, paused.(bool))

		if err != nil {
			return diag.FromErr(fmt.Errorf("error setting paused for zone ID %q: %w", zoneID, err))
		}
	}

	if ztype, ok := d.GetOkExists("type"); ok && d.HasChange("type") {
		_, err := client.ZoneSetType(ctx, zoneID, ztype.(string))

		if err != nil {
			return diag.FromErr(fmt.Errorf("error setting type for on zone ID %q: %w", zoneID, err))
		}
	}

	// In the cases where the zone isn't completely setup yet, we need to
	// check the `status` field and should it be pending, use the `LegacyID`
	// from `zone.PlanPending` instead to account for paid plans.
	if zone.Status == "pending" && zone.PlanPending.Name != "" {
		d.Set("plan", zone.PlanPending.LegacyID)
	}

	if change := d.HasChange("plan"); change {
		// If we're upgrading from a free plan, we need to use POST (not PUT) as the
		// the subscription needs to be created, not modified despite the resource
		// already existing.
		existingPlan, newPlan := d.GetChange("plan")
		wasFreePlan := existingPlan.(string) == "free"
		planID := newPlan.(string)

		if err := setRatePlan(ctx, client, zoneID, planID, wasFreePlan, d); err != nil {
			return diag.FromErr(err)
		}
	}

	return resourceCloudflareZoneRead(ctx, d, meta)
}

func resourceCloudflareZoneDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Id()

	log.Printf("[INFO] Deleting Cloudflare Zone: id %s", zoneID)

	_, err := client.DeleteZone(ctx, zoneID)

	if err != nil {
		return diag.FromErr(fmt.Errorf("error deleting Cloudflare Zone: %w", err))
	}

	return nil
}

func flattenMeta(d *schema.ResourceData, meta cloudflare.ZoneMeta) map[string]interface{} {
	cfg := map[string]interface{}{}

	cfg["wildcard_proxiable"] = meta.WildcardProxiable
	cfg["phishing_detected"] = meta.PhishingDetected

	log.Printf("[DEBUG] flattenMeta %#v", cfg)

	return cfg
}

// setRatePlan handles the internals of creating or updating a zone
// subscription rate plan.
func setRatePlan(ctx context.Context, client *cloudflare.API, zoneID, planID string, isNewPlan bool, d *schema.ResourceData) error {
	if isNewPlan {
		// A free rate plan is the default so no need to explicitly make another
		// HTTP call to set it.
		if ratePlans[planID].ID != planIDFree {
			if err := client.ZoneSetPlan(ctx, zoneID, ratePlans[planID].Name); err != nil {
				return fmt.Errorf("error setting plan %s for zone %q: %w", planID, zoneID, err)
			}
		}
	} else {
		if err := client.ZoneUpdatePlan(ctx, zoneID, ratePlans[planID].Name); err != nil {
			return fmt.Errorf("error updating plan %s for zone %q: %w", planID, zoneID, err)
		}
	}

	return resource.RetryContext(ctx, d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		zone, _ := client.ZoneDetails(ctx, zoneID)

		// This is a little confusing but due to the multiple views of
		// subscriptions, partner plans actually end up "appearing" like regular
		// rate plans to end users. To ensure we don't get stuck in an endless
		// loop, we need to compare the "name" of the plan to the "description"
		// of the rate plan. That way, even if we send `PARTNERS_ENT` to the
		// subscriptions service, we will compare it in Terraform as
		// "Enterprise Website" and know that we made the swap and just trust
		// that the rate plan identifier did the right thing.
		if zone.Plan.Name != ratePlans[planID].Description {
			return resource.RetryableError(fmt.Errorf("plan ID change has not yet propagated"))
		}

		return nil
	})
}

// zoneDiffFunc is a DiffSuppressFunc that accepts two strings and then converts
// them to unicode before performing the comparison whether or not the value has
// changed. This ensures that zones which could be either are evaluated
// consistently and align with what the Cloudflare API returns.
func zoneDiffFunc(k, old, new string, d *schema.ResourceData) bool {
	var p *idna.Profile
	p = idna.New()
	unicodeOld, _ := p.ToUnicode(old)
	unicodeNew, _ := p.ToUnicode(new)

	return unicodeOld == unicodeNew
}
