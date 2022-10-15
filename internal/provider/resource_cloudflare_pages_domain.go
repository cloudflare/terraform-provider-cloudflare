package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflarePagesDomain() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflarePagesDomainSchema(),
		CreateContext: resourceCloudflarePagesDomainCreate,
		ReadContext:   resourceCloudflarePagesDomainRead,
		DeleteContext: resourceCloudflarePagesDomainDelete,
		Description: heredoc.Doc(`
			Provides a resource for managing Cloudflare Pages domains.
		`),
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

func resourceCloudflarePagesDomainImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	client := meta.(*cloudflare.API)

	// split the id so we can look up
	idAttr := strings.SplitN(d.Id(), "/", 2)
	var accountID string
	var projectName string
	if len(idAttr) == 2 {
		accountID = idAttr[0]
		projectName = idAttr[1]
	} else {
		return nil, fmt.Errorf("invalid id %q specified, should be in format \"accountID/project-name\" for import", d.Id())
	}

	domains, err := client.GetPagesDomains(ctx, cloudflare.PagesDomainsParameters{AccountID: accountID, ProjectName: projectName})
	if err != nil {
		return nil, fmt.Errorf("unable to find domains for project %q: %w", d.Id(), err)
	}

	tflog.Info(ctx, fmt.Sprintf("Found domains: %+v", domains))

	resourceData := []*schema.ResourceData{}

	for _, domain := range domains {
		d.SetId(domain.ID)
		d.Set("account_id", accountID)
		d.Set("project_name", projectName)
		d.Set("domain", domain.Name)
		resourceCloudflarePagesDomainRead(ctx, d, meta)

		resourceData = append(resourceData, d)
	}

	return resourceData, nil
}
