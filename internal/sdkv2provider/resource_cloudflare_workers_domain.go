package sdkv2provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/MakeNowJust/heredoc/v2"
	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
)

func resourceCloudflareWorkerDomain() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareWorkerDomainSchema(),
		CreateContext: resourceCloudflareWorkerDomainCreate,
		ReadContext:   resourceCloudflareWorkerDomainRead,
		UpdateContext: resourceCloudflareWorkerDomainUpdate,
		DeleteContext: resourceCloudflareWorkerDomainDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceCloudflareWorkerDomainImport,
		},
		Description: heredoc.Doc(
			"Creates a Worker Custom Domain.",
		),
	}
}

func resourceCloudflareWorkerDomainCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)

	workerDomain, err := client.AttachWorkersDomain(ctx, cloudflare.AccountIdentifier(accountID), cloudflare.AttachWorkersDomainParams{
		ZoneID:      d.Get("zone_id").(string),
		Hostname:    d.Get("hostname").(string),
		Service:     d.Get("service").(string),
		Environment: d.Get("environment").(string),
	})

	if err != nil {
		return diag.FromErr(errors.Wrap(err, "error attaching worker domain"))
	}

	d.SetId(workerDomain.ID)
	return resourceCloudflareWorkerDomainRead(ctx, d, meta)
}

func resourceCloudflareWorkerDomainRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)

	workerDomain, err := client.GetWorkersDomain(ctx, cloudflare.AccountIdentifier((accountID)), d.Id())

	if err != nil {
		return diag.FromErr(fmt.Errorf("error reading worker domain %q", d.Id()))
	}

	d.Set("zone_id", workerDomain.ZoneID)
	d.Set("hostname", workerDomain.Hostname)
	d.Set("service", workerDomain.Service)
	d.Set("environment", workerDomain.Environment)

	return nil
}

func resourceCloudflareWorkerDomainUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)

	workerDomain, err := client.AttachWorkersDomain(ctx, cloudflare.AccountIdentifier(accountID), cloudflare.AttachWorkersDomainParams{
		ZoneID:      d.Get("zone_id").(string),
		Hostname:    d.Get("hostname").(string),
		Service:     d.Get("service").(string),
		Environment: d.Get("environment").(string),
	})

	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(workerDomain.ID)
	return nil
}

func resourceCloudflareWorkerDomainDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)

	err := client.DetachWorkersDomain(ctx, cloudflare.AccountIdentifier(accountID), d.Id())

	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceCloudflareWorkerDomainImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	idAttr := strings.SplitN(d.Id(), "/", 2)
	var accountID string
	var workersDomainID string
	if len(idAttr) == 2 {
		accountID = idAttr[0]
		workersDomainID = idAttr[1]
		d.Set(consts.AccountIDSchemaKey, accountID)
		d.SetId(workersDomainID)
	} else {
		return nil, fmt.Errorf("invalid id (%q) specified, should be in format \"account/workersDomainID\"", d.Id())
	}

	resourceCloudflareWorkerDomainRead(ctx, d, meta)

	return []*schema.ResourceData{d}, nil
}
