package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
)

func resourceCloudflareList() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareListSchema(),
		CreateContext: resourceCloudflareListCreate,
		ReadContext:   resourceCloudflareListRead,
		UpdateContext: resourceCloudflareListUpdate,
		DeleteContext: resourceCloudflareListDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceCloudflareListImport,
		},
	}
}

func resourceCloudflareListCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get("account_id").(string)

	list, err := client.CreateList(ctx, cloudflare.ListCreateParams{
		AccountID:   accountID,
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Kind:        d.Get("kind").(string),
	})
	if err != nil {
		return diag.FromErr(errors.Wrap(err, fmt.Sprintf("error creating List %s", d.Get("name").(string))))
	}

	d.SetId(list.ID)

	if items, ok := d.GetOk("item"); ok {
		items := buildListItemsCreateRequest(items.(*schema.Set).List())
		_, err = client.CreateListItems(ctx, cloudflare.ListCreateItemsParams{
			AccountID: accountID,
			ID:        d.Id(),
			Items:     items,
		})
		if err != nil {
			return diag.FromErr(errors.Wrap(err, fmt.Sprintf("error creating List Items")))
		}
	}

	return resourceCloudflareListRead(ctx, d, meta)
}

func resourceCloudflareListImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	attributes := strings.SplitN(d.Id(), "/", 2)

	if len(attributes) != 2 {
		return nil, fmt.Errorf("invalid id (\"%s\") specified, should be in format \"accountID/listID\"", d.Id())
	}

	accountID, listID := attributes[0], attributes[1]
	d.SetId(listID)
	d.Set("account_id", accountID)

	resourceCloudflareListRead(ctx, d, meta)

	return []*schema.ResourceData{d}, nil
}

func resourceCloudflareListRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get("account_id").(string)

	list, err := client.GetList(ctx, cloudflare.ListGetParams{
		AccountID: accountID,
		ID:        d.Id(),
	})
	if err != nil {
		if strings.Contains(err.Error(), "could not find list") {
			tflog.Info(ctx, fmt.Sprintf("List %s no longer exists", d.Id()))
			d.SetId("")
			return nil
		}
		return diag.FromErr(errors.Wrap(err, fmt.Sprintf("error reading List with ID %q", d.Id())))
	}

	d.Set("name", list.Name)
	d.Set("description", list.Description)
	d.Set("kind", list.Kind)

	items, err := client.ListListItems(ctx, cloudflare.ListListItemsParams{
		AccountID: accountID,
		ID:        d.Id(),
	})
	if err != nil {
		return diag.FromErr(errors.Wrap(err, fmt.Sprintf("error reading List Items")))
	}

	var itemData []map[string]interface{}
	var item map[string]interface{}

	for _, i := range items {
		item = make(map[string]interface{})
		item["value"] = []map[string]interface{}{
			{"ip": i.IP},
		}
		item["comment"] = i.Comment

		itemData = append(itemData, item)
	}

	d.Set("item", itemData)

	return nil
}

func resourceCloudflareListUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get("account_id").(string)

	_, err := client.UpdateList(ctx, cloudflare.ListUpdateParams{
		AccountID:   accountID,
		ID:          d.Id(),
		Description: d.Get("description").(string),
	})
	if err != nil {
		return diag.FromErr(errors.Wrap(err, fmt.Sprintf("error updating List description")))
	}

	if items, ok := d.GetOk("item"); ok {
		items := buildListItemsCreateRequest(items.(*schema.Set).List())
		_, err = client.ReplaceListItems(ctx, cloudflare.ListReplaceItemsParams{
			AccountID: accountID,
			ID:        d.Id(),
			Items:     items,
		})
		if err != nil {
			return diag.FromErr(errors.Wrap(err, fmt.Sprintf("error creating List Items")))
		}
	}

	return resourceCloudflareListRead(ctx, d, meta)
}

func resourceCloudflareListDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get("account_id").(string)

	_, err := client.DeleteList(ctx, cloudflare.ListDeleteParams{
		AccountID: accountID,
		ID:        d.Id(),
	})
	if err != nil {
		return diag.FromErr(errors.Wrap(err, fmt.Sprintf("error deleting List with ID %q", d.Id())))
	}

	return nil
}

func buildListItemsCreateRequest(items []interface{}) []cloudflare.ListItemCreateRequest {
	var listItems []cloudflare.ListItemCreateRequest

	for _, item := range items {
		value := item.(map[string]interface{})["value"].([]interface{})[0]

		listItems = append(listItems, cloudflare.ListItemCreateRequest{
			IP:      cloudflare.StringPtr(value.(map[string]interface{})["ip"].(string)),
			Comment: item.(map[string]interface{})["comment"].(string),
		})
	}

	return listItems
}
