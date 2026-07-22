// This file contains hand-written logic that extends the generated
// resource.go for cloudflare_zero_trust_list; it is not itself generated.
//
// resource.go hooks into this file in two places:
//   - Read() calls fetchItems to populate data.Items. The list metadata
//     endpoint (GatewayListService.Get) never returns items — they live
//     behind the separate, paginated GatewayListItemService.List endpoint —
//     so without this, Read() overwrites state's items with nothing on every
//     refresh, and the next plan always shows a spurious diff.
//   - ModifyPlan() calls suppressItemsDiff to avoid showing a "changed" plan
//     when config and refreshed state hold the same items, just in a
//     different order (items is a set) or freshly re-fetched from the API.

package zero_trust_list

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"sort"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// fetchItems retrieves every item of a gateway list via the dedicated items
// endpoint (GET /accounts/{account_id}/gateway/lists/{list_id}/items).
//
// Return values:
//   - (items, nil): success; items is always non-nil (an empty slice means
//     the list exists with zero items).
//   - (nil, nil)  : the list itself was deleted (404); the caller should
//     remove the resource from state.
//   - (nil, err): an unexpected error occurred.
func fetchItems(ctx context.Context, client *cloudflare.Client, accountID, listID string) ([]*ZeroTrustListItemsModel, error) {
	page, err := client.ZeroTrust.Gateway.Lists.Items.List(
		ctx,
		listID,
		zero_trust.GatewayListItemListParams{
			AccountID: cloudflare.F(accountID),
		},
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		var apiErr *cloudflare.Error
		if errors.As(err, &apiErr) && apiErr.StatusCode == 404 {
			return nil, nil
		}
		return nil, err
	}

	items := make([]*ZeroTrustListItemsModel, 0, len(page.Result))
	for _, item := range page.Result {
		items = append(items, &ZeroTrustListItemsModel{
			Value:       types.StringValue(item.Value),
			Description: gatewayItemDescription(item),
		})
	}
	return items, nil
}

// gatewayItemDescription maps a GatewayItem's description to a Terraform
// string, returning StringNull() when the API left the field null or absent
// so refreshed state matches configs that omit description entirely.
func gatewayItemDescription(item zero_trust.GatewayItem) types.String {
	if item.JSON.Description.IsNull() || item.JSON.Description.IsInvalid() {
		return types.StringNull()
	}
	return types.StringValue(item.Description)
}

// suppressItemsDiff normalizes the plan's items to state's whenever the two
// are semantically identical (same values/descriptions, any order): items is
// a set, so neither a config reorder nor items simply being re-fetched from
// the API in a different order should ever surface as a plan change.
func suppressItemsDiff(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	// Nothing to compare on create (state is null) or destroy (plan is null).
	if req.State.Raw.IsNull() || req.Plan.Raw.IsNull() ||
		!req.State.Raw.IsKnown() || !req.Plan.Raw.IsKnown() {
		return
	}

	var plan, state *ZeroTrustListModel

	// Local diagnostics: a decode failure here should fall back to
	// Terraform's own diffing instead of surfacing as a provider error.
	var planDiags, stateDiags diag.Diagnostics
	planDiags.Append(req.Plan.Get(ctx, &plan)...)
	stateDiags.Append(req.State.Get(ctx, &state)...)
	if planDiags.HasError() || stateDiags.HasError() || plan == nil || state == nil {
		return
	}

	if plan.Items == nil && state.Items == nil {
		return
	}

	emptyItems := []*ZeroTrustListItemsModel{}
	planItems, stateItems := plan.Items, state.Items
	if planItems == nil {
		planItems = &emptyItems
	}
	if stateItems == nil {
		stateItems = &emptyItems
	}

	// An unknown value can't be compared yet; let Terraform show the diff.
	for _, item := range *planItems {
		if item != nil && (item.Value.IsUnknown() || item.Description.IsUnknown()) {
			return
		}
	}

	if computeItemsHash(*planItems) != computeItemsHash(*stateItems) {
		return
	}

	filtered := make([]*ZeroTrustListItemsModel, 0, len(*stateItems))
	for _, item := range *stateItems {
		if item != nil {
			filtered = append(filtered, item)
		}
	}
	resp.Diagnostics.Append(resp.Plan.SetAttribute(ctx, path.Root("items"), &filtered)...)
}

// computeItemsHash returns a deterministic, order-independent SHA-256 hash of
// a set of list items: two sets holding the same (value, description) pairs
// hash identically regardless of order. Each field is length-prefixed and
// tagged with its null/unknown/string state so that, for example, a null
// description can never collide with an empty-string one.
func computeItemsHash(items []*ZeroTrustListItemsModel) [32]byte {
	encoded := make([]string, 0, len(items))
	for _, item := range items {
		if item == nil {
			continue
		}
		encoded = append(encoded, encodeItemField("v", item.Value)+"/"+encodeItemField("d", item.Description))
	}
	sort.Strings(encoded)

	h := sha256.New()
	for _, v := range encoded {
		h.Write([]byte(v))
		h.Write([]byte{0})
	}
	var result [32]byte
	copy(result[:], h.Sum(nil))
	return result
}

// encodeItemField renders a single string attribute for hashing, tagged so
// null, unknown, and empty-string values can never collide with one another.
func encodeItemField(tag string, v types.String) string {
	switch {
	case v.IsUnknown():
		return tag + ":unknown"
	case v.IsNull():
		return tag + ":null"
	default:
		raw := v.ValueString()
		return fmt.Sprintf("%s(%d):%s", tag, len(raw), raw)
	}
}
