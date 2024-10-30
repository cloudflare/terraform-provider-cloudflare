package sdkv2provider

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/MakeNowJust/heredoc/v2"
	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareTeamsList() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareTeamsListSchema(),
		CreateContext: resourceCloudflareTeamsListCreate,
		ReadContext:   resourceCloudflareTeamsListRead,
		UpdateContext: resourceCloudflareTeamsListUpdate,
		DeleteContext: resourceCloudflareTeamsListDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceCloudflareTeamsListImport,
		},
		Description: heredoc.Doc(`
			Provides a Cloudflare Teams List resource. Teams lists are
			referenced when creating secure web gateway policies or device
			posture rules.
		`),
		DeprecationMessage: "`cloudflare_teams_list` is now deprecated and will be removed in the next major version. Use `cloudflare_zero_trust_list` instead.",
	}
}

func resourceCloudflareZeroTrustList() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareTeamsListSchema(),
		CreateContext: resourceCloudflareTeamsListCreate,
		ReadContext:   resourceCloudflareTeamsListRead,
		UpdateContext: resourceCloudflareTeamsListUpdate,
		DeleteContext: resourceCloudflareTeamsListDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceCloudflareTeamsListImport,
		},
		Description: heredoc.Doc(`
			Provides a Cloudflare Teams List resource. Teams lists are
			referenced when creating secure web gateway policies or device
			posture rules.
		`),
	}
}

func resourceCloudflareTeamsListCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	accountID := d.Get(consts.AccountIDSchemaKey).(string)

	newTeamsList := cloudflare.CreateTeamsListParams{
		Name:        d.Get("name").(string),
		Type:        d.Get("type").(string),
		Description: d.Get("description").(string),
	}

	itemsWithoutDescription := d.Get("items").(*schema.Set).List()
	itemsWithDescriptionValues := d.Get("items_with_description").(*schema.Set).List()
	allItems := append([]interface{}{}, itemsWithoutDescription...)
	allItems = append(allItems, itemsWithDescriptionValues...)
	for _, v := range allItems {
		item, err := convertItemCFTeamsListItems(v)
		if err != nil {
			return diag.FromErr(fmt.Errorf("error creating Teams List for account %q: %w", accountID, err))
		}
		newTeamsList.Items = append(newTeamsList.Items, *item)
	}

	tflog.Debug(ctx, fmt.Sprintf("Creating Cloudflare Teams List from struct: %+v", newTeamsList))

	identifier := cloudflare.AccountIdentifier(accountID)
	list, err := client.CreateTeamsList(ctx, identifier, newTeamsList)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating Teams List for account %q: %w", accountID, err))
	}

	d.SetId(list.ID)

	return resourceCloudflareTeamsListRead(ctx, d, meta)
}

func resourceCloudflareTeamsListRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)

	identifier := cloudflare.AccountIdentifier(accountID)
	list, err := client.GetTeamsList(ctx, identifier, d.Id())

	if err != nil {
		var notFoundError *cloudflare.NotFoundError
		if errors.As(err, &notFoundError) {
			tflog.Info(ctx, fmt.Sprintf("Teams List %s no longer exists", d.Id()))
			d.SetId("")
			return nil
		}
		return diag.FromErr(fmt.Errorf("error finding Teams List %q: %w", d.Id(), err))
	}

	d.Set("name", list.Name)
	d.Set("type", list.Type)
	d.Set("description", list.Description)

	listItems, _, err := client.ListTeamsListItems(ctx, identifier, cloudflare.ListTeamsListItemsParams{
		ListID: d.Id(),
	})
	if err != nil {
		return diag.FromErr(fmt.Errorf("error finding Teams List %q: %w", d.Id(), err))
	}

	itemsWithoutDescription, itemsWithDescription := convertListItemsToSchema(listItems)
	// items with description and without description are processed in separate attributes,
	// so customers may mix and match these two formats instead of forcing them to adopt one style
	// The provider will stitch these fields together before processing
	// this was done to avoid having to specify all items in object format(which is clunky),
	// since terraform can not implement mixed types atm.
	d.Set("items", itemsWithoutDescription)
	d.Set("items_with_description", itemsWithDescription)

	return nil
}

func resourceCloudflareTeamsListUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	updatedTeamsList := cloudflare.UpdateTeamsListParams{
		ID:          d.Id(),
		Name:        d.Get("name").(string),
		Type:        d.Get("type").(string),
		Description: d.Get("description").(string),
	}

	tflog.Debug(ctx, fmt.Sprintf("Updating Cloudflare Teams List from struct: %+v", updatedTeamsList))

	accountID := d.Get(consts.AccountIDSchemaKey).(string)

	identifier := cloudflare.AccountIdentifier(accountID)
	teamsList, err := client.UpdateTeamsList(ctx, identifier, updatedTeamsList)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error updating Teams List for account %q: %w", accountID, err))
	}
	if teamsList.ID == "" {
		return diag.FromErr(fmt.Errorf("failed to find Teams List ID in update response; resource was empty"))
	}

	if d.HasChange("items") || d.HasChange("items_with_description") {
		oldItemsIface, newItemsIface := d.GetChange("items")
		oldItemsWithDescriptionIface, newItemsWithDescriptionIface := d.GetChange("items_with_description")

		oldItems := oldItemsIface.(*schema.Set).List()
		newItems := newItemsIface.(*schema.Set).List()
		oldItemsWithDescription := oldItemsWithDescriptionIface.(*schema.Set).List()
		newItemsWithDescription := newItemsWithDescriptionIface.(*schema.Set).List()

		convertedOldItems := []cloudflare.TeamsListItem{}
		convertedNewItems := []cloudflare.TeamsListItem{}

		for _, v := range oldItems {
			item, err := convertItemCFTeamsListItems(v)
			if err != nil {
				return diag.FromErr(fmt.Errorf("error creating Teams List for account %q: %w", accountID, err))
			}
			convertedOldItems = append(convertedOldItems, *item)
		}

		for _, v := range oldItemsWithDescription {
			item, err := convertItemCFTeamsListItems(v)
			if err != nil {
				return diag.FromErr(fmt.Errorf("error creating Teams List for account %q: %w", accountID, err))
			}
			convertedOldItems = append(convertedOldItems, *item)
		}

		for _, v := range newItems {
			item, err := convertItemCFTeamsListItems(v)
			if err != nil {
				return diag.FromErr(fmt.Errorf("error creating Teams List for account %q: %w", accountID, err))
			}
			convertedNewItems = append(convertedNewItems, *item)
		}

		for _, v := range newItemsWithDescription {
			item, err := convertItemCFTeamsListItems(v)
			if err != nil {
				return diag.FromErr(fmt.Errorf("error creating Teams List for account %q: %w", accountID, err))
			}
			convertedNewItems = append(convertedNewItems, *item)
		}

		patchTeamsList := cloudflare.PatchTeamsListParams{ID: d.Id()}
		setListItemDiff(&patchTeamsList, convertedOldItems, convertedNewItems)

		l, err := client.PatchTeamsList(ctx, identifier, patchTeamsList)

		if err != nil {
			return diag.FromErr(fmt.Errorf("error updating Teams List for account %q: %w", accountID, err))
		}

		teamsList.Items = l.Items
	}

	return resourceCloudflareTeamsListRead(ctx, d, meta)
}

func resourceCloudflareTeamsListDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	appID := d.Id()
	accountID := d.Get(consts.AccountIDSchemaKey).(string)

	tflog.Debug(ctx, fmt.Sprintf("Deleting Cloudflare Teams List using ID: %s", appID))

	identifier := cloudflare.AccountIdentifier(accountID)
	err := client.DeleteTeamsList(ctx, identifier, appID)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error deleting Teams List for account %q: %w", accountID, err))
	}

	resourceCloudflareTeamsListRead(ctx, d, meta)

	return nil
}

func resourceCloudflareTeamsListImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	attributes := strings.SplitN(d.Id(), "/", 2)

	if len(attributes) != 2 {
		return nil, fmt.Errorf("invalid id (\"%s\") specified, should be in format \"accountID/teamsListID\"", d.Id())
	}

	accountID, teamsListID := attributes[0], attributes[1]

	tflog.Debug(ctx, fmt.Sprintf("Importing Cloudflare Teams List: id %s for account %s", teamsListID, accountID))

	d.Set(consts.AccountIDSchemaKey, accountID)
	d.SetId(teamsListID)

	resourceCloudflareTeamsListRead(ctx, d, meta)

	return []*schema.ResourceData{d}, nil
}

func setListItemDiff(patchList *cloudflare.PatchTeamsListParams, oldItems, newItems []cloudflare.TeamsListItem) {
	counts := make(map[cloudflare.TeamsListItem]int)

	for _, item := range newItems {
		counts[item] += 1
	}
	for _, item := range oldItems {
		counts[item] -= 1
	}

	for item, val := range counts {
		if val > 0 {
			patchList.Append = append(patchList.Append, item)
		}
		if val < 0 {
			patchList.Remove = append(patchList.Remove, item.Value)
		}
	}
}

func convertItemCFTeamsListItems(item any) (*cloudflare.TeamsListItem, error) {
	switch item.(type) {
	case string:
		return &cloudflare.TeamsListItem{Description: "", Value: item.(string)}, nil
	case map[string]interface{}:
		return &cloudflare.TeamsListItem{Description: item.(map[string]interface{})["description"].(string), Value: item.(map[string]interface{})["value"].(string)}, nil
	}

	return nil, fmt.Errorf("invalid list item `%v`. Should be string OR {\"description\": .., \"value\": ..} object", item)
}

// this method returns array of list items without any description and map of items with description
// and value separate.
func convertListItemsToSchema(listItems []cloudflare.TeamsListItem) ([]string, []map[string]string) {
	itemValuesWithDescription := []map[string]string{}
	itemValuesWithoutDescription := []string{}
	// The API returns items in reverse order so we iterate backwards for correct ordering.
	for i := len(listItems) - 1; i >= 0; i-- {
		item := listItems[i]
		if item.Description != "" {
			itemValuesWithDescription = append(itemValuesWithDescription,
				map[string]string{"value": item.Value, "description": item.Description},
			)
		} else {
			itemValuesWithoutDescription = append(itemValuesWithoutDescription, item.Value)
		}
	}

	return itemValuesWithoutDescription, itemValuesWithDescription
}
