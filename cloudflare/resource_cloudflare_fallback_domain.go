package cloudflare

import (
	"context"
	"fmt"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareFallbackDomain() *schema.Resource {
	return &schema.Resource{
		Schema: resourceCloudflareFallbackDomainSchema(),
		ReadContext: resourceCloudflareFallbackDomainRead,
		CreateContext: resourceCloudflareFallbackDomainUpdate, // Intentionally identical to Update as the resource is always present
		UpdateContext: resourceCloudflareFallbackDomainUpdate,
		DeleteContext: resourceCloudflareFallbackDomainDelete,
		Importer: &schema.ResourceImporter{
			State: resourceCloudflareFallbackDomainImport,
		},
	}
}

func resourceCloudflareFallbackDomainRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get("account_id").(string)

	domain, err := client.ListFallbackDomains(context.Background(), accountID)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error finding Fallback Domains: %w", err))
	}

	if err := d.Set("domains", flattenFallbackDomains(domain)); err != nil {
		return diag.FromErr(fmt.Errorf("error setting domains attribute: %w", err))
	}

	return nil
}

func resourceCloudflareFallbackDomainUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get("account_id").(string)

	domainList := expandFallbackDomains(d.Get("domains").([]interface{}))

	newFallbackDomains, err := client.UpdateFallbackDomain(context.Background(), accountID, domainList)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error updating Fallback Domains: %w", err))
	}

	if err := d.Set("domains", flattenFallbackDomains(newFallbackDomains)); err != nil {
		return diag.FromErr(fmt.Errorf("error setting domain attribute: %w", err))
	}

	d.SetId(accountID)

	return resourceCloudflareFallbackDomainRead(d, meta)
}

func resourceCloudflareFallbackDomainDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get("account_id").(string)

	err := client.RestoreFallbackDomainDefaults(context.Background(), accountID)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return nil
}

func resourceCloudflareFallbackDomainImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	accountID := d.Id()

	if accountID == "" {
		return nil, fmt.Errorf("must provide account ID")
	}

	d.Set("account_id", accountID)
	d.SetId(accountID)

	readErr := resourceCloudflareFallbackDomainRead(d, meta)

	return []*schema.ResourceData{d}, readErr
}

// flattenFallbackDomains accepts the cloudflare.FallbackDomain struct and returns the
// schema representation for use in Terraform state.
func flattenFallbackDomains(domains []cloudflare.FallbackDomain) []interface{} {
	schemaDomains := make([]interface{}, 0)

	for _, d := range domains {
		schemaDomains = append(schemaDomains, map[string]interface{}{
			"suffix":      d.Suffix,
			"description": d.Description,
			"dns_server":  flattenStringList(d.DNSServer),
		})
	}

	return schemaDomains
}

// expandFallbackDomains accepts the schema representation of Fallback Domains and
// returns a fully qualified struct.
func expandFallbackDomains(domains []interface{}) []cloudflare.FallbackDomain {
	domainList := make([]cloudflare.FallbackDomain, 0)

	for _, domain := range domains {
		domainList = append(domainList, cloudflare.FallbackDomain{
			Suffix:      domain.(map[string]interface{})["suffix"].(string),
			Description: domain.(map[string]interface{})["description"].(string),
			DNSServer:   expandInterfaceToStringList(domain.(map[string]interface{})["dns_server"]),
		})
	}

	return domainList
}
