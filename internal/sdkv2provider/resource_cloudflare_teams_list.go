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
	}
}

func resourceCloudflareTeamsListCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	newTeamsList := cloudflare.CreateTeamsListParams{
		Name:        d.Get("name").(string),
		Type:        d.Get("type").(string),
		Description: d.Get("description").(string),
	}

	itemValues := d.Get("items").(*schema.Set).List()
	for _, v := range itemValues {
		newTeamsList.Items = append(newTeamsList.Items, cloudflare.TeamsListItem{Value: v.(string)})
	}

	tflog.Debug(ctx, fmt.Sprintf("Creating Cloudflare Teams List from struct: %+v", newTeamsList))

	accountID := d.Get(consts.AccountIDSchemaKey).(string)

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

	d.Set("items", convertListItemsToSchema(listItems))

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

	if d.HasChange("items") {
		oldItemsIface, newItemsIface := d.GetChange("items")
		oldItems := oldItemsIface.(*schema.Set).List()
		newItems := newItemsIface.(*schema.Set).List()
		patchTeamsList := cloudflare.PatchTeamsListParams{ID: d.Id()}
		setListItemDiff(&patchTeamsList, oldItems, newItems)

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

func setListItemDiff(patchList *cloudflare.PatchTeamsListParams, oldItems, newItems []interface{}) {
	counts := make(map[string]int)
	for _, val := range newItems {
		counts[val.(string)] += 1
	}
	for _, val := range oldItems {
		counts[val.(string)] -= 1
	}

	for key, val := range counts {
		if val > 0 {
			patchList.Append = append(patchList.Append, cloudflare.TeamsListItem{Value: key})
		}
		if val < 0 {
			patchList.Remove = append(patchList.Remove, key)
		}
	}
}

func convertListItemsToSchema(listItems []cloudflare.TeamsListItem) []string {
	itemValues := []string{}
	// The API returns items in reverse order so we iterate backwards for correct ordering.
	for i := len(listItems) - 1; i >= 0; i-- {
		item := listItems[i]
		itemValues = append(itemValues, item.Value)
	}

	return itemValues
}
