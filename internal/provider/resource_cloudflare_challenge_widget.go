package provider

import (
	"context"
	"fmt"
	"github.com/MakeNowJust/heredoc/v2"
	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strings"
)

func resourceCloudflareChallengeWidget() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceChallengeWidgetSchema(),
		ReadContext:   resourceCloudflareChallengeWidgetRead,
		CreateContext: resourceCloudflareChallengeWidgetCreate,
		UpdateContext: resourceCloudflareChallengeWidgetUpdate,
		DeleteContext: resourceCloudflareChallengeWidgetDelete,
		Description: heredoc.Doc(`
			Challenge widgets are used to verify that a user is a human and not a bot.
		`),
		Importer: &schema.ResourceImporter{
			StateContext: resourceCloudflareChallengeWidgetImport,
		},
	}
}

func resourceCloudflareChallengeWidgetRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get("account_id").(string)

	widget, err := client.GetChallengeWidget(ctx, cloudflare.AccountIdentifier(accountID), d.Id())
	if err != nil {
		return diag.FromErr(fmt.Errorf("error reading cloudflare challenge widget %q: %w", d.Id(), err))
	}

	d.Set("site_key", widget.SiteKey)
	d.Set("secret", widget.Secret)
	d.Set("name", widget.Name)
	d.Set("domains", widget.Domains)
	d.Set("type", widget.Type)

	return nil
}

func resourceCloudflareChallengeWidgetCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get("account_id").(string)
	domains := d.Get("domains").(*schema.Set)

	widget := cloudflare.ChallengeWidget{
		Name:    d.Get("name").(string),
		Type:    d.Get("type").(string),
		Domains: expandInterfaceToStringList(domains.List()),
	}
	widget, err := client.CreateChallengeWidget(ctx, cloudflare.AccountIdentifier(accountID), widget)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating cloudflare challenge widget %q: %w", d.Id(), err))
	}
	d.SetId(widget.SiteKey)
	return resourceCloudflareChallengeWidgetRead(ctx, d, meta)
}

func resourceCloudflareChallengeWidgetUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get("account_id").(string)

	widget := cloudflare.ChallengeWidget{
		SiteKey: d.Id(),
		Name:    d.Get("name").(string),
		Type:    d.Get("type").(string),
		Domains: d.Get("domains").([]string),
	}
	widget, err := client.UpdateChallengeWidget(ctx, cloudflare.AccountIdentifier(accountID), widget)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating cloudflare challenge widget %q: %w", d.Id(), err))
	}

	d.SetId(widget.SiteKey)
	return resourceCloudflareChallengeWidgetRead(ctx, d, meta)
}

func resourceCloudflareChallengeWidgetDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get("account_id").(string)

	err := client.DeleteChallengeWidget(ctx, cloudflare.AccountIdentifier(accountID), d.Id())
	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating cloudflare challenge widget %q: %w", d.Id(), err))
	}

	return nil
}

func resourceCloudflareChallengeWidgetImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	client := meta.(*cloudflare.API)

	// split the id so we can look up
	idAttr := strings.SplitN(d.Id(), "/", 2)
	var accountID string
	var siteKey string
	if len(idAttr) == 2 {
		accountID = idAttr[0]
		siteKey = idAttr[1]
	} else {
		return nil, fmt.Errorf("invalid id %q specified, should be in format \"accountID/site-key\" for import", d.Id())
	}

	widget, err := client.GetChallengeWidget(ctx, cloudflare.AccountIdentifier(accountID), siteKey)
	if err != nil {
		return nil, fmt.Errorf("unable to find challenge widget with site key %q: %w", d.Id(), err)
	}

	d.SetId(widget.SiteKey)
	d.Set("account_id", accountID)

	resourceCloudflareChallengeWidgetRead(ctx, d, meta)

	return []*schema.ResourceData{d}, nil
}
