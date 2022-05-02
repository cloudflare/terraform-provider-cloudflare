package cloudflare

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
)

func resourceCloudflareIPList() *schema.Resource {
	return &schema.Resource{
		Schema: resourceCloudflareIPListSchema(),
		CreateContext: resourceCloudflareIPListCreate,
		ReadContext: resourceCloudflareIPListRead,
		UpdateContext: resourceCloudflareIPListUpdate,
		DeleteContext: resourceCloudflareIPListDelete,
		Importer: &schema.ResourceImporter{
			State: resourceCloudflareIPListImport,
		},
	}
}

func resourceCloudflareIPListCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get("account_id").(string)

	list, err := client.CreateIPList(context.Background(), accountID, d.Get("name").(string), d.Get("description").(string), d.Get("kind").(string))
	if err != nil {
		return err.Wrap(err, fmt.Sprintf("error creating IP List %s", d.Get("name").(string)))
	}

	d.SetId(list.ID)

	if items, ok := d.GetOk("item"); ok {
		IPListItems := buildIPListItemsCreateRequest(items.(*schema.Set).List())
		_, err = client.CreateIPListItems(context.Background(), accountID, d.Id(), IPListItems)
		if err != nil {
			return err.Wrap(err, fmt.Sprintf("error creating IP List Items"))
		}
	}

	return resourceCloudflareIPListRead(d, meta)
}

func resourceCloudflareIPListImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	attributes := strings.SplitN(d.Id(), "/", 2)

	if len(attributes) != 2 {
		return nil, fmt.Errorf("invalid id (\"%s\") specified, should be in format \"accountID/listID\"", d.Id())
	}

	accountID, listID := attributes[0], attributes[1]
	d.SetId(listID)
	d.Set("account_id", accountID)

	resourceCloudflareIPListRead(d, meta)

	return []*schema.ResourceData{d}, nil
}

func resourceCloudflareIPListRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get("account_id").(string)

	list, err := client.GetIPList(context.Background(), accountID, d.Id())
	if err != nil {
		if strings.Contains(err.Error(), "could not find list") {
			log.Printf("[INFO] IP List %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return err.Wrap(err, fmt.Sprintf("error reading IP List with ID %q", d.Id()))
	}

	d.Set("name", list.Name)
	d.Set("description", list.Description)
	d.Set("kind", list.Kind)

	items, err := client.ListIPListItems(context.Background(), accountID, d.Id())
	if err != nil {
		return err.Wrap(err, fmt.Sprintf("error reading IP List Items"))
	}

	var itemData []map[string]interface{}
	var item map[string]interface{}

	for _, i := range items {
		item = make(map[string]interface{})
		item["value"] = i.IP
		item["comment"] = i.Comment

		itemData = append(itemData, item)
	}

	d.Set("item", itemData)

	return nil
}

func resourceCloudflareIPListUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get("account_id").(string)

	_, err := client.UpdateIPList(context.Background(), accountID, d.Id(), d.Get("description").(string))
	if err != nil {
		return err.Wrap(err, fmt.Sprintf("error updating IP List description"))
	}

	if items, ok := d.GetOk("item"); ok {
		IPListItems := buildIPListItemsCreateRequest(items.(*schema.Set).List())
		_, err = client.ReplaceIPListItems(context.Background(), accountID, d.Id(), IPListItems)
		if err != nil {
			return err.Wrap(err, fmt.Sprintf("error creating IP List Items"))
		}
	}

	return resourceCloudflareIPListRead(d, meta)
}

func resourceCloudflareIPListDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get("account_id").(string)

	_, err := client.DeleteIPList(context.Background(), accountID, d.Id())
	if err != nil {
		return err.Wrap(err, fmt.Sprintf("error deleting IP List with ID %q", d.Id()))
	}

	return nil
}

func buildIPListItemsCreateRequest(items []interface{}) []cloudflare.IPListItemCreateRequest {
	var IPListItems []cloudflare.IPListItemCreateRequest

	for _, item := range items {
		IPListItems = append(IPListItems, cloudflare.IPListItemCreateRequest{
			IP:      item.(map[string]interface{})["value"].(string),
			Comment: item.(map[string]interface{})["comment"].(string),
		})
	}

	return IPListItems
}
