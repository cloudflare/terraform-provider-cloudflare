package provider

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/MakeNowJust/heredoc/v2"
	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareAccessBookmark() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareAccessBookmarkSchema(),
		CreateContext: resourceCloudflareAccessBookmarkCreate,
		ReadContext:   resourceCloudflareAccessBookmarkRead,
		UpdateContext: resourceCloudflareAccessBookmarkUpdate,
		DeleteContext: resourceCloudflareAccessBookmarkDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceCloudflareAccessBookmarkImport,
		},
		Description: heredoc.Doc(`
			Provides a Cloudflare Access Bookmark resource. Access Bookmark
			applications are not protected behind Access but are displayed in
			the App Launcher.
		`),
		DeprecationMessage: "This resource is deprecated in favor of the cloudflare_access_application resource: https://registry.terraform.io/providers/cloudflare/cloudflare/latest/docs/resources/access_application.",
	}
}

func resourceCloudflareAccessBookmarkCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	newAccessBookmark := cloudflare.AccessBookmark{
		Name:               d.Get("name").(string),
		Domain:             d.Get("domain").(string),
		LogoURL:            d.Get("logo_url").(string),
		AppLauncherVisible: cloudflare.BoolPtr(d.Get("app_launcher_visible").(bool)),
	}

	tflog.Debug(ctx, fmt.Sprintf("Creating Cloudflare Access Bookmark from struct: %+v", newAccessBookmark))

	identifier, err := initIdentifier(d)
	if err != nil {
		return diag.FromErr(err)
	}

	var accessBookmark cloudflare.AccessBookmark
	if identifier.Type == AccountType {
		accessBookmark, err = client.CreateAccessBookmark(ctx, identifier.Value, newAccessBookmark)
	} else {
		accessBookmark, err = client.CreateZoneLevelAccessBookmark(ctx, identifier.Value, newAccessBookmark)
	}
	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating Access Bookmark for %s %q: %w", identifier.Type, identifier.Value, err))
	}

	d.SetId(accessBookmark.ID)

	return resourceCloudflareAccessBookmarkRead(ctx, d, meta)
}

func resourceCloudflareAccessBookmarkRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	identifier, err := initIdentifier(d)
	if err != nil {
		return diag.FromErr(err)
	}

	var accessBookmark cloudflare.AccessBookmark
	if identifier.Type == AccountType {
		accessBookmark, err = client.AccessBookmark(ctx, identifier.Value, d.Id())
	} else {
		accessBookmark, err = client.ZoneLevelAccessBookmark(ctx, identifier.Value, d.Id())
	}

	if err != nil {
		var notFoundError *cloudflare.NotFoundError
		if errors.As(err, &notFoundError) {
			tflog.Info(ctx, fmt.Sprintf("Access Bookmark %s no longer exists", d.Id()))
			d.SetId("")
			return nil
		}
		return diag.FromErr(fmt.Errorf("error finding Access Bookmark %q: %w", d.Id(), err))
	}

	d.Set("name", accessBookmark.Name)
	d.Set("domain", accessBookmark.Domain)
	d.Set("logo_url", accessBookmark.LogoURL)
	d.Set("app_launcher_visible", accessBookmark.AppLauncherVisible)

	return nil
}

func resourceCloudflareAccessBookmarkUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	updatedAccessBookmark := cloudflare.AccessBookmark{
		ID:                 d.Id(),
		Name:               d.Get("name").(string),
		Domain:             d.Get("domain").(string),
		LogoURL:            d.Get("logo_url").(string),
		AppLauncherVisible: cloudflare.BoolPtr(d.Get("app_launcher_visible").(bool)),
	}

	tflog.Debug(ctx, fmt.Sprintf("Updating Cloudflare Access Bookmark from struct: %+v", updatedAccessBookmark))

	identifier, err := initIdentifier(d)
	if err != nil {
		return diag.FromErr(err)
	}

	var accessBookmark cloudflare.AccessBookmark
	if identifier.Type == AccountType {
		accessBookmark, err = client.UpdateAccessBookmark(ctx, identifier.Value, updatedAccessBookmark)
	} else {
		accessBookmark, err = client.UpdateZoneLevelAccessBookmark(ctx, identifier.Value, updatedAccessBookmark)
	}
	if err != nil {
		return diag.FromErr(fmt.Errorf("error updating Access Bookmark for %s %q: %w", identifier.Type, identifier.Value, err))
	}

	if accessBookmark.ID == "" {
		return diag.FromErr(fmt.Errorf("failed to find Access Bookmark ID in update response; resource was empty"))
	}

	return resourceCloudflareAccessBookmarkRead(ctx, d, meta)
}

func resourceCloudflareAccessBookmarkDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	bookmarkID := d.Id()

	tflog.Debug(ctx, fmt.Sprintf("Deleting Cloudflare Access Bookmark using ID: %s", bookmarkID))

	identifier, err := initIdentifier(d)
	if err != nil {
		return diag.FromErr(err)
	}

	if identifier.Type == AccountType {
		err = client.DeleteAccessBookmark(ctx, identifier.Value, bookmarkID)
	} else {
		err = client.DeleteZoneLevelAccessBookmark(ctx, identifier.Value, bookmarkID)
	}
	if err != nil {
		return diag.FromErr(fmt.Errorf("error deleting Access Bookmark for %s %q: %w", identifier.Type, identifier.Value, err))
	}

	readErr := resourceCloudflareAccessBookmarkRead(ctx, d, meta)
	if readErr != nil {
		return readErr
	}

	return nil
}

func resourceCloudflareAccessBookmarkImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	attributes := strings.SplitN(d.Id(), "/", 2)

	if len(attributes) != 2 {
		return nil, fmt.Errorf("invalid id (\"%s\") specified, should be in format \"accountID/accessBookmarkID\"", d.Id())
	}

	accountID, accessBookmarkID := attributes[0], attributes[1]

	tflog.Debug(ctx, fmt.Sprintf("Importing Cloudflare Access Bookmark: id %s for account %s", accessBookmarkID, accountID))

	d.Set("account_id", accountID)
	d.SetId(accessBookmarkID)

	readErr := resourceCloudflareAccessBookmarkRead(ctx, d, meta)
	if readErr != nil {
		return nil, errors.New("failed to read Access Bookmark state")
	}

	return []*schema.ResourceData{d}, nil
}
