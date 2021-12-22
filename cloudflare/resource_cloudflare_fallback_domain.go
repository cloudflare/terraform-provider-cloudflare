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
		Read:   resourceCloudflareFallbackDomainRead,
		Create: resourceCloudflareFallbackDomainUpdate, // Intentionally identical to Update as the resource is always present
		Update: resourceCloudflareFallbackDomainUpdate,
		Delete: resourceCloudflareFallbackDomainDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceCloudflareFallbackDomainRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	accountID := d.Get("account_id").(string)

	domain, err := client.ListFallbackDomains(context.Background(), accountID)
	if err != nil {
		return fmt.Errorf("error finding Fallback Domains: %s", err)
	}

	if err := d.Set("domains", flattenFallbackDomains(domain)); err != nil {
		return fmt.Errorf("error setting domains attribute: %s", err)
	}

	return nil
}

func resourceCloudflareFallbackDomainUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	accountID := d.Get("account_id").(string)

	domainList, err := expandFallbackDomains(d.Get("domains").([]interface{}))
	if err != nil {
		return fmt.Errorf("error updating Fallback Domains: %s", err)
	}

	newFallbackDomains, err := client.UpdateFallbackDomain(context.Background(), accountID, domainList)
	if err != nil {
		return fmt.Errorf("error updating Fallback Domains: %s", err)
	}

	if err := d.Set("domains", flattenFallbackDomains(newFallbackDomains)); err != nil {
		return fmt.Errorf("error setting domain attribute: %s", err)
	}

	d.SetId(accountID)

	return resourceCloudflareFallbackDomainRead(d, meta)
}

func resourceCloudflareFallbackDomainDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	accountID := d.Get("account_id").(string)

	client.UpdateFallbackDomain(context.Background(), accountID, nil)

	d.SetId("")
	return nil
}

func resourceCloudflareFallbackDomainImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	accountID := d.Id()

	if (accountID == "") {
		return nil, fmt.Errorf("must provide account ID")
	}

	d.Set("account_id", accountID)
	d.SetId(accountID)

	resourceCloudflareFallbackDomainRead(d, meta)

	return []*schema.ResourceData{d}, nil
}

// flattenFallbackDomains accepts the cloudflare.FallbackDomain struct and returns the
// schema representation for use in Terraform state.
func flattenFallbackDomains(domains []cloudflare.FallbackDomain) []interface{} {
	schemaDomains := make([]interface{}, 0)

	for _, d := range domains {
		schemaDomains = append(schemaDomains, map[string]interface{}{
			"suffix":      d.Suffix,
			"description": d.Description,
			"dns_server":  d.DNSServer,

		})
	}

	return schemaDomains
}

// expandFallbackDomains accepts the schema representation of Fallback Domains and
// returns a fully qualified struct.
func expandFallbackDomains(domains []interface{}) ([]cloudflare.FallbackDomain, error) {
	domainList := make([]cloudflare.FallbackDomain, 0)

	for _, domain := range domains {
		domainList = append(domainList, cloudflare.FallbackDomain{
			Suffix:      domain.(map[string]interface{})["suffix"].(string),
			Description: domain.(map[string]interface{})["description"].(string),
			DNSServer:   domain.(map[string]interface{})["dns_server"].([]string),
		})
	}

	return domainList, nil
}