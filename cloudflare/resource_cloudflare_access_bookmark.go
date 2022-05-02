package cloudflare

import (
	"context"
	"fmt"
	"log"
	"strings"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareAccessBookmark() *schema.Resource {
	return &schema.Resource{
		Schema: resourceCloudflareAccessBookmarkSchema(),
		CreateContext: resourceCloudflareAccessBookmarkCreate,
		ReadContext: resourceCloudflareAccessBookmarkRead,
		UpdateContext: resourceCloudflareAccessBookmarkUpdate,
		DeleteContext: resourceCloudflareAccessBookmarkDelete,
		Importer: &schema.ResourceImporter{
			State: resourceCloudflareAccessBookmarkImport,
		},
	}
}

func resourceCloudflareAccessBookmarkCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)

	newAccessBookmark := cloudflare.AccessBookmark{
		Name:               d.Get("name").(string),
		Domain:             d.Get("domain").(string),
		LogoURL:            d.Get("logo_url").(string),
		AppLauncherVisible: d.Get("app_launcher_visible").(bool),
	}

	log.Printf("[DEBUG] Creating Cloudflare Access Bookmark from struct: %+v", newAccessBookmark)

	identifier, err := initIdentifier(d)
	if err != nil {
		return err
	}

	var accessBookmark cloudflare.AccessBookmark
	if identifier.Type == AccountType {
		accessBookmark, err = client.CreateAccessBookmark(context.Background(), identifier.Value, newAccessBookmark)
	} else {
		accessBookmark, err = client.CreateZoneLevelAccessBookmark(context.Background(), identifier.Value, newAccessBookmark)
	}
	if err != nil {
		return fmt.Errorf("error creating Access Bookmark for %s %q: %s", identifier.Type, identifier.Value, err)
	}

	d.SetId(accessBookmark.ID)

	return resourceCloudflareAccessBookmarkRead(d, meta)
}

func resourceCloudflareAccessBookmarkRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)

	identifier, err := initIdentifier(d)
	if err != nil {
		return err
	}

	var accessBookmark cloudflare.AccessBookmark
	if identifier.Type == AccountType {
		accessBookmark, err = client.AccessBookmark(context.Background(), identifier.Value, d.Id())
	} else {
		accessBookmark, err = client.ZoneLevelAccessBookmark(context.Background(), identifier.Value, d.Id())
	}

	if err != nil {
		if strings.Contains(err.Error(), "HTTP status 404") {
			log.Printf("[INFO] Access Bookmark %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("error finding Access Bookmark %q: %s", d.Id(), err)
	}

	d.Set("name", accessBookmark.Name)
	d.Set("domain", accessBookmark.Domain)
	d.Set("logo_url", accessBookmark.LogoURL)
	d.Set("app_launcher_visible", accessBookmark.AppLauncherVisible)

	return nil
}

func resourceCloudflareAccessBookmarkUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)

	updatedAccessBookmark := cloudflare.AccessBookmark{
		ID:                 d.Id(),
		Name:               d.Get("name").(string),
		Domain:             d.Get("domain").(string),
		LogoURL:            d.Get("logo_url").(string),
		AppLauncherVisible: d.Get("app_launcher_visible").(bool),
	}

	log.Printf("[DEBUG] Updating Cloudflare Access Bookmark from struct: %+v", updatedAccessBookmark)

	identifier, err := initIdentifier(d)
	if err != nil {
		return err
	}

	var accessBookmark cloudflare.AccessBookmark
	if identifier.Type == AccountType {
		accessBookmark, err = client.UpdateAccessBookmark(context.Background(), identifier.Value, updatedAccessBookmark)
	} else {
		accessBookmark, err = client.UpdateZoneLevelAccessBookmark(context.Background(), identifier.Value, updatedAccessBookmark)
	}
	if err != nil {
		return fmt.Errorf("error updating Access Bookmark for %s %q: %s", identifier.Type, identifier.Value, err)
	}

	if accessBookmark.ID == "" {
		return fmt.Errorf("failed to find Access Bookmark ID in update response; resource was empty")
	}

	return resourceCloudflareAccessBookmarkRead(d, meta)
}

func resourceCloudflareAccessBookmarkDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	bookmarkID := d.Id()

	log.Printf("[DEBUG] Deleting Cloudflare Access Bookmark using ID: %s", bookmarkID)

	identifier, err := initIdentifier(d)
	if err != nil {
		return err
	}

	if identifier.Type == AccountType {
		err = client.DeleteAccessBookmark(context.Background(), identifier.Value, bookmarkID)
	} else {
		err = client.DeleteZoneLevelAccessBookmark(context.Background(), identifier.Value, bookmarkID)
	}
	if err != nil {
		return fmt.Errorf("error deleting Access Bookmark for %s %q: %s", identifier.Type, identifier.Value, err)
	}

	resourceCloudflareAccessBookmarkRead(d, meta)

	return nil
}

func resourceCloudflareAccessBookmarkImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	attributes := strings.SplitN(d.Id(), "/", 2)

	if len(attributes) != 2 {
		return nil, fmt.Errorf("invalid id (\"%s\") specified, should be in format \"accountID/accessBookmarkID\"", d.Id())
	}

	accountID, accessBookmarkID := attributes[0], attributes[1]

	log.Printf("[DEBUG] Importing Cloudflare Access Bookmark: id %s for account %s", accessBookmarkID, accountID)

	d.Set("account_id", accountID)
	d.SetId(accessBookmarkID)

	resourceCloudflareAccessBookmarkRead(d, meta)

	return []*schema.ResourceData{d}, nil
}
