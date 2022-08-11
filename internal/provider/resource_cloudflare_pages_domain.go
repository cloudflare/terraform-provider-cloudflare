package provider

import (
	"context"
	"fmt"
	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflarePagesDomain() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflarePagesDomainSchema(),
		CreateContext: resourceCloudflarePagesDomainCreate,
		ReadContext:   resourceCloudflarePagesDomainRead,
		DeleteContext: resourceCloudflarePagesDomainDelete,
	}
}

func resourceCloudflarePagesDomainCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get("account_id").(string)
	projectName := d.Get("project_name").(string)
	domain := d.Get("domain").(string)

	params := cloudflare.PagesDomainParameters{
		AccountID:   accountID,
		ProjectName: projectName,
		DomainName:  domain,
	}

	r, err := client.PagesAddDomain(ctx, params)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating domain for project %q: %w", accountID, err))
	}
	d.SetId(r.ID)
	return resourceCloudflarePagesDomainRead(ctx, d, meta)
}

func resourceCloudflarePagesDomainRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get("account_id").(string)
	projectName := d.Get("project_name").(string)
	domain := d.Get("domain").(string)

	params := cloudflare.PagesDomainParameters{
		AccountID:   accountID,
		ProjectName: projectName,
		DomainName:  domain,
	}
	r, err := client.GetPagesDomain(ctx, params)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating domain for project %q: %w", accountID, err))
	}
	d.Set("status", r.Status)
	return nil
}

func resourceCloudflarePagesDomainDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get("account_id").(string)
	projectName := d.Get("project_name").(string)
	domain := d.Get("domain").(string)

	params := cloudflare.PagesDomainParameters{
		AccountID:   accountID,
		ProjectName: projectName,
		DomainName:  domain,
	}
	err := client.PagesDeleteDomain(ctx, params)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating domain for project %q: %w", accountID, err))
	}
	return nil
}
