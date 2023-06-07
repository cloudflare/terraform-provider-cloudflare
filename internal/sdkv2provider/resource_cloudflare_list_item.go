package sdkv2provider

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
)

func resourceCloudflareListItem() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareListItemSchema(),
		CreateContext: resourceCloudflareListItemCreate,
		ReadContext:   resourceCloudflareListItemRead,
		UpdateContext: resourceCloudflareListItemUpdate,
		DeleteContext: resourceCloudflareListItemDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceCloudflareListItemImport,
		},
		Description: heredoc.Doc(`
			Provides individual list items (IPs, Redirects) to be used in Edge Rules Engine
			across all zones within the same account.
		`),
	}
}

func resourceCloudflareListItemCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)
	listID := d.Get("list_id").(string)
	listItemType := listItemType(d)

	list, err := client.GetList(ctx, cloudflare.AccountIdentifier(accountID), listID)
	if err != nil {
		return diag.FromErr(fmt.Errorf("unable to find list with id %s: %w", listID, err))
	}

	if list.Kind != listItemType {
		return diag.FromErr(fmt.Errorf("items of type %s can not be added to lists of type %s", listItemType, list.Kind))
	}

	createListItemResponse, err := client.CreateListItem(ctx, cloudflare.AccountIdentifier(accountID), cloudflare.ListCreateItemParams{
		ID:   listID,
		Item: buildListItemCreateRequest(d),
	})
	if err != nil {
		return diag.FromErr(fmt.Errorf("unable to create list item on list id %s: %w", listID, err))
	}

	newestItem := mostRecentlyCreatedItem(createListItemResponse)
	d.SetId(newestItem.ID)

	return nil
}

func resourceCloudflareListItemImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	attributes := strings.SplitN(d.Id(), "/", 3)

	if len(attributes) != 3 {
		return nil, fmt.Errorf("invalid id (\"%s\") specified, should be in format \"accountID/listID/itemID\"", d.Id())
	}

	accountID, listID, itemID := attributes[0], attributes[1], attributes[2]
	d.SetId(itemID)
	d.Set(consts.AccountIDSchemaKey, accountID)
	d.Set("list_id", listID)

	resourceCloudflareListItemRead(ctx, d, meta)

	return []*schema.ResourceData{d}, nil
}

func resourceCloudflareListItemRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)
	listID := d.Get("list_id").(string)

	listItem, err := client.GetListItem(ctx, cloudflare.AccountIdentifier(accountID), listID, d.Id())
	if err != nil {
		var notFoundError *cloudflare.NotFoundError
		if errors.As(err, &notFoundError) {
			tflog.Info(ctx, fmt.Sprintf("List item %s no longer exists", d.Id()))
			d.SetId("")
			return nil
		}
		return diag.FromErr(errors.Wrap(err, fmt.Sprintf("error reading List Item with ID %q", d.Id())))
	}

	d.Set("comment", listItem.Comment)

	if listItem.IP != nil {
		d.Set("ip", listItem.IP)
	}

	if listItem.ASN != nil {
		asn := int(*listItem.ASN)
		d.Set("asn", asn)
	}

	if listItem.Hostname != nil {
		d.Set("url_hostname", listItem.Hostname.UrlHostname)
	}

	if listItem.Redirect != nil {
		optBoolToString := func(b *bool) string {
			if b != nil {
				switch *b {
				case true:
					return "enabled"
				case false:
					return "disabled"
				}
			}
			return ""
		}

		d.Set("source_url", listItem.Redirect.SourceUrl)
		d.Set("include_subdomains", optBoolToString(listItem.Redirect.IncludeSubdomains))
		d.Set("target_url", listItem.Redirect.TargetUrl)
		d.Set("status_code", listItem.Redirect.StatusCode)
		d.Set("preserve_query_string", optBoolToString(listItem.Redirect.PreserveQueryString))
		d.Set("subpath_matching", optBoolToString(listItem.Redirect.SubpathMatching))
		d.Set("preserve_path_suffix", optBoolToString(listItem.Redirect.PreservePathSuffix))
	}

	return nil
}

func resourceCloudflareListItemUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	cr := resourceCloudflareListItemCreate(ctx, d, meta)
	if cr != nil {
		return cr
	}

	return resourceCloudflareListItemRead(ctx, d, meta)
}

func resourceCloudflareListItemDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)
	listID := d.Get("list_id").(string)

	_, err := client.DeleteListItems(ctx, cloudflare.AccountIdentifier(accountID), cloudflare.ListDeleteItemsParams{
		ID: listID,
		Items: cloudflare.ListItemDeleteRequest{
			Items: []cloudflare.ListItemDeleteItemRequest{{ID: d.Id()}},
		},
	})
	if err != nil {
		return diag.FromErr(errors.Wrap(err, fmt.Sprintf("error removing List Item %s from list %s", listID, d.Id())))
	}

	return nil
}

func listItemType(d *schema.ResourceData) string {
	if _, ok := d.GetOk("ip"); ok {
		return "ip"
	}

	if _, ok := d.GetOk("redirect"); ok {
		return "redirect"
	}

	if _, ok := d.GetOk("hostname"); ok {
		return "hostname"
	}

	if _, ok := d.GetOk("asn"); ok {
		return "asn"
	}

	return ""
}

func buildListItemCreateRequest(d *schema.ResourceData) cloudflare.ListItemCreateRequest {
	itemType := listItemType(d)

	stringToOptBool := func(r map[string]interface{}, s string) *bool {
		switch r[s] {
		case "enabled":
			return cloudflare.BoolPtr(true)
		case "disabled":
			return cloudflare.BoolPtr(false)
		default:
			return nil
		}
	}

	request := cloudflare.ListItemCreateRequest{
		Comment: d.Get("comment").(string),
	}

	switch itemType {
	case "ip":
		request.IP = cloudflare.StringPtr(d.Get("ip").(string))
	case "asn":
		request.ASN = cloudflare.Uint32Ptr(uint32(d.Get("asn").(int)))
	case "hostname":
		hostname := d.Get("hostname").([]interface{})[0].(map[string]interface{})
		request.Hostname = &cloudflare.Hostname{
			UrlHostname: *cloudflare.StringPtr(hostname["url_hostname"].(string)),
		}
	case "redirect":
		redirect := d.Get("redirect").([]interface{})[0].(map[string]interface{})
		request.Redirect = &cloudflare.Redirect{
			SourceUrl: redirect["source_url"].(string),
			TargetUrl: redirect["target_url"].(string),
		}

		if value, ok := redirect["status_code"]; ok && value != 0 {
			request.Redirect.StatusCode = cloudflare.IntPtr(value.(int))
		}

		request.Redirect.IncludeSubdomains = stringToOptBool(redirect, "include_subdomains")
		request.Redirect.PreserveQueryString = stringToOptBool(redirect, "preserve_query_string")
		request.Redirect.SubpathMatching = stringToOptBool(redirect, "subpath_matching")
		request.Redirect.PreservePathSuffix = stringToOptBool(redirect, "preserve_path_suffix")
	}
	return request
}

func mostRecentlyCreatedItem(createListItemResponse []cloudflare.ListItem) cloudflare.ListItem {
	sort.Slice(createListItemResponse, func(i, j int) bool {
		return createListItemResponse[i].CreatedOn.After(*createListItemResponse[j].CreatedOn)
	})

	return createListItemResponse[0]
}
