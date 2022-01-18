package cloudflare

import (
	"context"
	"log"
	"fmt"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// use this since we cannot use const for lists
func default_domains() []string {
	return []string{
		"intranet", "internal", "private", "localdomain", "domain", "lan", "home",
		"host", "corp", "local", "localhost", "home.arpa", "invalid", "test",
	}
}

func resourceCloudflareFallbackDomain() *schema.Resource {
	return &schema.Resource{
		Schema: resourceCloudflareFallbackDomainSchema(),
		Read:   resourceCloudflareFallbackDomainRead,
		Create: resourceCloudflareFallbackDomainUpdate, // Intentionally identical to Update as the resource is always present
		Update: resourceCloudflareFallbackDomainUpdate,
		Delete: resourceCloudflareFallbackDomainDelete,
		Importer: &schema.ResourceImporter{
			State: resourceCloudflareFallbackDomainImport,
		},
		CustomizeDiff: resourceCloudflareFallbackDomainDiff,
	}
}

func resourceCloudflareFallbackDomainRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	accountID := d.Get("account_id").(string)

	domain, err := client.ListFallbackDomains(context.Background(), accountID)
	if err != nil {
		return fmt.Errorf("error finding Fallback Domains: %w", err)
	}

	if err := d.Set("domains", flattenFallbackDomains(domain)); err != nil {
		return fmt.Errorf("error setting domains attribute: %w", err)
	}

	return nil
}

func resourceCloudflareFallbackDomainUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	accountID := d.Get("account_id").(string)

	domainList := expandFallbackDomains(d.Get("domains").([]interface{}))
	log.Printf("[INFO] Updating Cloudflare Fallback Domain: %s", domainList)

	newFallbackDomains, err := client.UpdateFallbackDomain(context.Background(), accountID, domainList)
	if err != nil {
		return fmt.Errorf("error updating Fallback Domains: %w", err)
	}

	if err := d.Set("domains", flattenFallbackDomains(newFallbackDomains)); err != nil {
		return fmt.Errorf("error setting domain attribute: %w", err)
	}

	d.SetId(accountID)

	return resourceCloudflareFallbackDomainRead(d, meta)
}

func resourceCloudflareFallbackDomainDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	accountID := d.Get("account_id").(string)

	domainList := make([]cloudflare.FallbackDomain, 0)
	if d.Get("restore_default_domains_on_delete").(bool) == true {
		domainList = getDefaultDomains()
	}
	_, err := client.UpdateFallbackDomain(context.Background(), accountID, domainList)
	if err != nil {
		return err
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

func getDefaultDomains() []cloudflare.FallbackDomain {
	domainList := make([]cloudflare.FallbackDomain, 0)

	for _, domain := range default_domains() {
		domainList = append(domainList, cloudflare.FallbackDomain{
			Suffix: domain,
		})
	}

	return domainList
}

func resourceCloudflareFallbackDomainDiff(_ context.Context, diff *schema.ResourceDiff, meta interface{}) error {
	_, include_n := diff.GetChange("include_default_domains")
	_, domains_n := diff.GetChange("domains")
	domains_y := domains_n.([]interface{})

	var domainList []interface{}
	if include_n.(bool) {
		domainList = flattenFallbackDomains(getDefaultDomains())
	}
	for _, domain := range domains_y {
		domainList = append(domainList, domain.(map[string]interface{}))
	}
	if err := diff.SetNew("domains", domainList); err != nil {
		return fmt.Errorf("error including default domains: %w", err)
	}

	return nil
}
