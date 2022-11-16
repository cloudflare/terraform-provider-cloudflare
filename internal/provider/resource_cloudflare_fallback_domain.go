package provider

import (
	"context"
	"fmt"
	"strings"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareFallbackDomain() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareFallbackDomainSchema(),
		ReadContext:   resourceCloudflareFallbackDomainRead,
		CreateContext: resourceCloudflareFallbackDomainUpdate, // Intentionally identical to Update as the resource is always present
		UpdateContext: resourceCloudflareFallbackDomainUpdate,
		DeleteContext: resourceCloudflareFallbackDomainDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceCloudflareFallbackDomainImport,
		},
	}
}

func resourceCloudflareFallbackDomainRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get("account_id").(string)
	_, policyID := parseDevicePolicyID(d.Get("policy_id").(string))

	var domain []cloudflare.FallbackDomain
	var err error
	if policyID == "" {
		domain, err = client.ListFallbackDomains(ctx, accountID)
	} else {
		domain, err = client.ListFallbackDomainsDeviceSettingsPolicy(ctx, accountID, policyID)
	}
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
	_, policyID := parseDevicePolicyID(d.Get("policy_id").(string))

	domainList := expandFallbackDomains(d.Get("domains").(*schema.Set))

	var newFallbackDomains []cloudflare.FallbackDomain
	var err error
	if policyID == "" {
		d.SetId(accountID)
		newFallbackDomains, err = client.UpdateFallbackDomain(ctx, accountID, domainList)
	} else {
		d.SetId(fmt.Sprintf("%s/%s", accountID, policyID))
		newFallbackDomains, err = client.UpdateFallbackDomainDeviceSettingsPolicy(ctx, accountID, policyID, domainList)
	}
	if err != nil {
		return diag.FromErr(fmt.Errorf("error updating Fallback Domains: %w", err))
	}

	if err := d.Set("domains", flattenFallbackDomains(newFallbackDomains)); err != nil {
		return diag.FromErr(fmt.Errorf("error setting domain attribute: %w", err))
	}

	return resourceCloudflareFallbackDomainRead(ctx, d, meta)
}

func resourceCloudflareFallbackDomainDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get("account_id").(string)
	_, policyID := parseDevicePolicyID(d.Get("policy_id").(string))

	var err error
	if policyID == "" {
		err = client.RestoreFallbackDomainDefaults(ctx, accountID)
	} else {
		err = client.RestoreFallbackDomainDefaultsDeviceSettingsPolicy(ctx, accountID, policyID)
	}
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return nil
}

func resourceCloudflareFallbackDomainImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	accountID, policyID, err := parseDeviceSettingsIDImport(d.Id())
	if err != nil {
		return nil, err
	}

	if accountID == "" {
		return nil, fmt.Errorf("must provide account ID")
	}

	d.Set("account_id", accountID)
	if policyID == "default" {
		d.Set("policy_id", accountID)
		d.SetId(accountID)
	} else {
		d.Set("policy_id", fmt.Sprintf("%s/%s", accountID, policyID))
		d.SetId(fmt.Sprintf("%s/%s", accountID, policyID))
	}

	resourceCloudflareFallbackDomainRead(ctx, d, meta)

	return []*schema.ResourceData{d}, nil
}

// flattenFallbackDomains accepts the cloudflare.FallbackDomain struct and returns the
// schema representation for use in Terraform state.
func flattenFallbackDomains(domains []cloudflare.FallbackDomain) *schema.Set {
	schemaDomains := make([]interface{}, 0)

	for _, d := range domains {
		schemaDomains = append(schemaDomains, map[string]interface{}{
			"suffix":      d.Suffix,
			"description": d.Description,
			"dns_server":  flattenStringList(d.DNSServer),
		})
	}

	return schema.NewSet(HashByMapKey("suffix"), schemaDomains)
}

// expandFallbackDomains accepts the schema representation of Fallback Domains and
// returns a fully qualified struct.
func expandFallbackDomains(domains *schema.Set) []cloudflare.FallbackDomain {
	domainList := make([]cloudflare.FallbackDomain, 0)

	for _, domain := range domains.List() {
		domainList = append(domainList, cloudflare.FallbackDomain{
			Suffix:      domain.(map[string]interface{})["suffix"].(string),
			Description: domain.(map[string]interface{})["description"].(string),
			DNSServer:   expandInterfaceToStringList(domain.(map[string]interface{})["dns_server"]),
		})
	}

	return domainList
}

// parsePolicyID parses the account ID and policy ID from the ID with format
// `<accountTag>` or `<accountTag>/<policyID>` and returns (account id, policy id).
func parseDevicePolicyID(id string) (string, string) {
	attributes := strings.Split(id, "/")

	if len(attributes) == 1 {
		return attributes[0], ""
	}

	return attributes[0], attributes[1]
}
