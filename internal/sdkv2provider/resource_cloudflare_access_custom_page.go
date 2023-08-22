package sdkv2provider

import (
	"context"
	"errors"
	"fmt"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareAccessCustomPage() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareAccessCustomPageSchema(),
		CreateContext: resourceCloudflareAccessCustomPageCreate,
		ReadContext:   resourceCloudflareAccessCustomPageRead,
		UpdateContext: resourceCloudflareAccessCustomPageUpdate,
		DeleteContext: resourceCloudflareAccessCustomPageDelete,
		Description: heredoc.Doc(`
			Provides a resource to customize the pages your end users will see
			when trying to reach applications behind Cloudflare Access.
		`),
	}
}

func resourceCloudflareAccessCustomPageRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	identifier, err := initIdentifier(d)
	if err != nil {
		return diag.FromErr(err)
	}

	accessCustomPage, err := client.GetAccessCustomPage(ctx, identifier, d.Id())
	if err != nil {
		var notFoundError *cloudflare.NotFoundError
		if errors.As(err, &notFoundError) {
			tflog.Info(ctx, fmt.Sprintf("Access Custom Page %s no longer exists", d.Id()))
			d.SetId("")
			return nil
		}
		return diag.FromErr(fmt.Errorf("error fetching Access Custom Page: %w", err))
	}

	d.SetId(accessCustomPage.UID)
	d.Set("name", accessCustomPage.Name)
	d.Set("type", accessCustomPage.Type)
	d.Set("custom_html", accessCustomPage.CustomHTML)
	d.Set("app_count", accessCustomPage.AppCount)

	return nil
}

func resourceCloudflareAccessCustomPageCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	cpType := d.Get("type").(string)
	customPage := cloudflare.CreateAccessCustomPageParams{
		Name:       d.Get("name").(string),
		Type:       cloudflare.AccessCustomPageType(cpType),
		CustomHTML: d.Get("custom_html").(string),
	}

	tflog.Debug(ctx, fmt.Sprintf("Creating Access Custom Page: %+v", customPage))

	identifier, err := initIdentifier(d)
	if err != nil {
		return diag.FromErr(err)
	}

	accessCustomPage, err := client.CreateAccessCustomPage(ctx, identifier, customPage)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating Access Custom Page: %w", err))
	}

	d.SetId(accessCustomPage.UID)

	return resourceCloudflareAccessCustomPageRead(ctx, d, meta)
}

func resourceCloudflareAccessCustomPageUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	cpType := d.Get("type").(string)
	customPage := cloudflare.UpdateAccessCustomPageParams{
		Name:       d.Get("name").(string),
		Type:       cloudflare.AccessCustomPageType(cpType),
		CustomHTML: d.Get("custom_html").(string),
		UID:        d.Id(),
	}

	tflog.Debug(ctx, fmt.Sprintf("Updating Access Custom Page: %+v", customPage))

	identifier, err := initIdentifier(d)
	if err != nil {
		return diag.FromErr(err)
	}

	accessCustomPage, err := client.UpdateAccessCustomPage(ctx, identifier, customPage)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error updating Access Custom Page: %w", err))
	}

	if accessCustomPage.UID == "" {
		return diag.FromErr(fmt.Errorf("failed to find Access Custom Page ID in update response; resource was empty"))
	}

	return resourceCloudflareAccessCustomPageRead(ctx, d, meta)
}

func resourceCloudflareAccessCustomPageDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	identifier, err := initIdentifier(d)
	if err != nil {
		return diag.FromErr(err)
	}

	err = client.DeleteAccessCustomPage(ctx, identifier, d.Id())
	if err != nil {
		return diag.FromErr(fmt.Errorf("error deleting Access Custom Page: %w", err))
	}

	d.SetId("")

	return nil
}
