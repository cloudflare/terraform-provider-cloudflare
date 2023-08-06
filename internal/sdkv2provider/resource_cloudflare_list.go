package sdkv2provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
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
		Description: heredoc.Doc(`
			Provides Lists (IPs, Redirects, Hostname, ASNs) to be used in Edge
			Rules Engine across all zones within the same account.
		`),
	}
}

func resourceCloudflareListCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)

	list, err := client.CreateList(ctx, cloudflare.AccountIdentifier(accountID), cloudflare.ListCreateParams{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Kind:        d.Get("kind").(string),
	})
	if err != nil {
		return diag.FromErr(errors.Wrap(err, fmt.Sprintf("error creating List %s", d.Get("name").(string))))
	}

	d.SetId(list.ID)

	if itemData, ok := d.GetOk("item"); ok {
		items := buildListItemsCreateRequest(itemData.(*schema.Set).List())
		_, err = client.CreateListItems(ctx, cloudflare.AccountIdentifier(accountID), cloudflare.ListCreateItemsParams{
			ID:    d.Id(),
			Items: items,
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
	d.Set(consts.AccountIDSchemaKey, accountID)

	var itemData []map[string]interface{}
	var item map[string]interface{}
	item = make(map[string]interface{})

	item["comment"] = "Trigger import"

	itemData = append(itemData, item)
	d.Set("item", itemData)

	resourceCloudflareListRead(ctx, d, meta)

	return []*schema.ResourceData{d}, nil
}

func resourceCloudflareListRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)

	list, err := client.GetList(ctx, cloudflare.AccountIdentifier(accountID), d.Id())
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

	if !d.HasChange("items") && len(d.Get("item").(*schema.Set).List()) == 0 {
		return nil
	}

	items, err := client.ListListItems(ctx, cloudflare.AccountIdentifier(accountID), cloudflare.ListListItemsParams{
		ID: d.Id(),
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
		if i.ASN != nil {
			value["asn"] = *i.ASN
		}
		if i.Hostname != nil {
			value["hostname"] = []map[string]interface{}{{
				"url_hostname": i.Hostname.UrlHostname,
			}}
		}
		if i.Redirect != nil {
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
			statusCode := 0
			if i.Redirect.StatusCode != nil {
				statusCode = *i.Redirect.StatusCode
			}

			value["redirect"] = []map[string]interface{}{{
				"source_url":            i.Redirect.SourceUrl,
				"include_subdomains":    optBoolToString(i.Redirect.IncludeSubdomains),
				"target_url":            i.Redirect.TargetUrl,
				"status_code":           statusCode,
				"preserve_query_string": optBoolToString(i.Redirect.PreserveQueryString),
				"subpath_matching":      optBoolToString(i.Redirect.SubpathMatching),
				"preserve_path_suffix":  optBoolToString(i.Redirect.PreservePathSuffix),
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
	accountID := d.Get(consts.AccountIDSchemaKey).(string)

	_, err := client.UpdateList(ctx, cloudflare.AccountIdentifier(accountID), cloudflare.ListUpdateParams{
		ID:          d.Id(),
		Description: d.Get("description").(string),
	})
	if err != nil {
		return diag.FromErr(errors.Wrap(err, fmt.Sprintf("error updating List description")))
	}

	if itemData, ok := d.GetOk("item"); ok {
		items := buildListItemsCreateRequest(itemData.(*schema.Set).List())
		_, err = client.ReplaceListItems(ctx, cloudflare.AccountIdentifier(accountID), cloudflare.ListReplaceItemsParams{
			ID:    d.Id(),
			Items: items,
		})
		if err != nil {
			return diag.FromErr(errors.Wrap(err, fmt.Sprintf("error creating List Items")))
		}
	}

	return resourceCloudflareListRead(ctx, d, meta)
}

func resourceCloudflareListDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)

	_, err := client.DeleteList(ctx, cloudflare.AccountIdentifier(accountID), d.Id())
	if err != nil {
		return diag.FromErr(errors.Wrap(err, fmt.Sprintf("error deleting List with ID %q", d.Id())))
	}

	return nil
}

func buildListItemsCreateRequest(items []interface{}) []cloudflare.ListItemCreateRequest {
	var listItems []cloudflare.ListItemCreateRequest

	for _, item := range items {
		value := item.(map[string]interface{})["value"].([]interface{})[0].(map[string]interface{})

		var ip *string = nil

		if field, ok := value["ip"]; ok {
			if field, ok := field.(string); ok {
				ip = &field
			}
		}

		var asn *uint32

		if field, ok := value["asn"]; ok {
			if field, ok := field.(int); ok {
				f := uint32(field)
				asn = &f
			}
		}

		var hostname *cloudflare.Hostname

		if field, ok := value["hostname"]; ok {
			if field, ok := field.([]interface{}); ok && len(field) > 0 {
				if field, ok := field[0].(map[string]interface{}); ok {
					if field != nil {
						urlHostname := field["url_hostname"].(string)
						hostname = &cloudflare.Hostname{
							UrlHostname: urlHostname,
						}
					}
				}
			}
		}

		var redirect *cloudflare.Redirect = nil
		var r map[string]interface{} = nil

		if field, ok := value["redirect"]; ok {
			if field, ok := field.([]interface{}); ok && len(field) > 0 {
				if field, ok := field[0].(map[string]interface{}); ok {
					r = field
				}
			}
		}

		if r != nil {
			sourceUrl := r["source_url"].(string)
			targetUrl := r["target_url"].(string)

			var includeSubdomains *bool = nil
			var subpathMatching *bool = nil
			var statusCode *int = nil
			var preserveQueryString *bool = nil
			var preservePathSuffix *bool = nil

			stringToOptBool := func(s string) *bool {
				switch s {
				case "enabled":
					return cloudflare.BoolPtr(true)
				case "disabled":
					return cloudflare.BoolPtr(false)
				default:
					return nil
				}
			}

			includeSubdomains = stringToOptBool(r["include_subdomains"].(string))
			subpathMatching = stringToOptBool(r["subpath_matching"].(string))

			vint := r["status_code"].(int)
			if vint != 0 {
				statusCode = cloudflare.IntPtr(vint)
			}

			preserveQueryString = stringToOptBool(r["preserve_query_string"].(string))
			preservePathSuffix = stringToOptBool(r["preserve_path_suffix"].(string))

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

		payload := cloudflare.ListItemCreateRequest{
			Comment: item.(map[string]interface{})["comment"].(string),
		}

		if ip != nil && *ip != "" {
			payload.IP = ip
		}

		if redirect != nil {
			payload.Redirect = redirect
		}

		if asn != nil && *asn > 0 {
			payload.ASN = asn
		}

		if hostname != nil {
			payload.Hostname = hostname
		}

		listItems = append(listItems, payload)
	}

	return listItems
}
