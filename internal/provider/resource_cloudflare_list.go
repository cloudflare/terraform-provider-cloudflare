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
		Description: "Provides Lists (IPs, Redirects) to be used in Edge Rules Engine across all zones within the same account.",
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
		items := buildListItemsCreateRequest(d, items.([]interface{}))
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

		value := make(map[string]interface{})

		if i.IP != nil {
			value["ip"] = *i.IP
		}
		if i.Redirect != nil {
			value["redirect"] = []map[string]interface{}{{
				"source_url":            i.Redirect.SourceUrl,
				"include_subdomains":    i.Redirect.IncludeSubdomains,
				"target_url":            i.Redirect.TargetUrl,
				"status_code":           i.Redirect.StatusCode,
				"preserve_query_string": i.Redirect.PreserveQueryString,
				"subpath_matching":      i.Redirect.SubpathMatching,
				"preserve_path_suffix":  i.Redirect.PreservePathSuffix,
			}}
		}

		item["value"] = []map[string]interface{}{value}
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
		items := buildListItemsCreateRequest(d, items.([]interface{}))
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

func buildListItemsCreateRequest(resource *schema.ResourceData, items []interface{}) []cloudflare.ListItemCreateRequest {
	var listItems []cloudflare.ListItemCreateRequest

	for i, item := range items {
		value := item.(map[string]interface{})["value"].([]interface{})[0].(map[string]interface{})

		_, hasIP := resource.GetOkExists(fmt.Sprintf("item.%d.value.0.ip", i))

		var ip *string = nil
		if hasIP {
			maybeIP := value["ip"].(string)
			ip = &maybeIP
		}

		_, hasRedirect := resource.GetOkExists(fmt.Sprintf("item.%d.value.0.redirect", i))

		var redirect *cloudflare.Redirect = nil
		if hasRedirect {
			r := value["redirect"].([]interface{})[0].(map[string]interface{})

			sourceUrl := r["source_url"].(string)
			targetUrl := r["target_url"].(string)

			var includeSubdomains *bool = nil
			var subpathMatching *bool = nil
			var statusCode *int = nil
			var preserveQueryString *bool = nil
			var preservePathSuffix *bool = nil

			hasField := func(field string) bool {
				_, has := resource.GetOkExists(fmt.Sprintf("item.%d.value.0.redirect.0.%s", i, field))
				return has
			}

			if hasField("include_subdomains") {
				includeSubdomains = cloudflare.BoolPtr(r["include_subdomains"].(bool))
			}
			if hasField("subpath_matching") {
				subpathMatching = cloudflare.BoolPtr(r["subpath_matching"].(bool))
			}
			if hasField("status_code") {
				statusCode = cloudflare.IntPtr(r["status_code"].(int))
			}
			if hasField("preserve_query_string") {
				preserveQueryString = cloudflare.BoolPtr(r["preserve_query_string"].(bool))
			}
			if hasField("preserve_path_suffix") {
				preservePathSuffix = cloudflare.BoolPtr(r["preserve_path_suffix"].(bool))
			}

			redirect = &cloudflare.Redirect{
				SourceUrl:           sourceUrl,
				IncludeSubdomains:   includeSubdomains,
				TargetUrl:           targetUrl,
				StatusCode:          statusCode,
				PreserveQueryString: preserveQueryString,
				SubpathMatching:     subpathMatching,
				PreservePathSuffix:  preservePathSuffix,
			}
		}

		listItems = append(listItems, cloudflare.ListItemCreateRequest{
			IP:       ip,
			Redirect: redirect,
			Comment:  item.(map[string]interface{})["comment"].(string),
		})
	}

	return listItems
}
