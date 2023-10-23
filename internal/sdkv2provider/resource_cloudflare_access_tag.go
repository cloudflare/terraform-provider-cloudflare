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

func resourceCloudflareAccessTag() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareAccessTagSchema(),
		CreateContext: resourceCloudflareAccessTagCreate,
		ReadContext:   resourceCloudflareAccessTagRead,
		DeleteContext: resourceCloudflareAccessTagDelete,
		UpdateContext: schema.NoopContext,
		Description: heredoc.Doc(`
			Provides a resource to customize the pages your end users will see
			when trying to reach applications behind Cloudflare Access.
		`),
	}
}

func resourceCloudflareAccessTagRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	identifier, err := initIdentifier(d)
	if err != nil {
		return diag.FromErr(err)
	}

	accessTag, err := client.GetAccessTag(ctx, identifier, d.Id())
	if err != nil {
		var notFoundError *cloudflare.NotFoundError
		if errors.As(err, &notFoundError) {
			tflog.Info(ctx, fmt.Sprintf("Access Tag %s no longer exists", d.Id()))
			d.SetId("")
			return nil
		}
		return diag.FromErr(fmt.Errorf("error fetching Access Tag: %w", err))
	}

	d.Set("name", accessTag.Name)
	d.Set("app_count", accessTag.AppCount)

	return nil
}

func resourceCloudflareAccessTagCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	Tag := cloudflare.CreateAccessTagParams{
		Name: d.Get("name").(string),
	}

	tflog.Debug(ctx, fmt.Sprintf("Creating Access Tag: %+v", Tag))

	identifier, err := initIdentifier(d)
	if err != nil {
		return diag.FromErr(err)
	}

	accessTag, err := client.CreateAccessTag(ctx, identifier, Tag)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating Access Tag: %w", err))
	}

	d.SetId(accessTag.Name)

	return resourceCloudflareAccessTagRead(ctx, d, meta)
}

func resourceCloudflareAccessTagDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	identifier, err := initIdentifier(d)
	if err != nil {
		return diag.FromErr(err)
	}

	err = client.DeleteAccessTag(ctx, identifier, d.Id())
	if err != nil {
		return diag.FromErr(fmt.Errorf("error deleting Access Tag: %w", err))
	}

	d.SetId("")

	return nil
}
